// Code generated by protoc-gen-go.
// source: hapi/chart/metadata.proto
// DO NOT EDIT!

package chart

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Maintainer describes a Chart maintainer.
type Maintainer struct {
	// Name is a user name or organization name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Email is an optional email address to contact the named maintainer
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Maintainer) Reset()                    { *m = Maintainer{} }
func (m *Maintainer) String() string            { return proto.CompactTextString(m) }
func (*Maintainer) ProtoMessage()               {}
func (*Maintainer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// 	Metadata for a Chart file. This models the structure of a Chart.yaml file.
//
// 	Spec: https://github.com/kubernetes/helm/blob/master/docs/design/chart_format.md#the-chart-file
type Metadata struct {
	// The name of the chart
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The URL to a relevant project page, git repo, or contact person
	Home string `protobuf:"bytes,2,opt,name=home" json:"home,omitempty"`
	// Source is the URL to the source code of this chart
	Sources []string `protobuf:"bytes,3,rep,name=sources" json:"sources,omitempty"`
	// A SemVer 2 conformant version string of the chart
	Version string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// A one-sentence description of the chart
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	// A list of string keywords
	Keywords []string `protobuf:"bytes,6,rep,name=keywords" json:"keywords,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	Maintainers []*Maintainer `protobuf:"bytes,7,rep,name=maintainers" json:"maintainers,omitempty"`
	// The name of the template engine to use. Defaults to 'gotpl'.
	Engine string `protobuf:"bytes,8,opt,name=engine" json:"engine,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Metadata) GetMaintainers() []*Maintainer {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

func init() {
	proto.RegisterType((*Maintainer)(nil), "hapi.chart.Maintainer")
	proto.RegisterType((*Metadata)(nil), "hapi.chart.Metadata")
}

var fileDescriptor2 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4f, 0xc3, 0x40,
	0x0c, 0xc5, 0x15, 0xda, 0x7c, 0xe0, 0x6c, 0x16, 0xaa, 0x0c, 0x53, 0xd4, 0x89, 0x29, 0x95, 0x40,
	0x42, 0xcc, 0xec, 0x5d, 0x3a, 0xb2, 0x99, 0xc4, 0x22, 0x27, 0x48, 0x2e, 0xba, 0x3b, 0x40, 0xfc,
	0xe3, 0xcc, 0x5c, 0xdc, 0xaf, 0x0c, 0x1d, 0x22, 0xbd, 0xf7, 0x7e, 0x79, 0x3e, 0xd9, 0x70, 0xdb,
	0xf1, 0x68, 0x36, 0x4d, 0xc7, 0x2e, 0x6c, 0x7a, 0x09, 0xdc, 0x72, 0xe0, 0x7a, 0x74, 0x36, 0x58,
	0x84, 0x09, 0xd5, 0x8a, 0xd6, 0x4f, 0x00, 0x5b, 0x36, 0x43, 0x88, 0x9f, 0x38, 0x44, 0x58, 0x0e,
	0xdc, 0x0b, 0x25, 0x55, 0x72, 0x7f, 0xbd, 0x53, 0x8d, 0x37, 0x90, 0x4a, 0xcf, 0xe6, 0x93, 0xae,
	0x34, 0xdc, 0x9b, 0xf5, 0x5f, 0x02, 0xc5, 0xf6, 0x30, 0xf6, 0x62, 0x2d, 0x66, 0x9d, 0x8d, 0xd9,
	0xbe, 0xa5, 0x1a, 0x09, 0x72, 0x6f, 0xbf, 0x5c, 0x23, 0x9e, 0x16, 0xd5, 0x22, 0xc6, 0x47, 0x3b,
	0x91, 0x6f, 0x71, 0xde, 0xd8, 0x81, 0x96, 0x5a, 0x38, 0x5a, 0xac, 0xa0, 0x6c, 0xc5, 0x37, 0xce,
	0x8c, 0x61, 0xa2, 0xa9, 0xd2, 0x79, 0x84, 0x77, 0x50, 0x7c, 0xc8, 0xef, 0x8f, 0x75, 0xad, 0xa7,
	0x4c, 0xc7, 0x9e, 0x3c, 0x3e, 0x43, 0xd9, 0x9f, 0xd6, 0xf3, 0x94, 0x47, 0x5c, 0x3e, 0xac, 0xea,
	0xf3, 0x01, 0xea, 0xf3, 0xf6, 0xbb, 0xf9, 0xaf, 0xb8, 0x82, 0x4c, 0x86, 0xf7, 0xa8, 0xa9, 0xd0,
	0x27, 0x0f, 0xee, 0x25, 0x7f, 0x4d, 0xb5, 0xf8, 0x96, 0xe9, 0x31, 0x1f, 0xff, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x6f, 0x4a, 0x7b, 0xd0, 0x69, 0x01, 0x00, 0x00,
}
