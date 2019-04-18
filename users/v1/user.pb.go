// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users/v1/user.proto

package v1

import (
	fmt "fmt"
	v1 "github.com/VideoCoin/cloud-api/accounts/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
)

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
	Name                 string     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" gorm:"type:varchar(100)"`
	CreatedAt            *time.Time `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	Token                string     `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty" gorm:"type:varchar(255);index"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_622714e2df60ae10, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
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
	Name                 string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Account              *v1.AccountProfile `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_622714e2df60ae10, []int{1}
}
func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfile.Unmarshal(m, b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
}
func (m *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(m, src)
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

func (m *UserProfile) GetName() string {
	if m != nil {
		return m.Name
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

func init() { proto.RegisterFile("users/v1/user.proto", fileDescriptor_622714e2df60ae10) }

var fileDescriptor_622714e2df60ae10 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x63, 0x3b, 0x2e, 0xed, 0x45, 0xaa, 0xd0, 0xb5, 0x83, 0x95, 0xc1, 0x8e, 0x4c, 0x91,
	0x92, 0x21, 0xe7, 0x3a, 0x55, 0xab, 0xd2, 0x4a, 0x48, 0x0e, 0x13, 0x1b, 0xb2, 0x80, 0x81, 0x25,
	0x5c, 0xec, 0xab, 0x7b, 0xaa, 0xed, 0x33, 0xe7, 0xb3, 0xa9, 0x77, 0x24, 0x56, 0xc6, 0x7e, 0x0e,
	0x46, 0x26, 0x46, 0x46, 0x3e, 0x41, 0x40, 0x61, 0x62, 0xcd, 0x27, 0x40, 0x3d, 0xdb, 0xa4, 0x88,
	0x0c, 0xdd, 0xde, 0xc9, 0xef, 0xf7, 0xfc, 0xbf, 0xdf, 0x3d, 0xb0, 0x57, 0xe4, 0x84, 0xe7, 0x4e,
	0xe9, 0x3a, 0xb7, 0x05, 0xca, 0x38, 0x13, 0x0c, 0xc2, 0x20, 0x66, 0x45, 0x88, 0x70, 0x46, 0x91,
	0xfc, 0x8c, 0x4a, 0xb7, 0x3f, 0x8e, 0xa8, 0xb8, 0x2c, 0xe6, 0x28, 0x60, 0x89, 0x13, 0xb1, 0x88,
	0x39, 0xb2, 0x75, 0x5e, 0x5c, 0xc8, 0x93, 0x3c, 0xc8, 0xaa, 0x1e, 0xd1, 0xf7, 0xee, 0xb4, 0x93,
	0xb4, 0x64, 0x55, 0xc6, 0xd9, 0x75, 0x55, 0x43, 0xc1, 0x38, 0x22, 0xe9, 0xb8, 0xc4, 0x31, 0x0d,
	0xb1, 0x20, 0xce, 0x7f, 0x45, 0x33, 0xe2, 0xf4, 0xce, 0x88, 0xd7, 0x34, 0x24, 0xec, 0x19, 0xa3,
	0xa9, 0x23, 0xa3, 0x8d, 0x71, 0x46, 0x1d, 0x1c, 0x04, 0xac, 0x48, 0x85, 0x0c, 0xdf, 0xd4, 0x0d,
	0x69, 0x45, 0x8c, 0x45, 0x31, 0x59, 0x47, 0x14, 0x34, 0x21, 0xb9, 0xc0, 0x49, 0x56, 0x37, 0xd8,
	0x1f, 0x34, 0xd0, 0x7d, 0x95, 0x13, 0x0e, 0x3d, 0xa0, 0xd2, 0xd0, 0x50, 0x06, 0xca, 0x70, 0x67,
	0xea, 0xae, 0x16, 0xd6, 0xa3, 0x88, 0xf1, 0xe4, 0xcc, 0x16, 0x55, 0x46, 0xce, 0x4a, 0xcc, 0x83,
	0x4b, 0xcc, 0x87, 0x47, 0x27, 0xa3, 0xf3, 0x8c, 0xd3, 0x04, 0xf3, 0x6a, 0x76, 0x45, 0x2a, 0xfb,
	0xcb, 0xef, 0xaf, 0x9a, 0xce, 0xb5, 0x1b, 0xe5, 0xc0, 0x57, 0x69, 0x08, 0x4f, 0x80, 0x4e, 0x12,
	0x4c, 0x63, 0x43, 0x95, 0x53, 0x06, 0xab, 0x85, 0xb5, 0x57, 0x4f, 0x29, 0x52, 0xfa, 0xae, 0x20,
	0x33, 0x9a, 0x86, 0xe4, 0x5a, 0x52, 0x5d, 0xae, 0xbe, 0x55, 0xfc, 0xba, 0x1d, 0x7a, 0x60, 0x3b,
	0xc3, 0x79, 0xfe, 0x9e, 0xf1, 0xd0, 0xd0, 0x24, 0xfa, 0x78, 0xb5, 0xb0, 0x8c, 0x0d, 0x01, 0xdc,
	0xc3, 0xc3, 0x51, 0xcb, 0x3f, 0xd4, 0xfd, 0xbf, 0x18, 0x7c, 0x02, 0xba, 0x29, 0x4e, 0x88, 0xd1,
	0xbd, 0x3f, 0xae, 0xf9, 0x12, 0x81, 0xcf, 0x01, 0x08, 0x38, 0xc1, 0x82, 0x84, 0x33, 0x2c, 0x0c,
	0x7d, 0xa0, 0x0c, 0x7b, 0x93, 0x3e, 0xaa, 0xbd, 0xa1, 0xd6, 0x1b, 0x7a, 0xd9, 0x7a, 0x9b, 0xee,
	0x7e, 0xfa, 0x61, 0x29, 0xf2, 0xe6, 0x9f, 0x15, 0x75, 0x5b, 0xf1, 0x77, 0x1a, 0xda, 0x13, 0xf0,
	0x14, 0xe8, 0x82, 0x5d, 0x91, 0xd4, 0xd8, 0x92, 0x31, 0xec, 0xd5, 0xc2, 0x32, 0x37, 0xc4, 0x98,
	0x1c, 0x1f, 0x8f, 0xce, 0x6b, 0x17, 0x7e, 0x0d, 0xd8, 0x1f, 0x15, 0xd0, 0xbb, 0x7d, 0x86, 0x17,
	0x9c, 0x5d, 0xd0, 0x98, 0xc0, 0xdd, 0xf5, 0x6b, 0x48, 0xb5, 0xfb, 0xff, 0xa8, 0x6d, 0xc5, 0xc1,
	0xe6, 0xd6, 0x52, 0x5a, 0x73, 0x9d, 0xa7, 0xe0, 0x41, 0xb3, 0x02, 0x52, 0x46, 0x6f, 0x72, 0x80,
	0xd6, 0x3b, 0xdc, 0x2e, 0x47, 0xe9, 0x22, 0xaf, 0x2e, 0x9b, 0x1f, 0xfa, 0x2d, 0x34, 0xdd, 0xff,
	0xb6, 0x34, 0x3b, 0xdf, 0x97, 0x66, 0xe7, 0xe7, 0xd2, 0xec, 0xdc, 0xfc, 0x32, 0x3b, 0x6f, 0xd4,
	0xd2, 0x9d, 0x6f, 0x49, 0x11, 0x47, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x82, 0x83, 0x4d,
	0x25, 0x03, 0x00, 0x00,
}
