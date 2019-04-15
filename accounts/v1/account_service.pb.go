// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accounts/v1/account_service.proto

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

type AccountRequest struct {
	OwnerID              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_service_320cdabe64faa457, []int{0}
}
func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (dst *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(dst, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetOwnerID() string {
	if m != nil {
		return m.OwnerID
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountRequest)(nil), "cloud.api.account.v1.AccountRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
	Create(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountProfile, error)
	Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AccountProfile, error)
	Refresh(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AccountProfile, error)
	Key(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountKey, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Create(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountProfile, error) {
	out := new(AccountProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AccountProfile, error) {
	out := new(AccountProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Refresh(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AccountProfile, error) {
	out := new(AccountProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Key(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountKey, error) {
	out := new(AccountKey)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/Key", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	Health(context.Context, *empty.Empty) (*rpc.HealthStatus, error)
	Create(context.Context, *AccountRequest) (*AccountProfile, error)
	Get(context.Context, *empty.Empty) (*AccountProfile, error)
	Refresh(context.Context, *empty.Empty) (*AccountProfile, error)
	Key(context.Context, *AccountRequest) (*AccountKey, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Create(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Get(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Refresh(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Key_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Key(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/Key",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Key(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.account.v1.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _AccountService_Health_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccountService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccountService_Get_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _AccountService_Refresh_Handler,
		},
		{
			MethodName: "Key",
			Handler:    _AccountService_Key_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/v1/account_service.proto",
}

func init() {
	proto.RegisterFile("accounts/v1/account_service.proto", fileDescriptor_account_service_320cdabe64faa457)
}

var fileDescriptor_account_service_320cdabe64faa457 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x4f, 0xbb, 0xb0, 0x5b, 0x47, 0x10, 0x19, 0x8a, 0x68, 0x2a, 0xb1, 0x06, 0x11, 0x2f, 0x99,
	0x61, 0xf5, 0xe2, 0xb5, 0xad, 0x45, 0x97, 0x1e, 0x56, 0xb6, 0xd0, 0x83, 0x97, 0x32, 0x3b, 0x79,
	0x9b, 0x0c, 0xa4, 0xf3, 0xe2, 0x64, 0x12, 0x8d, 0x47, 0xbf, 0x82, 0x17, 0x0f, 0x7e, 0x20, 0x8f,
	0x82, 0x77, 0x91, 0xe8, 0x07, 0x91, 0x9d, 0x8c, 0xb6, 0x62, 0x5d, 0x16, 0x7a, 0x7b, 0x33, 0xbf,
	0x3f, 0xf3, 0x7b, 0xbf, 0x21, 0xf7, 0x85, 0x94, 0x58, 0x6b, 0x5b, 0xf1, 0x66, 0xcc, 0xfd, 0x7c,
	0x5a, 0x81, 0x69, 0x94, 0x04, 0x56, 0x1a, 0xb4, 0x48, 0xb7, 0x65, 0x81, 0x75, 0xca, 0x44, 0xa9,
	0x98, 0x27, 0xb0, 0x66, 0x1c, 0xde, 0xb9, 0x44, 0xd8, 0x0b, 0x42, 0x9e, 0x29, 0x9b, 0xd7, 0x73,
	0x26, 0xf1, 0x8c, 0x9f, 0xa8, 0x14, 0xf0, 0x00, 0x95, 0xe6, 0xce, 0x25, 0x11, 0xa5, 0xe2, 0xa6,
	0x94, 0x3c, 0x07, 0x51, 0xd8, 0xdc, 0x0b, 0x92, 0x0b, 0x82, 0x0c, 0x33, 0xe4, 0xee, 0x7a, 0x5e,
	0x2f, 0xdc, 0xc9, 0x1d, 0xdc, 0xe4, 0xe9, 0x7b, 0x17, 0xe8, 0xa0, 0x1b, 0x6c, 0x4b, 0x83, 0x6f,
	0xdb, 0x5e, 0x24, 0x93, 0x0c, 0x74, 0xd2, 0x88, 0x42, 0xa5, 0xc2, 0x02, 0xff, 0x67, 0xf0, 0x16,
	0x3b, 0x19, 0x62, 0x56, 0xc0, 0xf9, 0x43, 0x70, 0x56, 0xda, 0xd6, 0x83, 0x77, 0x3d, 0xb8, 0xcc,
	0x2a, 0xb4, 0x46, 0x2b, 0xac, 0x42, 0x5d, 0xf5, 0x68, 0xfc, 0x94, 0xdc, 0xd8, 0xeb, 0xd7, 0x9d,
	0xc1, 0xeb, 0x1a, 0x2a, 0x4b, 0x1f, 0x92, 0x2d, 0x7c, 0xa3, 0xc1, 0x9c, 0xaa, 0xf4, 0xf6, 0xc6,
	0xee, 0xc6, 0xa3, 0x6b, 0xfb, 0xd7, 0xbb, 0x6f, 0xf7, 0x46, 0xd3, 0xe5, 0xdd, 0xe4, 0xd9, 0x6c,
	0xe4, 0xc0, 0x49, 0xfa, 0xf8, 0xd3, 0xe0, 0x8f, 0xf4, 0xb8, 0x6f, 0x98, 0x4e, 0xc9, 0xf0, 0x85,
	0x6b, 0x82, 0xde, 0x62, 0xfd, 0xab, 0xec, 0x77, 0x24, 0x76, 0xb8, 0x8c, 0x14, 0xee, 0xb0, 0xf3,
	0xfa, 0x4d, 0x29, 0x59, 0x4f, 0x3f, 0xb6, 0xc2, 0xd6, 0x55, 0x7c, 0xf3, 0xfd, 0xd7, 0x9f, 0x1f,
	0x36, 0x09, 0xdd, 0xf2, 0x7d, 0xbe, 0xa3, 0x27, 0x64, 0x78, 0x60, 0x40, 0x58, 0xa0, 0x0f, 0xd8,
	0x65, 0xff, 0xc6, 0xfe, 0xce, 0x1e, 0xae, 0x66, 0xbd, 0x34, 0xb8, 0x50, 0x05, 0xc4, 0x01, 0x3d,
	0x24, 0x83, 0xe7, 0x60, 0xff, 0x9b, 0x72, 0x5d, 0x9b, 0x09, 0x19, 0xcd, 0x60, 0x61, 0xa0, 0xca,
	0xaf, 0x6c, 0x35, 0x25, 0x83, 0x23, 0x68, 0xd7, 0x5c, 0x73, 0x77, 0x25, 0xeb, 0x08, 0xda, 0x38,
	0xd8, 0xdf, 0xfe, 0xdc, 0x45, 0xc1, 0x97, 0x2e, 0x0a, 0xbe, 0x77, 0x51, 0xf0, 0xf1, 0x47, 0x14,
	0xbc, 0xda, 0x6c, 0xc6, 0xf3, 0xa1, 0x8b, 0xf7, 0xe4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03,
	0xaf, 0x12, 0x71, 0x29, 0x03, 0x00, 0x00,
}
