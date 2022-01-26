// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto-encoding/base128-varints/base128-varints.proto

package base128_varints

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Test2 struct {
	A                    *int32   `protobuf:"zigzag32,1,opt,name=a" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test2) Reset()         { *m = Test2{} }
func (m *Test2) String() string { return proto.CompactTextString(m) }
func (*Test2) ProtoMessage()    {}
func (*Test2) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01aa008283dcb9c, []int{0}
}

func (m *Test2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test2.Unmarshal(m, b)
}
func (m *Test2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test2.Marshal(b, m, deterministic)
}
func (m *Test2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test2.Merge(m, src)
}
func (m *Test2) XXX_Size() int {
	return xxx_messageInfo_Test2.Size(m)
}
func (m *Test2) XXX_DiscardUnknown() {
	xxx_messageInfo_Test2.DiscardUnknown(m)
}

var xxx_messageInfo_Test2 proto.InternalMessageInfo

func (m *Test2) GetA() int32 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func init() {
	proto.RegisterType((*Test2)(nil), "Test2")
}

func init() {
	proto.RegisterFile("proto-encoding/base128-varints/base128-varints.proto", fileDescriptor_d01aa008283dcb9c)
}

var fileDescriptor_d01aa008283dcb9c = []byte{
	// 84 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4d, 0xcd, 0x4b, 0xce, 0x4f, 0xc9, 0xcc, 0x4b, 0xd7, 0x4f, 0x4a, 0x2c, 0x4e, 0x35,
	0x34, 0xb2, 0xd0, 0x2d, 0x4b, 0x2c, 0xca, 0xcc, 0x2b, 0x29, 0x46, 0xe7, 0xeb, 0x81, 0x95, 0x2b,
	0x89, 0x72, 0xb1, 0x86, 0xa4, 0x16, 0x97, 0x18, 0x09, 0xf1, 0x70, 0x31, 0x26, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x08, 0x06, 0x31, 0x26, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x59, 0xe8, 0x07,
	0x4d, 0x00, 0x00, 0x00,
}
