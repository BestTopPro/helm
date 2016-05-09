// Code generated by protoc-gen-go.
// source: hapi/services/tiller.proto
// DO NOT EDIT!

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
	GetReleaseContentResponse
	UpdateReleaseRequest
	UpdateReleaseResponse
	InstallReleaseRequest
	InstallReleaseResponse
	UninstallReleaseRequest
	UninstallReleaseResponse
*/
package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hapi_chart3 "github.com/kubernetes/helm/pkg/proto/hapi/chart"
import hapi_chart "github.com/kubernetes/helm/pkg/proto/hapi/chart"
import hapi_release2 "github.com/kubernetes/helm/pkg/proto/hapi/release"
import hapi_release1 "github.com/kubernetes/helm/pkg/proto/hapi/release"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

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

// ListReleasesRequest requests a list of releases.
//
// Releases can be retrieved in chunks by setting limit and offset.
//
// Releases can be sorted according to a few pre-determined sort stategies.
type ListReleasesRequest struct {
	// Limit is the maximum number of releases to be returned.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset is the zero-based offset at which the returned release list begins.
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// SortBy is the sort field that the ListReleases server should sort data before returning.
	SortBy ListSort_SortBy `protobuf:"varint,3,opt,name=sort_by,json=sortBy,enum=hapi.services.tiller.ListSort_SortBy" json:"sort_by,omitempty"`
	// Filter is a regular expression used to filter which releases should be listed.
	//
	// Anything that matches the regexp will be included in the results.
	Filter string `protobuf:"bytes,4,opt,name=filter" json:"filter,omitempty"`
}

func (m *ListReleasesRequest) Reset()                    { *m = ListReleasesRequest{} }
func (m *ListReleasesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListReleasesRequest) ProtoMessage()               {}
func (*ListReleasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListSort struct {
}

func (m *ListSort) Reset()                    { *m = ListSort{} }
func (m *ListSort) String() string            { return proto.CompactTextString(m) }
func (*ListSort) ProtoMessage()               {}
func (*ListSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// ListReleasesResponse is a list of releases.
type ListReleasesResponse struct {
	// The expected total number of releases to be returned
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// The zero-based offset at which the list is positioned
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// The total number of queryable releases
	Total int64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	// The resulting releases
	Releases []*hapi_release2.Release `protobuf:"bytes,4,rep,name=releases" json:"releases,omitempty"`
}

func (m *ListReleasesResponse) Reset()                    { *m = ListReleasesResponse{} }
func (m *ListReleasesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListReleasesResponse) ProtoMessage()               {}
func (*ListReleasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListReleasesResponse) GetReleases() []*hapi_release2.Release {
	if m != nil {
		return m.Releases
	}
	return nil
}

// GetReleaseStatusRequest is a request to get the status of a release.
type GetReleaseStatusRequest struct {
	// Name is the name of the release
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetReleaseStatusRequest) Reset()                    { *m = GetReleaseStatusRequest{} }
func (m *GetReleaseStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReleaseStatusRequest) ProtoMessage()               {}
func (*GetReleaseStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// GetReleaseStatusResponse is the response indicating the status of the named release.
type GetReleaseStatusResponse struct {
	// Name is the name of the release.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Info contains information about the release.
	Info *hapi_release1.Info `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *GetReleaseStatusResponse) Reset()                    { *m = GetReleaseStatusResponse{} }
func (m *GetReleaseStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetReleaseStatusResponse) ProtoMessage()               {}
func (*GetReleaseStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetReleaseStatusResponse) GetInfo() *hapi_release1.Info {
	if m != nil {
		return m.Info
	}
	return nil
}

// GetReleaseContentRequest is a request to get the contents of a release.
type GetReleaseContentRequest struct {
	// The name of the release
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetReleaseContentRequest) Reset()                    { *m = GetReleaseContentRequest{} }
func (m *GetReleaseContentRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReleaseContentRequest) ProtoMessage()               {}
func (*GetReleaseContentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// GetReleaseContentResponse is a response containing the contents of a release.
type GetReleaseContentResponse struct {
	// The release content
	Release *hapi_release2.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
}

func (m *GetReleaseContentResponse) Reset()                    { *m = GetReleaseContentResponse{} }
func (m *GetReleaseContentResponse) String() string            { return proto.CompactTextString(m) }
func (*GetReleaseContentResponse) ProtoMessage()               {}
func (*GetReleaseContentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetReleaseContentResponse) GetRelease() *hapi_release2.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

// UpdateReleaseRequest updates a release.
type UpdateReleaseRequest struct {
}

func (m *UpdateReleaseRequest) Reset()                    { *m = UpdateReleaseRequest{} }
func (m *UpdateReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateReleaseRequest) ProtoMessage()               {}
func (*UpdateReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// UpdateReleaseResponse is the response to an update request.
type UpdateReleaseResponse struct {
}

func (m *UpdateReleaseResponse) Reset()                    { *m = UpdateReleaseResponse{} }
func (m *UpdateReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateReleaseResponse) ProtoMessage()               {}
func (*UpdateReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// InstallReleaseRequest is the request for an installation of a chart.
type InstallReleaseRequest struct {
	// Chart is the protobuf representation of a chart.
	Chart *hapi_chart3.Chart `protobuf:"bytes,1,opt,name=chart" json:"chart,omitempty"`
	// Values is a string containing (unparsed) TOML values.
	Values *hapi_chart.Config `protobuf:"bytes,2,opt,name=values" json:"values,omitempty"`
	// DryRun, if true, will run through the release logic, but neither create
	// a release object nor deploy to Kubernetes. The release object returned
	// in the response will be fake.
	DryRun bool `protobuf:"varint,3,opt,name=dry_run,json=dryRun" json:"dry_run,omitempty"`
}

func (m *InstallReleaseRequest) Reset()                    { *m = InstallReleaseRequest{} }
func (m *InstallReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*InstallReleaseRequest) ProtoMessage()               {}
func (*InstallReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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

// InstallReleaseResponse is the response from a release installation.
type InstallReleaseResponse struct {
	Release *hapi_release2.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
}

func (m *InstallReleaseResponse) Reset()                    { *m = InstallReleaseResponse{} }
func (m *InstallReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*InstallReleaseResponse) ProtoMessage()               {}
func (*InstallReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *InstallReleaseResponse) GetRelease() *hapi_release2.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

// UninstallReleaseRequest represents a request to uninstall a named release.
type UninstallReleaseRequest struct {
	// Name is the name of the release to delete.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *UninstallReleaseRequest) Reset()                    { *m = UninstallReleaseRequest{} }
func (m *UninstallReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*UninstallReleaseRequest) ProtoMessage()               {}
func (*UninstallReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// UninstallReleaseResponse represents a successful response to an uninstall request.
type UninstallReleaseResponse struct {
	// Release is the release that was marked deleted.
	Release *hapi_release2.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
}

func (m *UninstallReleaseResponse) Reset()                    { *m = UninstallReleaseResponse{} }
func (m *UninstallReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*UninstallReleaseResponse) ProtoMessage()               {}
func (*UninstallReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UninstallReleaseResponse) GetRelease() *hapi_release2.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func init() {
	proto.RegisterType((*ListReleasesRequest)(nil), "hapi.services.tiller.ListReleasesRequest")
	proto.RegisterType((*ListSort)(nil), "hapi.services.tiller.ListSort")
	proto.RegisterType((*ListReleasesResponse)(nil), "hapi.services.tiller.ListReleasesResponse")
	proto.RegisterType((*GetReleaseStatusRequest)(nil), "hapi.services.tiller.GetReleaseStatusRequest")
	proto.RegisterType((*GetReleaseStatusResponse)(nil), "hapi.services.tiller.GetReleaseStatusResponse")
	proto.RegisterType((*GetReleaseContentRequest)(nil), "hapi.services.tiller.GetReleaseContentRequest")
	proto.RegisterType((*GetReleaseContentResponse)(nil), "hapi.services.tiller.GetReleaseContentResponse")
	proto.RegisterType((*UpdateReleaseRequest)(nil), "hapi.services.tiller.UpdateReleaseRequest")
	proto.RegisterType((*UpdateReleaseResponse)(nil), "hapi.services.tiller.UpdateReleaseResponse")
	proto.RegisterType((*InstallReleaseRequest)(nil), "hapi.services.tiller.InstallReleaseRequest")
	proto.RegisterType((*InstallReleaseResponse)(nil), "hapi.services.tiller.InstallReleaseResponse")
	proto.RegisterType((*UninstallReleaseRequest)(nil), "hapi.services.tiller.UninstallReleaseRequest")
	proto.RegisterType((*UninstallReleaseResponse)(nil), "hapi.services.tiller.UninstallReleaseResponse")
	proto.RegisterEnum("hapi.services.tiller.ListSort_SortBy", ListSort_SortBy_name, ListSort_SortBy_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for ReleaseService service

type ReleaseServiceClient interface {
	// ListReleases retrieves release history.
	// TODO: Allow filtering the set of releases by
	// release status. By default, ListAllReleases returns the releases who
	// current status is "Active".
	ListReleases(ctx context.Context, in *ListReleasesRequest, opts ...grpc.CallOption) (ReleaseService_ListReleasesClient, error)
	// GetReleasesStatus retrieves status information for the specified release.
	GetReleaseStatus(ctx context.Context, in *GetReleaseStatusRequest, opts ...grpc.CallOption) (*GetReleaseStatusResponse, error)
	// GetReleaseContent retrieves the release content (chart + value) for the specifed release.
	GetReleaseContent(ctx context.Context, in *GetReleaseContentRequest, opts ...grpc.CallOption) (*GetReleaseContentResponse, error)
	// UpdateRelease updates release content.
	UpdateRelease(ctx context.Context, in *UpdateReleaseRequest, opts ...grpc.CallOption) (*UpdateReleaseResponse, error)
	// InstallRelease requests installation of a chart as a new release.
	InstallRelease(ctx context.Context, in *InstallReleaseRequest, opts ...grpc.CallOption) (*InstallReleaseResponse, error)
	// UninstallRelease requests deletion of a named release.
	UninstallRelease(ctx context.Context, in *UninstallReleaseRequest, opts ...grpc.CallOption) (*UninstallReleaseResponse, error)
}

type releaseServiceClient struct {
	cc *grpc.ClientConn
}

func NewReleaseServiceClient(cc *grpc.ClientConn) ReleaseServiceClient {
	return &releaseServiceClient{cc}
}

func (c *releaseServiceClient) ListReleases(ctx context.Context, in *ListReleasesRequest, opts ...grpc.CallOption) (ReleaseService_ListReleasesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ReleaseService_serviceDesc.Streams[0], c.cc, "/hapi.services.tiller.ReleaseService/ListReleases", opts...)
	if err != nil {
		return nil, err
	}
	x := &releaseServiceListReleasesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReleaseService_ListReleasesClient interface {
	Recv() (*ListReleasesResponse, error)
	grpc.ClientStream
}

type releaseServiceListReleasesClient struct {
	grpc.ClientStream
}

func (x *releaseServiceListReleasesClient) Recv() (*ListReleasesResponse, error) {
	m := new(ListReleasesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *releaseServiceClient) GetReleaseStatus(ctx context.Context, in *GetReleaseStatusRequest, opts ...grpc.CallOption) (*GetReleaseStatusResponse, error) {
	out := new(GetReleaseStatusResponse)
	err := grpc.Invoke(ctx, "/hapi.services.tiller.ReleaseService/GetReleaseStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) GetReleaseContent(ctx context.Context, in *GetReleaseContentRequest, opts ...grpc.CallOption) (*GetReleaseContentResponse, error) {
	out := new(GetReleaseContentResponse)
	err := grpc.Invoke(ctx, "/hapi.services.tiller.ReleaseService/GetReleaseContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) UpdateRelease(ctx context.Context, in *UpdateReleaseRequest, opts ...grpc.CallOption) (*UpdateReleaseResponse, error) {
	out := new(UpdateReleaseResponse)
	err := grpc.Invoke(ctx, "/hapi.services.tiller.ReleaseService/UpdateRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) InstallRelease(ctx context.Context, in *InstallReleaseRequest, opts ...grpc.CallOption) (*InstallReleaseResponse, error) {
	out := new(InstallReleaseResponse)
	err := grpc.Invoke(ctx, "/hapi.services.tiller.ReleaseService/InstallRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) UninstallRelease(ctx context.Context, in *UninstallReleaseRequest, opts ...grpc.CallOption) (*UninstallReleaseResponse, error) {
	out := new(UninstallReleaseResponse)
	err := grpc.Invoke(ctx, "/hapi.services.tiller.ReleaseService/UninstallRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReleaseService service

type ReleaseServiceServer interface {
	// ListReleases retrieves release history.
	// TODO: Allow filtering the set of releases by
	// release status. By default, ListAllReleases returns the releases who
	// current status is "Active".
	ListReleases(*ListReleasesRequest, ReleaseService_ListReleasesServer) error
	// GetReleasesStatus retrieves status information for the specified release.
	GetReleaseStatus(context.Context, *GetReleaseStatusRequest) (*GetReleaseStatusResponse, error)
	// GetReleaseContent retrieves the release content (chart + value) for the specifed release.
	GetReleaseContent(context.Context, *GetReleaseContentRequest) (*GetReleaseContentResponse, error)
	// UpdateRelease updates release content.
	UpdateRelease(context.Context, *UpdateReleaseRequest) (*UpdateReleaseResponse, error)
	// InstallRelease requests installation of a chart as a new release.
	InstallRelease(context.Context, *InstallReleaseRequest) (*InstallReleaseResponse, error)
	// UninstallRelease requests deletion of a named release.
	UninstallRelease(context.Context, *UninstallReleaseRequest) (*UninstallReleaseResponse, error)
}

func RegisterReleaseServiceServer(s *grpc.Server, srv ReleaseServiceServer) {
	s.RegisterService(&_ReleaseService_serviceDesc, srv)
}

func _ReleaseService_ListReleases_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListReleasesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReleaseServiceServer).ListReleases(m, &releaseServiceListReleasesServer{stream})
}

type ReleaseService_ListReleasesServer interface {
	Send(*ListReleasesResponse) error
	grpc.ServerStream
}

type releaseServiceListReleasesServer struct {
	grpc.ServerStream
}

func (x *releaseServiceListReleasesServer) Send(m *ListReleasesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ReleaseService_GetReleaseStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetReleaseStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ReleaseServiceServer).GetReleaseStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ReleaseService_GetReleaseContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetReleaseContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ReleaseServiceServer).GetReleaseContent(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ReleaseService_UpdateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ReleaseServiceServer).UpdateRelease(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ReleaseService_InstallRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InstallReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ReleaseServiceServer).InstallRelease(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ReleaseService_UninstallRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UninstallReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ReleaseServiceServer).UninstallRelease(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ReleaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hapi.services.tiller.ReleaseService",
	HandlerType: (*ReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReleaseStatus",
			Handler:    _ReleaseService_GetReleaseStatus_Handler,
		},
		{
			MethodName: "GetReleaseContent",
			Handler:    _ReleaseService_GetReleaseContent_Handler,
		},
		{
			MethodName: "UpdateRelease",
			Handler:    _ReleaseService_UpdateRelease_Handler,
		},
		{
			MethodName: "InstallRelease",
			Handler:    _ReleaseService_InstallRelease_Handler,
		},
		{
			MethodName: "UninstallRelease",
			Handler:    _ReleaseService_UninstallRelease_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListReleases",
			Handler:       _ReleaseService_ListReleases_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0x5e, 0xb6, 0x2e, 0xed, 0xce, 0x7e, 0x9b, 0x3a, 0xff, 0xba, 0x36, 0xe4, 0x0a, 0x59, 0x02,
	0xc6, 0x60, 0x29, 0x94, 0xfb, 0x49, 0xdd, 0xa8, 0x50, 0xb5, 0x52, 0x24, 0x97, 0x82, 0xc4, 0x4d,
	0x95, 0x75, 0x2e, 0x0b, 0x4a, 0x93, 0x12, 0xbb, 0x95, 0xfa, 0x00, 0x5c, 0xf2, 0x04, 0xbc, 0x08,
	0x8f, 0x87, 0x63, 0x3b, 0x51, 0xd3, 0x26, 0x50, 0xed, 0xc6, 0xd9, 0xf1, 0xf7, 0x1d, 0x7f, 0xe7,
	0xef, 0x0a, 0xf6, 0xbd, 0x3b, 0xf3, 0x9a, 0x8c, 0x46, 0x0b, 0x6f, 0x4c, 0x59, 0x93, 0x7b, 0xbe,
	0x4f, 0x23, 0x67, 0x16, 0x85, 0x3c, 0x44, 0xb5, 0x18, 0x73, 0x12, 0xcc, 0x51, 0x98, 0x5d, 0x97,
	0x1e, 0xe3, 0x7b, 0x37, 0xe2, 0xea, 0x54, 0x6c, 0xbb, 0xb1, 0x7a, 0x1f, 0x06, 0x13, 0xef, 0xab,
	0x06, 0x94, 0x44, 0x44, 0x7d, 0xea, 0x32, 0x9a, 0x7c, 0x33, 0x4e, 0x09, 0xe6, 0x05, 0x93, 0x50,
	0x01, 0xf8, 0x97, 0x01, 0xff, 0xf7, 0x3c, 0xc6, 0x89, 0x82, 0x18, 0xa1, 0xdf, 0xe7, 0x94, 0x71,
	0x54, 0x83, 0x7d, 0xdf, 0x9b, 0x7a, 0xdc, 0x32, 0x1e, 0x1b, 0x67, 0x7b, 0x44, 0x19, 0xa8, 0x0e,
	0x66, 0x38, 0x99, 0x30, 0xca, 0xad, 0x5d, 0x79, 0xad, 0x2d, 0x74, 0x09, 0x65, 0x16, 0x46, 0x7c,
	0x74, 0xbb, 0xb4, 0xf6, 0x04, 0x70, 0xdc, 0x7a, 0xe2, 0xe4, 0xe5, 0xe4, 0xc4, 0x4a, 0x03, 0x41,
	0x74, 0xe2, 0xe3, 0x6a, 0x49, 0x4c, 0x26, 0xbf, 0xf1, 0xbb, 0x13, 0xcf, 0xe7, 0x34, 0xb2, 0x4a,
	0xc2, 0xfd, 0x80, 0x68, 0x0b, 0x5f, 0x42, 0x25, 0x71, 0xc1, 0x2d, 0x30, 0x95, 0x17, 0x3a, 0x84,
	0xf2, 0xb0, 0x7f, 0xd3, 0xff, 0xf0, 0xb9, 0x5f, 0xdd, 0x41, 0x15, 0x28, 0xf5, 0xdb, 0xef, 0x3b,
	0x55, 0x03, 0x9d, 0xc0, 0x51, 0xaf, 0x3d, 0xf8, 0x38, 0x22, 0x9d, 0x5e, 0xa7, 0x3d, 0xe8, 0xbc,
	0xad, 0xee, 0xe2, 0x9f, 0x06, 0xd4, 0xb2, 0xd9, 0xb1, 0x59, 0x18, 0x30, 0x1a, 0xa7, 0x37, 0x0e,
	0xe7, 0x41, 0x9a, 0x9e, 0x34, 0x0a, 0xd3, 0x13, 0x6c, 0x1e, 0x72, 0xd7, 0x97, 0xc9, 0x09, 0xb6,
	0x34, 0xd0, 0x6b, 0xa8, 0xe8, 0x82, 0x32, 0x11, 0xf6, 0xde, 0xd9, 0x61, 0xeb, 0x54, 0x65, 0x9d,
	0x94, 0x5e, 0xab, 0x92, 0x94, 0x86, 0x2f, 0xa0, 0xf1, 0x8e, 0x26, 0xd1, 0x0c, 0xb8, 0xcb, 0xe7,
	0x69, 0xc1, 0x11, 0x94, 0x02, 0x77, 0x4a, 0x65, 0x40, 0x07, 0x44, 0xfe, 0x8d, 0x3f, 0x81, 0xb5,
	0x49, 0xd7, 0x19, 0xe4, 0xf0, 0xd1, 0x53, 0x28, 0xc5, 0xad, 0x95, 0xd1, 0x1f, 0xb6, 0x50, 0x36,
	0x9a, 0xae, 0x40, 0x88, 0xc4, 0xb1, 0xb3, 0xfa, 0xee, 0x75, 0x18, 0x70, 0x1a, 0xf0, 0xbf, 0xc5,
	0xd1, 0x83, 0x47, 0x39, 0x7c, 0x1d, 0x48, 0x13, 0xca, 0x5a, 0x42, 0xfa, 0x14, 0x56, 0x21, 0x61,
	0xe1, 0x3a, 0xd4, 0x86, 0xb3, 0x3b, 0x97, 0xd3, 0x04, 0x51, 0xca, 0xb8, 0x01, 0xa7, 0x6b, 0xf7,
	0x4a, 0x01, 0xff, 0x30, 0xe0, 0xb4, 0x1b, 0x30, 0x51, 0x73, 0x3f, 0xeb, 0x82, 0x9e, 0x89, 0x36,
	0xc6, 0x8b, 0xa0, 0x95, 0x4f, 0x94, 0xb2, 0xda, 0x96, 0xeb, 0xf8, 0x24, 0x0a, 0x47, 0xe7, 0x60,
	0x2e, 0x5c, 0x5f, 0xf8, 0x64, 0x6b, 0xa3, 0x99, 0x72, 0x8b, 0x88, 0x66, 0xa0, 0x06, 0x94, 0xef,
	0xa2, 0xe5, 0x28, 0x9a, 0x07, 0xb2, 0xdf, 0x15, 0x62, 0x0a, 0x93, 0xcc, 0x03, 0xdc, 0x85, 0xfa,
	0x7a, 0x18, 0x0f, 0xad, 0x81, 0x18, 0x84, 0x61, 0xe0, 0xe5, 0xe6, 0x94, 0xd7, 0x80, 0x1b, 0xb0,
	0x36, 0xe9, 0x0f, 0xd4, 0x6e, 0xfd, 0xde, 0x87, 0xe3, 0x64, 0xa6, 0xd4, 0x7e, 0x22, 0x0f, 0xfe,
	0x5b, 0x5d, 0x13, 0xf4, 0xbc, 0x78, 0x7d, 0xd7, 0xfe, 0x51, 0xd8, 0xe7, 0xdb, 0x50, 0x75, 0x23,
	0x77, 0x5e, 0x19, 0x88, 0x41, 0x75, 0x7d, 0xa6, 0xd1, 0x45, 0xfe, 0x1b, 0x05, 0xab, 0x62, 0x3b,
	0xdb, 0xd2, 0x13, 0x59, 0xb4, 0x80, 0x93, 0x8d, 0x01, 0x46, 0xff, 0x7c, 0x26, 0xbb, 0x19, 0x76,
	0x73, 0x6b, 0x7e, 0xaa, 0xfb, 0x0d, 0x8e, 0x32, 0x23, 0x8d, 0x0a, 0xaa, 0x95, 0xb7, 0x0f, 0xf6,
	0x8b, 0xad, 0xb8, 0xa9, 0xd6, 0x14, 0x8e, 0xb3, 0xd3, 0x89, 0x0a, 0x1e, 0xc8, 0x5d, 0x25, 0xfb,
	0xe5, 0x76, 0xe4, 0x54, 0x4e, 0xf4, 0x71, 0x7d, 0x24, 0x8b, 0xfa, 0x58, 0x30, 0xe9, 0x45, 0x7d,
	0x2c, 0x9a, 0x74, 0xbc, 0x73, 0x05, 0x5f, 0x2a, 0x09, 0xfb, 0xd6, 0x94, 0x3f, 0x60, 0x6f, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x16, 0x42, 0xc7, 0x5a, 0x07, 0x00, 0x00,
}
