// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transcoder/v1/transcoder.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v11 "github.com/VideoCoin/cloud-api/profiles/v1"
import v1 "github.com/VideoCoin/cloud-api/workorder/v1"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TranscoderStatus int32

const (
	TranscoderStatusAvailable TranscoderStatus = 0
	TranscoderStatusOffline   TranscoderStatus = 1
	TranscoderStatusError     TranscoderStatus = 2
	TranscoderStatusBusy      TranscoderStatus = 3
)

var TranscoderStatus_name = map[int32]string{
	0: "available",
	1: "offline",
	2: "error",
	3: "busy",
}
var TranscoderStatus_value = map[string]int32{
	"available": 0,
	"offline":   1,
	"error":     2,
	"busy":      3,
}

func (x TranscoderStatus) String() string {
	return proto.EnumName(TranscoderStatus_name, int32(x))
}
func (TranscoderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_403bea862d72b8de, []int{0}
}

type Transcoder struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CpuCores             int32    `protobuf:"varint,2,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	CpuMhz               float64  `protobuf:"fixed64,3,opt,name=cpu_mhz,json=cpuMhz,proto3" json:"cpu_mhz,omitempty"`
	TotalMemory          uint64   `protobuf:"varint,4,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Worker               []byte   `protobuf:"bytes,7,opt,name=worker,proto3" json:"worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transcoder) Reset()         { *m = Transcoder{} }
func (m *Transcoder) String() string { return proto.CompactTextString(m) }
func (*Transcoder) ProtoMessage()    {}
func (*Transcoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_403bea862d72b8de, []int{0}
}
func (m *Transcoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transcoder.Unmarshal(m, b)
}
func (m *Transcoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transcoder.Marshal(b, m, deterministic)
}
func (dst *Transcoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transcoder.Merge(dst, src)
}
func (m *Transcoder) XXX_Size() int {
	return xxx_messageInfo_Transcoder.Size(m)
}
func (m *Transcoder) XXX_DiscardUnknown() {
	xxx_messageInfo_Transcoder.DiscardUnknown(m)
}

var xxx_messageInfo_Transcoder proto.InternalMessageInfo

func (m *Transcoder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transcoder) GetCpuCores() int32 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *Transcoder) GetCpuMhz() float64 {
	if m != nil {
		return m.CpuMhz
	}
	return 0
}

func (m *Transcoder) GetTotalMemory() uint64 {
	if m != nil {
		return m.TotalMemory
	}
	return 0
}

func (m *Transcoder) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Transcoder) GetWorker() []byte {
	if m != nil {
		return m.Worker
	}
	return nil
}

type Assignment struct {
	Workorder            *v1.WorkOrder `protobuf:"bytes,1,opt,name=workorder,proto3" json:"workorder,omitempty"`
	Profile              *v11.Profile  `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_403bea862d72b8de, []int{1}
}
func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (dst *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(dst, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetWorkorder() *v1.WorkOrder {
	if m != nil {
		return m.Workorder
	}
	return nil
}

func (m *Assignment) GetProfile() *v11.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*Transcoder)(nil), "cloud.api.transcoder.v1.Transcoder")
	proto.RegisterType((*Assignment)(nil), "cloud.api.transcoder.v1.Assignment")
	proto.RegisterEnum("cloud.api.transcoder.v1.TranscoderStatus", TranscoderStatus_name, TranscoderStatus_value)
}

func init() {
	proto.RegisterFile("transcoder/v1/transcoder.proto", fileDescriptor_transcoder_403bea862d72b8de)
}

var fileDescriptor_transcoder_403bea862d72b8de = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xf5, 0xb8, 0x79, 0x90, 0x49, 0x85, 0xac, 0x51, 0x21, 0xae, 0x2b, 0x2c, 0x37, 0x62, 0x61,
	0x21, 0x62, 0x2b, 0x65, 0x85, 0x58, 0x25, 0x85, 0x15, 0xaa, 0x8a, 0x0c, 0x02, 0x89, 0x4d, 0xe5,
	0xd8, 0x93, 0x64, 0x54, 0xc7, 0xd7, 0x1a, 0x3f, 0x20, 0xfd, 0x02, 0x94, 0x2f, 0x60, 0x13, 0xb1,
	0x00, 0x7e, 0xa2, 0x2b, 0x96, 0xac, 0x10, 0x9f, 0x80, 0xc2, 0x8a, 0xbf, 0x40, 0x1e, 0x3b, 0x71,
	0x48, 0x17, 0xec, 0xee, 0xb9, 0xf7, 0x9c, 0xe3, 0x7b, 0x8f, 0x07, 0xeb, 0x09, 0x77, 0xc3, 0xd8,
	0x03, 0x9f, 0x72, 0x3b, 0xeb, 0xdb, 0x15, 0xb2, 0x22, 0x0e, 0x09, 0x90, 0x8e, 0x17, 0x40, 0xea,
	0x5b, 0x6e, 0xc4, 0xac, 0xad, 0x59, 0xd6, 0xd7, 0x7a, 0x13, 0x96, 0x4c, 0xd3, 0x91, 0xe5, 0xc1,
	0xcc, 0x9e, 0xc0, 0x04, 0x6c, 0xc1, 0x1f, 0xa5, 0x63, 0x81, 0x04, 0x10, 0x55, 0xe1, 0xa3, 0x3d,
	0xd9, 0xa2, 0xbf, 0x66, 0x3e, 0x85, 0x53, 0x60, 0xa1, 0x2d, 0xcc, 0x7b, 0x6e, 0xc4, 0xec, 0x77,
	0xc0, 0x2f, 0x81, 0x97, 0x5b, 0x6c, 0x40, 0x29, 0x7e, 0xfc, 0x1f, 0x71, 0xc4, 0x61, 0xcc, 0x02,
	0x1a, 0xe7, 0xda, 0x75, 0x5d, 0x4a, 0x07, 0x5b, 0x52, 0x1a, 0x66, 0x30, 0x8f, 0x38, 0xbc, 0x9f,
	0x17, 0xcb, 0x7a, 0xbd, 0x09, 0x0d, 0x7b, 0x99, 0x1b, 0x30, 0xdf, 0x4d, 0xa8, 0x7d, 0xa3, 0x28,
	0x2c, 0xba, 0x5f, 0x11, 0xc6, 0xaf, 0x36, 0xb7, 0x93, 0xdb, 0x58, 0x66, 0xbe, 0x8a, 0x0c, 0x64,
	0xb6, 0x1c, 0x99, 0xf9, 0xe4, 0x08, 0xb7, 0xbc, 0x28, 0xbd, 0xf0, 0x80, 0xd3, 0x58, 0x95, 0x0d,
	0x64, 0xd6, 0x9d, 0x5b, 0x5e, 0x94, 0x9e, 0xe6, 0x98, 0x74, 0x70, 0x33, 0x1f, 0xce, 0xa6, 0x57,
	0xea, 0x9e, 0x81, 0x4c, 0xe4, 0x34, 0xbc, 0x28, 0x3d, 0x9b, 0x5e, 0x91, 0x63, 0xbc, 0x9f, 0x40,
	0xe2, 0x06, 0x17, 0x33, 0x3a, 0x03, 0x3e, 0x57, 0x6b, 0x06, 0x32, 0x6b, 0x4e, 0x5b, 0xf4, 0xce,
	0x44, 0x8b, 0xdc, 0xc5, 0x8d, 0x38, 0x71, 0x93, 0x34, 0x56, 0x1b, 0xe2, 0x63, 0x25, 0xca, 0xfb,
	0x79, 0x40, 0x94, 0xab, 0x4d, 0x03, 0x99, 0xfb, 0x4e, 0x89, 0xba, 0x9f, 0x10, 0xc6, 0x83, 0x38,
	0x66, 0x93, 0x70, 0x46, 0xc3, 0x84, 0x3c, 0xc7, 0xad, 0x4d, 0x8e, 0x62, 0xdd, 0xf6, 0xc9, 0xb1,
	0x55, 0xfd, 0xcd, 0x2a, 0xe3, 0xac, 0x6f, 0xbd, 0x01, 0x7e, 0x79, 0x9e, 0x83, 0x21, 0xbe, 0xfe,
	0xf3, 0x6d, 0xaf, 0xbe, 0x40, 0xb2, 0x82, 0x9c, 0x4a, 0x4f, 0x9e, 0xe2, 0x66, 0x19, 0xac, 0x38,
	0xb1, 0x7d, 0xa2, 0x6f, 0x59, 0x6d, 0x22, 0xcf, 0xfa, 0xd6, 0x8b, 0xa2, 0xfe, 0xc7, 0x67, 0x2d,
	0x7d, 0xf0, 0x03, 0x61, 0xa5, 0x4a, 0xf2, 0x65, 0x71, 0xce, 0x43, 0xdc, 0x72, 0x33, 0x97, 0x05,
	0xee, 0x28, 0xa0, 0x8a, 0xa4, 0xdd, 0x5b, 0x2c, 0x8d, 0xc3, 0x5d, 0xd2, 0x60, 0x4d, 0x20, 0x26,
	0x6e, 0xc2, 0x78, 0x1c, 0xb0, 0x90, 0x2a, 0x48, 0x3b, 0x5a, 0x2c, 0x8d, 0xce, 0x2e, 0xf7, 0xbc,
	0x18, 0x93, 0xfb, 0xb8, 0x4e, 0x39, 0x07, 0xae, 0xc8, 0xda, 0xe1, 0x62, 0x69, 0xdc, 0xd9, 0xe5,
	0x3d, 0xcb, 0x87, 0xa4, 0x8b, 0x6b, 0xa3, 0x34, 0x9e, 0x2b, 0x7b, 0x9a, 0xba, 0x58, 0x1a, 0x07,
	0xbb, 0xa4, 0x61, 0x1a, 0xcf, 0x35, 0xf5, 0xc3, 0x67, 0x5d, 0xba, 0xfe, 0xa2, 0xdf, 0xd8, 0x7d,
	0x78, 0xf0, 0x7d, 0xa5, 0x4b, 0x3f, 0x57, 0xba, 0xf4, 0x6b, 0xa5, 0x4b, 0x1f, 0x7f, 0xeb, 0xd2,
	0x5b, 0x39, 0xeb, 0x8f, 0x1a, 0xe2, 0xdd, 0x3c, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x73,
	0xd2, 0x3b, 0x5c, 0x03, 0x00, 0x00,
}
