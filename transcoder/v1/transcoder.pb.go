// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transcoder/v1/transcoder.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import v11 "github.com/videocoin/cloud-api/profiles/v1"
import v1 "github.com/videocoin/cloud-api/workorder/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

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
	return fileDescriptor_transcoder_c0e17541c5c755ae, []int{0}
}

type Transcoder struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CpuCores             int32            `protobuf:"varint,2,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	CpuMhz               float64          `protobuf:"fixed64,3,opt,name=cpu_mhz,json=cpuMhz,proto3" json:"cpu_mhz,omitempty"`
	TotalMemory          uint64           `protobuf:"varint,4,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	Status               TranscoderStatus `protobuf:"varint,6,opt,name=status,proto3,enum=cloud.api.transcoder.v1.TranscoderStatus" json:"status,omitempty"`
	Worker               []byte           `protobuf:"bytes,7,opt,name=worker,proto3" json:"worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Transcoder) Reset()         { *m = Transcoder{} }
func (m *Transcoder) String() string { return proto.CompactTextString(m) }
func (*Transcoder) ProtoMessage()    {}
func (*Transcoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_c0e17541c5c755ae, []int{0}
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

func (m *Transcoder) GetStatus() TranscoderStatus {
	if m != nil {
		return m.Status
	}
	return TranscoderStatusAvailable
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
	return fileDescriptor_transcoder_c0e17541c5c755ae, []int{1}
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
	proto.RegisterFile("transcoder/v1/transcoder.proto", fileDescriptor_transcoder_c0e17541c5c755ae)
}

var fileDescriptor_transcoder_c0e17541c5c755ae = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdd, 0x8a, 0xd3, 0x40,
	0x1c, 0xc5, 0x33, 0xd9, 0x6e, 0x6a, 0xa7, 0xcb, 0x12, 0x06, 0xb5, 0xd9, 0x2c, 0x86, 0xd9, 0xe2,
	0x45, 0x14, 0x37, 0x21, 0xf5, 0x46, 0xf1, 0x42, 0xba, 0xe2, 0xe5, 0xb2, 0x12, 0x05, 0xc1, 0x9b,
	0x25, 0x1f, 0xd3, 0x76, 0xd8, 0x24, 0x13, 0x26, 0x99, 0x48, 0xf7, 0x05, 0x94, 0xbe, 0x43, 0xaf,
	0xf4, 0x29, 0x7c, 0x02, 0xaf, 0x64, 0x1f, 0x41, 0xfa, 0x24, 0x92, 0x69, 0xfa, 0x41, 0x44, 0xbc,
	0xfb, 0x9f, 0xf9, 0x9f, 0xdf, 0xe4, 0xe4, 0x30, 0xd0, 0x2a, 0x79, 0x90, 0x15, 0x11, 0x8b, 0x09,
	0x77, 0x2b, 0xcf, 0xdd, 0x29, 0x27, 0xe7, 0xac, 0x64, 0x68, 0x10, 0x25, 0x4c, 0xc4, 0x4e, 0x90,
	0x53, 0x67, 0x6f, 0x57, 0x79, 0xe6, 0xf9, 0x94, 0x96, 0x33, 0x11, 0x3a, 0x11, 0x4b, 0xdd, 0x29,
	0x9b, 0x32, 0x57, 0xfa, 0x43, 0x31, 0x91, 0x4a, 0x0a, 0x39, 0xad, 0xef, 0x31, 0x5f, 0xed, 0xd9,
	0x2b, 0x1a, 0x13, 0x16, 0x31, 0x9a, 0xb9, 0xf2, 0xf2, 0xf3, 0x20, 0xa7, 0xee, 0x67, 0xc6, 0x6f,
	0x18, 0x6f, 0x52, 0x6c, 0x45, 0x03, 0xbf, 0xfc, 0x0f, 0x9c, 0x73, 0x36, 0xa1, 0x09, 0x29, 0x6a,
	0x76, 0x33, 0xaf, 0xd1, 0xe1, 0x1d, 0x80, 0xf0, 0xc3, 0x36, 0x38, 0x3a, 0x86, 0x2a, 0x8d, 0x0d,
	0x80, 0x81, 0xdd, 0xf3, 0x55, 0x1a, 0xa3, 0x53, 0xd8, 0x8b, 0x72, 0x71, 0x1d, 0x31, 0x4e, 0x0a,
	0x43, 0xc5, 0xc0, 0x3e, 0xf4, 0xef, 0x45, 0xb9, 0x78, 0x53, 0x6b, 0x34, 0x80, 0xdd, 0x7a, 0x99,
	0xce, 0x6e, 0x8d, 0x03, 0x0c, 0x6c, 0xe0, 0x6b, 0x51, 0x2e, 0x2e, 0x67, 0xb7, 0xe8, 0x0c, 0x1e,
	0x95, 0xac, 0x0c, 0x92, 0xeb, 0x94, 0xa4, 0x8c, 0xcf, 0x8d, 0x0e, 0x06, 0x76, 0xc7, 0xef, 0xcb,
	0xb3, 0x4b, 0x79, 0x84, 0xc6, 0x50, 0x2b, 0xca, 0xa0, 0x14, 0x85, 0xa1, 0x61, 0x60, 0x1f, 0x8f,
	0x9e, 0x38, 0xff, 0x28, 0xd2, 0xd9, 0xa5, 0x7b, 0x2f, 0x01, 0xbf, 0x01, 0xd1, 0x43, 0xa8, 0xd5,
	0x45, 0x10, 0x6e, 0x74, 0x31, 0xb0, 0x8f, 0xfc, 0x46, 0x0d, 0xbf, 0x00, 0x08, 0xc7, 0x45, 0x41,
	0xa7, 0x59, 0x4a, 0xb2, 0x12, 0xbd, 0x86, 0xbd, 0x6d, 0x5f, 0xf2, 0xcf, 0xfa, 0xa3, 0xb3, 0xbd,
	0x8f, 0xed, 0xba, 0xac, 0x3c, 0xe7, 0x23, 0xe3, 0x37, 0x57, 0xb5, 0xf0, 0x77, 0x0c, 0x7a, 0x01,
	0xbb, 0x4d, 0x69, 0xb2, 0x81, 0xfe, 0xc8, 0xda, 0xc3, 0xb7, 0x75, 0x56, 0x9e, 0xf3, 0x6e, 0x3d,
	0xfb, 0x1b, 0xfb, 0xd3, 0x5f, 0x00, 0xea, 0xed, 0xf8, 0xe8, 0x19, 0xec, 0x05, 0x55, 0x40, 0x93,
	0x20, 0x4c, 0x88, 0xae, 0x98, 0x8f, 0x16, 0x4b, 0x7c, 0xd2, 0x36, 0x8d, 0x37, 0x06, 0x64, 0xc3,
	0x2e, 0x9b, 0x4c, 0x12, 0x9a, 0x11, 0x1d, 0x98, 0xa7, 0x8b, 0x25, 0x1e, 0xb4, 0xbd, 0x57, 0xeb,
	0x35, 0x7a, 0x0c, 0x0f, 0x09, 0xe7, 0x8c, 0xeb, 0xaa, 0x79, 0xb2, 0x58, 0xe2, 0x07, 0x6d, 0xdf,
	0xdb, 0x7a, 0x89, 0x86, 0xb0, 0x13, 0x8a, 0x62, 0xae, 0x1f, 0x98, 0xc6, 0x62, 0x89, 0xef, 0xb7,
	0x4d, 0x17, 0xa2, 0x98, 0x9b, 0xc6, 0xd7, 0x6f, 0x96, 0xf2, 0xe3, 0xbb, 0xf5, 0x57, 0xf6, 0x0b,
	0xfd, 0xe7, 0xca, 0x52, 0xee, 0x56, 0x96, 0xf2, 0x7b, 0x65, 0x29, 0x9f, 0xd4, 0xca, 0x0b, 0x35,
	0xf9, 0x8c, 0x9e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x54, 0xeb, 0xad, 0x28, 0x03, 0x00,
	0x00,
}
