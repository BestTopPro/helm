package helm

import (
	"testing"

	chartutil "github.com/kubernetes/helm/pkg/chart"
	"gopkg.in/yaml.v2"
)

func TestInstallReleaseOverrides(t *testing.T) {
	vals := `name = "mariner"`
	ch := "./testdata/albatross"
	ir, err := InstallRelease([]byte(vals), ch, true)
	if err != nil {
		t.Fatalf("Failed to release: %s", err)
	}

	if len(ir.Release.Manifest) == 0 {
		t.Fatalf("Expected a manifest.")
	}

	// Parse the result and see if the override worked
	d := map[string]interface{}{}
	if err := yaml.Unmarshal([]byte(ir.Release.Manifest), d); err != nil {
		t.Fatalf("Failed to unmarshal manifest: %s", err)
	}

	if d["name"] != "mariner" {
		t.Errorf("Unexpected name %q", d["name"])
	}

	if d["home"] != "nest" {
		t.Errorf("Unexpected home %q", d["home"])
	}
}

func TestOverridesToProto(t *testing.T) {
	override := []byte(`test = "foo"`)
	c := OverridesToProto(override)
	if c.Raw != string(override) {
		t.Errorf("Expected %q to match %q", c.Raw, override)
	}
}

func TestChartToProto(t *testing.T) {
	c, err := chartutil.LoadDir("./testdata/albatross")
	if err != nil {
		t.Fatalf("failed to load testdata chart: %s", err)
	}

	p, err := ChartToProto(c)
	if err != nil {
		t.Fatalf("failed to conver chart to proto: %s", err)
	}

	if p.Metadata.Name != c.Chartfile().Name {
		t.Errorf("Expected names to match.")
	}
}
