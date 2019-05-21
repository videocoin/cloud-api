// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transcoder/v1/transcoder.proto

package v1

import proto "github.com/gogo/protobuf/proto"
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
	return fileDescriptor_transcoder_ccab4349a41f64b0, []int{0}
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
	return fileDescriptor_transcoder_ccab4349a41f64b0, []int{0}
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
	return fileDescriptor_transcoder_ccab4349a41f64b0, []int{1}
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
	proto.RegisterFile("transcoder/v1/transcoder.proto", fileDescriptor_transcoder_ccab4349a41f64b0)
}

var fileDescriptor_transcoder_ccab4349a41f64b0 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xcd, 0x38, 0x2f, 0x32, 0xa9, 0x90, 0x35, 0x02, 0xe2, 0xba, 0xc2, 0x72, 0x23, 0x16, 0x16,
	0x22, 0xb6, 0x52, 0x56, 0x88, 0x55, 0x52, 0x58, 0xa1, 0xaa, 0xc8, 0x20, 0x90, 0xd8, 0x54, 0x13,
	0x7b, 0x92, 0x8c, 0xea, 0xf8, 0x5a, 0xe3, 0x07, 0xa4, 0x5f, 0x80, 0xf2, 0x0f, 0x11, 0x0b, 0xe0,
	0x27, 0xba, 0x62, 0xc9, 0x0a, 0xf1, 0x09, 0x28, 0x3b, 0xfe, 0x02, 0x79, 0xec, 0x3c, 0x48, 0x17,
	0xdd, 0xdd, 0x73, 0xef, 0x39, 0xc7, 0xf7, 0x1e, 0x0f, 0x36, 0x12, 0x41, 0xc3, 0xd8, 0x03, 0x9f,
	0x09, 0x27, 0xeb, 0x3b, 0x5b, 0x64, 0x47, 0x02, 0x12, 0x20, 0x1d, 0x2f, 0x80, 0xd4, 0xb7, 0x69,
	0xc4, 0xed, 0x9d, 0x59, 0xd6, 0xd7, 0x7b, 0x13, 0x9e, 0x4c, 0xd3, 0x91, 0xed, 0xc1, 0xcc, 0x99,
	0xc0, 0x04, 0x1c, 0xc9, 0x1f, 0xa5, 0x63, 0x89, 0x24, 0x90, 0x55, 0xe1, 0xa3, 0x3f, 0xdf, 0xa1,
	0xbf, 0xe3, 0x3e, 0x83, 0x53, 0xe0, 0xa1, 0x23, 0xcd, 0x7b, 0x34, 0xe2, 0xce, 0x47, 0x10, 0x97,
	0x20, 0xca, 0x2d, 0x36, 0xa0, 0x14, 0x3f, 0xbb, 0x45, 0x1c, 0x09, 0x18, 0xf3, 0x80, 0xc5, 0xb9,
	0x76, 0x5d, 0x97, 0xd2, 0xc1, 0x8e, 0x94, 0x85, 0x19, 0xcc, 0x23, 0x01, 0x9f, 0xe6, 0xc5, 0xb2,
	0x5e, 0x6f, 0xc2, 0xc2, 0x5e, 0x46, 0x03, 0xee, 0xd3, 0x84, 0x39, 0x37, 0x8a, 0xc2, 0xa2, 0xfb,
	0x1d, 0x61, 0xfc, 0x76, 0x73, 0x3b, 0xb9, 0x8b, 0x15, 0xee, 0x6b, 0xc8, 0x44, 0x56, 0xcb, 0x55,
	0xb8, 0x4f, 0x8e, 0x70, 0xcb, 0x8b, 0xd2, 0x0b, 0x0f, 0x04, 0x8b, 0x35, 0xc5, 0x44, 0x56, 0xdd,
	0xbd, 0xe3, 0x45, 0xe9, 0x69, 0x8e, 0x49, 0x07, 0x37, 0xf3, 0xe1, 0x6c, 0x7a, 0xa5, 0x55, 0x4d,
	0x64, 0x21, 0xb7, 0xe1, 0x45, 0xe9, 0xd9, 0xf4, 0x8a, 0x1c, 0xe3, 0x83, 0x04, 0x12, 0x1a, 0x5c,
	0xcc, 0xd8, 0x0c, 0xc4, 0x5c, 0xab, 0x99, 0xc8, 0xaa, 0xb9, 0x6d, 0xd9, 0x3b, 0x93, 0x2d, 0xf2,
	0x00, 0x37, 0xe2, 0x84, 0x26, 0x69, 0xac, 0x35, 0xe4, 0xc7, 0x4a, 0x94, 0xf7, 0xf3, 0x80, 0x98,
	0xd0, 0x9a, 0x26, 0xb2, 0x0e, 0xdc, 0x12, 0x75, 0xbf, 0x20, 0x8c, 0x07, 0x71, 0xcc, 0x27, 0xe1,
	0x8c, 0x85, 0x09, 0x79, 0x85, 0x5b, 0x9b, 0x1c, 0xe5, 0xba, 0xed, 0x93, 0x63, 0x7b, 0xfb, 0x37,
	0xb7, 0x19, 0x67, 0x7d, 0xfb, 0x3d, 0x88, 0xcb, 0xf3, 0x1c, 0x0c, 0xf1, 0xf5, 0xdf, 0x1f, 0xd5,
	0xfa, 0x02, 0x29, 0x2a, 0x72, 0xb7, 0x7a, 0xf2, 0x02, 0x37, 0xcb, 0x60, 0xe5, 0x89, 0xed, 0x13,
	0x63, 0xc7, 0x6a, 0x13, 0x79, 0xd6, 0xb7, 0x5f, 0x17, 0xf5, 0x7f, 0x3e, 0x6b, 0xe9, 0xe3, 0x5f,
	0x08, 0xab, 0xdb, 0x24, 0xdf, 0x14, 0xe7, 0x3c, 0xc1, 0x2d, 0x9a, 0x51, 0x1e, 0xd0, 0x51, 0xc0,
	0xd4, 0x8a, 0xfe, 0x70, 0xb1, 0x34, 0x0f, 0xf7, 0x49, 0x83, 0x35, 0x81, 0x58, 0xb8, 0x09, 0xe3,
	0x71, 0xc0, 0x43, 0xa6, 0x22, 0xfd, 0x68, 0xb1, 0x34, 0x3b, 0xfb, 0xdc, 0xf3, 0x62, 0x4c, 0x1e,
	0xe1, 0x3a, 0x13, 0x02, 0x84, 0xaa, 0xe8, 0x87, 0x8b, 0xa5, 0x79, 0x7f, 0x9f, 0xf7, 0x32, 0x1f,
	0x92, 0x2e, 0xae, 0x8d, 0xd2, 0x78, 0xae, 0x56, 0x75, 0x6d, 0xb1, 0x34, 0xef, 0xed, 0x93, 0x86,
	0x69, 0x3c, 0xd7, 0xb5, 0xcf, 0x5f, 0x8d, 0xca, 0xf5, 0x37, 0xe3, 0xc6, 0xee, 0x43, 0xf5, 0xe7,
	0xca, 0xa8, 0xfc, 0x5e, 0x19, 0x95, 0x3f, 0x2b, 0xa3, 0xf2, 0x41, 0xc9, 0xfa, 0xa3, 0x86, 0x7c,
	0x33, 0x4f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x07, 0x77, 0xaa, 0x23, 0x58, 0x03, 0x00, 0x00,
}
