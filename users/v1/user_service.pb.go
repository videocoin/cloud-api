// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users/v1/user_service.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import rpc "github.com/VideoCoin/cloud-api/rpc"
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{0}
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{1}
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{2}
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{3}
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{4}
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{5}
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{6}
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{7}
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
	return fileDescriptor_user_service_09a70ad1fddf880a, []int{8}
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
	proto.RegisterFile("users/v1/user_service.proto", fileDescriptor_user_service_09a70ad1fddf880a)
}

var fileDescriptor_user_service_09a70ad1fddf880a = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0x4e, 0x1b, 0x3b,
	0x14, 0xc7, 0xf3, 0x01, 0x01, 0x7c, 0xf9, 0x8a, 0x89, 0x60, 0x48, 0x2e, 0x81, 0x6b, 0x5d, 0xae,
	0x2e, 0xa8, 0xcc, 0x14, 0xba, 0xeb, 0x0e, 0x68, 0xd5, 0x2e, 0x50, 0x8b, 0x42, 0x3f, 0xa4, 0x6e,
	0x90, 0x99, 0x31, 0x13, 0x8b, 0xcc, 0x78, 0xb0, 0x3d, 0x01, 0x8a, 0xd8, 0x54, 0x7d, 0x82, 0x76,
	0xd3, 0x47, 0xea, 0xb2, 0x52, 0x5f, 0xa0, 0xa2, 0x7d, 0x90, 0x6a, 0x3c, 0x9e, 0x30, 0x49, 0x66,
	0x1a, 0x55, 0xdd, 0xf9, 0xd8, 0xc7, 0xff, 0xdf, 0x39, 0x3e, 0xe7, 0x18, 0x34, 0x42, 0x41, 0xb8,
	0xb0, 0xba, 0xdb, 0x56, 0xb4, 0x38, 0x16, 0x84, 0x77, 0xa9, 0x4d, 0xcc, 0x80, 0x33, 0xc9, 0x20,
	0xb4, 0x3b, 0x2c, 0x74, 0x4c, 0x1c, 0x50, 0x53, 0xb9, 0x99, 0xdd, 0xed, 0xfa, 0x96, 0x4b, 0x65,
	0x3b, 0x3c, 0x31, 0x6d, 0xe6, 0x59, 0x2e, 0x73, 0x99, 0xa5, 0x5c, 0x4f, 0xc2, 0x53, 0x65, 0x29,
	0x43, 0xad, 0x62, 0x89, 0xba, 0x95, 0x72, 0x7f, 0x45, 0x1d, 0xc2, 0xf6, 0x19, 0xf5, 0x2d, 0xa5,
	0xbb, 0x85, 0x03, 0x6a, 0xf1, 0xc0, 0xb6, 0xda, 0x04, 0x77, 0x64, 0x5b, 0x5f, 0x68, 0xb8, 0x8c,
	0xb9, 0x1d, 0x72, 0x27, 0x4b, 0xbc, 0x40, 0x5e, 0xe9, 0xc3, 0xbf, 0xf5, 0x61, 0x74, 0x13, 0xfb,
	0x3e, 0x93, 0x58, 0x52, 0xe6, 0x0b, 0x7d, 0xba, 0xd0, 0x97, 0x4b, 0xbc, 0x89, 0xde, 0x17, 0x41,
	0x75, 0x9f, 0x13, 0x2c, 0xc9, 0x4b, 0x41, 0x78, 0x8b, 0x9c, 0x87, 0x44, 0x48, 0x58, 0x03, 0xe3,
	0xc4, 0xc3, 0xb4, 0x63, 0x14, 0xd7, 0x8a, 0xff, 0x4f, 0xb5, 0x62, 0x03, 0xd6, 0xc1, 0x64, 0x80,
	0x85, 0xb8, 0x60, 0xdc, 0x31, 0x4a, 0xea, 0xa0, 0x67, 0x43, 0x08, 0xc6, 0x7c, 0xec, 0x11, 0xa3,
	0xac, 0xf6, 0xd5, 0x1a, 0x6e, 0x80, 0x79, 0x9b, 0xf9, 0xa7, 0x94, 0x7b, 0xc7, 0xbd, 0x7b, 0x63,
	0xea, 0x7c, 0x4e, 0xef, 0x1f, 0xea, 0x6d, 0xf4, 0x08, 0xcc, 0x1f, 0x30, 0x97, 0xfa, 0x7f, 0x14,
	0x04, 0xda, 0x00, 0xd5, 0x94, 0x8a, 0x08, 0x98, 0x2f, 0x48, 0x24, 0x23, 0xd9, 0x19, 0xf1, 0x13,
	0x19, 0x65, 0xa0, 0x15, 0xf0, 0x57, 0x9a, 0x35, 0x0b, 0x4a, 0xd4, 0xd1, 0x1e, 0x25, 0xea, 0x20,
	0x0c, 0x8c, 0x16, 0x11, 0x44, 0x26, 0x01, 0xa6, 0x7d, 0xd3, 0x11, 0x14, 0x07, 0x9e, 0x21, 0x2b,
	0xe5, 0x52, 0x76, 0xca, 0xf7, 0x81, 0x71, 0x24, 0x31, 0x97, 0x2d, 0x62, 0xb3, 0x2e, 0xe1, 0x57,
	0x23, 0x53, 0x47, 0xe7, 0x00, 0x6a, 0xe7, 0x01, 0xdf, 0xe1, 0xfc, 0x7e, 0x59, 0xab, 0xac, 0x20,
	0xcb, 0xd9, 0x41, 0x6e, 0x80, 0xea, 0xeb, 0x36, 0x95, 0xa4, 0x43, 0x85, 0x4c, 0xbf, 0x28, 0x95,
	0xc4, 0x13, 0x46, 0x71, 0xad, 0x1c, 0x11, 0x95, 0x81, 0x76, 0xc0, 0xe2, 0x01, 0x63, 0x67, 0x61,
	0xb0, 0x77, 0xb5, 0xeb, 0x38, 0x9c, 0x08, 0x91, 0x44, 0x68, 0x80, 0x09, 0x1c, 0xef, 0xe8, 0x18,
	0x13, 0x73, 0xe7, 0xc3, 0x54, 0x5c, 0x86, 0xa3, 0x78, 0xae, 0xe0, 0x73, 0x50, 0x79, 0xaa, 0xba,
	0x1d, 0x2e, 0x9a, 0x71, 0x2f, 0x9b, 0x49, 0xa3, 0x9b, 0x8f, 0xa3, 0x46, 0xaf, 0x37, 0xcc, 0xbb,
	0xa1, 0xe3, 0x81, 0x6d, 0xc6, 0xee, 0x47, 0x12, 0xcb, 0x50, 0xa0, 0xf9, 0x77, 0x5f, 0x7f, 0x7c,
	0x2c, 0x01, 0x38, 0xa9, 0x67, 0xe6, 0x2d, 0x0c, 0x40, 0x25, 0xee, 0x6e, 0xb8, 0x6e, 0x0e, 0x4f,
	0xab, 0x39, 0xd4, 0xf9, 0xf5, 0x4c, 0xb7, 0xa1, 0xa6, 0x42, 0x86, 0x22, 0x41, 0x34, 0xa3, 0x66,
	0x4d, 0x8f, 0x94, 0x78, 0x58, 0xdc, 0x84, 0x1e, 0x18, 0x57, 0xee, 0xf0, 0xdf, 0x11, 0x4a, 0xbf,
	0xc5, 0x5b, 0x52, 0xbc, 0x2a, 0x9a, 0x4e, 0x78, 0x38, 0x94, 0xed, 0x08, 0xf7, 0x0c, 0x54, 0x0e,
	0x98, 0xcb, 0x42, 0x99, 0xfb, 0x62, 0x39, 0xfb, 0xa8, 0xa6, 0x24, 0x67, 0x37, 0xfb, 0x24, 0xa1,
	0x04, 0x33, 0x7d, 0x8d, 0x0f, 0xef, 0x65, 0x05, 0x98, 0x37, 0x1b, 0xb9, 0xb0, 0x15, 0x05, 0x5b,
	0x42, 0x30, 0x0d, 0xb3, 0x78, 0x24, 0x13, 0x65, 0xf1, 0x02, 0x94, 0x9f, 0x90, 0xfc, 0x14, 0x56,
	0xb3, 0x62, 0x88, 0xb0, 0x87, 0x9c, 0x9d, 0xd2, 0x0e, 0x49, 0x72, 0x81, 0xd3, 0xe9, 0x72, 0x40,
	0x01, 0x26, 0x77, 0x6d, 0x49, 0xbb, 0x51, 0xf9, 0x73, 0x25, 0x92, 0xc8, 0x47, 0x32, 0xd6, 0x15,
	0x63, 0x15, 0xd5, 0xd3, 0x0c, 0xeb, 0x9a, 0x3a, 0x37, 0x16, 0xd6, 0x94, 0x28, 0x95, 0x4b, 0x30,
	0xd3, 0x37, 0xd6, 0xd9, 0x0f, 0x98, 0x37, 0xf9, 0xb9, 0x0f, 0xf8, 0x8f, 0xa2, 0x37, 0xd0, 0x62,
	0x42, 0xe7, 0xfa, 0xb2, 0x25, 0x22, 0xa9, 0x88, 0x4c, 0xc0, 0x84, 0x56, 0x84, 0xff, 0x65, 0x17,
	0x6d, 0xf0, 0xef, 0xc8, 0xa5, 0xd5, 0x15, 0xad, 0x86, 0xe6, 0x06, 0x68, 0x31, 0x66, 0xaa, 0xf7,
	0x25, 0xe4, 0x56, 0x2c, 0xb3, 0xad, 0x87, 0x7e, 0x12, 0xb4, 0xac, 0x38, 0x0b, 0xb0, 0x9a, 0x70,
	0x2e, 0x7a, 0xca, 0x97, 0x60, 0x6e, 0xe0, 0x3b, 0x81, 0x9b, 0xd9, 0xb3, 0x92, 0xf5, 0xe7, 0x8c,
	0x7a, 0x47, 0xb8, 0xdc, 0x6b, 0xc4, 0xf8, 0x9e, 0x75, 0xad, 0x17, 0x37, 0x7b, 0xb5, 0xcf, 0xb7,
	0xcd, 0xc2, 0x97, 0xdb, 0x66, 0xe1, 0xdb, 0x6d, 0xb3, 0xf0, 0xe9, 0x7b, 0xb3, 0xf0, 0xa6, 0xd4,
	0xdd, 0x3e, 0xa9, 0x28, 0xa1, 0x07, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x45, 0x4b, 0x82, 0xc2,
	0x12, 0x08, 0x00, 0x00,
}
