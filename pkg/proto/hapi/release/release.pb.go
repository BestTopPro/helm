// Code generated by protoc-gen-go.
// source: hapi/release/release.proto
// DO NOT EDIT!

package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hapi_chart "github.com/kubernetes/helm/pkg/proto/hapi/chart"
import hapi_chart3 "github.com/kubernetes/helm/pkg/proto/hapi/chart"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Release describes a deployment of a chart, together with the chart
// and the variables used to deploy that chart.
type Release struct {
	// Name is the name of the release
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Info provides information about a release
	Info *Info `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	// Chart is the chart that was released.
	Chart *hapi_chart3.Chart `protobuf:"bytes,3,opt,name=chart" json:"chart,omitempty"`
	// Config is the set of extra Values added to the chart.
	// These values override the default values inside of the chart.
	Config *hapi_chart.Config `protobuf:"bytes,4,opt,name=config" json:"config,omitempty"`
	// Manifest is the string representation of the rendered template.
	Manifest string `protobuf:"bytes,5,opt,name=manifest" json:"manifest,omitempty"`
}

func (m *Release) Reset()                    { *m = Release{} }
func (m *Release) String() string            { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()               {}
func (*Release) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Release) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Release) GetChart() *hapi_chart3.Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

func (m *Release) GetConfig() *hapi_chart.Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Release)(nil), "hapi.release.Release")
}

var fileDescriptor1 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0xbd, 0x0e, 0x82, 0x30,
	0x14, 0x85, 0x83, 0xf2, 0x23, 0xd5, 0xc5, 0x3b, 0x68, 0xc3, 0x64, 0x1c, 0xd4, 0x38, 0x94, 0x44,
	0xdf, 0x40, 0x27, 0xd7, 0x8e, 0x6e, 0x95, 0x14, 0x69, 0x22, 0x2d, 0x01, 0x9e, 0xcd, 0xe7, 0x93,
	0xf6, 0x56, 0x83, 0xcb, 0x85, 0xde, 0xef, 0xcb, 0xe9, 0x29, 0xc9, 0x2a, 0xd1, 0xa8, 0xbc, 0x95,
	0x2f, 0x29, 0x3a, 0xf9, 0xfd, 0xb2, 0xa6, 0x35, 0xbd, 0x81, 0x85, 0x65, 0xcc, 0xef, 0xb2, 0xf5,
	0x9f, 0xa9, 0x74, 0x69, 0x50, 0xf3, 0xa0, 0xa8, 0x44, 0xdb, 0xe7, 0x85, 0xd1, 0xa5, 0x7a, 0x7a,
	0xb0, 0x1a, 0x03, 0x3b, 0x71, 0xbf, 0x7d, 0x07, 0x24, 0xe1, 0x98, 0x03, 0x40, 0x42, 0x2d, 0x6a,
	0x49, 0x83, 0x4d, 0x70, 0x48, 0xb9, 0xfb, 0x87, 0x1d, 0x09, 0x6d, 0x3c, 0x9d, 0x0c, 0xbb, 0xf9,
	0x09, 0xd8, 0xb8, 0x06, 0xbb, 0x0d, 0x84, 0x3b, 0x0e, 0x7b, 0x12, 0xb9, 0x58, 0x3a, 0x75, 0xe2,
	0x12, 0x45, 0xbc, 0xe9, 0x6a, 0x27, 0x47, 0x0e, 0x47, 0x12, 0x63, 0x31, 0x1a, 0x8e, 0x23, 0xbd,
	0xe9, 0x08, 0xf7, 0x06, 0x64, 0x64, 0x56, 0x0b, 0xad, 0x4a, 0xd9, 0xf5, 0x34, 0x72, 0xa5, 0x7e,
	0xe7, 0x4b, 0x7a, 0x4f, 0x7c, 0x8d, 0x47, 0xec, 0x9e, 0x72, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0xf3, 0x60, 0x0b, 0x40, 0x01, 0x00, 0x00,
}
