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
	"io"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/spf13/pflag"

	"k8s.io/helm/pkg/chart"
	"k8s.io/helm/pkg/chart/loader"
)

type ShowOutputFormat string

const (
	ShowAll    ShowOutputFormat = "all"
	ShowChart  ShowOutputFormat = "chart"
	ShowValues ShowOutputFormat = "values"
	ShowReadme ShowOutputFormat = "readme"
)

var readmeFileNames = []string{"readme.md", "readme.txt", "readme"}

func (o ShowOutputFormat) String() string {
	return string(o)
}

func ParseShowOutputFormat(s string) (out ShowOutputFormat, err error) {
	switch s {
	case ShowAll.String():
		out, err = ShowAll, nil
	case ShowChart.String():
		out, err = ShowChart, nil
	case ShowValues.String():
		out, err = ShowValues, nil
	case ShowReadme.String():
		out, err = ShowReadme, nil
	default:
		out, err = "", ErrInvalidFormatType
	}
	return
}

// Show is the action for checking a given release's information.
//
// It provides the implementation of 'helm show' and its respective subcommands.
type Show struct {
	Out          io.Writer
	OutputFormat ShowOutputFormat
	ChartPathOptions
}

// NewShow creates a new Show object with the given configuration.
func NewShow(out io.Writer, output ShowOutputFormat) *Show {
	return &Show{
		Out:          out,
		OutputFormat: output,
	}
}

func (s *Show) AddFlags(f *pflag.FlagSet) {
	s.ChartPathOptions.AddFlags(f)
}

// Run executes 'helm show' against the given release.
func (s *Show) Run(chartpath string) error {
	chrt, err := loader.Load(chartpath)
	if err != nil {
		return err
	}
	cf, err := yaml.Marshal(chrt.Metadata)
	if err != nil {
		return err
	}

	if s.OutputFormat == ShowChart || s.OutputFormat == ShowAll {
		fmt.Fprintln(s.Out, string(cf))
	}

	if (s.OutputFormat == ShowValues || s.OutputFormat == ShowAll) && chrt.Values != nil {
		if s.OutputFormat == ShowAll {
			fmt.Fprintln(s.Out, "---")
		}
		b, err := yaml.Marshal(chrt.Values)
		if err != nil {
			return err
		}
		fmt.Fprintln(s.Out, string(b))
	}

	if s.OutputFormat == ShowReadme || s.OutputFormat == ShowAll {
		if s.OutputFormat == ShowAll {
			fmt.Fprintln(s.Out, "---")
		}
		readme := findReadme(chrt.Files)
		if readme == nil {
			return nil
		}
		fmt.Fprintln(s.Out, string(readme.Data))
	}
	return nil
}

func findReadme(files []*chart.File) (file *chart.File) {
	for _, file := range files {
		for _, n := range readmeFileNames {
			if strings.EqualFold(file.Name, n) {
				return file
			}
		}
	}
	return nil
}
