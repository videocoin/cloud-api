// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dispatcher/v1/dispatcher_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	v12 "github.com/videocoin/cloud-api/miners/v1"
	v11 "github.com/videocoin/cloud-api/syncer/v1"
	v1 "github.com/videocoin/cloud-api/validator/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TaskPendingRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskPendingRequest) Reset()         { *m = TaskPendingRequest{} }
func (m *TaskPendingRequest) String() string { return proto.CompactTextString(m) }
func (*TaskPendingRequest) ProtoMessage()    {}
func (*TaskPendingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_008866cbcf1626a3, []int{0}
}
func (m *TaskPendingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskPendingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskPendingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskPendingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskPendingRequest.Merge(m, src)
}
func (m *TaskPendingRequest) XXX_Size() int {
	return m.Size()
}
func (m *TaskPendingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskPendingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskPendingRequest proto.InternalMessageInfo

func (m *TaskPendingRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (*TaskPendingRequest) XXX_MessageName() string {
	return "cloud.api.dispatcher.v1.TaskPendingRequest"
}

type TaskRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_008866cbcf1626a3, []int{1}
}
func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return m.Size()
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *TaskRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (*TaskRequest) XXX_MessageName() string {
	return "cloud.api.dispatcher.v1.TaskRequest"
}

type RegistrationRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationRequest) Reset()         { *m = RegistrationRequest{} }
func (m *RegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*RegistrationRequest) ProtoMessage()    {}
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_008866cbcf1626a3, []int{2}
}
func (m *RegistrationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegistrationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationRequest.Merge(m, src)
}
func (m *RegistrationRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationRequest proto.InternalMessageInfo

func (m *RegistrationRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (*RegistrationRequest) XXX_MessageName() string {
	return "cloud.api.dispatcher.v1.RegistrationRequest"
}
func init() {
	proto.RegisterType((*TaskPendingRequest)(nil), "cloud.api.dispatcher.v1.TaskPendingRequest")
	golang_proto.RegisterType((*TaskPendingRequest)(nil), "cloud.api.dispatcher.v1.TaskPendingRequest")
	proto.RegisterType((*TaskRequest)(nil), "cloud.api.dispatcher.v1.TaskRequest")
	golang_proto.RegisterType((*TaskRequest)(nil), "cloud.api.dispatcher.v1.TaskRequest")
	proto.RegisterType((*RegistrationRequest)(nil), "cloud.api.dispatcher.v1.RegistrationRequest")
	golang_proto.RegisterType((*RegistrationRequest)(nil), "cloud.api.dispatcher.v1.RegistrationRequest")
}

func init() {
	proto.RegisterFile("dispatcher/v1/dispatcher_service.proto", fileDescriptor_008866cbcf1626a3)
}
func init() {
	golang_proto.RegisterFile("dispatcher/v1/dispatcher_service.proto", fileDescriptor_008866cbcf1626a3)
}

var fileDescriptor_008866cbcf1626a3 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x9b, 0x6a, 0x8c, 0xce, 0xfc, 0x11, 0x78, 0xd2, 0xa8, 0x0a, 0x64, 0x23, 0x42, 0x08,
	0x04, 0x73, 0x54, 0xb8, 0x05, 0x01, 0x6d, 0x61, 0xea, 0xc5, 0xa4, 0x28, 0x43, 0x20, 0x10, 0x62,
	0x72, 0x63, 0x2f, 0xb3, 0x9a, 0xda, 0x21, 0x76, 0x23, 0xf5, 0xe5, 0x10, 0x97, 0xbb, 0xe4, 0x09,
	0x26, 0x94, 0xbd, 0x08, 0xb2, 0x9d, 0xb5, 0x0d, 0xb0, 0x76, 0x93, 0x76, 0xe7, 0x93, 0xf3, 0x7d,
	0xbf, 0x9c, 0x63, 0x9f, 0x03, 0x1e, 0x11, 0x26, 0x53, 0xac, 0xa2, 0x43, 0x9a, 0xf9, 0x79, 0xdb,
	0x9f, 0x45, 0xfb, 0x92, 0x66, 0x39, 0x8b, 0x28, 0x4a, 0x33, 0xa1, 0x04, 0xbc, 0x13, 0x25, 0x62,
	0x4c, 0x10, 0x4e, 0x19, 0x9a, 0x69, 0x50, 0xde, 0x6e, 0xdd, 0x8d, 0x85, 0x88, 0x13, 0xea, 0x1b,
	0xd9, 0x60, 0x7c, 0xe0, 0xd3, 0x51, 0xaa, 0x26, 0xd6, 0xd5, 0xba, 0x57, 0x26, 0x71, 0xca, 0x7c,
	0xcc, 0xb9, 0x50, 0x58, 0x31, 0xc1, 0x65, 0x99, 0xdd, 0x8e, 0x99, 0x3a, 0x1c, 0x0f, 0x50, 0x24,
	0x46, 0x7e, 0x2c, 0x62, 0x31, 0x63, 0xe8, 0xc8, 0x04, 0xe6, 0x54, 0xca, 0x9b, 0xd5, 0x52, 0x15,
	0x96, 0xc3, 0x32, 0xd3, 0x9d, 0x03, 0xe5, 0x8c, 0x50, 0x11, 0x09, 0xc6, 0x7d, 0x53, 0xf1, 0xb6,
	0xfe, 0x75, 0x8e, 0x13, 0x46, 0xb0, 0x12, 0xc6, 0x37, 0x0d, 0xaa, 0x1d, 0xb6, 0x5e, 0x2d, 0x81,
	0xc8, 0x09, 0x8f, 0xec, 0x9f, 0xed, 0xe9, 0x2f, 0xfb, 0xcb, 0x25, 0xf6, 0x11, 0xe3, 0x34, 0x93,
	0xda, 0x6e, 0x4e, 0x55, 0xb7, 0xf7, 0x1a, 0xc0, 0x0f, 0x58, 0x0e, 0x03, 0xca, 0x09, 0xe3, 0x71,
	0x48, 0xbf, 0x8f, 0xa9, 0x54, 0xf0, 0x09, 0x58, 0x8b, 0x12, 0x46, 0xb9, 0xda, 0x67, 0xa4, 0xe9,
	0x6c, 0x39, 0x8f, 0xd7, 0x3a, 0xd7, 0x8b, 0xe3, 0xcd, 0x46, 0xd7, 0x7c, 0xec, 0xf7, 0xc2, 0x86,
	0x4d, 0xf7, 0x89, 0x17, 0x80, 0x6b, 0x1a, 0x70, 0x71, 0x27, 0xdc, 0x00, 0x75, 0x46, 0x9a, 0x75,
	0xa3, 0x59, 0x2d, 0x8e, 0x37, 0xeb, 0xfd, 0x5e, 0x58, 0x67, 0xc4, 0x7b, 0x03, 0xd6, 0x43, 0x1a,
	0x33, 0xa9, 0x32, 0xf3, 0x68, 0x17, 0x27, 0x3f, 0xff, 0x71, 0x05, 0xdc, 0xee, 0x4d, 0xdf, 0x6c,
	0xcf, 0x36, 0x0c, 0xbf, 0x81, 0x9b, 0x3b, 0x54, 0x95, 0x9d, 0xea, 0x9a, 0xe1, 0x53, 0x74, 0xc6,
	0x70, 0xa1, 0x7f, 0xef, 0xa4, 0x75, 0x7f, 0xa1, 0xd8, 0xab, 0xc1, 0x00, 0x5c, 0xdd, 0xa1, 0xca,
	0x80, 0x1f, 0x2e, 0xd4, 0x9e, 0x9b, 0xf8, 0x15, 0xac, 0xef, 0xe2, 0x6c, 0xa8, 0xa3, 0xb7, 0xb2,
	0x2b, 0x46, 0x69, 0x42, 0x15, 0x25, 0x97, 0x45, 0xff, 0x0c, 0x6e, 0xcd, 0xe8, 0xef, 0x31, 0x4b,
	0x2e, 0x0f, 0xfd, 0x09, 0xdc, 0xf8, 0x68, 0xa7, 0x9d, 0x06, 0x99, 0x10, 0x07, 0xf0, 0xd9, 0x9c,
	0x63, 0xba, 0x07, 0xda, 0x50, 0x91, 0x9d, 0xf2, 0x37, 0x90, 0x5d, 0x5f, 0x74, 0xba, 0x97, 0xe8,
	0x9d, 0xde, 0x6d, 0xaf, 0x06, 0x3b, 0x60, 0x65, 0x6f, 0xc2, 0x23, 0xb8, 0x35, 0xc7, 0xb3, 0x5b,
	0xa1, 0x61, 0x3a, 0xb5, 0x9c, 0xb1, 0x0b, 0x56, 0x02, 0xc6, 0xe3, 0x0a, 0xc3, 0x2e, 0x89, 0x66,
	0x04, 0x73, 0x4f, 0xfe, 0x60, 0x81, 0x42, 0xa6, 0x82, 0x4b, 0xea, 0xd5, 0x60, 0x08, 0x1a, 0x76,
	0x5c, 0x69, 0x56, 0x69, 0xb3, 0x7a, 0x31, 0xff, 0x99, 0xe8, 0xb3, 0x4b, 0xec, 0x34, 0x8f, 0x0a,
	0xd7, 0xf9, 0x55, 0xb8, 0xce, 0xef, 0xc2, 0x75, 0x7e, 0x9e, 0xb8, 0xce, 0xd1, 0x89, 0xeb, 0x7c,
	0xa9, 0xe7, 0xed, 0xc1, 0xaa, 0xd1, 0xbe, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xa2, 0xc9,
	0xd2, 0x3f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DispatcherServiceClient is the client API for DispatcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DispatcherServiceClient interface {
	GetPendingTask(ctx context.Context, in *TaskPendingRequest, opts ...grpc.CallOption) (*Task, error)
	GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error)
	MarkTaskAsCompleted(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error)
	MarkTaskAsFailed(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error)
	ValidateProof(ctx context.Context, in *v1.ValidateProofRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Sync(ctx context.Context, in *v11.SyncRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Ping(ctx context.Context, in *v12.PingRequest, opts ...grpc.CallOption) (*v12.PingResponse, error)
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type dispatcherServiceClient struct {
	cc *grpc.ClientConn
}

func NewDispatcherServiceClient(cc *grpc.ClientConn) DispatcherServiceClient {
	return &dispatcherServiceClient{cc}
}

func (c *dispatcherServiceClient) GetPendingTask(ctx context.Context, in *TaskPendingRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/GetPendingTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) MarkTaskAsCompleted(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/MarkTaskAsCompleted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) MarkTaskAsFailed(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/MarkTaskAsFailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) ValidateProof(ctx context.Context, in *v1.ValidateProofRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/ValidateProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) Sync(ctx context.Context, in *v11.SyncRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) Ping(ctx context.Context, in *v12.PingRequest, opts ...grpc.CallOption) (*v12.PingResponse, error) {
	out := new(v12.PingResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServiceServer is the server API for DispatcherService service.
type DispatcherServiceServer interface {
	GetPendingTask(context.Context, *TaskPendingRequest) (*Task, error)
	GetTask(context.Context, *TaskRequest) (*Task, error)
	MarkTaskAsCompleted(context.Context, *TaskRequest) (*Task, error)
	MarkTaskAsFailed(context.Context, *TaskRequest) (*Task, error)
	ValidateProof(context.Context, *v1.ValidateProofRequest) (*types.Empty, error)
	Sync(context.Context, *v11.SyncRequest) (*types.Empty, error)
	Ping(context.Context, *v12.PingRequest) (*v12.PingResponse, error)
	Register(context.Context, *RegistrationRequest) (*types.Empty, error)
}

// UnimplementedDispatcherServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDispatcherServiceServer struct {
}

func (*UnimplementedDispatcherServiceServer) GetPendingTask(ctx context.Context, req *TaskPendingRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingTask not implemented")
}
func (*UnimplementedDispatcherServiceServer) GetTask(ctx context.Context, req *TaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedDispatcherServiceServer) MarkTaskAsCompleted(ctx context.Context, req *TaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkTaskAsCompleted not implemented")
}
func (*UnimplementedDispatcherServiceServer) MarkTaskAsFailed(ctx context.Context, req *TaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkTaskAsFailed not implemented")
}
func (*UnimplementedDispatcherServiceServer) ValidateProof(ctx context.Context, req *v1.ValidateProofRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateProof not implemented")
}
func (*UnimplementedDispatcherServiceServer) Sync(ctx context.Context, req *v11.SyncRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (*UnimplementedDispatcherServiceServer) Ping(ctx context.Context, req *v12.PingRequest) (*v12.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedDispatcherServiceServer) Register(ctx context.Context, req *RegistrationRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterDispatcherServiceServer(s *grpc.Server, srv DispatcherServiceServer) {
	s.RegisterService(&_DispatcherService_serviceDesc, srv)
}

func _DispatcherService_GetPendingTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskPendingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetPendingTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/GetPendingTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetPendingTask(ctx, req.(*TaskPendingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_MarkTaskAsCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).MarkTaskAsCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/MarkTaskAsCompleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).MarkTaskAsCompleted(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_MarkTaskAsFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).MarkTaskAsFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/MarkTaskAsFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).MarkTaskAsFailed(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_ValidateProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ValidateProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).ValidateProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/ValidateProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).ValidateProof(ctx, req.(*v1.ValidateProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).Sync(ctx, req.(*v11.SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v12.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).Ping(ctx, req.(*v12.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).Register(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DispatcherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.dispatcher.v1.DispatcherService",
	HandlerType: (*DispatcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPendingTask",
			Handler:    _DispatcherService_GetPendingTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _DispatcherService_GetTask_Handler,
		},
		{
			MethodName: "MarkTaskAsCompleted",
			Handler:    _DispatcherService_MarkTaskAsCompleted_Handler,
		},
		{
			MethodName: "MarkTaskAsFailed",
			Handler:    _DispatcherService_MarkTaskAsFailed_Handler,
		},
		{
			MethodName: "ValidateProof",
			Handler:    _DispatcherService_ValidateProof_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _DispatcherService_Sync_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _DispatcherService_Ping_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _DispatcherService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dispatcher/v1/dispatcher_service.proto",
}

func (m *TaskPendingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskPendingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskPendingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintDispatcherService(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintDispatcherService(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintDispatcherService(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegistrationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegistrationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegistrationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintDispatcherService(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDispatcherService(dAtA []byte, offset int, v uint64) int {
	offset -= sovDispatcherService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TaskPendingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovDispatcherService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TaskRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovDispatcherService(uint64(l))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovDispatcherService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegistrationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovDispatcherService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDispatcherService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDispatcherService(x uint64) (n int) {
	return sovDispatcherService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TaskPendingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDispatcherService
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
			return fmt.Errorf("proto: TaskPendingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskPendingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDispatcherService
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
				return ErrInvalidLengthDispatcherService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDispatcherService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDispatcherService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDispatcherService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDispatcherService
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
func (m *TaskRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDispatcherService
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
			return fmt.Errorf("proto: TaskRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDispatcherService
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
				return ErrInvalidLengthDispatcherService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDispatcherService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDispatcherService
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
				return ErrInvalidLengthDispatcherService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDispatcherService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDispatcherService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDispatcherService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDispatcherService
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
func (m *RegistrationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDispatcherService
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
			return fmt.Errorf("proto: RegistrationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegistrationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDispatcherService
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
				return ErrInvalidLengthDispatcherService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDispatcherService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDispatcherService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDispatcherService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDispatcherService
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
func skipDispatcherService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDispatcherService
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
					return 0, ErrIntOverflowDispatcherService
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
					return 0, ErrIntOverflowDispatcherService
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
				return 0, ErrInvalidLengthDispatcherService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDispatcherService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDispatcherService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDispatcherService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDispatcherService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDispatcherService = fmt.Errorf("proto: unexpected end of group")
)
