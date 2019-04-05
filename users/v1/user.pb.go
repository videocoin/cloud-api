// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users/v1/user.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "github.com/VideoCoin/cloud-api/accounts/v1"
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

type User struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"type:varchar(36);primary_key"`
	Email                string     `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" gorm:"unique_index"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" gorm:"type:varchar(100)"`
	CreatedAt            *time.Time `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	Token                string     `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty" gorm:"type:varchar(255);index"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45337a0acfc12558, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserProfile struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string             `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Account              *v1.AccountProfile `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45337a0acfc12558, []int{1}
}
func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfile.Unmarshal(m, b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
}
func (dst *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(dst, src)
}
func (m *UserProfile) XXX_Size() int {
	return xxx_messageInfo_UserProfile.Size(m)
}
func (m *UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfile proto.InternalMessageInfo

func (m *UserProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserProfile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserProfile) GetAccount() *v1.AccountProfile {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "cloud.api.users.v1.User")
	proto.RegisterType((*UserProfile)(nil), "cloud.api.users.v1.UserProfile")
}

func init() { proto.RegisterFile("users/v1/user.proto", fileDescriptor_user_45337a0acfc12558) }

var fileDescriptor_user_45337a0acfc12558 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0xce, 0xd2, 0x40,
	0x1c, 0xc5, 0xdb, 0x0a, 0xea, 0x37, 0x24, 0x2e, 0xe6, 0x23, 0xb1, 0x21, 0xa6, 0x25, 0xd5, 0x44,
	0x58, 0x30, 0xa5, 0x10, 0x94, 0x40, 0xa2, 0x01, 0x2f, 0x60, 0x1a, 0x75, 0xe1, 0x86, 0x0c, 0xed,
	0x50, 0x26, 0xb4, 0x9d, 0x3a, 0x9d, 0x56, 0xb9, 0x85, 0x4b, 0xf7, 0x5e, 0xc6, 0xa5, 0x27, 0xa8,
	0x06, 0x6f, 0xd0, 0x13, 0x18, 0xa6, 0xad, 0xf8, 0x25, 0xec, 0xde, 0xa4, 0xef, 0xfd, 0xdf, 0xfc,
	0x7f, 0x1d, 0x70, 0x9b, 0xa5, 0x84, 0xa7, 0x76, 0xee, 0xd8, 0x67, 0x81, 0x12, 0xce, 0x04, 0x83,
	0xd0, 0x0b, 0x59, 0xe6, 0x23, 0x9c, 0x50, 0x24, 0x3f, 0xa3, 0xdc, 0xe9, 0x8d, 0x02, 0x2a, 0xf6,
	0xd9, 0x16, 0x79, 0x2c, 0xb2, 0x03, 0x16, 0x30, 0x5b, 0x5a, 0xb7, 0xd9, 0x4e, 0x9e, 0xe4, 0x41,
	0xaa, 0x6a, 0x44, 0xcf, 0x0c, 0x18, 0x0b, 0x42, 0x72, 0x71, 0x09, 0x1a, 0x91, 0x54, 0xe0, 0x28,
	0xa9, 0x0d, 0xf3, 0xff, 0xe6, 0x7d, 0xa0, 0x3e, 0x61, 0x6f, 0x18, 0x8d, 0x6d, 0x59, 0x3c, 0xc2,
	0x09, 0xb5, 0xb1, 0xe7, 0xb1, 0x2c, 0x16, 0xf2, 0x6a, 0xb5, 0xae, 0x92, 0xd6, 0x77, 0x0d, 0xb4,
	0xde, 0xa7, 0x84, 0xc3, 0x97, 0x40, 0xa3, 0xbe, 0xae, 0xf6, 0xd5, 0xc1, 0xcd, 0xfa, 0x79, 0x59,
	0x98, 0x4f, 0x03, 0xc6, 0xa3, 0x85, 0x25, 0x8e, 0x09, 0x59, 0xe4, 0x98, 0x7b, 0x7b, 0xcc, 0x07,
	0xd3, 0x17, 0xc3, 0x65, 0xc2, 0x69, 0x84, 0xf9, 0x71, 0x73, 0x20, 0x47, 0xcb, 0xd5, 0xa8, 0x0f,
	0x47, 0xa0, 0x4d, 0x22, 0x4c, 0x43, 0x5d, 0x93, 0xd9, 0xc7, 0x65, 0x61, 0xde, 0x56, 0xd9, 0x2c,
	0xa6, 0x9f, 0x32, 0xb2, 0xa1, 0xb1, 0x4f, 0xbe, 0x58, 0x6e, 0xe5, 0x82, 0x73, 0xf0, 0x30, 0xc1,
	0x69, 0xfa, 0x99, 0x71, 0x5f, 0xbf, 0x27, 0x13, 0x4f, 0xca, 0xc2, 0xd4, 0xaf, 0xb4, 0x39, 0xe3,
	0xf1, 0xd0, 0x72, 0xff, 0xb9, 0xe1, 0x6b, 0x00, 0x3c, 0x4e, 0xb0, 0x20, 0xfe, 0x06, 0x0b, 0xbd,
	0xd5, 0x57, 0x07, 0x9d, 0x49, 0x0f, 0x55, 0x68, 0x50, 0x83, 0x06, 0xbd, 0x6b, 0xd0, 0xac, 0x5b,
	0x5f, 0x7f, 0x99, 0xaa, 0x7b, 0x53, 0x67, 0x56, 0x02, 0xce, 0x41, 0x5b, 0xb0, 0x03, 0x89, 0xf5,
	0xb6, 0xec, 0xb5, 0xca, 0xc2, 0x34, 0xae, 0xf4, 0x4e, 0x66, 0xb3, 0xe1, 0xb2, 0xb9, 0xb4, 0x0c,
	0x58, 0x29, 0xe8, 0x9c, 0x21, 0xbd, 0xe5, 0x6c, 0x47, 0x43, 0x02, 0x1f, 0x5d, 0x58, 0x49, 0x04,
	0xdd, 0x3b, 0x08, 0x9a, 0x4d, 0x5f, 0x81, 0x07, 0x35, 0x6b, 0xb9, 0x68, 0x67, 0xf2, 0x0c, 0x5d,
	0x9e, 0x42, 0xf3, 0x17, 0x72, 0x07, 0xad, 0x2a, 0x59, 0x0f, 0x77, 0x9b, 0xd0, 0xba, 0xfb, 0xe3,
	0x64, 0x28, 0x3f, 0x4f, 0x86, 0xf2, 0xfb, 0x64, 0x28, 0xdf, 0xfe, 0x18, 0xca, 0x47, 0x2d, 0x77,
	0xb6, 0xf7, 0xe5, 0xa6, 0xd3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x85, 0x0f, 0xa9, 0x6c,
	0x02, 0x00, 0x00,
}
