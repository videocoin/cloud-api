// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accounts/v1/account_service.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import rpc "github.com/VideoCoin/cloud-api/rpc"
import _ "github.com/gogo/protobuf/gogoproto"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateAccountRequest struct {
	OwnerID              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_service_147e4cfec372a6aa, []int{0}
}
func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (dst *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(dst, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetOwnerID() string {
	if m != nil {
		return m.OwnerID
	}
	return ""
}

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
	return fileDescriptor_account_service_147e4cfec372a6aa, []int{1}
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

type UpdateBalanceRequest struct {
	TransactionHash      string   `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBalanceRequest) Reset()         { *m = UpdateBalanceRequest{} }
func (m *UpdateBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBalanceRequest) ProtoMessage()    {}
func (*UpdateBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_service_147e4cfec372a6aa, []int{2}
}
func (m *UpdateBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBalanceRequest.Unmarshal(m, b)
}
func (m *UpdateBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBalanceRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBalanceRequest.Merge(dst, src)
}
func (m *UpdateBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBalanceRequest.Size(m)
}
func (m *UpdateBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBalanceRequest proto.InternalMessageInfo

func (m *UpdateBalanceRequest) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

type TxListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxListRequest) Reset()         { *m = TxListRequest{} }
func (m *TxListRequest) String() string { return proto.CompactTextString(m) }
func (*TxListRequest) ProtoMessage()    {}
func (*TxListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_service_147e4cfec372a6aa, []int{3}
}
func (m *TxListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxListRequest.Unmarshal(m, b)
}
func (m *TxListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxListRequest.Marshal(b, m, deterministic)
}
func (dst *TxListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxListRequest.Merge(dst, src)
}
func (m *TxListRequest) XXX_Size() int {
	return xxx_messageInfo_TxListRequest.Size(m)
}
func (m *TxListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxListRequest proto.InternalMessageInfo

type TxListResponse struct {
	Items                []*TxResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TxListResponse) Reset()         { *m = TxListResponse{} }
func (m *TxListResponse) String() string { return proto.CompactTextString(m) }
func (*TxListResponse) ProtoMessage()    {}
func (*TxListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_service_147e4cfec372a6aa, []int{4}
}
func (m *TxListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxListResponse.Unmarshal(m, b)
}
func (m *TxListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxListResponse.Marshal(b, m, deterministic)
}
func (dst *TxListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxListResponse.Merge(dst, src)
}
func (m *TxListResponse) XXX_Size() int {
	return xxx_messageInfo_TxListResponse.Size(m)
}
func (m *TxListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxListResponse proto.InternalMessageInfo

func (m *TxListResponse) GetItems() []*TxResponse {
	if m != nil {
		return m.Items
	}
	return nil
}

type TxResponse struct {
	Type                 string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	EthTxHash            string               `protobuf:"bytes,3,opt,name=eth_tx_hash,json=ethTxHash,proto3" json:"eth_tx_hash,omitempty"`
	AmountEth            float64              `protobuf:"fixed64,4,opt,name=amount_eth,json=amountEth,proto3" json:"amount_eth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TxResponse) Reset()         { *m = TxResponse{} }
func (m *TxResponse) String() string { return proto.CompactTextString(m) }
func (*TxResponse) ProtoMessage()    {}
func (*TxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_service_147e4cfec372a6aa, []int{5}
}
func (m *TxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxResponse.Unmarshal(m, b)
}
func (m *TxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxResponse.Marshal(b, m, deterministic)
}
func (dst *TxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResponse.Merge(dst, src)
}
func (m *TxResponse) XXX_Size() int {
	return xxx_messageInfo_TxResponse.Size(m)
}
func (m *TxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxResponse proto.InternalMessageInfo

func (m *TxResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TxResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TxResponse) GetEthTxHash() string {
	if m != nil {
		return m.EthTxHash
	}
	return ""
}

func (m *TxResponse) GetAmountEth() float64 {
	if m != nil {
		return m.AmountEth
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateAccountRequest)(nil), "cloud.api.account.v1.CreateAccountRequest")
	proto.RegisterType((*AccountRequest)(nil), "cloud.api.account.v1.AccountRequest")
	proto.RegisterType((*UpdateBalanceRequest)(nil), "cloud.api.account.v1.UpdateBalanceRequest")
	proto.RegisterType((*TxListRequest)(nil), "cloud.api.account.v1.TxListRequest")
	proto.RegisterType((*TxListResponse)(nil), "cloud.api.account.v1.TxListResponse")
	proto.RegisterType((*TxResponse)(nil), "cloud.api.account.v1.TxResponse")
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
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*AccountProfile, error)
	GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountProfile, error)
	UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*AccountProfile, error)
	GetTxList(ctx context.Context, in *TxListRequest, opts ...grpc.CallOption) (*TxListResponse, error)
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

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*AccountProfile, error) {
	out := new(AccountProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountProfile, error) {
	out := new(AccountProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*AccountProfile, error) {
	out := new(AccountProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/UpdateBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetTxList(ctx context.Context, in *TxListRequest, opts ...grpc.CallOption) (*TxListResponse, error) {
	out := new(TxListResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.account.v1.AccountService/GetTxList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	Health(context.Context, *empty.Empty) (*rpc.HealthStatus, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*AccountProfile, error)
	GetAccount(context.Context, *AccountRequest) (*AccountProfile, error)
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*AccountProfile, error)
	GetTxList(context.Context, *TxListRequest) (*TxListResponse, error)
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

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/UpdateBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateBalance(ctx, req.(*UpdateBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetTxList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetTxList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.account.v1.AccountService/GetTxList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetTxList(ctx, req.(*TxListRequest))
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
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AccountService_GetAccount_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _AccountService_UpdateBalance_Handler,
		},
		{
			MethodName: "GetTxList",
			Handler:    _AccountService_GetTxList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/v1/account_service.proto",
}

func init() {
	proto.RegisterFile("accounts/v1/account_service.proto", fileDescriptor_account_service_147e4cfec372a6aa)
}

var fileDescriptor_account_service_147e4cfec372a6aa = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe3, 0xb6, 0xb4, 0xcd, 0x44, 0xfd, 0xd0, 0x12, 0x20, 0xb8, 0x25, 0x0d, 0x06, 0xa1,
	0x50, 0xa9, 0xb6, 0x5a, 0x24, 0x04, 0x1c, 0x90, 0xda, 0x52, 0xb5, 0x95, 0x90, 0x8a, 0xdc, 0xc0,
	0xa1, 0x97, 0x68, 0x63, 0x6f, 0xe3, 0x95, 0x62, 0xaf, 0xf1, 0x8e, 0x43, 0xca, 0x11, 0x89, 0x27,
	0xe0, 0x82, 0xb8, 0xf0, 0x3a, 0x1c, 0x91, 0xb8, 0x23, 0x14, 0x78, 0x10, 0x94, 0x5d, 0x9b, 0xa6,
	0xc5, 0x45, 0xc0, 0x6d, 0x3d, 0xf3, 0x9f, 0xef, 0x9f, 0xe1, 0x26, 0xf5, 0x3c, 0x91, 0x46, 0x28,
	0x9d, 0xfe, 0xba, 0x93, 0xbd, 0xdb, 0x92, 0x25, 0x7d, 0xee, 0x31, 0x3b, 0x4e, 0x04, 0x0a, 0x52,
	0xf5, 0x7a, 0x22, 0xf5, 0x6d, 0x1a, 0x73, 0x3b, 0x13, 0xd8, 0xfd, 0x75, 0x73, 0xad, 0xcb, 0x31,
	0x48, 0x3b, 0xb6, 0x27, 0x42, 0xa7, 0x2b, 0xba, 0xc2, 0x51, 0xe2, 0x4e, 0x7a, 0xac, 0xbe, 0xd4,
	0x87, 0x7a, 0xe9, 0x24, 0xe6, 0x52, 0x57, 0x88, 0x6e, 0x8f, 0x9d, 0xaa, 0x58, 0x18, 0xe3, 0x49,
	0xe6, 0x5c, 0x39, 0xef, 0x44, 0x1e, 0x32, 0x89, 0x34, 0x8c, 0x33, 0xc1, 0x72, 0x26, 0xa0, 0x31,
	0x77, 0x68, 0x14, 0x09, 0xa4, 0xc8, 0x45, 0x24, 0x33, 0xaf, 0x33, 0xd6, 0xca, 0x0b, 0xee, 0x33,
	0xb1, 0x2d, 0x78, 0xe4, 0xa8, 0xae, 0xd7, 0x46, 0x11, 0x49, 0xec, 0x39, 0x01, 0xa3, 0x3d, 0x0c,
	0xb2, 0x80, 0xeb, 0x05, 0x43, 0x6b, 0x97, 0xf5, 0x18, 0xaa, 0xdb, 0x09, 0xa3, 0xc8, 0x36, 0xb5,
	0xd9, 0x65, 0x2f, 0x53, 0x26, 0x91, 0xdc, 0x81, 0x59, 0xf1, 0x2a, 0x62, 0x49, 0x9b, 0xfb, 0x35,
	0xa3, 0x61, 0x34, 0xcb, 0x5b, 0x95, 0xe1, 0xd7, 0x95, 0x99, 0x83, 0x91, 0x6d, 0xff, 0x89, 0x3b,
	0xa3, 0x9c, 0xfb, 0xbe, 0xf5, 0x00, 0xe6, 0xff, 0x33, 0x72, 0x13, 0xaa, 0xcf, 0x63, 0x9f, 0x22,
	0xdb, 0xa2, 0x3d, 0x1a, 0x79, 0x2c, 0x8f, 0xbf, 0x0b, 0x8b, 0x98, 0xd0, 0x48, 0x52, 0x6f, 0x34,
	0x73, 0x3b, 0xa0, 0x32, 0xd0, 0x79, 0xdc, 0x85, 0x31, 0xfb, 0x1e, 0x95, 0x81, 0xb5, 0x00, 0x73,
	0xad, 0xc1, 0x53, 0x2e, 0xf3, 0xda, 0xd6, 0x1e, 0xcc, 0xe7, 0x06, 0x19, 0x8b, 0x48, 0x32, 0x72,
	0x1f, 0x2e, 0x71, 0x64, 0xa1, 0xac, 0x19, 0x8d, 0xc9, 0x66, 0x65, 0xa3, 0x61, 0x17, 0x1d, 0xd7,
	0x6e, 0x0d, 0xf2, 0x00, 0x57, 0xcb, 0xad, 0x0f, 0x06, 0xc0, 0xa9, 0x95, 0x10, 0x98, 0xc2, 0x93,
	0x98, 0x65, 0x8d, 0xa8, 0x37, 0x79, 0x08, 0xe0, 0xa9, 0xd5, 0xf9, 0x6d, 0x8a, 0xb5, 0x89, 0x86,
	0xd1, 0xac, 0x6c, 0x98, 0xb6, 0xbe, 0x9c, 0x9d, 0x9f, 0xd6, 0x6e, 0xe5, 0xa7, 0x75, 0xcb, 0x99,
	0x7a, 0x13, 0x49, 0x1d, 0x2a, 0x0c, 0x83, 0x36, 0x0e, 0xf4, 0x78, 0x93, 0x2a, 0x6b, 0x99, 0x61,
	0xd0, 0x1a, 0x8c, 0x06, 0x23, 0x37, 0x00, 0x68, 0xa8, 0xd0, 0x64, 0x18, 0xd4, 0xa6, 0x1a, 0x46,
	0xd3, 0x70, 0xcb, 0xda, 0xb2, 0x83, 0xc1, 0xc6, 0xc7, 0xa9, 0x5f, 0x5b, 0x3f, 0xd4, 0xe8, 0x92,
	0x03, 0x98, 0xde, 0x53, 0x27, 0x27, 0x57, 0x7f, 0x6b, 0x61, 0x67, 0x84, 0x9e, 0xb9, 0x34, 0x36,
	0x7a, 0x12, 0x7b, 0xb6, 0x96, 0x1f, 0x22, 0xc5, 0x54, 0x5a, 0x8b, 0x6f, 0xbe, 0xfc, 0x78, 0x37,
	0x01, 0x64, 0x36, 0x03, 0xe7, 0x35, 0xf1, 0x60, 0xee, 0x0c, 0x18, 0x64, 0xb5, 0x78, 0x75, 0x45,
	0xf4, 0x98, 0xb7, 0x8b, 0xb5, 0x99, 0xea, 0x59, 0x22, 0x8e, 0x79, 0x8f, 0x59, 0x25, 0x72, 0x04,
	0xb0, 0xcb, 0x30, 0xaf, 0xf0, 0xe7, 0xa8, 0x7f, 0xcd, 0xfd, 0xd6, 0x80, 0xb9, 0x33, 0x80, 0x5d,
	0x34, 0x41, 0x11, 0x85, 0x7f, 0x59, 0xc5, 0x52, 0x6b, 0x5b, 0x36, 0xaf, 0xa9, 0x3f, 0xf5, 0xf4,
	0xe7, 0x72, 0x3a, 0x3a, 0xdb, 0x23, 0x63, 0x95, 0x48, 0x28, 0xef, 0x32, 0xd4, 0x58, 0x92, 0x5b,
	0x17, 0xf1, 0x37, 0x46, 0xf1, 0x45, 0xb5, 0xcf, 0x92, 0x6d, 0x2d, 0xa9, 0xda, 0x57, 0xc8, 0xe5,
	0xf3, 0xb5, 0x71, 0x20, 0xb7, 0xaa, 0x9f, 0x86, 0xf5, 0xd2, 0xe7, 0x61, 0xbd, 0xf4, 0x6d, 0x58,
	0x2f, 0xbd, 0xff, 0x5e, 0x2f, 0x1d, 0x4d, 0xf4, 0xd7, 0x3b, 0xd3, 0x0a, 0x89, 0x7b, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x8c, 0xed, 0x4e, 0x66, 0x05, 0x05, 0x00, 0x00,
}
