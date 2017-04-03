// Code generated by protoc-gen-go.
// source: hapi/release/hook.proto
// DO NOT EDIT!

/*
Package release is a generated protocol buffer package.

It is generated from these files:
	hapi/release/hook.proto
	hapi/release/info.proto
	hapi/release/release.proto
	hapi/release/status.proto
	hapi/release/test_run.proto
	hapi/release/test_suite.proto

It has these top-level messages:
	Hook
	Info
	Release
	Status
	TestRun
	TestSuite
*/
package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Hook_Event int32

const (
	Hook_UNKNOWN              Hook_Event = 0
	Hook_PRE_INSTALL          Hook_Event = 1
	Hook_POST_INSTALL         Hook_Event = 2
	Hook_PRE_DELETE           Hook_Event = 3
	Hook_POST_DELETE          Hook_Event = 4
	Hook_PRE_UPGRADE          Hook_Event = 5
	Hook_POST_UPGRADE         Hook_Event = 6
	Hook_PRE_ROLLBACK         Hook_Event = 7
	Hook_POST_ROLLBACK        Hook_Event = 8
	Hook_RELEASE_TEST_SUCCESS Hook_Event = 9
	Hook_RELEASE_TEST_FAILURE Hook_Event = 10
)

var Hook_Event_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "PRE_INSTALL",
	2:  "POST_INSTALL",
	3:  "PRE_DELETE",
	4:  "POST_DELETE",
	5:  "PRE_UPGRADE",
	6:  "POST_UPGRADE",
	7:  "PRE_ROLLBACK",
	8:  "POST_ROLLBACK",
	9:  "RELEASE_TEST_SUCCESS",
	10: "RELEASE_TEST_FAILURE",
}
var Hook_Event_value = map[string]int32{
	"UNKNOWN":              0,
	"PRE_INSTALL":          1,
	"POST_INSTALL":         2,
	"PRE_DELETE":           3,
	"POST_DELETE":          4,
	"PRE_UPGRADE":          5,
	"POST_UPGRADE":         6,
	"PRE_ROLLBACK":         7,
	"POST_ROLLBACK":        8,
	"RELEASE_TEST_SUCCESS": 9,
	"RELEASE_TEST_FAILURE": 10,
}

func (x Hook_Event) String() string {
	return proto.EnumName(Hook_Event_name, int32(x))
}
func (Hook_Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Hook defines a hook object.
type Hook struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Kind is the Kubernetes kind.
	Kind string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	// Path is the chart-relative path to the template.
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	// Manifest is the manifest contents.
	Manifest string `protobuf:"bytes,4,opt,name=manifest" json:"manifest,omitempty"`
	// Events are the events that this hook fires on.
	Events []Hook_Event `protobuf:"varint,5,rep,packed,name=events,enum=hapi.release.Hook_Event" json:"events,omitempty"`
	// LastRun indicates the date/time this was last run.
	LastRun *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=last_run,json=lastRun" json:"last_run,omitempty"`
	// Weight indicates the sort order for execution among similar Hook types
	Weight int `protobuf:"bytes,7,opt,name=weight" json:"weight,omitempty"`
}

func (m *Hook) Reset()                    { *m = Hook{} }
func (m *Hook) String() string            { return proto.CompactTextString(m) }
func (*Hook) ProtoMessage()               {}
func (*Hook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hook) GetLastRun() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastRun
	}
	return nil
}

func init() {
	proto.RegisterType((*Hook)(nil), "hapi.release.Hook")
	proto.RegisterEnum("hapi.release.Hook_Event", Hook_Event_name, Hook_Event_value)
}

func init() { proto.RegisterFile("hapi/release/hook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0xdd, 0x6e, 0xa2, 0x40,
	0x18, 0x86, 0x17, 0x41, 0xd0, 0xd1, 0x75, 0x67, 0x27, 0x9b, 0xec, 0xc4, 0x93, 0x35, 0x1e, 0x79,
	0x34, 0x6c, 0x6c, 0x7a, 0x01, 0xa8, 0xd3, 0xd6, 0x48, 0xd0, 0x0c, 0x90, 0x26, 0x3d, 0x21, 0x98,
	0x8e, 0x4a, 0x14, 0x86, 0x08, 0xf6, 0x72, 0x7a, 0x55, 0xbd, 0xa0, 0x66, 0x86, 0x9f, 0x34, 0xe9,
	0xd9, 0xc7, 0xf3, 0x3e, 0x7c, 0x33, 0xef, 0x80, 0xbf, 0xa7, 0x38, 0x4f, 0xec, 0x2b, 0xbf, 0xf0,
	0xb8, 0xe0, 0xf6, 0x49, 0x88, 0x33, 0xc9, 0xaf, 0xa2, 0x14, 0x68, 0x28, 0x03, 0x52, 0x07, 0xe3,
	0x7f, 0x47, 0x21, 0x8e, 0x17, 0x6e, 0xab, 0x6c, 0x7f, 0x3b, 0xd8, 0x65, 0x92, 0xf2, 0xa2, 0x8c,
	0xd3, 0xbc, 0xd2, 0xa7, 0xef, 0x3a, 0x30, 0x9e, 0x84, 0x38, 0x23, 0x04, 0x8c, 0x2c, 0x4e, 0x39,
	0xd6, 0x26, 0xda, 0xac, 0xcf, 0xd4, 0x2c, 0xd9, 0x39, 0xc9, 0x5e, 0x71, 0xa7, 0x62, 0x72, 0x96,
	0x2c, 0x8f, 0xcb, 0x13, 0xd6, 0x2b, 0x26, 0x67, 0x34, 0x06, 0xbd, 0x34, 0xce, 0x92, 0x03, 0x2f,
	0x4a, 0x6c, 0x28, 0xde, 0x7e, 0xa3, 0xff, 0xc0, 0xe4, 0x6f, 0x3c, 0x2b, 0x0b, 0xdc, 0x9d, 0xe8,
	0xb3, 0xd1, 0x1c, 0x93, 0xaf, 0x17, 0x24, 0xf2, 0x6c, 0x42, 0xa5, 0xc0, 0x6a, 0x0f, 0xdd, 0x83,
	0xde, 0x25, 0x2e, 0xca, 0xe8, 0x7a, 0xcb, 0xb0, 0x39, 0xd1, 0x66, 0x83, 0xf9, 0x98, 0x54, 0x35,
	0x48, 0x53, 0x83, 0x04, 0x4d, 0x0d, 0x66, 0x49, 0x97, 0xdd, 0xb2, 0xe9, 0x87, 0x06, 0xba, 0x6a,
	0x11, 0x1a, 0x00, 0x2b, 0xf4, 0x36, 0xde, 0xf6, 0xd9, 0x83, 0x3f, 0xd0, 0x2f, 0x30, 0xd8, 0x31,
	0x1a, 0xad, 0x3d, 0x3f, 0x70, 0x5c, 0x17, 0x6a, 0x08, 0x82, 0xe1, 0x6e, 0xeb, 0x07, 0x2d, 0xe9,
	0xa0, 0x11, 0x00, 0x52, 0x59, 0x51, 0x97, 0x06, 0x14, 0xea, 0xea, 0x17, 0x69, 0xd4, 0xc0, 0x68,
	0x76, 0x84, 0xbb, 0x47, 0xe6, 0xac, 0x28, 0xec, 0xb6, 0x3b, 0x1a, 0x62, 0x2a, 0xc2, 0x68, 0xc4,
	0xb6, 0xae, 0xbb, 0x70, 0x96, 0x1b, 0x68, 0xa1, 0xdf, 0xe0, 0xa7, 0x72, 0x5a, 0xd4, 0x43, 0x18,
	0xfc, 0x61, 0xd4, 0xa5, 0x8e, 0x4f, 0xa3, 0x80, 0xfa, 0x41, 0xe4, 0x87, 0xcb, 0x25, 0xf5, 0x7d,
	0xd8, 0xff, 0x96, 0x3c, 0x38, 0x6b, 0x37, 0x64, 0x14, 0x82, 0x45, 0xff, 0xc5, 0xaa, 0xdf, 0x6a,
	0x6f, 0xaa, 0xfa, 0x77, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x41, 0x62, 0x57, 0xfc, 0x01,
	0x00, 0x00,
}
