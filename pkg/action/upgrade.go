/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package action

import (
	"bytes"
	"fmt"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"k8s.io/client-go/discovery"

	"k8s.io/helm/pkg/chart"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/engine"
	"k8s.io/helm/pkg/hooks"
	"k8s.io/helm/pkg/kube"
	"k8s.io/helm/pkg/release"
	"k8s.io/helm/pkg/releaseutil"
	"k8s.io/helm/pkg/version"
)

// Upgrade is the action for upgrading releases.
//
// It provides the implementation of 'helm upgrade'.
type Upgrade struct {
	cfg *Configuration

	ChartPathOptions
	ValueOptions

	Install   bool
	Devel     bool
	Namespace string
	Timeout   int64
	Wait      bool
	// Values is a string containing (unparsed) YAML values.
	Values       map[string]interface{}
	DisableHooks bool
	DryRun       bool
	Force        bool
	ResetValues  bool
	ReuseValues  bool
	// Recreate will (if true) recreate pods after a rollback.
	Recreate bool
	// MaxHistory limits the maximum number of revisions saved per release
	MaxHistory int
}

// NewUpgrade creates a new Upgrade object with the given configuration.
func NewUpgrade(cfg *Configuration) *Upgrade {
	return &Upgrade{
		cfg: cfg,
	}
}

func (u *Upgrade) AddFlags(f *pflag.FlagSet) {
	f.BoolVarP(&u.Install, "install", "i", false, "if a release by this name doesn't already exist, run an install")
	f.BoolVar(&u.Devel, "devel", false, "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored.")
	f.BoolVar(&u.DryRun, "dry-run", false, "simulate an upgrade")
	f.BoolVar(&u.Recreate, "recreate-pods", false, "performs pods restart for the resource if applicable")
	f.BoolVar(&u.Force, "force", false, "force resource update through delete/recreate if needed")
	f.BoolVar(&u.DisableHooks, "no-hooks", false, "disable pre/post upgrade hooks")
	f.Int64Var(&u.Timeout, "timeout", 300, "time in seconds to wait for any individual Kubernetes operation (like Jobs for hooks)")
	f.BoolVar(&u.ResetValues, "reset-values", false, "when upgrading, reset the values to the ones built into the chart")
	f.BoolVar(&u.ReuseValues, "reuse-values", false, "when upgrading, reuse the last release's values and merge in any overrides from the command line via --set and -f. If '--reset-values' is specified, this is ignored.")
	f.BoolVar(&u.Wait, "wait", false, "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment are in a ready state before marking the release as successful. It will wait for as long as --timeout")
	f.IntVar(&u.MaxHistory, "history-max", 0, "limit the maximum number of revisions saved per release. Use 0 for no limit.")
	u.ChartPathOptions.AddFlags(f)
	u.ValueOptions.AddFlags(f)
}

// Run executes the upgrade on the given release.
func (u *Upgrade) Run(name string, chart *chart.Chart) (*release.Release, error) {
	if err := chartutil.ProcessDependencies(chart, u.Values); err != nil {
		return nil, err
	}

	if err := validateReleaseName(name); err != nil {
		return nil, errors.Errorf("upgradeRelease: Release name is invalid: %s", name)
	}
	u.cfg.Log("preparing upgrade for %s", name)
	currentRelease, upgradedRelease, err := u.prepareUpgrade(name, chart)
	if err != nil {
		return nil, err
	}

	u.cfg.Releases.MaxHistory = u.MaxHistory

	if !u.DryRun {
		u.cfg.Log("creating upgraded release for %s", name)
		if err := u.cfg.Releases.Create(upgradedRelease); err != nil {
			return nil, err
		}
	}

	u.cfg.Log("performing update for %s", name)
	res, err := u.performUpgrade(currentRelease, upgradedRelease)
	if err != nil {
		return res, err
	}

	if !u.DryRun {
		u.cfg.Log("updating status for upgraded release for %s", name)
		if err := u.cfg.Releases.Update(upgradedRelease); err != nil {
			return res, err
		}
	}

	return res, nil
}

func validateReleaseName(releaseName string) error {
	if releaseName == "" {
		return errMissingRelease
	}

	if !ValidName.MatchString(releaseName) || (len(releaseName) > releaseNameMaxLen) {
		return errInvalidName
	}

	return nil
}

// prepareUpgrade builds an upgraded release for an upgrade operation.
func (u *Upgrade) prepareUpgrade(name string, chart *chart.Chart) (*release.Release, *release.Release, error) {
	if chart == nil {
		return nil, nil, errMissingChart
	}

	// finds the deployed release with the given name
	currentRelease, err := u.cfg.Releases.Deployed(name)
	if err != nil {
		return nil, nil, err
	}

	// determine if values will be reused
	if err := u.reuseValues(chart, currentRelease); err != nil {
		return nil, nil, err
	}

	// finds the non-deleted release with the given name
	lastRelease, err := u.cfg.Releases.Last(name)
	if err != nil {
		return nil, nil, err
	}

	// Increment revision count. This is passed to templates, and also stored on
	// the release object.
	revision := lastRelease.Version + 1

	options := chartutil.ReleaseOptions{
		Name:      name,
		IsUpgrade: true,
	}

	caps, err := newCapabilities(u.cfg.Discovery)
	if err != nil {
		return nil, nil, err
	}
	valuesToRender, err := chartutil.ToRenderValues(chart, u.Values, options, caps)
	if err != nil {
		return nil, nil, err
	}

	hooks, manifestDoc, notesTxt, err := u.renderResources(chart, valuesToRender, caps.APIVersions)
	if err != nil {
		return nil, nil, err
	}

	// Store an upgraded release.
	upgradedRelease := &release.Release{
		Name:      name,
		Namespace: currentRelease.Namespace,
		Chart:     chart,
		Config:    u.Values,
		Info: &release.Info{
			FirstDeployed: currentRelease.Info.FirstDeployed,
			LastDeployed:  Timestamper(),
			Status:        release.StatusPendingUpgrade,
			Description:   "Preparing upgrade", // This should be overwritten later.
		},
		Version:  revision,
		Manifest: manifestDoc.String(),
		Hooks:    hooks,
	}

	if len(notesTxt) > 0 {
		upgradedRelease.Info.Notes = notesTxt
	}
	err = validateManifest(u.cfg.KubeClient, currentRelease.Namespace, manifestDoc.Bytes())
	return currentRelease, upgradedRelease, err
}

func (u *Upgrade) performUpgrade(originalRelease, upgradedRelease *release.Release) (*release.Release, error) {
	if u.DryRun {
		u.cfg.Log("dry run for %s", upgradedRelease.Name)
		upgradedRelease.Info.Description = "Dry run complete"
		return upgradedRelease, nil
	}

	// pre-upgrade hooks
	if !u.DisableHooks {
		if err := u.execHook(upgradedRelease.Hooks, hooks.PreUpgrade); err != nil {
			return upgradedRelease, err
		}
	} else {
		u.cfg.Log("upgrade hooks disabled for %s", upgradedRelease.Name)
	}
	if err := u.upgradeRelease(originalRelease, upgradedRelease); err != nil {
		msg := fmt.Sprintf("Upgrade %q failed: %s", upgradedRelease.Name, err)
		u.cfg.Log("warning: %s", msg)
		upgradedRelease.Info.Status = release.StatusFailed
		upgradedRelease.Info.Description = msg
		u.cfg.recordRelease(originalRelease, true)
		u.cfg.recordRelease(upgradedRelease, true)
		return upgradedRelease, err
	}

	// post-upgrade hooks
	if !u.DisableHooks {
		if err := u.execHook(upgradedRelease.Hooks, hooks.PostUpgrade); err != nil {
			return upgradedRelease, err
		}
	}

	originalRelease.Info.Status = release.StatusSuperseded
	u.cfg.recordRelease(originalRelease, true)

	upgradedRelease.Info.Status = release.StatusDeployed
	upgradedRelease.Info.Description = "Upgrade complete"

	return upgradedRelease, nil
}

// upgradeRelease performs an upgrade from current to target release
func (u *Upgrade) upgradeRelease(current, target *release.Release) error {
	cm := bytes.NewBufferString(current.Manifest)
	tm := bytes.NewBufferString(target.Manifest)
	return u.cfg.KubeClient.Update(target.Namespace, cm, tm, u.Force, u.Recreate, u.Timeout, u.Wait)
}

// reuseValues copies values from the current release to a new release if the
// new release does not have any values.
//
// If the request already has values, or if there are no values in the current
// release, this does nothing.
//
// This is skipped if the u.ResetValues flag is set, in which case the
// request values are not altered.
func (u *Upgrade) reuseValues(chart *chart.Chart, current *release.Release) error {
	if u.ResetValues {
		// If ResetValues is set, we comletely ignore current.Config.
		u.cfg.Log("resetting values to the chart's original version")
		return nil
	}

	// If the ReuseValues flag is set, we always copy the old values over the new config's values.
	if u.ReuseValues {
		u.cfg.Log("reusing the old release's values")

		// We have to regenerate the old coalesced values:
		oldVals, err := chartutil.CoalesceValues(current.Chart, current.Config)
		if err != nil {
			return errors.Wrap(err, "failed to rebuild old values")
		}

		u.Values = chartutil.CoalesceTables(current.Config, u.Values)

		chart.Values = oldVals

		return nil
	}

	if len(u.Values) == 0 && len(current.Config) > 0 {
		u.cfg.Log("copying values from %s (v%d) to new release.", current.Name, current.Version)
		u.Values = current.Config
	}
	return nil
}

func newCapabilities(dc discovery.DiscoveryInterface) (*chartutil.Capabilities, error) {
	kubeVersion, err := dc.ServerVersion()
	if err != nil {
		return nil, err
	}

	apiVersions, err := GetVersionSet(dc)
	if err != nil {
		return nil, errors.Wrap(err, "could not get apiVersions from Kubernetes")
	}

	return &chartutil.Capabilities{
		KubeVersion: kubeVersion,
		APIVersions: apiVersions,
	}, nil
}

func (u *Upgrade) renderResources(ch *chart.Chart, values chartutil.Values, vs chartutil.VersionSet) ([]*release.Hook, *bytes.Buffer, string, error) {
	if ch.Metadata.KubeVersion != "" {
		cap, _ := values["Capabilities"].(*chartutil.Capabilities)
		gitVersion := cap.KubeVersion.String()
		k8sVersion := strings.Split(gitVersion, "+")[0]
		if !version.IsCompatibleRange(ch.Metadata.KubeVersion, k8sVersion) {
			return nil, nil, "", errors.Errorf("chart requires kubernetesVersion: %s which is incompatible with Kubernetes %s", ch.Metadata.KubeVersion, k8sVersion)
		}
	}

	u.cfg.Log("rendering %s chart using values", ch.Name())
	files, err := engine.Render(ch, values)
	if err != nil {
		return nil, nil, "", err
	}

	// NOTES.txt gets rendered like all the other files, but because it's not a hook nor a resource,
	// pull it out of here into a separate file so that we can actually use the output of the rendered
	// text file. We have to spin through this map because the file contains path information, so we
	// look for terminating NOTES.txt. We also remove it from the files so that we don't have to skip
	// it in the sortHooks.
	notes := ""
	for k, v := range files {
		if strings.HasSuffix(k, notesFileSuffix) {
			// Only apply the notes if it belongs to the parent chart
			// Note: Do not use filePath.Join since it creates a path with \ which is not expected
			if k == path.Join(ch.Name(), "templates", notesFileSuffix) {
				notes = v
			}
			delete(files, k)
		}
	}

	// Sort hooks, manifests, and partials. Only hooks and manifests are returned,
	// as partials are not used after renderer.Render. Empty manifests are also
	// removed here.
	hooks, manifests, err := releaseutil.SortManifests(files, vs, releaseutil.InstallOrder)
	if err != nil {
		// By catching parse errors here, we can prevent bogus releases from going
		// to Kubernetes.
		//
		// We return the files as a big blob of data to help the user debug parser
		// errors.
		b := bytes.NewBuffer(nil)
		for name, content := range files {
			if len(strings.TrimSpace(content)) == 0 {
				continue
			}
			b.WriteString("\n---\n# Source: " + name + "\n")
			b.WriteString(content)
		}
		return nil, b, "", err
	}

	// Aggregate all valid manifests into one big doc.
	b := bytes.NewBuffer(nil)
	for _, m := range manifests {
		b.WriteString("\n---\n# Source: " + m.Name + "\n")
		b.WriteString(m.Content)
	}

	return hooks, b, notes, nil
}

func validateManifest(c kube.KubernetesClient, ns string, manifest []byte) error {
	r := bytes.NewReader(manifest)
	_, err := c.BuildUnstructured(ns, r)
	return err
}

// execHook executes all of the hooks for the given hook event.
func (u *Upgrade) execHook(hs []*release.Hook, hook string) error {
	timeout := u.Timeout
	executingHooks := []*release.Hook{}

	for _, h := range hs {
		for _, e := range h.Events {
			if string(e) == hook {
				executingHooks = append(executingHooks, h)
			}
		}
	}

	sort.Sort(hookByWeight(executingHooks))

	for _, h := range executingHooks {
		if err := deleteHookByPolicy(u.cfg, u.Namespace, h, hooks.BeforeHookCreation, hook); err != nil {
			return err
		}

		b := bytes.NewBufferString(h.Manifest)
		if err := u.cfg.KubeClient.Create(u.Namespace, b, timeout, false); err != nil {
			return errors.Wrapf(err, "warning: Hook %s %s failed", hook, h.Path)
		}
		b.Reset()
		b.WriteString(h.Manifest)

		if err := u.cfg.KubeClient.WatchUntilReady(u.Namespace, b, timeout, false); err != nil {
			// If a hook is failed, checkout the annotation of the hook to determine whether the hook should be deleted
			// under failed condition. If so, then clear the corresponding resource object in the hook
			if err := deleteHookByPolicy(u.cfg, u.Namespace, h, hooks.HookFailed, hook); err != nil {
				return err
			}
			return err
		}
	}

	// If all hooks are succeeded, checkout the annotation of each hook to determine whether the hook should be deleted
	// under succeeded condition. If so, then clear the corresponding resource object in each hook
	for _, h := range executingHooks {
		if err := deleteHookByPolicy(u.cfg, u.Namespace, h, hooks.HookSucceeded, hook); err != nil {
			return err
		}
		h.LastRun = time.Now()
	}

	return nil
}
