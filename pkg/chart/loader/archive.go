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

package loader

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"

	"helm.sh/helm/pkg/chart"
)

// FileLoader loads a chart from a file
type FileLoader string

// Load loads a chart
func (l FileLoader) Load() (*chart.Chart, error) {
	return LoadFile(string(l))
}

// LoadFile loads from an archive file.
func LoadFile(name string) (*chart.Chart, error) {
	if fi, err := os.Stat(name); err != nil {
		return nil, err
	} else if fi.IsDir() {
		return nil, errors.New("cannot load a directory")
	}

	raw, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer raw.Close()

	isLibChart, err := IsArchiveLibraryChart(raw)
	if err != nil {
		return nil, err
	}
	_, _ = raw.Seek(0, 0)
	return LoadArchive(raw, isLibChart)
}

// IsArchiveLibraryChart return true if it is a library chart
func IsArchiveLibraryChart(in io.Reader) (bool, error) {
	var isLibChart = false
	_, err := traverse(in, false, &isLibChart)
	if err != nil {
		return false, err
	}
	return isLibChart, nil
}

// LoadArchive loads from a reader containing a compressed tar archive.
func LoadArchive(in io.Reader, isLibChart bool) (*chart.Chart, error) {
	return traverse(in, true, &isLibChart)
}

func traverse(in io.Reader, loadFiles bool, isLibChart *bool) (*chart.Chart, error) {
	unzipped, err := gzip.NewReader(in)
	if err != nil {
		return &chart.Chart{}, err
	}
	defer unzipped.Close()

	files := []*BufferedFile{}
	tr := tar.NewReader(unzipped)
	for {
		b := bytes.NewBuffer(nil)
		hd, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return &chart.Chart{}, err
		}

		if hd.FileInfo().IsDir() {
			// Use this instead of hd.Typeflag because we don't have to do any
			// inference chasing.
			continue
		}

		// Archive could contain \ if generated on Windows
		delimiter := "/"
		if strings.ContainsRune(hd.Name, '\\') {
			delimiter = "\\"
		}

		parts := strings.Split(hd.Name, delimiter)
		n := strings.Join(parts[1:], delimiter)

		// Normalize the path to the / delimiter
		n = strings.ReplaceAll(n, delimiter, "/")

		if parts[0] == "Chart.yaml" {
			return nil, errors.New("chart yaml not in base directory")
		}

		if _, err := io.Copy(b, tr); err != nil {
			return &chart.Chart{}, err
		}

		if !loadFiles && n == "Chart.yaml" {
			var err error
			*isLibChart, err = IsLibraryChart(b.Bytes())
			if err != nil {
				return nil, errors.Wrapf(err, "cannot load Chart.yaml")
			}
			return nil, nil
		}

		if IsFileValid(n, *isLibChart) {
			files = append(files, &BufferedFile{Name: n, Data: b.Bytes()})
		}

		b.Reset()
	}

	if loadFiles && len(files) == 0 {
		return nil, errors.New("no files in chart archive")
	}

	if !loadFiles {
		return nil, errors.New("cannot find the chart type")
	}

	return LoadFiles(files)
}
