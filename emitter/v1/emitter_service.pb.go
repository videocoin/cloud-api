// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emitter/v1/emitter_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	rpc "github.com/videocoin/cloud-api/rpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StreamRequest struct {
	PipelineId           string   `protobuf:"bytes,1,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StreamId             uint64   `protobuf:"varint,3,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	ClientAddress        string   `protobuf:"bytes,4,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_595d714f170d55af, []int{0}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return m.Size()
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetPipelineId() string {
	if m != nil {
		return m.PipelineId
	}
	return ""
}

func (m *StreamRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *StreamRequest) GetStreamId() uint64 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

func (m *StreamRequest) GetClientAddress() string {
	if m != nil {
		return m.ClientAddress
	}
	return ""
}

func (*StreamRequest) XXX_MessageName() string {
	return "cloud.api.streams.v1.StreamRequest"
}
func init() {
	proto.RegisterType((*StreamRequest)(nil), "cloud.api.streams.v1.StreamRequest")
	golang_proto.RegisterType((*StreamRequest)(nil), "cloud.api.streams.v1.StreamRequest")
}

func init() { proto.RegisterFile("emitter/v1/emitter_service.proto", fileDescriptor_595d714f170d55af) }
func init() {
	golang_proto.RegisterFile("emitter/v1/emitter_service.proto", fileDescriptor_595d714f170d55af)
}

var fileDescriptor_595d714f170d55af = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x51, 0x41, 0xcb, 0x13, 0x31,
	0x10, 0x35, 0xfd, 0x3e, 0x6a, 0x1b, 0x6d, 0x91, 0x20, 0x5a, 0xb6, 0xb2, 0x96, 0x8a, 0xd0, 0x4b,
	0x13, 0xaa, 0xbf, 0xa0, 0x4a, 0xc1, 0x0a, 0x2a, 0xb4, 0x37, 0x2f, 0x25, 0xdd, 0x8c, 0xdb, 0xc0,
	0xee, 0x26, 0x26, 0xd9, 0x05, 0x3d, 0x7a, 0xf7, 0xe4, 0xd5, 0x1f, 0xe3, 0xb1, 0x47, 0xc1, 0x3f,
	0x20, 0xad, 0x3f, 0x44, 0x36, 0xd9, 0x52, 0x05, 0xbd, 0xf5, 0x96, 0x99, 0xf7, 0xe6, 0xcd, 0xbc,
	0x17, 0x3c, 0x82, 0x5c, 0x3a, 0x07, 0x86, 0x55, 0x33, 0xd6, 0x3c, 0x37, 0x16, 0x4c, 0x25, 0x13,
	0xa0, 0xda, 0x28, 0xa7, 0xc8, 0xdd, 0x24, 0x53, 0xa5, 0xa0, 0x5c, 0x4b, 0x6a, 0x9d, 0x01, 0x9e,
	0x5b, 0x5a, 0xcd, 0xa2, 0x61, 0xaa, 0x54, 0x9a, 0x01, 0xf3, 0x9c, 0x6d, 0xf9, 0x8e, 0x41, 0xae,
	0xdd, 0x87, 0x30, 0x12, 0x3d, 0x68, 0x40, 0xae, 0x25, 0xe3, 0x45, 0xa1, 0x1c, 0x77, 0x52, 0x15,
	0xb6, 0x41, 0xa7, 0xa9, 0x74, 0xbb, 0x72, 0x4b, 0x13, 0x95, 0xb3, 0x54, 0xa5, 0xea, 0xac, 0x51,
	0x57, 0xbe, 0xf0, 0xaf, 0x86, 0xce, 0xfe, 0xa0, 0x57, 0x52, 0x80, 0x4a, 0x94, 0x2c, 0x98, 0x3f,
	0x6a, 0x5a, 0x2f, 0x30, 0x3a, 0x61, 0x3b, 0xe0, 0x99, 0xdb, 0x85, 0x81, 0xf1, 0x67, 0x84, 0x7b,
	0x6b, 0x7f, 0xe9, 0x0a, 0xde, 0x97, 0x60, 0x1d, 0x79, 0x88, 0x6f, 0x69, 0xa9, 0x21, 0x93, 0x05,
	0x6c, 0xa4, 0x18, 0xa0, 0x11, 0x9a, 0x74, 0x57, 0xf8, 0xd4, 0x5a, 0x0a, 0x72, 0x1f, 0xdf, 0x2c,
	0x2d, 0x98, 0x1a, 0x6c, 0x79, 0xb0, 0x5d, 0x97, 0x4b, 0x41, 0x86, 0xb8, 0x1b, 0x4c, 0xd7, 0xd0,
	0xd5, 0x08, 0x4d, 0xae, 0x57, 0x9d, 0xd0, 0x58, 0x0a, 0xf2, 0x18, 0xf7, 0x93, 0x4c, 0x42, 0xe1,
	0x36, 0x5c, 0x08, 0x03, 0xd6, 0x0e, 0xae, 0xfd, 0x70, 0x2f, 0x74, 0xe7, 0xa1, 0xf9, 0xe4, 0xeb,
	0x15, 0xee, 0x2f, 0x42, 0xb4, 0xeb, 0x90, 0x2c, 0x79, 0x83, 0xdb, 0x2f, 0xfc, 0xc9, 0xe4, 0x1e,
	0x0d, 0x59, 0xd1, 0x53, 0x08, 0x74, 0x51, 0x07, 0x19, 0x0d, 0xe9, 0x39, 0x76, 0xa3, 0x13, 0x1a,
	0xe8, 0x6b, 0xc7, 0x5d, 0x69, 0xc7, 0x77, 0x3e, 0xfd, 0xf8, 0xf5, 0xa5, 0x85, 0x49, 0xa7, 0x31,
	0xfe, 0x91, 0xbc, 0xc6, 0xbd, 0xc6, 0x6c, 0x70, 0x4e, 0x1e, 0xd1, 0x7f, 0x7d, 0x1b, 0xfd, 0x2b,
	0x97, 0xe8, 0x3f, 0xcb, 0xc7, 0x37, 0x6a, 0xbd, 0xb9, 0xd6, 0x46, 0x55, 0x70, 0x19, 0xbd, 0x57,
	0xf8, 0xf6, 0x73, 0x03, 0xdc, 0x5d, 0x48, 0xee, 0x25, 0xee, 0x2e, 0x0a, 0x71, 0x11, 0xad, 0x67,
	0x83, 0xfd, 0x21, 0x46, 0xdf, 0x0f, 0x31, 0xfa, 0x79, 0x88, 0xd1, 0xb7, 0x63, 0x8c, 0xf6, 0xc7,
	0x18, 0xbd, 0x6d, 0x55, 0xb3, 0x6d, 0xdb, 0x73, 0x9f, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x2b,
	0xaa, 0x36, 0x30, 0x24, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmitterServiceClient is the client API for EmitterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmitterServiceClient interface {
	Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
	RequestStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*types.Empty, error)
	ApproveStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*types.Empty, error)
	CreateStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*types.Empty, error)
	EndStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type emitterServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmitterServiceClient(cc *grpc.ClientConn) EmitterServiceClient {
	return &emitterServiceClient{cc}
}

func (c *emitterServiceClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.EmitterService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emitterServiceClient) RequestStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.EmitterService/RequestStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emitterServiceClient) ApproveStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.EmitterService/ApproveStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emitterServiceClient) CreateStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.EmitterService/CreateStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emitterServiceClient) EndStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.streams.v1.EmitterService/EndStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmitterServiceServer is the server API for EmitterService service.
type EmitterServiceServer interface {
	Health(context.Context, *types.Empty) (*rpc.HealthStatus, error)
	RequestStream(context.Context, *StreamRequest) (*types.Empty, error)
	ApproveStream(context.Context, *StreamRequest) (*types.Empty, error)
	CreateStream(context.Context, *StreamRequest) (*types.Empty, error)
	EndStream(context.Context, *StreamRequest) (*types.Empty, error)
}

// UnimplementedEmitterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmitterServiceServer struct {
}

func (*UnimplementedEmitterServiceServer) Health(ctx context.Context, req *types.Empty) (*rpc.HealthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedEmitterServiceServer) RequestStream(ctx context.Context, req *StreamRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestStream not implemented")
}
func (*UnimplementedEmitterServiceServer) ApproveStream(ctx context.Context, req *StreamRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveStream not implemented")
}
func (*UnimplementedEmitterServiceServer) CreateStream(ctx context.Context, req *StreamRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStream not implemented")
}
func (*UnimplementedEmitterServiceServer) EndStream(ctx context.Context, req *StreamRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndStream not implemented")
}

func RegisterEmitterServiceServer(s *grpc.Server, srv EmitterServiceServer) {
	s.RegisterService(&_EmitterService_serviceDesc, srv)
}

func _EmitterService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmitterServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.EmitterService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmitterServiceServer).Health(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmitterService_RequestStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmitterServiceServer).RequestStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.EmitterService/RequestStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmitterServiceServer).RequestStream(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmitterService_ApproveStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmitterServiceServer).ApproveStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.EmitterService/ApproveStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmitterServiceServer).ApproveStream(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmitterService_CreateStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmitterServiceServer).CreateStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.EmitterService/CreateStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmitterServiceServer).CreateStream(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmitterService_EndStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmitterServiceServer).EndStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.streams.v1.EmitterService/EndStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmitterServiceServer).EndStream(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmitterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.streams.v1.EmitterService",
	HandlerType: (*EmitterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _EmitterService_Health_Handler,
		},
		{
			MethodName: "RequestStream",
			Handler:    _EmitterService_RequestStream_Handler,
		},
		{
			MethodName: "ApproveStream",
			Handler:    _EmitterService_ApproveStream_Handler,
		},
		{
			MethodName: "CreateStream",
			Handler:    _EmitterService_CreateStream_Handler,
		},
		{
			MethodName: "EndStream",
			Handler:    _EmitterService_EndStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emitter/v1/emitter_service.proto",
}

func (m *StreamRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PipelineId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEmitterService(dAtA, i, uint64(len(m.PipelineId)))
		i += copy(dAtA[i:], m.PipelineId)
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEmitterService(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if m.StreamId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEmitterService(dAtA, i, uint64(m.StreamId))
	}
	if len(m.ClientAddress) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEmitterService(dAtA, i, uint64(len(m.ClientAddress)))
		i += copy(dAtA[i:], m.ClientAddress)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintEmitterService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StreamRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PipelineId)
	if l > 0 {
		n += 1 + l + sovEmitterService(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovEmitterService(uint64(l))
	}
	if m.StreamId != 0 {
		n += 1 + sovEmitterService(uint64(m.StreamId))
	}
	l = len(m.ClientAddress)
	if l > 0 {
		n += 1 + l + sovEmitterService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEmitterService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEmitterService(x uint64) (n int) {
	return sovEmitterService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmitterService
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
			return fmt.Errorf("proto: StreamRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PipelineId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmitterService
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
				return ErrInvalidLengthEmitterService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmitterService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PipelineId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmitterService
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
				return ErrInvalidLengthEmitterService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmitterService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamId", wireType)
			}
			m.StreamId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmitterService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmitterService
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
				return ErrInvalidLengthEmitterService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmitterService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEmitterService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEmitterService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEmitterService
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
func skipEmitterService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEmitterService
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
					return 0, ErrIntOverflowEmitterService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEmitterService
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
				return 0, ErrInvalidLengthEmitterService
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthEmitterService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEmitterService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEmitterService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthEmitterService
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEmitterService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEmitterService   = fmt.Errorf("proto: integer overflow")
)
