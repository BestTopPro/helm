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

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	"k8s.io/helm/pkg/action"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/downloader"
	"k8s.io/helm/pkg/getter"
)

const packageDesc = `
This command packages a chart into a versioned chart archive file. If a path
is given, this will look at that path for a chart (which must contain a
Chart.yaml file) and then package that directory.

If no path is given, this will look in the present working directory for a
Chart.yaml file, and (if found) build the current directory into a chart.

Versioned chart archives are used by Helm package repositories.
`

func newPackageCmd(out io.Writer) *cobra.Command {
	client := action.NewPackage()

	cmd := &cobra.Command{
		Use:   "package [CHART_PATH] [...]",
		Short: "package a chart directory into a chart archive",
		Long:  packageDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.Errorf("need at least one argument, the path to the chart")
			}
			if client.Sign {
				if client.Key == "" {
					return errors.New("--key is required for signing a package")
				}
				if client.Keyring == "" {
					return errors.New("--keyring is required for signing a package")
				}
			}
			if err := client.ValueOptions.MergeValues(settings); err != nil {
				return err
			}

			for i := 0; i < len(args); i++ {
				path, err := filepath.Abs(args[i])
				if err != nil {
					return err
				}

				if client.DependencyUpdate {
					downloadManager := &downloader.Manager{
						Out:       ioutil.Discard,
						ChartPath: path,
						HelmHome:  settings.Home,
						Keyring:   client.Keyring,
						Getters:   getter.All(settings),
						Debug:     settings.Debug,
					}

					if err := downloadManager.Update(); err != nil {
						return err
					}
				}
				p, err := client.Run(path)
				if err != nil {
					return err
				}
				fmt.Fprintf(out, "Successfully packaged chart and saved it to: %s\n", p)
			}
			return nil
		},
	}

	client.AddFlags(cmd.Flags())

	return cmd
}
