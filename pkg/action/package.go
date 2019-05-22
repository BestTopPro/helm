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
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh/terminal"

	"helm.sh/helm/pkg/chart"
	"helm.sh/helm/pkg/chart/loader"
	"helm.sh/helm/pkg/chartutil"
	"helm.sh/helm/pkg/provenance"
)

// Package is the action for packaging a chart.
//
// It provides the implementation of 'helm package'.
type Package struct {
	ValueOptions

	Sign             bool
	Key              string
	Keyring          string
	Version          string
	AppVersion       string
	Destination      string
	DependencyUpdate bool
}

// NewPackage creates a new Package object with the given configuration.
func NewPackage() *Package {
	return &Package{}
}

// Run executes 'helm package' against the given chart and returns the path to the packaged chart.
func (p *Package) Run(path string) (string, error) {
	ch, err := loader.LoadDir(path)
	if err != nil {
		return "", err
	}

	combinedVals, err := chartutil.CoalesceValues(ch, p.ValueOptions.rawValues)
	if err != nil {
		return "", err
	}
	ch.Values = combinedVals

	// If version is set, modify the version.
	if p.Version != "" {
		if err := setVersion(ch, p.Version); err != nil {
			return "", err
		}
	}

	if p.AppVersion != "" {
		ch.Metadata.AppVersion = p.AppVersion
	}

	if reqs := ch.Metadata.Dependencies; reqs != nil {
		if err := CheckDependencies(ch, reqs); err != nil {
			return "", err
		}
	}

	var dest string
	if p.Destination == "." {
		// Save to the current working directory.
		dest, err = os.Getwd()
		if err != nil {
			return "", err
		}
	} else {
		// Otherwise save to set destination
		dest = p.Destination
	}

	name, err := chartutil.Save(ch, dest)
	if err != nil {
		return "", errors.Wrap(err, "failed to save")
	}

	if p.Sign {
		err = p.Clearsign(name)
	}

	return "", err
}

func setVersion(ch *chart.Chart, ver string) error {
	// Verify that version is a Version, and error out if it is not.
	if _, err := semver.NewVersion(ver); err != nil {
		return err
	}

	// Set the version field on the chart.
	ch.Metadata.Version = ver
	return nil
}

func (p *Package) Clearsign(filename string) error {
	// Load keyring
	signer, err := provenance.NewFromKeyring(p.Keyring, p.Key)
	if err != nil {
		return err
	}

	if err := signer.DecryptKey(promptUser); err != nil {
		return err
	}

	sig, err := signer.ClearSign(filename)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename+".prov", []byte(sig), 0755)
}

// promptUser implements provenance.PassphraseFetcher
func promptUser(name string) ([]byte, error) {
	fmt.Printf("Password for key %q >  ", name)
	pw, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	return pw, err
}
