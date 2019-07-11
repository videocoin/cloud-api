// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: miners/v1/miner_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	rpc "github.com/videocoin/cloud-api/rpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a3d7c5a419ed72c, []int{0}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (*CreateResponse) XXX_MessageName() string {
	return "cloud.api.miners.v1.CreateResponse"
}

type Request struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hash                 string      `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Status               MinerStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cloud.api.miners.v1.MinerStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a3d7c5a419ed72c, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Request) GetStatus() MinerStatus {
	if m != nil {
		return m.Status
	}
	return MinerStatusOffline
}

func (*Request) XXX_MessageName() string {
	return "cloud.api.miners.v1.Request"
}

type Response struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               MinerStatus `protobuf:"varint,2,opt,name=status,proto3,enum=cloud.api.miners.v1.MinerStatus" json:"status,omitempty"`
	CpuIdle              int64       `protobuf:"varint,3,opt,name=cpu_idle,json=cpuIdle,proto3" json:"cpu_idle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a3d7c5a419ed72c, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response) GetStatus() MinerStatus {
	if m != nil {
		return m.Status
	}
	return MinerStatusOffline
}

func (m *Response) GetCpuIdle() int64 {
	if m != nil {
		return m.CpuIdle
	}
	return 0
}

func (*Response) XXX_MessageName() string {
	return "cloud.api.miners.v1.Response"
}

type ListRequest struct {
	Status               MinerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=cloud.api.miners.v1.MinerStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a3d7c5a419ed72c, []int{3}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetStatus() MinerStatus {
	if m != nil {
		return m.Status
	}
	return MinerStatusOffline
}

func (*ListRequest) XXX_MessageName() string {
	return "cloud.api.miners.v1.ListRequest"
}

type ListResponse struct {
	Items                []*Miner `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a3d7c5a419ed72c, []int{4}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetItems() []*Miner {
	if m != nil {
		return m.Items
	}
	return nil
}

func (*ListResponse) XXX_MessageName() string {
	return "cloud.api.miners.v1.ListResponse"
}
func init() {
	proto.RegisterType((*CreateResponse)(nil), "cloud.api.miners.v1.CreateResponse")
	golang_proto.RegisterType((*CreateResponse)(nil), "cloud.api.miners.v1.CreateResponse")
	proto.RegisterType((*Request)(nil), "cloud.api.miners.v1.Request")
	golang_proto.RegisterType((*Request)(nil), "cloud.api.miners.v1.Request")
	proto.RegisterType((*Response)(nil), "cloud.api.miners.v1.Response")
	golang_proto.RegisterType((*Response)(nil), "cloud.api.miners.v1.Response")
	proto.RegisterType((*ListRequest)(nil), "cloud.api.miners.v1.ListRequest")
	golang_proto.RegisterType((*ListRequest)(nil), "cloud.api.miners.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "cloud.api.miners.v1.ListResponse")
	golang_proto.RegisterType((*ListResponse)(nil), "cloud.api.miners.v1.ListResponse")
}

func init() { proto.RegisterFile("miners/v1/miner_service.proto", fileDescriptor_1a3d7c5a419ed72c) }
func init() {
	golang_proto.RegisterFile("miners/v1/miner_service.proto", fileDescriptor_1a3d7c5a419ed72c)
}

var fileDescriptor_1a3d7c5a419ed72c = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0x6e, 0xd2, 0xfd, 0xba, 0xfe, 0xbc, 0x51, 0x55, 0xae, 0x36, 0x75, 0xe9, 0x16, 0x95, 0x20,
	0xa1, 0x0a, 0x69, 0x36, 0x2d, 0x17, 0xc4, 0x09, 0x81, 0xd0, 0x40, 0x62, 0x42, 0xca, 0x04, 0x07,
	0x2e, 0x25, 0x4d, 0x4c, 0x6a, 0x91, 0xc6, 0x26, 0x76, 0x82, 0x06, 0xda, 0x85, 0xaf, 0xc0, 0x17,
	0xe2, 0xd8, 0x23, 0x12, 0x5f, 0x00, 0x75, 0x7c, 0x10, 0x64, 0xc7, 0x55, 0xbb, 0xd2, 0x22, 0x4d,
	0xbb, 0xbd, 0xf6, 0xfb, 0xf8, 0x79, 0xde, 0x3f, 0x8f, 0xc1, 0xd1, 0x84, 0xa6, 0x24, 0x13, 0xb8,
	0xe8, 0x63, 0x1d, 0x0d, 0x05, 0xc9, 0x0a, 0x1a, 0x12, 0xc4, 0x33, 0x26, 0x19, 0x6c, 0x85, 0x09,
	0xcb, 0x23, 0x14, 0x70, 0x8a, 0x4a, 0x20, 0x2a, 0xfa, 0xce, 0xde, 0xca, 0x9b, 0x12, 0xeb, 0xe0,
	0x98, 0xca, 0x71, 0x3e, 0x42, 0x21, 0x9b, 0xe0, 0x82, 0x46, 0x84, 0x85, 0x8c, 0xa6, 0x58, 0x13,
	0x1c, 0x07, 0x9c, 0xe2, 0x8c, 0x87, 0x78, 0x4c, 0x82, 0x44, 0x8e, 0xcd, 0x83, 0x4e, 0xcc, 0x58,
	0x9c, 0x10, 0xac, 0x4f, 0xa3, 0xfc, 0x3d, 0x26, 0x13, 0x2e, 0xcf, 0x4d, 0xf2, 0xd0, 0x24, 0xd5,
	0xcb, 0x20, 0x4d, 0x99, 0x0c, 0x24, 0x65, 0xa9, 0x30, 0xd9, 0xe3, 0x25, 0xad, 0x98, 0xc5, 0x6c,
	0xc1, 0xa1, 0x4e, 0xfa, 0xa0, 0x23, 0x03, 0x77, 0x57, 0x95, 0x3e, 0x65, 0x01, 0xe7, 0xaa, 0x1d,
	0x7d, 0xe3, 0x0d, 0x40, 0xe3, 0x69, 0x46, 0x02, 0x49, 0x7c, 0x22, 0x38, 0x4b, 0x05, 0x81, 0x0d,
	0x60, 0xd3, 0xa8, 0x6d, 0x75, 0xad, 0xde, 0xff, 0xbe, 0x4d, 0x23, 0xd8, 0x04, 0xd5, 0x0f, 0xe4,
	0xbc, 0x6d, 0xeb, 0x0b, 0x15, 0x7a, 0x31, 0xd8, 0xf6, 0xc9, 0xc7, 0x9c, 0x08, 0xf9, 0x17, 0x18,
	0x82, 0xad, 0x71, 0x20, 0xc6, 0x06, 0xad, 0x63, 0xf8, 0x10, 0xd4, 0x84, 0x0c, 0x64, 0x2e, 0xda,
	0xd5, 0xae, 0xd5, 0x6b, 0x0c, 0xba, 0x68, 0xcd, 0x68, 0xd1, 0xa9, 0x8a, 0xce, 0x34, 0xce, 0x37,
	0x78, 0x8f, 0x81, 0xfa, 0xc6, 0xb2, 0x16, 0xac, 0xf6, 0xf5, 0x58, 0xe1, 0x01, 0xa8, 0x87, 0x3c,
	0x1f, 0xd2, 0x28, 0x21, 0xba, 0xa2, 0xaa, 0xbf, 0x1d, 0xf2, 0xfc, 0x45, 0x94, 0x10, 0xef, 0x04,
	0xec, 0xbc, 0xa4, 0x42, 0xce, 0xbb, 0x5b, 0x68, 0x58, 0xd7, 0xac, 0xfc, 0x31, 0xd8, 0x2d, 0x89,
	0x4c, 0xf5, 0xf7, 0xc1, 0x7f, 0x54, 0x92, 0x89, 0x22, 0xaa, 0xf6, 0x76, 0x06, 0xce, 0x66, 0x22,
	0xbf, 0x04, 0x0e, 0xa6, 0x5b, 0xe0, 0x96, 0xbe, 0x10, 0x67, 0xa5, 0x2f, 0xe1, 0x2b, 0x50, 0x7b,
	0xae, 0x4d, 0x04, 0xf7, 0x51, 0xb9, 0x55, 0x34, 0xdf, 0x2a, 0x7a, 0xa6, 0xfc, 0xe3, 0x74, 0x96,
	0x68, 0x33, 0x1e, 0xa2, 0x12, 0x5e, 0x96, 0xe6, 0x35, 0xbf, 0xfe, 0xfc, 0xfd, 0xcd, 0x06, 0xb0,
	0x6e, 0xac, 0xf8, 0x19, 0xbe, 0x03, 0xb5, 0x72, 0xf7, 0x1b, 0x09, 0xef, 0xac, 0xad, 0xf3, 0xaa,
	0x61, 0xbc, 0x03, 0x4d, 0xdc, 0xf2, 0x1a, 0xda, 0xb1, 0xf3, 0xaf, 0x21, 0x1e, 0x59, 0xf7, 0xe0,
	0x10, 0x54, 0x4f, 0x88, 0x84, 0x87, 0x6b, 0x69, 0xcc, 0x94, 0x9d, 0xa3, 0x0d, 0x59, 0x43, 0xdf,
	0xd1, 0xf4, 0x7b, 0xb0, 0x75, 0x95, 0x1e, 0x7f, 0xa1, 0xd1, 0x05, 0x64, 0xa0, 0xfe, 0x26, 0x48,
	0x68, 0xa4, 0x9a, 0xb8, 0x91, 0xca, 0x5d, 0xad, 0xd2, 0x85, 0xee, 0xaa, 0x8a, 0xb2, 0xf0, 0x05,
	0x2e, 0xe6, 0x22, 0x21, 0xd8, 0x52, 0x8b, 0x85, 0xeb, 0xad, 0xb0, 0x64, 0x1e, 0xe7, 0xf6, 0x3f,
	0x10, 0x46, 0x74, 0x5f, 0x8b, 0x36, 0xe1, 0xca, 0xe4, 0xe0, 0x29, 0xd8, 0x7d, 0xcd, 0x95, 0x5c,
	0xb9, 0xba, 0x9b, 0x75, 0x56, 0x79, 0xd2, 0x9e, 0xce, 0xdc, 0xca, 0x8f, 0x99, 0x5b, 0xf9, 0x35,
	0x73, 0x2b, 0xdf, 0x2f, 0x5d, 0x6b, 0x7a, 0xe9, 0x5a, 0x6f, 0xed, 0xa2, 0x3f, 0xaa, 0xe9, 0x7d,
	0x3f, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe2, 0xed, 0x4d, 0x0c, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MinersServiceClient is the client API for MinersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MinersServiceClient interface {
	Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
	Create(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Validate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	UpdateStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type minersServiceClient struct {
	cc *grpc.ClientConn
}

func NewMinersServiceClient(cc *grpc.ClientConn) MinersServiceClient {
	return &minersServiceClient{cc}
}

func (c *minersServiceClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.miners.v1.MinersService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minersServiceClient) Create(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.miners.v1.MinersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minersServiceClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cloud.api.miners.v1.MinersService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minersServiceClient) Validate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cloud.api.miners.v1.MinersService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minersServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.miners.v1.MinersService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minersServiceClient) UpdateStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cloud.api.miners.v1.MinersService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MinersServiceServer is the server API for MinersService service.
type MinersServiceServer interface {
	Health(context.Context, *types.Empty) (*rpc.HealthStatus, error)
	Create(context.Context, *types.Empty) (*CreateResponse, error)
	Get(context.Context, *Request) (*Response, error)
	Validate(context.Context, *Request) (*Response, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	UpdateStatus(context.Context, *Request) (*Response, error)
}

// UnimplementedMinersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMinersServiceServer struct {
}

func (*UnimplementedMinersServiceServer) Health(ctx context.Context, req *types.Empty) (*rpc.HealthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedMinersServiceServer) Create(ctx context.Context, req *types.Empty) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedMinersServiceServer) Get(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedMinersServiceServer) Validate(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (*UnimplementedMinersServiceServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedMinersServiceServer) UpdateStatus(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}

func RegisterMinersServiceServer(s *grpc.Server, srv MinersServiceServer) {
	s.RegisterService(&_MinersService_serviceDesc, srv)
}

func _MinersService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinersServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.miners.v1.MinersService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinersServiceServer).Health(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.miners.v1.MinersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinersServiceServer).Create(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinersService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinersServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.miners.v1.MinersService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinersServiceServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinersService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinersServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.miners.v1.MinersService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinersServiceServer).Validate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinersService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinersServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.miners.v1.MinersService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinersServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinersService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinersServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.miners.v1.MinersService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinersServiceServer).UpdateStatus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MinersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.miners.v1.MinersService",
	HandlerType: (*MinersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _MinersService_Health_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MinersService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MinersService_Get_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _MinersService_Validate_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MinersService_List_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _MinersService_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "miners/v1/miner_service.proto",
}
