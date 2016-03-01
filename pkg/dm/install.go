package dm

import (
	"github.com/kubernetes/deployment-manager/pkg/format"
	"github.com/kubernetes/deployment-manager/pkg/kubectl"
)

// Install uses kubectl to install the base DM.
//
// Returns the string output received from the operation, and an error if the
// command failed.
func Install(runner kubectl.Runner) (string, error) {
	o, err := runner.Create([]byte(InstallYAML))
	return string(o), err
}

// IsInstalled checks whether DM has been installed.
func IsInstalled(runner kubectl.Runner) bool {
	// Basically, we test "all-or-nothing" here: if this returns without error
	// we know that we have both the namespace and the manager API server.
	out, err := runner.GetByKind("rc", "manager-rc", "dm")
	if err != nil {
		format.Err("Installation not found: %s %s", out, err)
		return false
	}
	return true
}

// InstallYAML is the installation YAML for DM.
const InstallYAML = `
######################################################################
# Copyright 2015 The Kubernetes Authors All rights reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
######################################################################

---
apiVersion: v1
kind: Namespace
metadata:
  labels:
    app: dm
    name: dm-namespace
  name: dm
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: dm
    name: expandybird-service
  name: expandybird-service
  namespace: dm
spec:
  ports:
  - name: expandybird
    port: 8081
    targetPort: 8080
  selector:
    app: dm
    name: expandybird
---
apiVersion: v1
kind: ReplicationController
metadata:
  labels:
    app: dm
    name: expandybird-rc
  name: expandybird-rc
  namespace: dm
spec:
  replicas: 2
  selector:
    app: dm
    name: expandybird
  template:
    metadata:
      labels:
        app: dm
        name: expandybird
    spec:
      containers:
      - env: []
        image: gcr.io/dm-k8s-testing/expandybird:latest
        name: expandybird
        ports:
        - containerPort: 8080
          name: expandybird
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: dm
    name: resourcifier-service
  name: resourcifier-service
  namespace: dm
spec:
  ports:
  - name: resourcifier
    port: 8082
    targetPort: 8080
  selector:
    app: dm
    name: resourcifier
---
apiVersion: v1
kind: ReplicationController
metadata:
  labels:
    app: dm
    name: resourcifier-rc
  name: resourcifier-rc
  namespace: dm
spec:
  replicas: 2
  selector:
    app: dm
    name: resourcifier
  template:
    metadata:
      labels:
        app: dm
        name: resourcifier
    spec:
      containers:
      - env: []
        image: gcr.io/dm-k8s-testing/resourcifier:latest
        name: resourcifier
        ports:
        - containerPort: 8080
          name: resourcifier
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: dm
    name: manager-service
  name: manager-service
  namespace: dm
spec:
  ports:
  - name: manager
    port: 8080
    targetPort: 8080
  selector:
    app: dm
    name: manager
---
apiVersion: v1
kind: ReplicationController
metadata:
  labels:
    app: dm
    name: manager-rc
  name: manager-rc
  namespace: dm
spec:
  replicas: 1
  selector:
    app: dm
    name: manager
  template:
    metadata:
      labels:
        app: dm
        name: manager
    spec:
      containers:
      - env: []
        image: gcr.io/dm-k8s-testing/manager:latest
        name: manager
        ports:
        - containerPort: 8080
          name: manager
`
