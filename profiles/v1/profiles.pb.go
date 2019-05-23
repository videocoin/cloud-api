// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profiles/v1/profiles.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
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

type ProfileId int32

const (
	ProfileIdNone ProfileId = 0
	ProfileIdSD   ProfileId = 1
	ProfileIdHD   ProfileId = 2
	ProfileIdFHD  ProfileId = 3
)

var ProfileId_name = map[int32]string{
	0: "profile_id_none",
	1: "profile_id_sd",
	2: "profile_id_hd",
	3: "profile_id_fhd",
}
var ProfileId_value = map[string]int32{
	"profile_id_none": 0,
	"profile_id_sd":   1,
	"profile_id_hd":   2,
	"profile_id_fhd":  3,
}

func (x ProfileId) String() string {
	return proto.EnumName(ProfileId_name, int32(x))
}
func (ProfileId) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_profiles_6d9a90093cd7000a, []int{0}
}

type Profiles struct {
	Profiles             []*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Profiles) Reset()         { *m = Profiles{} }
func (m *Profiles) String() string { return proto.CompactTextString(m) }
func (*Profiles) ProtoMessage()    {}
func (*Profiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_profiles_6d9a90093cd7000a, []int{0}
}
func (m *Profiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profiles.Unmarshal(m, b)
}
func (m *Profiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profiles.Marshal(b, m, deterministic)
}
func (dst *Profiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profiles.Merge(dst, src)
}
func (m *Profiles) XXX_Size() int {
	return xxx_messageInfo_Profiles.Size(m)
}
func (m *Profiles) XXX_DiscardUnknown() {
	xxx_messageInfo_Profiles.DiscardUnknown(m)
}

var xxx_messageInfo_Profiles proto.InternalMessageInfo

func (m *Profiles) GetProfiles() []*Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

type Profile struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bitrate              uint32   `protobuf:"varint,3,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	Width                uint32   `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Fps                  float32  `protobuf:"fixed32,5,opt,name=fps,proto3" json:"fps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_profiles_6d9a90093cd7000a, []int{1}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (dst *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(dst, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

func (m *Profile) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Profile) GetFps() float32 {
	if m != nil {
		return m.Fps
	}
	return 0
}

func init() {
	proto.RegisterType((*Profiles)(nil), "cloud.api.profiles.v1.Profiles")
	proto.RegisterType((*Profile)(nil), "cloud.api.profiles.v1.Profile")
	proto.RegisterEnum("cloud.api.profiles.v1.ProfileId", ProfileId_name, ProfileId_value)
}

func init() {
	proto.RegisterFile("profiles/v1/profiles.proto", fileDescriptor_profiles_6d9a90093cd7000a)
}

var fileDescriptor_profiles_6d9a90093cd7000a = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4e, 0xc2, 0x40,
	0x14, 0x87, 0x3b, 0x2d, 0x08, 0x3c, 0x04, 0x86, 0x89, 0x26, 0x93, 0x2e, 0x9a, 0x09, 0x31, 0xa6,
	0x31, 0xb1, 0x0d, 0xba, 0x73, 0x69, 0x08, 0xc1, 0x8d, 0x31, 0x75, 0xe7, 0x86, 0x50, 0xa6, 0xa5,
	0x93, 0x40, 0xa7, 0x81, 0x82, 0x57, 0x30, 0xdc, 0x81, 0x95, 0x1e, 0xc1, 0x95, 0x27, 0x70, 0xe9,
	0x11, 0x0c, 0x27, 0x31, 0x0e, 0x6d, 0x25, 0xc4, 0xdd, 0xef, 0x7b, 0xf3, 0xbd, 0x3f, 0xc9, 0x80,
	0x99, 0xcc, 0x65, 0x28, 0xa6, 0xc1, 0xc2, 0x5d, 0x75, 0xdd, 0x3c, 0x3b, 0xc9, 0x5c, 0xa6, 0x92,
	0x9c, 0x8e, 0xa7, 0x72, 0xc9, 0x9d, 0x51, 0x22, 0x9c, 0xe2, 0x65, 0xd5, 0x35, 0x2f, 0x27, 0x22,
	0x8d, 0x96, 0xbe, 0x33, 0x96, 0x33, 0x77, 0x22, 0x27, 0xd2, 0x55, 0xb6, 0xbf, 0x0c, 0x15, 0x29,
	0x50, 0x69, 0x37, 0xa5, 0xd3, 0x87, 0xea, 0x43, 0xd6, 0x4d, 0x6e, 0xa0, 0x9a, 0x4f, 0xa2, 0x88,
	0x19, 0x76, 0xfd, 0xca, 0x72, 0xfe, 0x5d, 0xe2, 0x64, 0x2d, 0x5e, 0xe1, 0x77, 0x24, 0x54, 0xb2,
	0x22, 0x69, 0x82, 0x2e, 0x38, 0x45, 0x0c, 0xd9, 0x65, 0x4f, 0x17, 0x9c, 0x10, 0x28, 0xc5, 0xa3,
	0x59, 0x40, 0x75, 0x86, 0xec, 0x9a, 0xa7, 0x32, 0xa1, 0x50, 0xf1, 0x45, 0x3a, 0x1f, 0xa5, 0x01,
	0x35, 0x18, 0xb2, 0x1b, 0x5e, 0x8e, 0xe4, 0x04, 0xca, 0xcf, 0x82, 0xa7, 0x11, 0x2d, 0xa9, 0xfa,
	0x0e, 0x08, 0x06, 0x23, 0x4c, 0x16, 0xb4, 0xcc, 0x90, 0xad, 0x7b, 0xbf, 0xf1, 0xe2, 0x1d, 0x41,
	0x2d, 0xdb, 0x78, 0xc7, 0xc9, 0x39, 0xb4, 0xb2, 0x53, 0x86, 0x82, 0x0f, 0x63, 0x19, 0x07, 0x58,
	0x33, 0xdb, 0xeb, 0x0d, 0x6b, 0x14, 0xce, 0xbd, 0x8c, 0x03, 0xd2, 0x81, 0xc6, 0x9e, 0xb7, 0xe0,
	0x18, 0x99, 0xad, 0xf5, 0x86, 0xd5, 0x0b, 0xeb, 0xb1, 0x77, 0xe0, 0x44, 0x1c, 0xeb, 0x07, 0xce,
	0xa0, 0x47, 0xce, 0xa0, 0xb9, 0xe7, 0x84, 0x11, 0xc7, 0x86, 0x89, 0xd7, 0x1b, 0x76, 0x5c, 0x48,
	0xfd, 0x41, 0xcf, 0x6c, 0xbf, 0xbc, 0x5a, 0xda, 0xc7, 0x9b, 0xf5, 0x77, 0xe8, 0x2d, 0xfe, 0xdc,
	0x5a, 0xda, 0xd7, 0xd6, 0xd2, 0xbe, 0xb7, 0x96, 0xf6, 0xa4, 0xaf, 0xba, 0xfe, 0x91, 0xfa, 0x88,
	0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x18, 0xf5, 0x86, 0xec, 0x01, 0x00, 0x00,
}
