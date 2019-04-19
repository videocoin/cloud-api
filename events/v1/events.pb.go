// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: events/v1/events.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payload              []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty" gorm:"payload;type:json"sql:"type:json"`
	CreatedAt            *time.Time `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_17ed077b2d8d2864, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "cloud.api.events.v1.Event")
}

func init() { proto.RegisterFile("events/v1/events.proto", fileDescriptor_events_17ed077b2d8d2864) }

var fileDescriptor_events_17ed077b2d8d2864 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2d, 0x4b, 0xcd,
	0x2b, 0x29, 0xd6, 0x2f, 0x33, 0xd4, 0x87, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84,
	0x93, 0x73, 0xf2, 0x4b, 0x53, 0xf4, 0x12, 0x0b, 0x32, 0xf5, 0xa0, 0xe2, 0x65, 0x86, 0x52, 0xba,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa,
	0x60, 0xb5, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xcc, 0x90, 0x92, 0x4f, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0x45, 0xa8, 0x2a, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x80,
	0x28, 0x50, 0x5a, 0xcc, 0xc8, 0xc5, 0xea, 0x0a, 0x32, 0x5d, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31,
	0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x72, 0xe5, 0x62, 0x2f, 0x48,
	0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x71, 0xd2, 0xfe, 0x74, 0x4f,
	0x5e, 0x3d, 0x3d, 0xbf, 0x28, 0xd7, 0x4a, 0x09, 0x2a, 0x61, 0x5d, 0x52, 0x59, 0x90, 0x6a, 0x95,
	0x55, 0x9c, 0x9f, 0xa7, 0x54, 0x5c, 0x98, 0x63, 0xa5, 0x84, 0xe0, 0x06, 0xc1, 0xf4, 0x0a, 0xd9,
	0x73, 0x71, 0x25, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0xa6, 0xc4, 0x27, 0x96, 0x48, 0x30, 0x2b, 0x30,
	0x6a, 0x70, 0x1b, 0x49, 0xe9, 0x41, 0x9c, 0xa6, 0x07, 0x73, 0x9a, 0x5e, 0x08, 0xcc, 0x69, 0x4e,
	0x2c, 0x13, 0xee, 0xcb, 0x33, 0x06, 0x71, 0x42, 0xf5, 0x38, 0x96, 0x38, 0x89, 0x9c, 0x78, 0x24,
	0xc7, 0x70, 0xe1, 0x91, 0x1c, 0xc3, 0x83, 0x47, 0x72, 0x0c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0x31,
	0x95, 0x19, 0x26, 0xb1, 0x81, 0xb5, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x20, 0x23,
	0xf2, 0x41, 0x01, 0x00, 0x00,
}
