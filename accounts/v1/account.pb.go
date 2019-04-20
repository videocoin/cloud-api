// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accounts/v1/account.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type Account struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"type:varchar(36);primary_key"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" gorm:"type:varchar(42)"`
	Key                  string   `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty" gorm:"type:varchar(512)"`
	Balance              float64  `protobuf:"fixed64,5,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_a21cc970c4f82246, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Account) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Account) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Account) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type AccountProfile struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Balance              float64  `protobuf:"fixed64,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountProfile) Reset()         { *m = AccountProfile{} }
func (m *AccountProfile) String() string { return proto.CompactTextString(m) }
func (*AccountProfile) ProtoMessage()    {}
func (*AccountProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_a21cc970c4f82246, []int{1}
}
func (m *AccountProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountProfile.Unmarshal(m, b)
}
func (m *AccountProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountProfile.Marshal(b, m, deterministic)
}
func (dst *AccountProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountProfile.Merge(dst, src)
}
func (m *AccountProfile) XXX_Size() int {
	return xxx_messageInfo_AccountProfile.Size(m)
}
func (m *AccountProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountProfile.DiscardUnknown(m)
}

var xxx_messageInfo_AccountProfile proto.InternalMessageInfo

func (m *AccountProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccountProfile) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountProfile) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type AccountKey struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountKey) Reset()         { *m = AccountKey{} }
func (m *AccountKey) String() string { return proto.CompactTextString(m) }
func (*AccountKey) ProtoMessage()    {}
func (*AccountKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_a21cc970c4f82246, []int{2}
}
func (m *AccountKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountKey.Unmarshal(m, b)
}
func (m *AccountKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountKey.Marshal(b, m, deterministic)
}
func (dst *AccountKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountKey.Merge(dst, src)
}
func (m *AccountKey) XXX_Size() int {
	return xxx_messageInfo_AccountKey.Size(m)
}
func (m *AccountKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountKey.DiscardUnknown(m)
}

var xxx_messageInfo_AccountKey proto.InternalMessageInfo

func (m *AccountKey) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "cloud.api.account.v1.Account")
	proto.RegisterType((*AccountProfile)(nil), "cloud.api.account.v1.AccountProfile")
	proto.RegisterType((*AccountKey)(nil), "cloud.api.account.v1.AccountKey")
}

func init() { proto.RegisterFile("accounts/v1/account.proto", fileDescriptor_account_a21cc970c4f82246) }

var fileDescriptor_account_a21cc970c4f82246 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x8b, 0xd3, 0x40,
	0x1c, 0xc6, 0xf3, 0xb2, 0xdb, 0xb0, 0x23, 0x2c, 0x65, 0x58, 0x70, 0x5c, 0x24, 0x5d, 0xb2, 0x22,
	0xc9, 0x42, 0x13, 0xb2, 0xab, 0x22, 0xf5, 0x62, 0x7b, 0x13, 0x2f, 0x12, 0x3c, 0x79, 0x29, 0x93,
	0xcc, 0x34, 0x1d, 0x9a, 0x66, 0xc2, 0xe4, 0x05, 0xf3, 0x29, 0xbc, 0xf6, 0xb3, 0x78, 0xf2, 0xe8,
	0xd1, 0x4f, 0x50, 0xa4, 0x9e, 0xbc, 0xf6, 0x13, 0x48, 0x26, 0x29, 0xa9, 0x6c, 0x73, 0x7a, 0x86,
	0x79, 0x9e, 0xdf, 0xff, 0xff, 0x4c, 0xc0, 0x33, 0x1c, 0x45, 0xbc, 0x4c, 0x8b, 0xdc, 0xab, 0x7c,
	0xaf, 0xd3, 0x6e, 0x26, 0x78, 0xc1, 0xe1, 0x55, 0x94, 0xf0, 0x92, 0xb8, 0x38, 0x63, 0xee, 0xe1,
	0xa2, 0xf2, 0xaf, 0xc7, 0x31, 0x2b, 0x96, 0x65, 0xe8, 0x46, 0x7c, 0xed, 0xc5, 0x3c, 0xe6, 0x9e,
	0x34, 0x87, 0xe5, 0x42, 0x9e, 0xe4, 0x41, 0xaa, 0x16, 0x72, 0x3d, 0x3d, 0xb2, 0xd3, 0xb4, 0xe2,
	0x75, 0x26, 0xf8, 0xd7, 0xba, 0x0d, 0x45, 0xe3, 0x98, 0xa6, 0xe3, 0x0a, 0x27, 0x8c, 0xe0, 0x82,
	0x7a, 0x8f, 0x44, 0x8b, 0xb0, 0xbe, 0x69, 0xc0, 0x98, 0xb6, 0x0b, 0xc0, 0x29, 0xd0, 0x18, 0x41,
	0xea, 0x8d, 0x6a, 0x5f, 0xcc, 0xfc, 0xfd, 0x76, 0x74, 0x1b, 0x73, 0xb1, 0x9e, 0x58, 0x45, 0x9d,
	0xd1, 0x49, 0x85, 0x45, 0xb4, 0xc4, 0xc2, 0x7e, 0x78, 0xe3, 0xbc, 0xcb, 0x04, 0x5b, 0x63, 0x51,
	0xcf, 0x57, 0xb4, 0xb6, 0xbe, 0xff, 0xfd, 0xa1, 0x9f, 0x0b, 0x7d, 0xa3, 0xbe, 0x08, 0x34, 0x46,
	0xe0, 0x2d, 0x30, 0xca, 0x9c, 0x8a, 0x39, 0x23, 0x48, 0x93, 0x1c, 0x70, 0x64, 0x18, 0x34, 0x57,
	0x1f, 0x08, 0x7c, 0x0f, 0x0c, 0x4c, 0x88, 0xa0, 0x79, 0x8e, 0x74, 0x69, 0x7a, 0xb9, 0xdf, 0x8e,
	0x9e, 0x9e, 0x18, 0xf6, 0xea, 0xde, 0x91, 0x03, 0x06, 0xe2, 0x6c, 0x68, 0xa3, 0xbb, 0xe0, 0x10,
	0x83, 0x2e, 0xd0, 0x57, 0xb4, 0x46, 0x67, 0x32, 0xfd, 0x7c, 0xbf, 0x1d, 0xa1, 0x13, 0xe9, 0xd7,
	0xfe, 0xbd, 0x63, 0x05, 0x8d, 0x11, 0xde, 0x01, 0x23, 0xc4, 0x09, 0x4e, 0x23, 0x8a, 0xce, 0x6f,
	0x54, 0x5b, 0x9d, 0x0d, 0x1b, 0xec, 0x13, 0x78, 0xe1, 0x28, 0xdd, 0x17, 0x1c, 0x0c, 0xd6, 0x67,
	0x70, 0xd9, 0x3d, 0xc8, 0x27, 0xc1, 0x17, 0x2c, 0xa1, 0xf0, 0xb2, 0x7f, 0x17, 0x59, 0x12, 0xf5,
	0xfb, 0xcb, 0x92, 0xfd, 0x5e, 0xa8, 0x9f, 0xd3, 0x34, 0x53, 0x7b, 0xea, 0x5b, 0x00, 0x3a, 0xea,
	0x47, 0x5a, 0x1f, 0x13, 0xd4, 0xff, 0x09, 0xc3, 0xb6, 0x59, 0xcb, 0x6d, 0xe4, 0xec, 0xea, 0xe7,
	0xce, 0x54, 0x7e, 0xed, 0x4c, 0xe5, 0xf7, 0xce, 0x54, 0x36, 0x7f, 0x4c, 0xe5, 0x8b, 0x56, 0xf9,
	0xe1, 0x40, 0xfe, 0xbe, 0x87, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0x7c, 0xb6, 0x8b, 0x63,
	0x02, 0x00, 0x00,
}
