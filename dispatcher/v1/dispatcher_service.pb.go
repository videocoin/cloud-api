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
	rpc "github.com/videocoin/cloud-api/rpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

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
		n, err := m.MarshalTo(b)
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
		n, err := m.MarshalTo(b)
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
func init() {
	proto.RegisterType((*TaskPendingRequest)(nil), "cloud.api.dispatcher.v1.TaskPendingRequest")
	golang_proto.RegisterType((*TaskPendingRequest)(nil), "cloud.api.dispatcher.v1.TaskPendingRequest")
	proto.RegisterType((*TaskRequest)(nil), "cloud.api.dispatcher.v1.TaskRequest")
	golang_proto.RegisterType((*TaskRequest)(nil), "cloud.api.dispatcher.v1.TaskRequest")
}

func init() {
	proto.RegisterFile("dispatcher/v1/dispatcher_service.proto", fileDescriptor_008866cbcf1626a3)
}
func init() {
	golang_proto.RegisterFile("dispatcher/v1/dispatcher_service.proto", fileDescriptor_008866cbcf1626a3)
}

var fileDescriptor_008866cbcf1626a3 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xbe, 0xc9, 0xa2, 0xf6, 0x8e, 0x3f, 0xe8, 0x2c, 0xae, 0x21, 0x57, 0x73, 0x25, 0x88, 0x28,
	0x7a, 0x67, 0xa8, 0x3e, 0x80, 0xd0, 0x56, 0x6a, 0x57, 0x86, 0x56, 0x14, 0x5c, 0x58, 0xa6, 0x33,
	0xd3, 0x64, 0x68, 0x9a, 0x89, 0xc9, 0x24, 0xa0, 0x4b, 0x5f, 0xc1, 0x17, 0x72, 0xd9, 0xa5, 0xe0,
	0xbe, 0x48, 0xea, 0xde, 0x57, 0x90, 0x4c, 0xd2, 0xa6, 0x45, 0x5a, 0x75, 0x97, 0x93, 0xef, 0x27,
	0xe7, 0x7c, 0x5f, 0xc0, 0x03, 0x26, 0xd2, 0x98, 0x28, 0x1a, 0xf0, 0x04, 0xe7, 0x1d, 0xdc, 0x4c,
	0x93, 0x94, 0x27, 0xb9, 0xa0, 0x1c, 0xc5, 0x89, 0x54, 0x12, 0xde, 0xa6, 0xa1, 0xcc, 0x18, 0x22,
	0xb1, 0x40, 0x0d, 0x07, 0xe5, 0x1d, 0x1b, 0xfb, 0x42, 0x05, 0xd9, 0x14, 0x51, 0xb9, 0xc0, 0xb9,
	0x60, 0x5c, 0x52, 0x29, 0x22, 0xac, 0xd9, 0x97, 0x24, 0x16, 0x38, 0x89, 0x29, 0x0e, 0x38, 0x09,
	0x55, 0x50, 0x39, 0xd9, 0xe7, 0xbe, 0x94, 0x7e, 0xc8, 0xb1, 0x9e, 0xa6, 0xd9, 0x0c, 0xf3, 0x45,
	0xac, 0x3e, 0xd6, 0xe0, 0x9d, 0x1a, 0x2c, 0x95, 0x24, 0x8a, 0xa4, 0x22, 0x4a, 0xc8, 0x28, 0xad,
	0xd1, 0xcb, 0x9d, 0x6f, 0xf9, 0xd2, 0x97, 0x8d, 0x47, 0x39, 0xe9, 0x41, 0x3f, 0xd5, 0x74, 0x6b,
	0xff, 0x36, 0x45, 0xd2, 0x79, 0x8d, 0xf4, 0xfe, 0xb2, 0x74, 0x4e, 0x42, 0xc1, 0x88, 0x92, 0x5a,
	0xb7, 0x1d, 0xf6, 0x23, 0x71, 0x9f, 0x03, 0xf8, 0x9a, 0xa4, 0x73, 0x8f, 0x47, 0x4c, 0x44, 0xfe,
	0x88, 0x7f, 0xc8, 0x78, 0xaa, 0xe0, 0x23, 0x70, 0x4a, 0x43, 0xc1, 0x23, 0x35, 0x11, 0xcc, 0x32,
	0xee, 0x19, 0x0f, 0x4f, 0xbb, 0xd7, 0x8a, 0xd5, 0x45, 0xbb, 0xa7, 0x5f, 0x0e, 0xfb, 0xa3, 0x76,
	0x05, 0x0f, 0x99, 0xeb, 0x81, 0xab, 0xa5, 0xc1, 0xff, 0x2b, 0xe1, 0x19, 0x30, 0x05, 0xb3, 0x4c,
	0xcd, 0x69, 0x15, 0xab, 0x0b, 0x73, 0xd8, 0x1f, 0x99, 0x82, 0x3d, 0xfd, 0x65, 0x82, 0x5b, 0xfd,
	0xed, 0xd1, 0xe3, 0x6a, 0x5d, 0xf8, 0x0a, 0xb4, 0x5e, 0xea, 0x06, 0xe0, 0x19, 0xaa, 0xf2, 0x45,
	0x9b, 0xe0, 0xd0, 0x8b, 0x32, 0x7c, 0xfb, 0x1c, 0x35, 0xf5, 0x26, 0x31, 0x45, 0x15, 0x7d, 0xac,
	0x88, 0xca, 0x52, 0xf7, 0xe6, 0xe7, 0xef, 0x3f, 0xbf, 0x98, 0x00, 0xb6, 0xeb, 0x1e, 0x3f, 0xc1,
	0xf7, 0xe0, 0xc6, 0x80, 0xab, 0xfa, 0xf0, 0xf2, 0x04, 0xf8, 0x18, 0x1d, 0xf8, 0x3f, 0xd0, 0x9f,
	0x11, 0xd9, 0x77, 0x8f, 0x92, 0xdd, 0x13, 0xe8, 0x81, 0x2b, 0x03, 0xae, 0xb4, 0xf1, 0xfd, 0xa3,
	0xdc, 0x7f, 0x76, 0x7c, 0x0b, 0xae, 0xbf, 0xa9, 0x6a, 0xe4, 0x5e, 0x22, 0xe5, 0x0c, 0x3e, 0xd9,
	0x51, 0x6c, 0x0b, 0x2e, 0x05, 0x7b, 0xb4, 0x8d, 0xff, 0x81, 0xdc, 0xdc, 0x93, 0xae, 0xb5, 0x2c,
	0x1c, 0xe3, 0x5b, 0xe1, 0x18, 0x3f, 0x0a, 0xc7, 0xf8, 0xba, 0x76, 0x8c, 0xe5, 0xda, 0x31, 0xde,
	0x99, 0x79, 0x67, 0xda, 0xd2, 0xdc, 0x67, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x76, 0x91,
	0xbc, 0x62, 0x03, 0x00, 0x00,
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
	Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
	GetPendingTask(ctx context.Context, in *TaskPendingRequest, opts ...grpc.CallOption) (*Task, error)
	GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error)
	ValidateProof(ctx context.Context, in *v1.ValidateProofRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type dispatcherServiceClient struct {
	cc *grpc.ClientConn
}

func NewDispatcherServiceClient(cc *grpc.ClientConn) DispatcherServiceClient {
	return &dispatcherServiceClient{cc}
}

func (c *dispatcherServiceClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func (c *dispatcherServiceClient) ValidateProof(ctx context.Context, in *v1.ValidateProofRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.dispatcher.v1.DispatcherService/ValidateProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServiceServer is the server API for DispatcherService service.
type DispatcherServiceServer interface {
	Health(context.Context, *types.Empty) (*rpc.HealthStatus, error)
	GetPendingTask(context.Context, *TaskPendingRequest) (*Task, error)
	GetTask(context.Context, *TaskRequest) (*Task, error)
	ValidateProof(context.Context, *v1.ValidateProofRequest) (*types.Empty, error)
}

// UnimplementedDispatcherServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDispatcherServiceServer struct {
}

func (*UnimplementedDispatcherServiceServer) Health(ctx context.Context, req *types.Empty) (*rpc.HealthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedDispatcherServiceServer) GetPendingTask(ctx context.Context, req *TaskPendingRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingTask not implemented")
}
func (*UnimplementedDispatcherServiceServer) GetTask(ctx context.Context, req *TaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedDispatcherServiceServer) ValidateProof(ctx context.Context, req *v1.ValidateProofRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateProof not implemented")
}

func RegisterDispatcherServiceServer(s *grpc.Server, srv DispatcherServiceServer) {
	s.RegisterService(&_DispatcherService_serviceDesc, srv)
}

func _DispatcherService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.dispatcher.v1.DispatcherService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).Health(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
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

var _DispatcherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.dispatcher.v1.DispatcherService",
	HandlerType: (*DispatcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _DispatcherService_Health_Handler,
		},
		{
			MethodName: "GetPendingTask",
			Handler:    _DispatcherService_GetPendingTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _DispatcherService_GetTask_Handler,
		},
		{
			MethodName: "ValidateProof",
			Handler:    _DispatcherService_ValidateProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dispatcher/v1/dispatcher_service.proto",
}

func (m *TaskPendingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskPendingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDispatcherService(dAtA, i, uint64(len(m.ClientID)))
		i += copy(dAtA[i:], m.ClientID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TaskRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDispatcherService(dAtA, i, uint64(len(m.ClientID)))
		i += copy(dAtA[i:], m.ClientID)
	}
	if len(m.ID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDispatcherService(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDispatcherService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
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
func skipDispatcherService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
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
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
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
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDispatcherService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDispatcherService
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
				next, err := skipDispatcherService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDispatcherService
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
	ErrInvalidLengthDispatcherService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDispatcherService   = fmt.Errorf("proto: integer overflow")
)
