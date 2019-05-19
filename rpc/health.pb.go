// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/health.proto

package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HealthStatus struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthStatus) Reset()         { *m = HealthStatus{} }
func (m *HealthStatus) String() string { return proto.CompactTextString(m) }
func (*HealthStatus) ProtoMessage()    {}
func (*HealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_b88ecd57ea43e522, []int{0}
}
func (m *HealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthStatus.Unmarshal(m, b)
}
func (m *HealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthStatus.Marshal(b, m, deterministic)
}
func (dst *HealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthStatus.Merge(dst, src)
}
func (m *HealthStatus) XXX_Size() int {
	return xxx_messageInfo_HealthStatus.Size(m)
}
func (m *HealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HealthStatus proto.InternalMessageInfo

func (m *HealthStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthStatus)(nil), "cloud.api.rpc.HealthStatus")
}

func init() { proto.RegisterFile("rpc/health.proto", fileDescriptor_health_b88ecd57ea43e522) }

var fileDescriptor_health_b88ecd57ea43e522 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2a, 0x48, 0xd6,
	0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0xce,
	0xc9, 0x2f, 0x4d, 0xd1, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x2a, 0x48, 0x56, 0x52, 0xe3, 0xe2, 0xf1,
	0x00, 0x4b, 0x07, 0x97, 0x24, 0x96, 0x94, 0x16, 0x0b, 0x89, 0x71, 0xb1, 0x15, 0x83, 0x59, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x13, 0x6b, 0x14, 0x73, 0x51, 0x41, 0x72, 0x12,
	0x1b, 0xd8, 0x10, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xaa, 0x90, 0xb8, 0x58, 0x00,
	0x00, 0x00,
}
