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

package registry // import "k8s.io/helm/pkg/registry"

import (
	"strings"

	"github.com/containerd/containerd/reference"
)

type (
	// Reference defines the main components of a reference specification
	Reference struct {
		*reference.Spec
	}
)

// ParseReference converts a string to a Reference
func ParseReference(s string) (*Reference, error) {
	spec, err := reference.Parse(s)
	if err != nil {
		return nil, err
	}
	ref := Reference{&spec}
	return &ref, nil
}

// Repo returns a reference's repo minus the hostname
func (ref *Reference) Repo() string {
	return strings.TrimPrefix(strings.TrimPrefix(ref.Locator, ref.Hostname()), "/")
}
