// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/services/tiller.proto

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	hapi/services/tiller.proto

It has these top-level messages:
	ListReleasesRequest
	ListSort
	ListReleasesResponse
	GetReleaseStatusRequest
	GetReleaseStatusResponse
	GetReleaseContentRequest
	UpdateReleaseRequest
	RollbackReleaseRequest
	InstallReleaseRequest
	UninstallReleaseRequest
	UninstallReleaseResponse
	GetHistoryRequest
	TestReleaseRequest
	TestReleaseResponse
*/
package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hapi_chart3 "k8s.io/helm/pkg/proto/hapi/chart"
import hapi_chart "k8s.io/helm/pkg/proto/hapi/chart"
import hapi_release5 "k8s.io/helm/pkg/proto/hapi/release"
import hapi_release4 "k8s.io/helm/pkg/proto/hapi/release"
import hapi_release1 "k8s.io/helm/pkg/proto/hapi/release"
import hapi_release3 "k8s.io/helm/pkg/proto/hapi/release"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SortBy defines sort operations.
type ListSort_SortBy int32

const (
	ListSort_UNKNOWN       ListSort_SortBy = 0
	ListSort_NAME          ListSort_SortBy = 1
	ListSort_LAST_RELEASED ListSort_SortBy = 2
)

var ListSort_SortBy_name = map[int32]string{
	0: "UNKNOWN",
	1: "NAME",
	2: "LAST_RELEASED",
}
var ListSort_SortBy_value = map[string]int32{
	"UNKNOWN":       0,
	"NAME":          1,
	"LAST_RELEASED": 2,
}

func (x ListSort_SortBy) String() string {
	return proto.EnumName(ListSort_SortBy_name, int32(x))
}
func (ListSort_SortBy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// SortOrder defines sort orders to augment sorting operations.
type ListSort_SortOrder int32

const (
	ListSort_ASC  ListSort_SortOrder = 0
	ListSort_DESC ListSort_SortOrder = 1
)

var ListSort_SortOrder_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}
var ListSort_SortOrder_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x ListSort_SortOrder) String() string {
	return proto.EnumName(ListSort_SortOrder_name, int32(x))
}
func (ListSort_SortOrder) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

// ListReleasesRequest requests a list of releases.
//
// Releases can be retrieved in chunks by setting limit and offset.
//
// Releases can be sorted according to a few pre-determined sort stategies.
type ListReleasesRequest struct {
	// Limit is the maximum number of releases to be returned.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset is the last release name that was seen. The next listing
	// operation will start with the name after this one.
	// Example: If list one returns albert, bernie, carl, and sets 'next: dennis'.
	// dennis is the offset. Supplying 'dennis' for the next request should
	// cause the next batch to return a set of results starting with 'dennis'.
	Offset string `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
	// SortBy is the sort field that the ListReleases server should sort data before returning.
	SortBy ListSort_SortBy `protobuf:"varint,3,opt,name=sort_by,json=sortBy,enum=hapi.services.tiller.ListSort_SortBy" json:"sort_by,omitempty"`
	// Filter is a regular expression used to filter which releases should be listed.
	//
	// Anything that matches the regexp will be included in the results.
	Filter string `protobuf:"bytes,4,opt,name=filter" json:"filter,omitempty"`
	// SortOrder is the ordering directive used for sorting.
	SortOrder   ListSort_SortOrder          `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,enum=hapi.services.tiller.ListSort_SortOrder" json:"sort_order,omitempty"`
	StatusCodes []hapi_release3.Status_Code `protobuf:"varint,6,rep,packed,name=status_codes,json=statusCodes,enum=hapi.release.Status_Code" json:"status_codes,omitempty"`
	// Namespace is the filter to select releases only from a specific namespace.
	Namespace string `protobuf:"bytes,7,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ListReleasesRequest) Reset()                    { *m = ListReleasesRequest{} }
func (m *ListReleasesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListReleasesRequest) ProtoMessage()               {}
func (*ListReleasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListReleasesRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListReleasesRequest) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *ListReleasesRequest) GetSortBy() ListSort_SortBy {
	if m != nil {
		return m.SortBy
	}
	return ListSort_UNKNOWN
}

func (m *ListReleasesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListReleasesRequest) GetSortOrder() ListSort_SortOrder {
	if m != nil {
		return m.SortOrder
	}
	return ListSort_ASC
}

func (m *ListReleasesRequest) GetStatusCodes() []hapi_release3.Status_Code {
	if m != nil {
		return m.StatusCodes
	}
	return nil
}

func (m *ListReleasesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// ListSort defines sorting fields on a release list.
type ListSort struct {
}

func (m *ListSort) Reset()                    { *m = ListSort{} }
func (m *ListSort) String() string            { return proto.CompactTextString(m) }
func (*ListSort) ProtoMessage()               {}
func (*ListSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// ListReleasesResponse is a list of releases.
type ListReleasesResponse struct {
	// Count is the expected total number of releases to be returned.
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// Next is the name of the next release. If this is other than an empty
	// string, it means there are more results.
	Next string `protobuf:"bytes,2,opt,name=next" json:"next,omitempty"`
	// Total is the total number of queryable releases.
	Total int64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	// Releases is the list of found release objects.
	Releases []*hapi_release5.Release `protobuf:"bytes,4,rep,name=releases" json:"releases,omitempty"`
}

func (m *ListReleasesResponse) Reset()                    { *m = ListReleasesResponse{} }
func (m *ListReleasesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListReleasesResponse) ProtoMessage()               {}
func (*ListReleasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListReleasesResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListReleasesResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *ListReleasesResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListReleasesResponse) GetReleases() []*hapi_release5.Release {
	if m != nil {
		return m.Releases
	}
	return nil
}

// GetReleaseStatusRequest is a request to get the status of a release.
type GetReleaseStatusRequest struct {
	// Name is the name of the release
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Version is the version of the release
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *GetReleaseStatusRequest) Reset()                    { *m = GetReleaseStatusRequest{} }
func (m *GetReleaseStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReleaseStatusRequest) ProtoMessage()               {}
func (*GetReleaseStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetReleaseStatusRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetReleaseStatusRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// GetReleaseStatusResponse is the response indicating the status of the named release.
type GetReleaseStatusResponse struct {
	// Name is the name of the release.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Info contains information about the release.
	Info *hapi_release4.Info `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	// Namespace the release was released into
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *GetReleaseStatusResponse) Reset()                    { *m = GetReleaseStatusResponse{} }
func (m *GetReleaseStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetReleaseStatusResponse) ProtoMessage()               {}
func (*GetReleaseStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetReleaseStatusResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetReleaseStatusResponse) GetInfo() *hapi_release4.Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *GetReleaseStatusResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// GetReleaseContentRequest is a request to get the contents of a release.
type GetReleaseContentRequest struct {
	// The name of the release
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Version is the version of the release
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *GetReleaseContentRequest) Reset()                    { *m = GetReleaseContentRequest{} }
func (m *GetReleaseContentRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReleaseContentRequest) ProtoMessage()               {}
func (*GetReleaseContentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetReleaseContentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetReleaseContentRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// UpdateReleaseRequest updates a release.
type UpdateReleaseRequest struct {
	// The name of the release
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Chart is the protobuf representation of a chart.
	Chart *hapi_chart3.Chart `protobuf:"bytes,2,opt,name=chart" json:"chart,omitempty"`
	// Values is a string containing (unparsed) YAML values.
	Values *hapi_chart.Config `protobuf:"bytes,3,opt,name=values" json:"values,omitempty"`
	// dry_run, if true, will run through the release logic, but neither create
	DryRun bool `protobuf:"varint,4,opt,name=dry_run,json=dryRun" json:"dry_run,omitempty"`
	// DisableHooks causes the server to skip running any hooks for the upgrade.
	DisableHooks bool `protobuf:"varint,5,opt,name=disable_hooks,json=disableHooks" json:"disable_hooks,omitempty"`
	// Performs pods restart for resources if applicable
	Recreate bool `protobuf:"varint,6,opt,name=recreate" json:"recreate,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `protobuf:"varint,7,opt,name=timeout" json:"timeout,omitempty"`
	// ResetValues will cause Tiller to ignore stored values, resetting to default values.
	ResetValues bool `protobuf:"varint,8,opt,name=reset_values,json=resetValues" json:"reset_values,omitempty"`
	// wait, if true, will wait until all Pods, PVCs, and Services are in a ready state
	// before marking the release as successful. It will wait for as long as timeout
	Wait bool `protobuf:"varint,9,opt,name=wait" json:"wait,omitempty"`
	// ReuseValues will cause Tiller to reuse the values from the last release.
	// This is ignored if reset_values is set.
	ReuseValues bool `protobuf:"varint,10,opt,name=reuse_values,json=reuseValues" json:"reuse_values,omitempty"`
	// Force resource update through delete/recreate if needed.
	Force bool `protobuf:"varint,11,opt,name=force" json:"force,omitempty"`
}

func (m *UpdateReleaseRequest) Reset()                    { *m = UpdateReleaseRequest{} }
func (m *UpdateReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateReleaseRequest) ProtoMessage()               {}
func (*UpdateReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateReleaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateReleaseRequest) GetChart() *hapi_chart3.Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

func (m *UpdateReleaseRequest) GetValues() *hapi_chart.Config {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *UpdateReleaseRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *UpdateReleaseRequest) GetDisableHooks() bool {
	if m != nil {
		return m.DisableHooks
	}
	return false
}

func (m *UpdateReleaseRequest) GetRecreate() bool {
	if m != nil {
		return m.Recreate
	}
	return false
}

func (m *UpdateReleaseRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *UpdateReleaseRequest) GetResetValues() bool {
	if m != nil {
		return m.ResetValues
	}
	return false
}

func (m *UpdateReleaseRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

func (m *UpdateReleaseRequest) GetReuseValues() bool {
	if m != nil {
		return m.ReuseValues
	}
	return false
}

func (m *UpdateReleaseRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

type RollbackReleaseRequest struct {
	// The name of the release
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// dry_run, if true, will run through the release logic but no create
	DryRun bool `protobuf:"varint,2,opt,name=dry_run,json=dryRun" json:"dry_run,omitempty"`
	// DisableHooks causes the server to skip running any hooks for the rollback
	DisableHooks bool `protobuf:"varint,3,opt,name=disable_hooks,json=disableHooks" json:"disable_hooks,omitempty"`
	// Version is the version of the release to deploy.
	Version int32 `protobuf:"varint,4,opt,name=version" json:"version,omitempty"`
	// Performs pods restart for resources if applicable
	Recreate bool `protobuf:"varint,5,opt,name=recreate" json:"recreate,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `protobuf:"varint,6,opt,name=timeout" json:"timeout,omitempty"`
	// wait, if true, will wait until all Pods, PVCs, and Services are in a ready state
	// before marking the release as successful. It will wait for as long as timeout
	Wait bool `protobuf:"varint,7,opt,name=wait" json:"wait,omitempty"`
	// Force resource update through delete/recreate if needed.
	Force bool `protobuf:"varint,8,opt,name=force" json:"force,omitempty"`
}

func (m *RollbackReleaseRequest) Reset()                    { *m = RollbackReleaseRequest{} }
func (m *RollbackReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*RollbackReleaseRequest) ProtoMessage()               {}
func (*RollbackReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RollbackReleaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RollbackReleaseRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *RollbackReleaseRequest) GetDisableHooks() bool {
	if m != nil {
		return m.DisableHooks
	}
	return false
}

func (m *RollbackReleaseRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RollbackReleaseRequest) GetRecreate() bool {
	if m != nil {
		return m.Recreate
	}
	return false
}

func (m *RollbackReleaseRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RollbackReleaseRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

func (m *RollbackReleaseRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

// InstallReleaseRequest is the request for an installation of a chart.
type InstallReleaseRequest struct {
	// Chart is the protobuf representation of a chart.
	Chart *hapi_chart3.Chart `protobuf:"bytes,1,opt,name=chart" json:"chart,omitempty"`
	// Values is a string containing (unparsed) YAML values.
	Values *hapi_chart.Config `protobuf:"bytes,2,opt,name=values" json:"values,omitempty"`
	// DryRun, if true, will run through the release logic, but neither create
	// a release object nor deploy to Kubernetes. The release object returned
	// in the response will be fake.
	DryRun bool `protobuf:"varint,3,opt,name=dry_run,json=dryRun" json:"dry_run,omitempty"`
	// Name is the candidate release name. This must be unique to the
	// namespace, otherwise the server will return an error. If it is not
	// supplied, the server will autogenerate one.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// DisableHooks causes the server to skip running any hooks for the install.
	DisableHooks bool `protobuf:"varint,5,opt,name=disable_hooks,json=disableHooks" json:"disable_hooks,omitempty"`
	// Namepace is the kubernetes namespace of the release.
	Namespace string `protobuf:"bytes,6,opt,name=namespace" json:"namespace,omitempty"`
	// ReuseName requests that Tiller re-uses a name, instead of erroring out.
	ReuseName bool `protobuf:"varint,7,opt,name=reuse_name,json=reuseName" json:"reuse_name,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `protobuf:"varint,8,opt,name=timeout" json:"timeout,omitempty"`
	// wait, if true, will wait until all Pods, PVCs, and Services are in a ready state
	// before marking the release as successful. It will wait for as long as timeout
	Wait bool `protobuf:"varint,9,opt,name=wait" json:"wait,omitempty"`
}

func (m *InstallReleaseRequest) Reset()                    { *m = InstallReleaseRequest{} }
func (m *InstallReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*InstallReleaseRequest) ProtoMessage()               {}
func (*InstallReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *InstallReleaseRequest) GetChart() *hapi_chart3.Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

func (m *InstallReleaseRequest) GetValues() *hapi_chart.Config {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *InstallReleaseRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *InstallReleaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstallReleaseRequest) GetDisableHooks() bool {
	if m != nil {
		return m.DisableHooks
	}
	return false
}

func (m *InstallReleaseRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *InstallReleaseRequest) GetReuseName() bool {
	if m != nil {
		return m.ReuseName
	}
	return false
}

func (m *InstallReleaseRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *InstallReleaseRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

// UninstallReleaseRequest represents a request to uninstall a named release.
type UninstallReleaseRequest struct {
	// Name is the name of the release to delete.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// DisableHooks causes the server to skip running any hooks for the uninstall.
	DisableHooks bool `protobuf:"varint,2,opt,name=disable_hooks,json=disableHooks" json:"disable_hooks,omitempty"`
	// Purge removes the release from the store and make its name free for later use.
	Purge bool `protobuf:"varint,3,opt,name=purge" json:"purge,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *UninstallReleaseRequest) Reset()                    { *m = UninstallReleaseRequest{} }
func (m *UninstallReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*UninstallReleaseRequest) ProtoMessage()               {}
func (*UninstallReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UninstallReleaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UninstallReleaseRequest) GetDisableHooks() bool {
	if m != nil {
		return m.DisableHooks
	}
	return false
}

func (m *UninstallReleaseRequest) GetPurge() bool {
	if m != nil {
		return m.Purge
	}
	return false
}

func (m *UninstallReleaseRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// UninstallReleaseResponse represents a successful response to an uninstall request.
type UninstallReleaseResponse struct {
	// Release is the release that was marked deleted.
	Release *hapi_release5.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
	// Info is an uninstall message
	Info string `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *UninstallReleaseResponse) Reset()                    { *m = UninstallReleaseResponse{} }
func (m *UninstallReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*UninstallReleaseResponse) ProtoMessage()               {}
func (*UninstallReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UninstallReleaseResponse) GetRelease() *hapi_release5.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *UninstallReleaseResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

// GetHistoryRequest requests a release's history.
type GetHistoryRequest struct {
	// The name of the release.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The maximum number of releases to include.
	Max int32 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
}

func (m *GetHistoryRequest) Reset()                    { *m = GetHistoryRequest{} }
func (m *GetHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryRequest) ProtoMessage()               {}
func (*GetHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetHistoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetHistoryRequest) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

// TestReleaseRequest is a request to get the status of a release.
type TestReleaseRequest struct {
	// Name is the name of the release
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	// cleanup specifies whether or not to attempt pod deletion after test completes
	Cleanup bool `protobuf:"varint,3,opt,name=cleanup" json:"cleanup,omitempty"`
}

func (m *TestReleaseRequest) Reset()                    { *m = TestReleaseRequest{} }
func (m *TestReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*TestReleaseRequest) ProtoMessage()               {}
func (*TestReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TestReleaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestReleaseRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *TestReleaseRequest) GetCleanup() bool {
	if m != nil {
		return m.Cleanup
	}
	return false
}

// TestReleaseResponse represents a message from executing a test
type TestReleaseResponse struct {
	Msg    string                       `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Status hapi_release1.TestRun_Status `protobuf:"varint,2,opt,name=status,enum=hapi.release.TestRun_Status" json:"status,omitempty"`
}

func (m *TestReleaseResponse) Reset()                    { *m = TestReleaseResponse{} }
func (m *TestReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*TestReleaseResponse) ProtoMessage()               {}
func (*TestReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TestReleaseResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TestReleaseResponse) GetStatus() hapi_release1.TestRun_Status {
	if m != nil {
		return m.Status
	}
	return hapi_release1.TestRun_UNKNOWN
}

func init() {
	proto.RegisterType((*ListReleasesRequest)(nil), "hapi.services.tiller.ListReleasesRequest")
	proto.RegisterType((*ListSort)(nil), "hapi.services.tiller.ListSort")
	proto.RegisterType((*ListReleasesResponse)(nil), "hapi.services.tiller.ListReleasesResponse")
	proto.RegisterType((*GetReleaseStatusRequest)(nil), "hapi.services.tiller.GetReleaseStatusRequest")
	proto.RegisterType((*GetReleaseStatusResponse)(nil), "hapi.services.tiller.GetReleaseStatusResponse")
	proto.RegisterType((*GetReleaseContentRequest)(nil), "hapi.services.tiller.GetReleaseContentRequest")
	proto.RegisterType((*UpdateReleaseRequest)(nil), "hapi.services.tiller.UpdateReleaseRequest")
	proto.RegisterType((*RollbackReleaseRequest)(nil), "hapi.services.tiller.RollbackReleaseRequest")
	proto.RegisterType((*InstallReleaseRequest)(nil), "hapi.services.tiller.InstallReleaseRequest")
	proto.RegisterType((*UninstallReleaseRequest)(nil), "hapi.services.tiller.UninstallReleaseRequest")
	proto.RegisterType((*UninstallReleaseResponse)(nil), "hapi.services.tiller.UninstallReleaseResponse")
	proto.RegisterType((*GetHistoryRequest)(nil), "hapi.services.tiller.GetHistoryRequest")
	proto.RegisterType((*TestReleaseRequest)(nil), "hapi.services.tiller.TestReleaseRequest")
	proto.RegisterType((*TestReleaseResponse)(nil), "hapi.services.tiller.TestReleaseResponse")
	proto.RegisterEnum("hapi.services.tiller.ListSort_SortBy", ListSort_SortBy_name, ListSort_SortBy_value)
	proto.RegisterEnum("hapi.services.tiller.ListSort_SortOrder", ListSort_SortOrder_name, ListSort_SortOrder_value)
}

func init() { proto.RegisterFile("hapi/services/tiller.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 948 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xe7, 0x7c, 0xf6, 0xd9, 0x1e, 0xa7, 0x91, 0xb3, 0x75, 0x93, 0x6b, 0x28, 0xc8, 0x1c, 0x02,
	0x2c, 0x1e, 0x1c, 0x61, 0x78, 0x41, 0x42, 0x48, 0xa9, 0x6b, 0x25, 0x15, 0xc1, 0x95, 0xd6, 0x0d,
	0x48, 0x08, 0xb0, 0x2e, 0xf6, 0x38, 0x39, 0xf5, 0x7c, 0x6b, 0x76, 0xf7, 0x42, 0xfd, 0xca, 0x1b,
	0x1f, 0x85, 0x6f, 0xc1, 0x77, 0x81, 0x0f, 0x82, 0xf6, 0x9f, 0xb9, 0x4b, 0xdc, 0x10, 0xfa, 0x72,
	0xde, 0xd9, 0xf9, 0xb3, 0xbf, 0xf9, 0xcd, 0xec, 0xac, 0xe1, 0xf0, 0x2a, 0x5e, 0x25, 0x47, 0x02,
	0xf9, 0x75, 0x32, 0x43, 0x71, 0x24, 0x93, 0x34, 0x45, 0xde, 0x5f, 0x71, 0x26, 0x19, 0xe9, 0x28,
	0x5d, 0xdf, 0xe9, 0xfa, 0x46, 0x77, 0xb8, 0xaf, 0x3d, 0x66, 0x57, 0x31, 0x97, 0xe6, 0x6b, 0xac,
	0x0f, 0x0f, 0x8a, 0xfb, 0x2c, 0x5b, 0x24, 0x97, 0x56, 0x61, 0x8e, 0xe0, 0x98, 0x62, 0x2c, 0xd0,
	0xfd, 0x96, 0x9c, 0x9c, 0x2e, 0xc9, 0x16, 0xcc, 0x2a, 0xde, 0x2d, 0x29, 0x24, 0x0a, 0x39, 0xe5,
	0x79, 0x66, 0x95, 0x8f, 0x4b, 0x4a, 0x21, 0x63, 0x99, 0x0b, 0xa3, 0x8a, 0xfe, 0xac, 0xc0, 0xc3,
	0xb3, 0x44, 0x48, 0x6a, 0x94, 0x82, 0xe2, 0x2f, 0x39, 0x0a, 0x49, 0x3a, 0x50, 0x4b, 0x93, 0x65,
	0x22, 0x43, 0xaf, 0xeb, 0xf5, 0x7c, 0x6a, 0x04, 0xb2, 0x0f, 0x01, 0x5b, 0x2c, 0x04, 0xca, 0xb0,
	0xd2, 0xf5, 0x7a, 0x4d, 0x6a, 0x25, 0xf2, 0x35, 0xd4, 0x05, 0xe3, 0x72, 0x7a, 0xb1, 0x0e, 0xfd,
	0xae, 0xd7, 0xdb, 0x1d, 0x7c, 0xd4, 0xdf, 0xc6, 0x45, 0x5f, 0x9d, 0x34, 0x61, 0x5c, 0xf6, 0xd5,
	0xe7, 0xe9, 0x9a, 0x06, 0x42, 0xff, 0xaa, 0xb8, 0x8b, 0x24, 0x95, 0xc8, 0xc3, 0xaa, 0x89, 0x6b,
	0x24, 0x72, 0x02, 0xa0, 0xe3, 0x32, 0x3e, 0x47, 0x1e, 0xd6, 0x74, 0xe8, 0xde, 0x3d, 0x42, 0xbf,
	0x50, 0xf6, 0xb4, 0x29, 0xdc, 0x92, 0x7c, 0x05, 0x3b, 0x26, 0xed, 0xe9, 0x8c, 0xcd, 0x51, 0x84,
	0x41, 0xd7, 0xef, 0xed, 0x0e, 0x1e, 0x9b, 0x50, 0x8e, 0xe2, 0x89, 0x21, 0x66, 0xc8, 0xe6, 0x48,
	0x5b, 0xc6, 0x5c, 0xad, 0x05, 0x79, 0x02, 0xcd, 0x2c, 0x5e, 0xa2, 0x58, 0xc5, 0x33, 0x0c, 0xeb,
	0x1a, 0xe1, 0xbf, 0x1b, 0xd1, 0xcf, 0xd0, 0x70, 0x87, 0x47, 0x03, 0x08, 0x4c, 0x6a, 0xa4, 0x05,
	0xf5, 0xf3, 0xf1, 0x37, 0xe3, 0x17, 0xdf, 0x8f, 0xdb, 0xef, 0x90, 0x06, 0x54, 0xc7, 0xc7, 0xdf,
	0x8e, 0xda, 0x1e, 0xd9, 0x83, 0x07, 0x67, 0xc7, 0x93, 0x97, 0x53, 0x3a, 0x3a, 0x1b, 0x1d, 0x4f,
	0x46, 0xcf, 0xda, 0x95, 0xe8, 0x7d, 0x68, 0x6e, 0x30, 0x93, 0x3a, 0xf8, 0xc7, 0x93, 0xa1, 0x71,
	0x79, 0x36, 0x9a, 0x0c, 0xdb, 0x5e, 0xf4, 0xbb, 0x07, 0x9d, 0x72, 0x89, 0xc4, 0x8a, 0x65, 0x02,
	0x55, 0x8d, 0x66, 0x2c, 0xcf, 0x36, 0x35, 0xd2, 0x02, 0x21, 0x50, 0xcd, 0xf0, 0xb5, 0xab, 0x90,
	0x5e, 0x2b, 0x4b, 0xc9, 0x64, 0x9c, 0xea, 0xea, 0xf8, 0xd4, 0x08, 0xe4, 0x33, 0x68, 0xd8, 0xd4,
	0x45, 0x58, 0xed, 0xfa, 0xbd, 0xd6, 0xe0, 0x51, 0x99, 0x10, 0x7b, 0x22, 0xdd, 0x98, 0x45, 0x27,
	0x70, 0x70, 0x82, 0x0e, 0x89, 0xe1, 0xcb, 0x75, 0x8c, 0x3a, 0x37, 0x5e, 0xa2, 0x06, 0xa3, 0xce,
	0x8d, 0x97, 0x48, 0x42, 0xa8, 0x5f, 0x23, 0x17, 0x09, 0xcb, 0x34, 0x9c, 0x1a, 0x75, 0x62, 0x24,
	0x21, 0xbc, 0x1d, 0xc8, 0xe6, 0xb5, 0x2d, 0xd2, 0xc7, 0x50, 0x55, 0xdd, 0xae, 0xc3, 0xb4, 0x06,
	0xa4, 0x8c, 0xf3, 0x79, 0xb6, 0x60, 0x54, 0xeb, 0xcb, 0xa5, 0xf2, 0x6f, 0x96, 0xea, 0xb4, 0x78,
	0xea, 0x90, 0x65, 0x12, 0x33, 0xf9, 0x76, 0xf8, 0xff, 0xaa, 0x40, 0xe7, 0x7c, 0x35, 0x8f, 0x25,
	0x3a, 0x92, 0xee, 0x08, 0xf3, 0x09, 0xd4, 0xf4, 0x3d, 0xb7, 0xe8, 0xf7, 0x0c, 0x7a, 0x33, 0x0c,
	0x86, 0xea, 0x4b, 0x8d, 0x9e, 0x7c, 0x0a, 0xc1, 0x75, 0x9c, 0xe6, 0x28, 0x34, 0xf4, 0x4d, 0x9e,
	0xd6, 0x52, 0x0f, 0x09, 0x6a, 0x2d, 0xc8, 0x01, 0xd4, 0xe7, 0x7c, 0xad, 0x6e, 0xb9, 0xbe, 0x34,
	0x0d, 0x1a, 0xcc, 0xf9, 0x9a, 0xe6, 0x19, 0xf9, 0x10, 0x1e, 0xcc, 0x13, 0x11, 0x5f, 0xa4, 0x38,
	0xbd, 0x62, 0xec, 0x95, 0xd0, 0xf7, 0xa6, 0x41, 0x77, 0xec, 0xe6, 0xa9, 0xda, 0x23, 0x87, 0xaa,
	0xf6, 0x33, 0x8e, 0xb1, 0xc4, 0x30, 0xd0, 0xfa, 0x8d, 0xac, 0xb2, 0x96, 0xc9, 0x12, 0x59, 0x2e,
	0x75, 0xb3, 0xfb, 0xd4, 0x89, 0xe4, 0x03, 0xd8, 0xe1, 0x28, 0x50, 0x4e, 0x2d, 0xca, 0x86, 0xf6,
	0x6c, 0xe9, 0xbd, 0xef, 0x0c, 0x2c, 0x02, 0xd5, 0x5f, 0xe3, 0x44, 0x86, 0x4d, 0xad, 0xd2, 0x6b,
	0xe3, 0x96, 0x0b, 0x74, 0x6e, 0xe0, 0xdc, 0x72, 0x81, 0xd6, 0xad, 0x03, 0xb5, 0x05, 0xe3, 0x33,
	0x0c, 0x5b, 0x5a, 0x67, 0x84, 0xe8, 0x6f, 0x0f, 0xf6, 0x29, 0x4b, 0xd3, 0x8b, 0x78, 0xf6, 0xea,
	0x1e, 0x3c, 0x17, 0x28, 0xa9, 0xdc, 0x4d, 0x89, 0xbf, 0x85, 0x92, 0x42, 0xb1, 0xab, 0xa5, 0x62,
	0x97, 0xc8, 0xaa, 0xbd, 0x99, 0xac, 0xa0, 0x4c, 0x96, 0x63, 0xa2, 0x5e, 0x60, 0x62, 0x93, 0x66,
	0xa3, 0x98, 0xe6, 0x1f, 0x15, 0x78, 0xf4, 0x3c, 0x13, 0x32, 0x4e, 0xd3, 0x1b, 0x59, 0x6e, 0x3a,
	0xc7, 0xbb, 0x77, 0xe7, 0x54, 0xfe, 0x4f, 0xe7, 0xf8, 0x25, 0x9a, 0x1c, 0xa7, 0xd5, 0x02, 0xa7,
	0xf7, 0xea, 0xa6, 0xd2, 0xad, 0x0b, 0x6e, 0xdc, 0x3a, 0xf2, 0x1e, 0x80, 0x29, 0xbf, 0x0e, 0x6e,
	0xe8, 0x68, 0xea, 0x9d, 0xb1, 0xbd, 0x64, 0x8e, 0xc1, 0xc6, 0x76, 0x06, 0x0b, 0xbd, 0x14, 0xfd,
	0xe6, 0xc1, 0xc1, 0x79, 0x96, 0x6c, 0x65, 0x6b, 0x5b, 0x4f, 0xdc, 0xc2, 0x5f, 0xd9, 0x82, 0xbf,
	0x03, 0xb5, 0x55, 0xce, 0x2f, 0xd1, 0xf2, 0x61, 0x84, 0x22, 0xb0, 0x6a, 0x09, 0x58, 0x34, 0x85,
	0xf0, 0x36, 0x06, 0x3b, 0xbd, 0x8e, 0xa0, 0x6e, 0xe7, 0x92, 0x2d, 0xda, 0x1b, 0x86, 0xaa, 0xb3,
	0x52, 0xa8, 0x37, 0xa3, 0xad, 0x69, 0xc6, 0x58, 0xf4, 0x25, 0xec, 0x9d, 0xa0, 0x3c, 0x4d, 0x84,
	0x64, 0x7c, 0x7d, 0x57, 0x7a, 0x6d, 0xf0, 0x97, 0xf1, 0x6b, 0x3b, 0x9d, 0xd4, 0x32, 0xfa, 0x11,
	0xc8, 0x4b, 0xdc, 0xbc, 0x16, 0xff, 0x31, 0xdd, 0x5c, 0x7e, 0x95, 0x32, 0xf1, 0x21, 0xd4, 0x67,
	0x29, 0xc6, 0x59, 0xbe, 0xb2, 0x8c, 0x38, 0x31, 0xfa, 0x09, 0x1e, 0x96, 0xa2, 0xdb, 0xa4, 0x15,
	0x0c, 0x71, 0x69, 0xa3, 0xab, 0x25, 0xf9, 0x02, 0x02, 0xf3, 0x84, 0xea, 0xd8, 0xbb, 0x83, 0x27,
	0x65, 0x16, 0x74, 0x90, 0x3c, 0xb3, 0x6f, 0x2e, 0xb5, 0xb6, 0x4f, 0xe1, 0x87, 0x86, 0x7b, 0xd8,
	0x2f, 0x02, 0xfd, 0x0f, 0xe5, 0xf3, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x31, 0x61, 0x13, 0xb4,
	0x73, 0x09, 0x00, 0x00,
}
