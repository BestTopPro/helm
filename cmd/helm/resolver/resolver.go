/*
Copyright 2016 The Kubernetes Authors All rights reserved.
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

package resolver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/semver"

	"k8s.io/helm/cmd/helm/helmpath"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/provenance"
	"k8s.io/helm/pkg/repo"
)

// Resolver resolves dependencies from semantic version ranges to a particular version.
type Resolver struct {
	chartpath string
	helmhome  helmpath.Home
}

// New creates a new resolver for a given chart and a given helm home.
func New(chartpath string, helmhome helmpath.Home) *Resolver {
	return &Resolver{
		chartpath: chartpath,
		helmhome:  helmhome,
	}
}

// Resolve resolves dependencies and returns a lock file with the resolution.
func (r *Resolver) Resolve(reqs *chartutil.Requirements, repoNames map[string]string) (*chartutil.RequirementsLock, error) {
	d, err := HashReq(reqs)
	if err != nil {
		return nil, err
	}

	// Now we clone the dependencies, locking as we go.
	locked := make([]*chartutil.Dependency, len(reqs.Dependencies))
	missing := []string{}
	for i, d := range reqs.Dependencies {
		constraint, err := semver.NewConstraint(d.Version)
		if err != nil {
			return nil, fmt.Errorf("dependency %q has an invalid version/constraint format: %s", d.Name, err)
		}

		repoIndex, err := repo.NewChartRepositoryIndexFromFile(r.helmhome.CacheIndex(repoNames[d.Name]))
		if err != nil {
			return nil, fmt.Errorf("no cached repo found. (try 'helm repo update'). %s", err)
		}

		vs, ok := repoIndex.Entries[d.Name]
		if !ok {
			return nil, fmt.Errorf("%s chart not found in repo %s", d.Name, d.Repository)
		}

		locked[i] = &chartutil.Dependency{
			Name:       d.Name,
			Repository: d.Repository,
		}
		found := false
		// The version are already sorted and hence the first one to satisfy the constraint is used
		for _, ver := range vs {
			v, err := semver.NewVersion(ver.Version)
			if err != nil || len(ver.URLs) == 0 {
				// Not a legit entry.
				continue
			}
			if constraint.Check(v) {
				found = true
				locked[i].Version = v.Original()
				break
			}
		}

		if !found {
			missing = append(missing, d.Name)
		}
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("Can't get a valid version for repositories %s. Try changing the version constraint in requirements.yaml", strings.Join(missing, ", "))
	}
	return &chartutil.RequirementsLock{
		Generated:    time.Now(),
		Digest:       d,
		Dependencies: locked,
	}, nil
}

// HashReq generates a hash of the requirements.
//
// This should be used only to compare against another hash generated by this
// function.
func HashReq(req *chartutil.Requirements) (string, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	s, err := provenance.Digest(bytes.NewBuffer(data))
	return "sha256:" + s, err
}
