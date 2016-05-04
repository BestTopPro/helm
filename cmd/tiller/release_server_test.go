package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kubernetes/helm/cmd/tiller/environment"
	"github.com/kubernetes/helm/pkg/proto/hapi/chart"
	"github.com/kubernetes/helm/pkg/proto/hapi/release"
	"github.com/kubernetes/helm/pkg/proto/hapi/services"
	"github.com/kubernetes/helm/pkg/storage"
	"github.com/kubernetes/helm/pkg/timeconv"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func rsFixture() *releaseServer {
	return &releaseServer{
		env: mockEnvironment(),
	}
}

func releaseMock() *release.Release {
	date := timestamp.Timestamp{Seconds: 242085845, Nanos: 0}
	return &release.Release{
		Name: "angry-panda",
		Info: &release.Info{
			FirstDeployed: &date,
			LastDeployed:  &date,
			Status:        &release.Status{Code: release.Status_DEPLOYED},
		},
		Chart: &chart.Chart{
			Metadata: &chart.Metadata{
				Name:    "foo",
				Version: "0.1.0-beta.1",
			},
			Templates: []*chart.Template{
				{Name: "foo.tpl", Data: []byte("Hello")},
			},
		},
		Config: &chart.Config{Raw: `name = "value"`},
	}
}

func TestInstallRelease(t *testing.T) {
	c := context.Background()
	rs := rsFixture()

	req := &services.InstallReleaseRequest{
		Chart: &chart.Chart{
			Metadata: &chart.Metadata{Name: "hello"},
			Templates: []*chart.Template{
				{Name: "hello", Data: []byte("hello: world")},
			},
		},
	}
	res, err := rs.InstallRelease(c, req)
	if err != nil {
		t.Errorf("Failed install: %s", err)
	}
	if res.Release.Name == "" {
		t.Errorf("Expected release name.")
	}

	rel, err := rs.env.Releases.Read(res.Release.Name)
	if err != nil {
		t.Errorf("Expected release for %s (%v).", res.Release.Name, rs.env.Releases)
	}

	t.Logf("rel: %v", rel)

	if len(res.Release.Manifest) == 0 {
		t.Errorf("No manifest returned: %v", res.Release)
	}

	if len(rel.Manifest) == 0 {
		t.Errorf("Expected manifest in %v", res)
	}

	if !strings.Contains(rel.Manifest, "---\n# Source: hello\nhello: world") {
		t.Errorf("unexpected output: %s", rel.Manifest)
	}
}

func TestInstallReleaseDryRun(t *testing.T) {
	c := context.Background()
	rs := rsFixture()

	req := &services.InstallReleaseRequest{
		Chart: &chart.Chart{
			Metadata: &chart.Metadata{Name: "hello"},
			Templates: []*chart.Template{
				{Name: "hello", Data: []byte("hello: world")},
				{Name: "goodbye", Data: []byte("goodbye: world")},
				{Name: "empty", Data: []byte("")},
			},
		},
		DryRun: true,
	}
	res, err := rs.InstallRelease(c, req)
	if err != nil {
		t.Errorf("Failed install: %s", err)
	}
	if res.Release.Name == "" {
		t.Errorf("Expected release name.")
	}

	if !strings.Contains(res.Release.Manifest, "---\n# Source: hello\nhello: world") {
		t.Errorf("unexpected output: %s", res.Release.Manifest)
	}

	if !strings.Contains(res.Release.Manifest, "---\n# Source: goodbye\ngoodbye: world") {
		t.Errorf("unexpected output: %s", res.Release.Manifest)
	}

	if strings.Contains(res.Release.Manifest, "empty") {
		t.Errorf("Should not contain template data for an empty file. %s", res.Release.Manifest)
	}

	if _, err := rs.env.Releases.Read(res.Release.Name); err == nil {
		t.Errorf("Expected no stored release.")
	}
}

func TestUninstallRelease(t *testing.T) {
	c := context.Background()
	rs := rsFixture()
	rs.env.Releases.Create(&release.Release{
		Name: "angry-panda",
		Info: &release.Info{
			FirstDeployed: timeconv.Now(),
			Status: &release.Status{
				Code: release.Status_DEPLOYED,
			},
		},
	})

	req := &services.UninstallReleaseRequest{
		Name: "angry-panda",
	}

	res, err := rs.UninstallRelease(c, req)
	if err != nil {
		t.Errorf("Failed uninstall: %s", err)
	}

	if res.Release.Name != "angry-panda" {
		t.Errorf("Expected angry-panda, got %q", res.Release.Name)
	}

	if res.Release.Info.Status.Code != release.Status_DELETED {
		t.Errorf("Expected status code to be DELETED, got %d", res.Release.Info.Status.Code)
	}

	if res.Release.Info.Deleted.Seconds <= 0 {
		t.Errorf("Expected valid UNIX date, got %d", res.Release.Info.Deleted.Seconds)
	}
}

func TestGetReleaseContent(t *testing.T) {
	c := context.Background()
	rs := rsFixture()
	rel := releaseMock()
	if err := rs.env.Releases.Create(rel); err != nil {
		t.Fatalf("Could not store mock release: %s", err)
	}

	res, err := rs.GetReleaseContent(c, &services.GetReleaseContentRequest{Name: rel.Name})
	if err != nil {
		t.Errorf("Error getting release content: %s", err)
	}

	if res.Release.Chart.Metadata.Name != rel.Chart.Metadata.Name {
		t.Errorf("Expected %q, got %q", rel.Chart.Metadata.Name, res.Release.Chart.Metadata.Name)
	}
}

func TestGetReleaseStatus(t *testing.T) {
	c := context.Background()
	rs := rsFixture()
	rel := releaseMock()
	if err := rs.env.Releases.Create(rel); err != nil {
		t.Fatalf("Could not store mock release: %s", err)
	}

	res, err := rs.GetReleaseStatus(c, &services.GetReleaseStatusRequest{Name: rel.Name})
	if err != nil {
		t.Errorf("Error getting release content: %s", err)
	}

	if res.Info.Status.Code != release.Status_DEPLOYED {
		t.Errorf("Expected %d, got %d", release.Status_DEPLOYED, res.Info.Status.Code)
	}
}

func TestListReleases(t *testing.T) {
	rs := rsFixture()
	num := 7
	for i := 0; i < num; i++ {
		rel := releaseMock()
		rel.Name = fmt.Sprintf("rel-%d", i)
		if err := rs.env.Releases.Create(rel); err != nil {
			t.Fatalf("Could not store mock release: %s", err)
		}
	}

	mrs := &mockListServer{}
	if err := rs.ListReleases(&services.ListReleasesRequest{Offset: 0, Limit: 64}, mrs); err != nil {
		t.Fatalf("Failed listing: %s", err)
	}

	if len(mrs.val.Releases) != num {
		t.Errorf("Expected %d releases, got %d", num, len(mrs.val.Releases))
	}
}

func mockEnvironment() *environment.Environment {
	e := environment.New()
	e.Releases = storage.NewMemory()
	e.KubeClient = &environment.PrintingKubeClient{Out: os.Stdout}
	return e
}

type mockListServer struct {
	val *services.ListReleasesResponse
}

func (l *mockListServer) Send(res *services.ListReleasesResponse) error {
	l.val = res
	return nil
}

func (l *mockListServer) Context() context.Context       { return context.TODO() }
func (l *mockListServer) SendMsg(v interface{}) error    { return nil }
func (l *mockListServer) RecvMsg(v interface{}) error    { return nil }
func (l *mockListServer) SendHeader(m metadata.MD) error { return nil }
func (l *mockListServer) SetTrailer(m metadata.MD)       {}
