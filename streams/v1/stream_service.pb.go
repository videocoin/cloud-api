// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: streams/v1/stream_service.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import rpc "github.com/VideoCoin/cloud-api/rpc"
import proto1 "github.com/VideoCoin/common/proto"
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

type CreateStreamRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStreamRequest) Reset()         { *m = CreateStreamRequest{} }
func (m *CreateStreamRequest) String() string { return proto.CompactTextString(m) }
func (*CreateStreamRequest) ProtoMessage()    {}
func (*CreateStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_service_37769884d46aa806, []int{0}
}
func (m *CreateStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStreamRequest.Unmarshal(m, b)
}
func (m *CreateStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStreamRequest.Marshal(b, m, deterministic)
}
func (dst *CreateStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStreamRequest.Merge(dst, src)
}
func (m *CreateStreamRequest) XXX_Size() int {
	return xxx_messageInfo_CreateStreamRequest.Size(m)
}
func (m *CreateStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStreamRequest proto.InternalMessageInfo

func (m *CreateStreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_service_37769884d46aa806, []int{1}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListResponse struct {
	Items                []*StreamProfile `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_service_37769884d46aa806, []int{2}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetItems() []*StreamProfile {
	if m != nil {
		return m.Items
	}
	return nil
}

type UpdateStreamRequest struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               StreamStatus     `protobuf:"varint,3,opt,name=status,proto3,enum=cloud.api.streams.v1.StreamStatus" json:"status,omitempty"`
	ProfileId            proto1.ProfileId `protobuf:"varint,4,opt,name=profile_id,json=profileId,proto3,enum=proto.ProfileId" json:"profile_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateStreamRequest) Reset()         { *m = UpdateStreamRequest{} }
func (m *UpdateStreamRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStreamRequest) ProtoMessage()    {}
func (*UpdateStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_service_37769884d46aa806, []int{3}
}
func (m *UpdateStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStreamRequest.Unmarshal(m, b)
}
func (m *UpdateStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStreamRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStreamRequest.Merge(dst, src)
}
func (m *UpdateStreamRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStreamRequest.Size(m)
}
func (m *UpdateStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStreamRequest proto.InternalMessageInfo

func (m *UpdateStreamRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateStreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateStreamRequest) GetStatus() StreamStatus {
	if m != nil {
		return m.Status
	}
	return StreamStatusIdle
}

func (m *UpdateStreamRequest) GetProfileId() proto1.ProfileId {
	if m != nil {
		return m.ProfileId
	}
	return proto1.ProfileIdNone
}

func init() {
	proto.RegisterType((*CreateStreamRequest)(nil), "cloud.api.streams.v1.CreateStreamRequest")
	proto.RegisterType((*StreamRequest)(nil), "cloud.api.streams.v1.StreamRequest")
	proto.RegisterType((*ListResponse)(nil), "cloud.api.streams.v1.ListResponse")
	proto.RegisterType((*UpdateStreamRequest)(nil), "cloud.api.streams.v1.UpdateStreamRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
	Create(ctx context.Context, in *CreateStreamRequest, opts ...grpc.CallOption) (*StreamProfile, error)
	Get(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamProfile, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListResponse, error)
	Update(ctx context.Context, in *UpdateStreamRequest, opts ...grpc.CallOption) (*StreamProfile, error)
	Start(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamProfile, error)
	Stop(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamProfile, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.StreamService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) Create(ctx context.Context, in *CreateStreamRequest, opts ...grpc.CallOption) (*StreamProfile, error) {
	out := new(StreamProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.StreamService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) Get(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamProfile, error) {
	out := new(StreamProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.StreamService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.StreamService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) Update(ctx context.Context, in *UpdateStreamRequest, opts ...grpc.CallOption) (*StreamProfile, error) {
	out := new(StreamProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.StreamService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) Start(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamProfile, error) {
	out := new(StreamProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.StreamService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) Stop(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamProfile, error) {
	out := new(StreamProfile)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.StreamService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	Health(context.Context, *empty.Empty) (*rpc.HealthStatus, error)
	Create(context.Context, *CreateStreamRequest) (*StreamProfile, error)
	Get(context.Context, *StreamRequest) (*StreamProfile, error)
	List(context.Context, *empty.Empty) (*ListResponse, error)
	Update(context.Context, *UpdateStreamRequest) (*StreamProfile, error)
	Start(context.Context, *StreamRequest) (*StreamProfile, error)
	Stop(context.Context, *StreamRequest) (*StreamProfile, error)
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.StreamService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.StreamService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Create(ctx, req.(*CreateStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.StreamService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Get(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.StreamService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.StreamService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Update(ctx, req.(*UpdateStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.StreamService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Start(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.StreamService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Stop(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.streams.v1.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _StreamService_Health_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _StreamService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StreamService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StreamService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StreamService_Update_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _StreamService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _StreamService_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streams/v1/stream_service.proto",
}

func init() {
	proto.RegisterFile("streams/v1/stream_service.proto", fileDescriptor_stream_service_37769884d46aa806)
}

var fileDescriptor_stream_service_37769884d46aa806 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x9b, 0xb4, 0x8b, 0x98, 0x81, 0x31, 0xbc, 0x6a, 0x2b, 0xd9, 0xe8, 0x86, 0xb9, 0x0c,
	0xa4, 0xda, 0xda, 0xe0, 0xc2, 0x6e, 0x6c, 0x42, 0x30, 0x09, 0x09, 0x94, 0x0a, 0x0e, 0x70, 0x98,
	0xbc, 0xc4, 0xcb, 0x2c, 0x35, 0xb1, 0x89, 0x9d, 0x88, 0x81, 0xb8, 0x70, 0xe0, 0x05, 0xb8, 0xf0,
	0x04, 0x3c, 0x04, 0x27, 0x8e, 0x1c, 0x91, 0x78, 0x01, 0x54, 0xb8, 0xf0, 0x16, 0xa8, 0x76, 0x0a,
	0xa5, 0x4d, 0x2b, 0x26, 0xed, 0x52, 0xfd, 0x62, 0x7f, 0x7f, 0xfe, 0xfc, 0xfe, 0x7c, 0x0b, 0xd6,
	0x95, 0xce, 0x18, 0x4d, 0x14, 0x29, 0xb6, 0x88, 0x0d, 0x0f, 0x14, 0xcb, 0x0a, 0x1e, 0x32, 0x2c,
	0x33, 0xa1, 0x05, 0x6c, 0x86, 0x3d, 0x91, 0x47, 0x98, 0x4a, 0x8e, 0x4b, 0x29, 0x2e, 0xb6, 0xfc,
	0xd5, 0x58, 0x88, 0xb8, 0xc7, 0x88, 0xd1, 0x1c, 0xe6, 0x47, 0x84, 0x25, 0x52, 0x9f, 0xd8, 0x14,
	0x7f, 0xad, 0xbc, 0xa4, 0x92, 0x13, 0x9a, 0xa6, 0x42, 0x53, 0xcd, 0x45, 0xaa, 0xca, 0xdb, 0x4e,
	0xcc, 0xf5, 0x71, 0x7e, 0x88, 0x43, 0x91, 0x90, 0x58, 0xc4, 0xe2, 0xef, 0x1b, 0x83, 0x2f, 0xf3,
	0x61, 0xa2, 0x52, 0x4e, 0x46, 0xe4, 0x4f, 0x79, 0xc4, 0xc4, 0x9e, 0xe0, 0x29, 0x31, 0x45, 0x75,
	0x06, 0x80, 0x4c, 0x86, 0xe4, 0x98, 0xd1, 0x9e, 0x3e, 0x9e, 0x9d, 0x20, 0x92, 0x44, 0xa4, 0x96,
	0x35, 0xf8, 0x3d, 0xe2, 0xbd, 0xb2, 0x43, 0xff, 0xee, 0x48, 0x02, 0x4b, 0x0b, 0x71, 0x22, 0x33,
	0xf1, 0xf2, 0xc4, 0x4a, 0xc3, 0x4e, 0xcc, 0xd2, 0x4e, 0x41, 0x7b, 0x3c, 0xa2, 0x9a, 0x91, 0x89,
	0xa0, 0x7c, 0x62, 0x65, 0x62, 0x8a, 0xf6, 0x02, 0xdd, 0x06, 0x4b, 0x7b, 0x19, 0xa3, 0x9a, 0x75,
	0xcd, 0x69, 0xc0, 0x5e, 0xe4, 0x4c, 0x69, 0x78, 0x15, 0x34, 0x52, 0x9a, 0xb0, 0x96, 0xb3, 0xe1,
	0x6c, 0xce, 0xef, 0xce, 0x7f, 0xfa, 0xf5, 0xb9, 0xde, 0xc8, 0xdc, 0x45, 0x27, 0x30, 0xc7, 0x68,
	0x1d, 0x5c, 0xfc, 0x57, 0xbf, 0x00, 0x5c, 0x1e, 0x59, 0x75, 0xe0, 0xf2, 0x08, 0xed, 0x83, 0x0b,
	0x0f, 0xb9, 0xd2, 0x01, 0x53, 0x52, 0xa4, 0x8a, 0xc1, 0x3b, 0x60, 0x8e, 0x6b, 0x96, 0xa8, 0x96,
	0xb3, 0x51, 0xdf, 0x3c, 0xbf, 0x7d, 0x1d, 0x57, 0x2d, 0x0d, 0xdb, 0x37, 0x1f, 0xdb, 0xe6, 0x03,
	0x9b, 0x81, 0x3e, 0x3a, 0x60, 0xe9, 0x89, 0x8c, 0x26, 0x4a, 0x1c, 0x43, 0x42, 0x58, 0x96, 0xec,
	0x9a, 0x13, 0x13, 0xc3, 0x1d, 0xe0, 0x29, 0x4d, 0x75, 0xae, 0x5a, 0xf5, 0x0d, 0x67, 0x73, 0x61,
	0x1b, 0xcd, 0xe2, 0x76, 0x8d, 0x32, 0x28, 0x33, 0x20, 0x01, 0xa0, 0x5c, 0xc3, 0x01, 0x8f, 0x5a,
	0x0d, 0x93, 0xbf, 0x68, 0xa7, 0x86, 0xcb, 0x12, 0xf7, 0xa3, 0x60, 0x5e, 0x0e, 0xc3, 0xed, 0x77,
	0xde, 0x70, 0x2a, 0x5d, 0x6b, 0x50, 0xf8, 0x08, 0x78, 0x0f, 0xcc, 0xe6, 0xe1, 0x32, 0xb6, 0x96,
	0xc3, 0x43, 0x2f, 0xe1, 0x7b, 0x03, 0x3f, 0xfa, 0xab, 0x23, 0x05, 0x65, 0x32, 0xc4, 0x56, 0x6e,
	0x2b, 0x41, 0x8b, 0x6f, 0xbf, 0xfd, 0x7c, 0xef, 0x02, 0x78, 0xae, 0xf4, 0xcf, 0x2b, 0xa8, 0x81,
	0x67, 0xb7, 0x05, 0x6f, 0x54, 0x77, 0x52, 0xb1, 0x4b, 0xff, 0x7f, 0x86, 0x8d, 0x7c, 0xc3, 0x6a,
	0xa2, 0x4b, 0xe6, 0x4f, 0xf1, 0xc7, 0x25, 0x6a, 0xc7, 0xb9, 0x09, 0x13, 0x50, 0xbf, 0xcf, 0x34,
	0x9c, 0xf9, 0xce, 0xa9, 0x60, 0x6b, 0x06, 0xb6, 0x0c, 0x9b, 0x63, 0x30, 0xf2, 0x9a, 0x47, 0x6f,
	0xe0, 0x73, 0xd0, 0x18, 0x78, 0x67, 0xea, 0xcc, 0xa6, 0x2c, 0x71, 0xd4, 0x6f, 0x68, 0xc5, 0x10,
	0x2e, 0xc3, 0xf1, 0x76, 0x60, 0x01, 0x3c, 0x6b, 0xa6, 0x69, 0x13, 0xac, 0xb0, 0xda, 0xa9, 0x9a,
	0xf2, 0xab, 0x9b, 0xca, 0xc1, 0x5c, 0x57, 0xd3, 0xec, 0x2c, 0xa7, 0x88, 0x0c, 0x70, 0x0d, 0xf9,
	0x55, 0x40, 0xa2, 0x0c, 0x4d, 0x81, 0x46, 0x57, 0x0b, 0x79, 0x86, 0xd4, 0x6b, 0x86, 0xba, 0x8a,
	0xae, 0x4c, 0xa1, 0x0a, 0xb9, 0xdb, 0xfc, 0xd2, 0x6f, 0xd7, 0xbe, 0xf6, 0xdb, 0xb5, 0xef, 0xfd,
	0x76, 0xed, 0xc3, 0x8f, 0x76, 0xed, 0x99, 0x5b, 0x6c, 0x1d, 0x7a, 0x66, 0x8d, 0xb7, 0x7e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0xf4, 0xb8, 0xb0, 0xd1, 0x05, 0x00, 0x00,
}
