// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users/v1/user.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/videocoin/cloud-api/accounts/v1"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TokenType int32

const (
	TokenTypeRegular TokenType = 0
	TokenTypeAPI     TokenType = 1
)

var TokenType_name = map[int32]string{
	0: "TOKEN_TYPE_REGULAR",
	1: "TOKEN_TYPE_API",
}

var TokenType_value = map[string]int32{
	"TOKEN_TYPE_REGULAR": 0,
	"TOKEN_TYPE_API":     1,
}

func (x TokenType) String() string {
	return proto.EnumName(TokenType_name, int32(x))
}

func (TokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_622714e2df60ae10, []int{0}
}

type UserRole int32

const (
	UserRoleRegular UserRole = 0
	UserRoleMiner   UserRole = 1
	UserRoleQa      UserRole = 3
	UserRoleManager UserRole = 6
	UserRoleSuper   UserRole = 9
)

var UserRole_name = map[int32]string{
	0: "USER_ROLE_REGULAR",
	1: "USER_ROLE_MINER",
	3: "USER_ROLE_QA",
	6: "USER_ROLE_MANAGER",
	9: "USER_ROLE_SUPER",
}

var UserRole_value = map[string]int32{
	"USER_ROLE_REGULAR": 0,
	"USER_ROLE_MINER":   1,
	"USER_ROLE_QA":      3,
	"USER_ROLE_MANAGER": 6,
	"USER_ROLE_SUPER":   9,
}

func (x UserRole) String() string {
	return proto.EnumName(UserRole_name, int32(x))
}

func (UserRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_622714e2df60ae10, []int{1}
}

type UserUIRole int32

const (
	UserUIRoleBoth      UserUIRole = 0
	UserUIRoleMiner     UserUIRole = 1
	UserUIRolePublisher UserUIRole = 3
)

var UserUIRole_name = map[int32]string{
	0: "USER_ROLE_UI_BOTH",
	1: "USER_ROLE_UI_MINER",
	3: "USER_ROLE_UI_PUBLISHER",
}

var UserUIRole_value = map[string]int32{
	"USER_ROLE_UI_BOTH":      0,
	"USER_ROLE_UI_MINER":     1,
	"USER_ROLE_UI_PUBLISHER": 3,
}

func (x UserUIRole) String() string {
	return proto.EnumName(UserUIRole_name, int32(x))
}

func (UserUIRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_622714e2df60ae10, []int{2}
}

type UserProfile struct {
	ID                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string     `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName            string     `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string     `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	IsActive             bool       `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Role                 UserRole   `protobuf:"varint,6,opt,name=role,proto3,enum=cloud.api.users.v1.UserRole" json:"role,omitempty"`
	UiRole               UserUIRole `protobuf:"varint,7,opt,name=ui_role,json=uiRole,proto3,enum=cloud.api.users.v1.UserUIRole" json:"ui_role,omitempty"`
	Country              string     `protobuf:"bytes,8,opt,name=country,proto3" json:"country,omitempty"`
	Region               string     `protobuf:"bytes,9,opt,name=region,proto3" json:"region,omitempty"`
	City                 string     `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	Zip                  string     `protobuf:"bytes,11,opt,name=zip,proto3" json:"zip,omitempty"`
	Address_1            string     `protobuf:"bytes,12,opt,name=address_1,json=address1,proto3" json:"address_1,omitempty"`
	Address_2            string     `protobuf:"bytes,13,opt,name=address_2,json=address2,proto3" json:"address_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_622714e2df60ae10, []int{0}
}
func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(m, src)
}
func (m *UserProfile) XXX_Size() int {
	return m.Size()
}
func (m *UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfile proto.InternalMessageInfo

func (m *UserProfile) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UserProfile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserProfile) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserProfile) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserProfile) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *UserProfile) GetRole() UserRole {
	if m != nil {
		return m.Role
	}
	return UserRoleRegular
}

func (m *UserProfile) GetUiRole() UserUIRole {
	if m != nil {
		return m.UiRole
	}
	return UserUIRoleBoth
}

func (m *UserProfile) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *UserProfile) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UserProfile) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserProfile) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *UserProfile) GetAddress_1() string {
	if m != nil {
		return m.Address_1
	}
	return ""
}

func (m *UserProfile) GetAddress_2() string {
	if m != nil {
		return m.Address_2
	}
	return ""
}

func (*UserProfile) XXX_MessageName() string {
	return "cloud.api.users.v1.UserProfile"
}
func init() {
	proto.RegisterEnum("cloud.api.users.v1.TokenType", TokenType_name, TokenType_value)
	golang_proto.RegisterEnum("cloud.api.users.v1.TokenType", TokenType_name, TokenType_value)
	proto.RegisterEnum("cloud.api.users.v1.UserRole", UserRole_name, UserRole_value)
	golang_proto.RegisterEnum("cloud.api.users.v1.UserRole", UserRole_name, UserRole_value)
	proto.RegisterEnum("cloud.api.users.v1.UserUIRole", UserUIRole_name, UserUIRole_value)
	golang_proto.RegisterEnum("cloud.api.users.v1.UserUIRole", UserUIRole_name, UserUIRole_value)
	proto.RegisterType((*UserProfile)(nil), "cloud.api.users.v1.UserProfile")
	golang_proto.RegisterType((*UserProfile)(nil), "cloud.api.users.v1.UserProfile")
}

func init() { proto.RegisterFile("users/v1/user.proto", fileDescriptor_622714e2df60ae10) }
func init() { golang_proto.RegisterFile("users/v1/user.proto", fileDescriptor_622714e2df60ae10) }

var fileDescriptor_622714e2df60ae10 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0xda, 0x4a,
	0x14, 0xc7, 0x63, 0x48, 0x08, 0x4c, 0x12, 0xe2, 0x4c, 0xa2, 0x5c, 0xcb, 0x37, 0xd7, 0xb1, 0xae,
	0xaa, 0x2a, 0xa5, 0x01, 0x97, 0x64, 0xd1, 0x6e, 0x41, 0xb5, 0x12, 0xd4, 0x40, 0xc8, 0x00, 0x8b,
	0x76, 0x63, 0x0d, 0x66, 0xe2, 0x8c, 0x6a, 0x3c, 0xd6, 0xd8, 0xa6, 0x4a, 0x9f, 0xa0, 0x62, 0xd7,
	0x07, 0x60, 0xd5, 0x6e, 0xfb, 0x02, 0x5d, 0x75, 0x99, 0x65, 0x9f, 0xa0, 0xaa, 0x88, 0xd4, 0xe7,
	0xa8, 0x3c, 0x98, 0xaf, 0x7e, 0xac, 0x7c, 0xce, 0xf9, 0xff, 0xce, 0xfc, 0xcf, 0x99, 0x01, 0xb0,
	0x1b, 0x05, 0x84, 0x07, 0xc6, 0xa0, 0x6c, 0xc4, 0x41, 0xc9, 0xe7, 0x2c, 0x64, 0x10, 0xda, 0x2e,
	0x8b, 0x7a, 0x25, 0xec, 0xd3, 0x92, 0x90, 0x4b, 0x83, 0xb2, 0xfa, 0xcc, 0xa1, 0xe1, 0x4d, 0xd4,
	0x2d, 0xd9, 0xac, 0x6f, 0x0c, 0x68, 0x8f, 0x30, 0x9b, 0x51, 0xcf, 0x10, 0x60, 0x11, 0xfb, 0xd4,
	0xc0, 0xb6, 0xcd, 0x22, 0x2f, 0x14, 0x47, 0x25, 0xf1, 0xe4, 0x34, 0xf5, 0xd0, 0x61, 0xcc, 0x71,
	0x89, 0x21, 0xb2, 0x6e, 0x74, 0x6d, 0x84, 0xb4, 0x4f, 0x82, 0x10, 0xf7, 0xfd, 0x04, 0x38, 0x48,
	0x00, 0x71, 0x8c, 0xe7, 0xb1, 0x10, 0x87, 0x94, 0x79, 0x41, 0xa2, 0x1e, 0x8b, 0x8f, 0x5d, 0x74,
	0x88, 0x57, 0x0c, 0xde, 0x60, 0xc7, 0x21, 0xdc, 0x60, 0xbe, 0x20, 0xfe, 0x40, 0x17, 0x17, 0xc6,
	0x74, 0x98, 0xc3, 0xe6, 0xae, 0x71, 0x26, 0x12, 0x11, 0x4d, 0xf0, 0xff, 0xdf, 0xa7, 0xc1, 0x46,
	0x27, 0x20, 0xbc, 0xc9, 0xd9, 0x35, 0x75, 0x09, 0xdc, 0x07, 0x29, 0xda, 0x53, 0x24, 0x5d, 0x3a,
	0xca, 0x55, 0x33, 0xe3, 0x6f, 0x87, 0xa9, 0xda, 0x73, 0x94, 0xa2, 0x3d, 0xb8, 0x07, 0xd6, 0x48,
	0x1f, 0x53, 0x57, 0x49, 0xc5, 0x12, 0x9a, 0x24, 0xf0, 0x3f, 0x00, 0xae, 0x29, 0x0f, 0x42, 0xcb,
	0xc3, 0x7d, 0xa2, 0xa4, 0x85, 0x94, 0x13, 0x95, 0x06, 0xee, 0x13, 0xf8, 0x2f, 0xc8, 0xb9, 0x78,
	0xaa, 0xae, 0x0a, 0x35, 0x1b, 0x17, 0xa6, 0x22, 0x0d, 0x2c, 0x6c, 0x87, 0x74, 0x40, 0x94, 0x35,
	0x5d, 0x3a, 0xca, 0xa2, 0x2c, 0x0d, 0x2a, 0x22, 0x87, 0x4f, 0xc0, 0x2a, 0x67, 0x2e, 0x51, 0x32,
	0xba, 0x74, 0x94, 0x3f, 0x39, 0x28, 0xfd, 0xfe, 0x1e, 0xa5, 0x78, 0x6a, 0xc4, 0x5c, 0x82, 0x04,
	0x09, 0x9f, 0x82, 0xf5, 0x88, 0x5a, 0xa2, 0x69, 0x5d, 0x34, 0x69, 0x7f, 0x6b, 0xea, 0xd4, 0x44,
	0x5b, 0x26, 0xa2, 0xf1, 0x17, 0x2a, 0x60, 0x5d, 0x3c, 0x16, 0xbf, 0x55, 0xb2, 0x62, 0xc4, 0x69,
	0x0a, 0xf7, 0x41, 0x86, 0x13, 0x87, 0x32, 0x4f, 0xc9, 0x09, 0x21, 0xc9, 0x20, 0x04, 0xab, 0x36,
	0x0d, 0x6f, 0x15, 0x20, 0xaa, 0x22, 0x86, 0x32, 0x48, 0xbf, 0xa5, 0xbe, 0xb2, 0x21, 0x4a, 0x71,
	0x18, 0xef, 0x87, 0x7b, 0x3d, 0x4e, 0x82, 0xc0, 0x2a, 0x2b, 0x9b, 0x93, 0xe5, 0x93, 0x42, 0x79,
	0x51, 0x3c, 0x51, 0xb6, 0x96, 0xc4, 0x93, 0x02, 0x07, 0xb9, 0x36, 0x7b, 0x4d, 0xbc, 0xf6, 0xad,
	0x4f, 0xe0, 0x31, 0x80, 0xed, 0xcb, 0x17, 0x66, 0xc3, 0x6a, 0xbf, 0x6c, 0x9a, 0x16, 0x32, 0xcf,
	0x3a, 0x17, 0x15, 0x24, 0xaf, 0xa8, 0x7b, 0xc3, 0x91, 0x2e, 0xcf, 0x30, 0x44, 0x9c, 0xc8, 0xc5,
	0x1c, 0x3e, 0x00, 0xf9, 0x05, 0xba, 0xd2, 0xac, 0xc9, 0x92, 0x2a, 0x0f, 0x47, 0xfa, 0xe6, 0x8c,
	0xac, 0x34, 0x6b, 0xea, 0xce, 0xbb, 0x0f, 0xda, 0xca, 0xe7, 0x8f, 0xda, 0xdc, 0xa6, 0xf0, 0x43,
	0x02, 0xd9, 0xe9, 0x8d, 0xc2, 0x02, 0xd8, 0xe9, 0xb4, 0x4c, 0x64, 0xa1, 0xcb, 0x8b, 0x45, 0xcb,
	0xdd, 0xe1, 0x48, 0xdf, 0x9e, 0x5d, 0x7b, 0xe2, 0xf8, 0x10, 0x6c, 0xcf, 0xd9, 0x7a, 0xad, 0x61,
	0x22, 0x59, 0x52, 0x77, 0x86, 0x23, 0x7d, 0x6b, 0x4a, 0xd6, 0xa9, 0x47, 0x38, 0xd4, 0xc1, 0xe6,
	0x9c, 0xbb, 0xaa, 0xc8, 0x69, 0x35, 0x3f, 0x1c, 0xe9, 0x60, 0x0a, 0x5d, 0xe1, 0x65, 0xd7, 0x7a,
	0xa5, 0x51, 0x39, 0x33, 0x91, 0x9c, 0x59, 0x76, 0xad, 0x63, 0x0f, 0x3b, 0xe4, 0x17, 0xd7, 0x56,
	0xa7, 0x69, 0x22, 0x39, 0xb7, 0xec, 0xda, 0x8a, 0x7c, 0xc2, 0x55, 0x39, 0xd9, 0x74, 0xb6, 0x5b,
	0xe1, 0x93, 0x04, 0xc0, 0xfc, 0x57, 0x00, 0x1f, 0x2d, 0x9a, 0x76, 0x6a, 0x56, 0xf5, 0xb2, 0x7d,
	0x2e, 0xaf, 0xa8, 0x70, 0x38, 0xd2, 0xf3, 0x73, 0xac, 0xca, 0xc2, 0x1b, 0xf8, 0x18, 0xc0, 0x25,
	0x74, 0xba, 0xec, 0x6c, 0xc0, 0x09, 0x3b, 0x59, 0xf7, 0x14, 0xec, 0x2f, 0xc1, 0xcd, 0x4e, 0xf5,
	0xa2, 0xd6, 0x3a, 0x37, 0x91, 0x9c, 0x56, 0xff, 0x19, 0x8e, 0xf4, 0xdd, 0x79, 0x43, 0x33, 0xea,
	0xba, 0x34, 0xb8, 0x21, 0x5c, 0x85, 0xc9, 0xb4, 0x0b, 0x03, 0x56, 0x95, 0xbb, 0xb1, 0x26, 0x7d,
	0x1d, 0x6b, 0xd2, 0xf7, 0xb1, 0x26, 0x7d, 0xb9, 0xd7, 0xa4, 0xbb, 0x7b, 0x4d, 0x7a, 0x95, 0x1a,
	0x94, 0xbb, 0x19, 0xf1, 0x0f, 0x3e, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x37, 0x8a, 0x24, 0x25,
	0xc2, 0x04, 0x00, 0x00,
}

func (m *UserProfile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserProfile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserProfile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Address_2) > 0 {
		i -= len(m.Address_2)
		copy(dAtA[i:], m.Address_2)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Address_2)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Address_1) > 0 {
		i -= len(m.Address_1)
		copy(dAtA[i:], m.Address_1)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Address_1)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Zip) > 0 {
		i -= len(m.Zip)
		copy(dAtA[i:], m.Zip)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Zip)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.City) > 0 {
		i -= len(m.City)
		copy(dAtA[i:], m.City)
		i = encodeVarintUser(dAtA, i, uint64(len(m.City)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Region) > 0 {
		i -= len(m.Region)
		copy(dAtA[i:], m.Region)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Region)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Country) > 0 {
		i -= len(m.Country)
		copy(dAtA[i:], m.Country)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Country)))
		i--
		dAtA[i] = 0x42
	}
	if m.UiRole != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.UiRole))
		i--
		dAtA[i] = 0x38
	}
	if m.Role != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x30
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.LastName) > 0 {
		i -= len(m.LastName)
		copy(dAtA[i:], m.LastName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.LastName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FirstName) > 0 {
		i -= len(m.FirstName)
		copy(dAtA[i:], m.FirstName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.FirstName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintUser(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserProfile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.IsActive {
		n += 2
	}
	if m.Role != 0 {
		n += 1 + sovUser(uint64(m.Role))
	}
	if m.UiRole != 0 {
		n += 1 + sovUser(uint64(m.UiRole))
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Region)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Zip)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Address_1)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Address_2)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserProfile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserProfile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserProfile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= UserRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UiRole", wireType)
			}
			m.UiRole = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UiRole |= UserUIRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Region", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Region = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Zip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address_1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address_1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address_2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address_2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
