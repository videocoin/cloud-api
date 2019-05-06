// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users/v1/user_service.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import rpc "github.com/VideoCoin/cloud-api/rpc"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,4,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{0}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

type LoginUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserRequest) Reset()         { *m = LoginUserRequest{} }
func (m *LoginUserRequest) String() string { return proto.CompactTextString(m) }
func (*LoginUserRequest) ProtoMessage()    {}
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{1}
}
func (m *LoginUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserRequest.Unmarshal(m, b)
}
func (m *LoginUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserRequest.Marshal(b, m, deterministic)
}
func (dst *LoginUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserRequest.Merge(dst, src)
}
func (m *LoginUserRequest) XXX_Size() int {
	return xxx_messageInfo_LoginUserRequest.Size(m)
}
func (m *LoginUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserRequest proto.InternalMessageInfo

func (m *LoginUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginUserResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserResponse) Reset()         { *m = LoginUserResponse{} }
func (m *LoginUserResponse) String() string { return proto.CompactTextString(m) }
func (*LoginUserResponse) ProtoMessage()    {}
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{2}
}
func (m *LoginUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserResponse.Unmarshal(m, b)
}
func (m *LoginUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserResponse.Marshal(b, m, deterministic)
}
func (dst *LoginUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserResponse.Merge(dst, src)
}
func (m *LoginUserResponse) XXX_Size() int {
	return xxx_messageInfo_LoginUserResponse.Size(m)
}
func (m *LoginUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserResponse proto.InternalMessageInfo

func (m *LoginUserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{3}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (dst *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(dst, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ResetPasswordUserRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,2,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordUserRequest) Reset()         { *m = ResetPasswordUserRequest{} }
func (m *ResetPasswordUserRequest) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordUserRequest) ProtoMessage()    {}
func (*ResetPasswordUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{4}
}
func (m *ResetPasswordUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordUserRequest.Unmarshal(m, b)
}
func (m *ResetPasswordUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordUserRequest.Marshal(b, m, deterministic)
}
func (dst *ResetPasswordUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordUserRequest.Merge(dst, src)
}
func (m *ResetPasswordUserRequest) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordUserRequest.Size(m)
}
func (m *ResetPasswordUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordUserRequest proto.InternalMessageInfo

func (m *ResetPasswordUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetPasswordUserRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

type StartRecoveryUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRecoveryUserRequest) Reset()         { *m = StartRecoveryUserRequest{} }
func (m *StartRecoveryUserRequest) String() string { return proto.CompactTextString(m) }
func (*StartRecoveryUserRequest) ProtoMessage()    {}
func (*StartRecoveryUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{5}
}
func (m *StartRecoveryUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRecoveryUserRequest.Unmarshal(m, b)
}
func (m *StartRecoveryUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRecoveryUserRequest.Marshal(b, m, deterministic)
}
func (dst *StartRecoveryUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRecoveryUserRequest.Merge(dst, src)
}
func (m *StartRecoveryUserRequest) XXX_Size() int {
	return xxx_messageInfo_StartRecoveryUserRequest.Size(m)
}
func (m *StartRecoveryUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRecoveryUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRecoveryUserRequest proto.InternalMessageInfo

func (m *StartRecoveryUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RecoverUserRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,3,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoverUserRequest) Reset()         { *m = RecoverUserRequest{} }
func (m *RecoverUserRequest) String() string { return proto.CompactTextString(m) }
func (*RecoverUserRequest) ProtoMessage()    {}
func (*RecoverUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{6}
}
func (m *RecoverUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoverUserRequest.Unmarshal(m, b)
}
func (m *RecoverUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoverUserRequest.Marshal(b, m, deterministic)
}
func (dst *RecoverUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoverUserRequest.Merge(dst, src)
}
func (m *RecoverUserRequest) XXX_Size() int {
	return xxx_messageInfo_RecoverUserRequest.Size(m)
}
func (m *RecoverUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoverUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecoverUserRequest proto.InternalMessageInfo

func (m *RecoverUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RecoverUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RecoverUserRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

type WhitelistResponse struct {
	Items                []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhitelistResponse) Reset()         { *m = WhitelistResponse{} }
func (m *WhitelistResponse) String() string { return proto.CompactTextString(m) }
func (*WhitelistResponse) ProtoMessage()    {}
func (*WhitelistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{7}
}
func (m *WhitelistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhitelistResponse.Unmarshal(m, b)
}
func (m *WhitelistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhitelistResponse.Marshal(b, m, deterministic)
}
func (dst *WhitelistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistResponse.Merge(dst, src)
}
func (m *WhitelistResponse) XXX_Size() int {
	return xxx_messageInfo_WhitelistResponse.Size(m)
}
func (m *WhitelistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistResponse proto.InternalMessageInfo

func (m *WhitelistResponse) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

type LookupByAddressRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupByAddressRequest) Reset()         { *m = LookupByAddressRequest{} }
func (m *LookupByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*LookupByAddressRequest) ProtoMessage()    {}
func (*LookupByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_90500abb06766970, []int{8}
}
func (m *LookupByAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupByAddressRequest.Unmarshal(m, b)
}
func (m *LookupByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupByAddressRequest.Marshal(b, m, deterministic)
}
func (dst *LookupByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupByAddressRequest.Merge(dst, src)
}
func (m *LookupByAddressRequest) XXX_Size() int {
	return xxx_messageInfo_LookupByAddressRequest.Size(m)
}
func (m *LookupByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupByAddressRequest proto.InternalMessageInfo

func (m *LookupByAddressRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "cloud.api.users.v1.CreateUserRequest")
	proto.RegisterType((*LoginUserRequest)(nil), "cloud.api.users.v1.LoginUserRequest")
	proto.RegisterType((*LoginUserResponse)(nil), "cloud.api.users.v1.LoginUserResponse")
	proto.RegisterType((*UserRequest)(nil), "cloud.api.users.v1.UserRequest")
	proto.RegisterType((*ResetPasswordUserRequest)(nil), "cloud.api.users.v1.ResetPasswordUserRequest")
	proto.RegisterType((*StartRecoveryUserRequest)(nil), "cloud.api.users.v1.StartRecoveryUserRequest")
	proto.RegisterType((*RecoverUserRequest)(nil), "cloud.api.users.v1.RecoverUserRequest")
	proto.RegisterType((*WhitelistResponse)(nil), "cloud.api.users.v1.WhitelistResponse")
	proto.RegisterType((*LookupByAddressRequest)(nil), "cloud.api.users.v1.LookupByAddressRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	Login(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	ResetPassword(ctx context.Context, in *ResetPasswordUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserProfile, error)
	Activate(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserProfile, error)
	StartRecovery(ctx context.Context, in *StartRecoveryUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Recover(ctx context.Context, in *RecoverUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Whitelist(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*WhitelistResponse, error)
	LookupByAddress(ctx context.Context, in *LookupByAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Activate(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) StartRecovery(ctx context.Context, in *StartRecoveryUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/StartRecovery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Recover(ctx context.Context, in *RecoverUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/Recover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Whitelist(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*WhitelistResponse, error) {
	out := new(WhitelistResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/Whitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LookupByAddress(ctx context.Context, in *LookupByAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.users.v1.UserService/LookupByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Health(context.Context, *empty.Empty) (*rpc.HealthStatus, error)
	Create(context.Context, *CreateUserRequest) (*LoginUserResponse, error)
	Login(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	Logout(context.Context, *empty.Empty) (*empty.Empty, error)
	ResetPassword(context.Context, *ResetPasswordUserRequest) (*empty.Empty, error)
	Get(context.Context, *empty.Empty) (*UserProfile, error)
	Activate(context.Context, *UserRequest) (*UserProfile, error)
	StartRecovery(context.Context, *StartRecoveryUserRequest) (*empty.Empty, error)
	Recover(context.Context, *RecoverUserRequest) (*empty.Empty, error)
	Whitelist(context.Context, *empty.Empty) (*WhitelistResponse, error)
	LookupByAddress(context.Context, *LookupByAddressRequest) (*empty.Empty, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Activate(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_StartRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRecoveryUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).StartRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/StartRecovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).StartRecovery(ctx, req.(*StartRecoveryUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Recover(ctx, req.(*RecoverUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Whitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Whitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/Whitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Whitelist(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LookupByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LookupByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.users.v1.UserService/LookupByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LookupByAddress(ctx, req.(*LookupByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.users.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _UserService_Health_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _UserService_Activate_Handler,
		},
		{
			MethodName: "StartRecovery",
			Handler:    _UserService_StartRecovery_Handler,
		},
		{
			MethodName: "Recover",
			Handler:    _UserService_Recover_Handler,
		},
		{
			MethodName: "Whitelist",
			Handler:    _UserService_Whitelist_Handler,
		},
		{
			MethodName: "LookupByAddress",
			Handler:    _UserService_LookupByAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/v1/user_service.proto",
}

func init() {
	proto.RegisterFile("users/v1/user_service.proto", fileDescriptor_user_service_90500abb06766970)
}

var fileDescriptor_user_service_90500abb06766970 = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0x8e, 0xf3, 0xb5, 0xcd, 0xec, 0x76, 0x9b, 0xcc, 0x46, 0x5d, 0x6f, 0x42, 0xd3, 0x65, 0x44,
	0xd1, 0x6e, 0x44, 0x6c, 0x15, 0x38, 0xc1, 0x01, 0xb5, 0x2b, 0x04, 0x87, 0x0a, 0x56, 0x2e, 0x1f,
	0xd2, 0x5e, 0xca, 0xd4, 0x9e, 0x3a, 0xa3, 0x26, 0x1e, 0x33, 0x33, 0x76, 0x1b, 0xaa, 0x5e, 0xb8,
	0x70, 0x87, 0x0b, 0x3f, 0x83, 0x33, 0x27, 0x8e, 0x1c, 0x91, 0xf8, 0x03, 0xa8, 0x70, 0xe1, 0x5f,
	0x20, 0x8f, 0xc7, 0xa9, 0x93, 0xd8, 0x0a, 0x3d, 0xec, 0x6d, 0x3e, 0x9e, 0xf7, 0x79, 0xde, 0xf7,
	0x9d, 0x79, 0x5e, 0xd0, 0x8f, 0x04, 0xe1, 0xc2, 0x8e, 0xf7, 0xed, 0x64, 0x71, 0x22, 0x08, 0x8f,
	0xa9, 0x4b, 0xac, 0x90, 0x33, 0xc9, 0x20, 0x74, 0x27, 0x2c, 0xf2, 0x2c, 0x1c, 0x52, 0x4b, 0xc1,
	0xac, 0x78, 0xbf, 0x37, 0xf2, 0xa9, 0x1c, 0x47, 0xa7, 0x96, 0xcb, 0xa6, 0xb6, 0xcf, 0x7c, 0x66,
	0x2b, 0xe8, 0x69, 0x74, 0xa6, 0x76, 0x6a, 0xa3, 0x56, 0x29, 0x45, 0xef, 0x20, 0x07, 0x27, 0x41,
	0xcc, 0x66, 0x21, 0x67, 0x97, 0xb3, 0x34, 0xc8, 0x1d, 0xf9, 0x24, 0x18, 0xc5, 0x78, 0x42, 0x3d,
	0x2c, 0x89, 0xbd, 0xb2, 0xd0, 0x14, 0x76, 0x8e, 0xe2, 0x2b, 0xea, 0x11, 0xf6, 0x82, 0xd1, 0xc0,
	0x56, 0xa9, 0x8d, 0x70, 0x48, 0x6d, 0x1e, 0xba, 0xf6, 0x98, 0xe0, 0x89, 0x1c, 0xeb, 0x80, 0xbe,
	0xcf, 0x98, 0x3f, 0x21, 0xb7, 0x99, 0x91, 0x69, 0x28, 0x67, 0xfa, 0xf2, 0x0d, 0x7d, 0x99, 0x44,
	0xe2, 0x20, 0x60, 0x12, 0x4b, 0xca, 0x02, 0xa1, 0x6f, 0x1f, 0x2d, 0xb4, 0x23, 0x3d, 0x44, 0xbf,
	0x18, 0xa0, 0xf3, 0x82, 0x13, 0x2c, 0xc9, 0x97, 0x82, 0x70, 0x87, 0x7c, 0x1b, 0x11, 0x21, 0xe1,
	0x2e, 0x68, 0x90, 0x29, 0xa6, 0x13, 0xd3, 0x78, 0x6a, 0x3c, 0x6b, 0x1d, 0xb6, 0x7e, 0xfd, 0xf7,
	0xb7, 0x5a, 0x9d, 0x57, 0xbf, 0x31, 0x9c, 0xf4, 0x1c, 0xee, 0x81, 0x8d, 0x10, 0x0b, 0x71, 0xc1,
	0xb8, 0x67, 0x56, 0x17, 0x30, 0xed, 0x86, 0x33, 0xbf, 0x82, 0x3b, 0xa0, 0x1e, 0xe0, 0x29, 0x31,
	0x6b, 0x8b, 0x90, 0x9a, 0xa3, 0x8e, 0xe1, 0xfb, 0xa0, 0xed, 0xb2, 0xe0, 0x8c, 0xf2, 0xe9, 0xc9,
	0x9c, 0xad, 0xbe, 0xcc, 0xb6, 0xa5, 0x21, 0x2f, 0x35, 0x02, 0xbd, 0x02, 0xed, 0x23, 0xe6, 0xd3,
	0xe0, 0x35, 0x24, 0x8c, 0x9e, 0x83, 0x4e, 0x8e, 0x5b, 0x84, 0x2c, 0x10, 0x04, 0x76, 0x41, 0x43,
	0xb2, 0x73, 0x12, 0xa4, 0xe4, 0x4e, 0xba, 0x41, 0x3b, 0xe0, 0x7e, 0x3e, 0x83, 0x87, 0xa0, 0x4a,
	0x3d, 0x8d, 0xa8, 0x52, 0x0f, 0x5d, 0x00, 0xd3, 0x21, 0x82, 0xc8, 0x2c, 0xed, 0x3c, 0x36, 0x9f,
	0x8c, 0x51, 0xde, 0xbd, 0xa2, 0xf6, 0x54, 0xd7, 0xb6, 0xe7, 0x43, 0x60, 0x1e, 0x4b, 0xcc, 0xa5,
	0x43, 0x5c, 0x16, 0x13, 0x3e, 0xbb, 0x4b, 0x9b, 0xd0, 0x0f, 0x06, 0x80, 0x3a, 0x30, 0x1f, 0x57,
	0xd8, 0x81, 0xff, 0xfb, 0x09, 0x8a, 0xca, 0xa8, 0xad, 0x2d, 0xe3, 0x39, 0xe8, 0x7c, 0x3d, 0xa6,
	0x92, 0x4c, 0xa8, 0x90, 0xf9, 0x97, 0xa0, 0x92, 0x4c, 0x85, 0x69, 0x3c, 0xad, 0x25, 0x79, 0xa8,
	0x0d, 0xfa, 0x08, 0x6c, 0x1f, 0x31, 0x76, 0x1e, 0x85, 0x87, 0xb3, 0x03, 0xcf, 0xe3, 0x44, 0x88,
	0xdb, 0x46, 0xdf, 0xc3, 0xe9, 0x89, 0xae, 0xf8, 0x7e, 0xa2, 0xd8, 0xe4, 0xf5, 0xf6, 0x33, 0x73,
	0xe8, 0x64, 0x77, 0xef, 0xfe, 0xd8, 0x4a, 0xdf, 0xf2, 0x38, 0x9d, 0x10, 0xf0, 0x73, 0xd0, 0xfc,
	0x54, 0x99, 0x0e, 0x6e, 0x5b, 0xa9, 0xa5, 0xac, 0xcc, 0x6f, 0xd6, 0xc7, 0x89, 0xdf, 0x7a, 0x7d,
	0xeb, 0x76, 0x7c, 0xf0, 0xd0, 0xb5, 0x52, 0xf8, 0xb1, 0xc4, 0x32, 0x12, 0xa8, 0xfd, 0xfd, 0x9f,
	0xff, 0xfc, 0x54, 0x05, 0x70, 0x43, 0x5b, 0xf7, 0x3b, 0x18, 0x82, 0x66, 0x6a, 0x32, 0xb8, 0x67,
	0xad, 0xce, 0x1d, 0x6b, 0xc5, 0x80, 0xbd, 0x42, 0xd8, 0xca, 0xcf, 0x44, 0xa6, 0x52, 0x82, 0x68,
	0x53, 0x59, 0x5e, 0x3b, 0x5b, 0x7c, 0x60, 0x0c, 0xe1, 0x14, 0x34, 0x14, 0x1c, 0xbe, 0xb5, 0x86,
	0xe9, 0x4e, 0x7a, 0x8f, 0x95, 0x5e, 0x07, 0x3d, 0xc8, 0xf4, 0x70, 0x24, 0xc7, 0x89, 0xdc, 0x67,
	0xa0, 0x79, 0xc4, 0x7c, 0x16, 0xc9, 0xd2, 0x8e, 0x95, 0x9c, 0xa3, 0xae, 0xa2, 0x7c, 0x38, 0x5c,
	0xa0, 0x84, 0x12, 0x6c, 0x2e, 0xb8, 0x07, 0xbe, 0x53, 0x94, 0x60, 0x99, 0xc1, 0x4a, 0xc5, 0x76,
	0x94, 0xd8, 0x63, 0x04, 0xf3, 0x62, 0x36, 0x4f, 0x68, 0x92, 0x2a, 0xbe, 0x00, 0xb5, 0x4f, 0x48,
	0x79, 0x09, 0xbb, 0x45, 0x39, 0x24, 0xb2, 0x2f, 0x39, 0x3b, 0xa3, 0x13, 0x92, 0xd5, 0x02, 0x1f,
	0xe4, 0x9f, 0x03, 0x0a, 0xb0, 0x71, 0xe0, 0x4a, 0x1a, 0x27, 0xcf, 0x5f, 0x4a, 0x91, 0x65, 0xbe,
	0x56, 0x63, 0x4f, 0x69, 0xec, 0xa2, 0x5e, 0x5e, 0xc3, 0xbe, 0xa2, 0xde, 0xb5, 0x8d, 0xb5, 0x4a,
	0x52, 0xca, 0x25, 0xd8, 0x5c, 0x98, 0x02, 0xc5, 0x0d, 0x2c, 0x1b, 0x14, 0xa5, 0x0d, 0x7c, 0x53,
	0xa9, 0xf7, 0xd1, 0x76, 0xa6, 0xce, 0x75, 0xb0, 0x2d, 0x12, 0xaa, 0x44, 0x99, 0x80, 0x7b, 0x9a,
	0x11, 0xbe, 0x5d, 0xfc, 0x68, 0xcb, 0xe3, 0xa5, 0x54, 0xad, 0xa7, 0xd4, 0xba, 0x68, 0x6b, 0x49,
	0x2d, 0x95, 0x69, 0xcd, 0xe7, 0x43, 0xe9, 0x8b, 0x15, 0x7e, 0xeb, 0x95, 0xb1, 0x82, 0x9e, 0x28,
	0x9d, 0x47, 0xb0, 0x93, 0xe9, 0x5c, 0xcc, 0x99, 0x2f, 0xc1, 0xd6, 0xd2, 0x6c, 0x81, 0xc3, 0x62,
	0xaf, 0x14, 0x0d, 0xa0, 0x75, 0x7d, 0x84, 0x4f, 0xe6, 0x1f, 0x31, 0x8d, 0xb3, 0xaf, 0xf4, 0xe2,
	0xfa, 0xb0, 0xfb, 0xfb, 0xcd, 0xa0, 0xf2, 0xc7, 0xcd, 0xa0, 0xf2, 0xd7, 0xcd, 0xa0, 0xf2, 0xf3,
	0xdf, 0x83, 0xca, 0xab, 0x6a, 0xbc, 0x7f, 0xda, 0x54, 0x44, 0xef, 0xfd, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x94, 0x1d, 0x38, 0xad, 0xdc, 0x08, 0x00, 0x00,
}
