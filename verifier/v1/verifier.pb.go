// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: verifier/v1/verifier.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import rpc "github.com/VideoCoin/cloud-api/rpc"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "github.com/mwitkow/go-proto-validators"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type VerifyRequest struct {
	StreamId             int64    `protobuf:"varint,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	Bitrate              uint32   `protobuf:"varint,2,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	InputId              uint64   `protobuf:"varint,3,opt,name=input_id,json=inputId,proto3" json:"input_id,omitempty"`
	OutputId             uint64   `protobuf:"varint,4,opt,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
	SourceChunkUrl       string   `protobuf:"bytes,5,opt,name=source_chunk_url,json=sourceChunkUrl,proto3" json:"source_chunk_url,omitempty"`
	ResultChunkUrl       string   `protobuf:"bytes,6,opt,name=result_chunk_url,json=resultChunkUrl,proto3" json:"result_chunk_url,omitempty"`
	TxHash               string   `protobuf:"bytes,7,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_verifier_a3ca281ef9feb3f4, []int{0}
}
func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (dst *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(dst, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetStreamId() int64 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

func (m *VerifyRequest) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

func (m *VerifyRequest) GetInputId() uint64 {
	if m != nil {
		return m.InputId
	}
	return 0
}

func (m *VerifyRequest) GetOutputId() uint64 {
	if m != nil {
		return m.OutputId
	}
	return 0
}

func (m *VerifyRequest) GetSourceChunkUrl() string {
	if m != nil {
		return m.SourceChunkUrl
	}
	return ""
}

func (m *VerifyRequest) GetResultChunkUrl() string {
	if m != nil {
		return m.ResultChunkUrl
	}
	return ""
}

func (m *VerifyRequest) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyRequest)(nil), "cloud.api.verifier.v1.VerifyRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VerifierServiceClient is the client API for VerifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerifierServiceClient interface {
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
}

type verifierServiceClient struct {
	cc *grpc.ClientConn
}

func NewVerifierServiceClient(cc *grpc.ClientConn) VerifierServiceClient {
	return &verifierServiceClient{cc}
}

func (c *verifierServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.verifier.v1.VerifierService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifierServiceClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.verifier.v1.VerifierService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifierServiceServer is the server API for VerifierService service.
type VerifierServiceServer interface {
	Verify(context.Context, *VerifyRequest) (*types.Empty, error)
	Health(context.Context, *types.Empty) (*rpc.HealthStatus, error)
}

func RegisterVerifierServiceServer(s *grpc.Server, srv VerifierServiceServer) {
	s.RegisterService(&_VerifierService_serviceDesc, srv)
}

func _VerifierService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.verifier.v1.VerifierService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifierService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.verifier.v1.VerifierService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServiceServer).Health(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.verifier.v1.VerifierService",
	HandlerType: (*VerifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _VerifierService_Verify_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _VerifierService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verifier/v1/verifier.proto",
}

func init() {
	proto.RegisterFile("verifier/v1/verifier.proto", fileDescriptor_verifier_a3ca281ef9feb3f4)
}

var fileDescriptor_verifier_a3ca281ef9feb3f4 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x6f, 0xba, 0x6b, 0xda, 0x1d, 0xd8, 0xb5, 0x8c, 0xa8, 0xd9, 0x54, 0x4a, 0x29, 0x1e, 0x82,
	0xd0, 0x19, 0xaa, 0xe0, 0xc1, 0xa3, 0x8b, 0xb0, 0x3d, 0x09, 0x59, 0xec, 0xc1, 0x4b, 0x99, 0x26,
	0xb3, 0xc9, 0xb0, 0x69, 0x26, 0xce, 0xbc, 0xc9, 0x6e, 0x3d, 0xfa, 0x15, 0xfc, 0x4c, 0x82, 0x47,
	0xc1, 0x2f, 0x20, 0xc5, 0x9b, 0x5f, 0x42, 0x32, 0x93, 0xd6, 0x0a, 0x7a, 0x7b, 0xef, 0xf7, 0xe7,
	0x25, 0xf3, 0x7e, 0x0f, 0x85, 0x35, 0x57, 0xe2, 0x5a, 0x70, 0x45, 0xeb, 0x19, 0xdd, 0xd5, 0xa4,
	0x52, 0x12, 0x24, 0x7e, 0x98, 0x14, 0xd2, 0xa4, 0x84, 0x55, 0x82, 0xec, 0x99, 0x7a, 0x16, 0x0e,
	0x33, 0x29, 0xb3, 0x82, 0x53, 0x2b, 0x5a, 0x99, 0x6b, 0xca, 0xd7, 0x15, 0x6c, 0x9c, 0x27, 0x7c,
	0xd2, 0x92, 0xac, 0x12, 0x94, 0x95, 0xa5, 0x04, 0x06, 0x42, 0x96, 0xba, 0x65, 0xa7, 0x99, 0x80,
	0xdc, 0xac, 0x48, 0x22, 0xd7, 0x34, 0x93, 0x99, 0xfc, 0x33, 0xa3, 0xe9, 0x6c, 0x63, 0xab, 0x56,
	0xfe, 0xf2, 0x40, 0xbe, 0xbe, 0x15, 0x70, 0x23, 0x6f, 0x69, 0x26, 0xa7, 0x96, 0x9c, 0xd6, 0xac,
	0x10, 0x29, 0x03, 0xa9, 0x34, 0xdd, 0x97, 0xad, 0x8f, 0x1e, 0xf8, 0x16, 0x22, 0xe5, 0xf2, 0x42,
	0x8a, 0x92, 0xda, 0xd7, 0x4c, 0x9b, 0x1f, 0x53, 0x55, 0x42, 0x73, 0xce, 0x0a, 0xc8, 0x9d, 0x61,
	0xf2, 0xcb, 0x43, 0xa7, 0x8b, 0xe6, 0x89, 0x9b, 0x98, 0x7f, 0x30, 0x5c, 0x03, 0x1e, 0xa2, 0x13,
	0x0d, 0x8a, 0xb3, 0xf5, 0x52, 0xa4, 0x81, 0x37, 0xf6, 0xa2, 0xa3, 0xb8, 0xef, 0x80, 0x79, 0x8a,
	0x03, 0xd4, 0x5b, 0x09, 0x50, 0x0c, 0x78, 0xd0, 0x1d, 0x7b, 0xd1, 0x69, 0xbc, 0x6b, 0xf1, 0x39,
	0xea, 0x8b, 0xb2, 0x32, 0xd0, 0xb8, 0x8e, 0xc6, 0x5e, 0x74, 0x1c, 0xf7, 0x6c, 0x3f, 0x4f, 0x9b,
	0x89, 0xd2, 0x40, 0xcb, 0x1d, 0x5b, 0xae, 0xef, 0x80, 0x79, 0x8a, 0x23, 0x34, 0xd0, 0xd2, 0xa8,
	0x84, 0x2f, 0x93, 0xdc, 0x94, 0x37, 0x4b, 0xa3, 0x8a, 0xe0, 0xde, 0xd8, 0x8b, 0x4e, 0xe2, 0x33,
	0x87, 0x5f, 0x34, 0xf0, 0x3b, 0x55, 0x34, 0x4a, 0xc5, 0xb5, 0x29, 0xe0, 0x40, 0xe9, 0x3b, 0xa5,
	0xc3, 0xf7, 0xca, 0xc7, 0xa8, 0x07, 0x77, 0xcb, 0x9c, 0xe9, 0x3c, 0xe8, 0x59, 0x81, 0x0f, 0x77,
	0x97, 0x4c, 0xe7, 0xcf, 0xbf, 0x78, 0xe8, 0xfe, 0xa2, 0x0d, 0xf4, 0x8a, 0xab, 0x5a, 0x24, 0x1c,
	0x33, 0xe4, 0xbb, 0x05, 0xe0, 0xa7, 0xe4, 0x9f, 0xb1, 0x93, 0xbf, 0xf6, 0x13, 0x3e, 0x22, 0x2e,
	0x68, 0xb2, 0x4b, 0x90, 0xbc, 0x69, 0xae, 0x60, 0x72, 0xfe, 0xe9, 0xfb, 0xcf, 0xcf, 0xdd, 0x07,
	0x93, 0x33, 0x7b, 0x01, 0xbb, 0x9b, 0xda, 0xbc, 0xf2, 0x9e, 0xe1, 0xb7, 0xc8, 0xbf, 0xb4, 0x4b,
	0xc7, 0xff, 0x31, 0x87, 0xc3, 0x83, 0x4f, 0xab, 0x2a, 0x21, 0x4e, 0x7e, 0x05, 0x0c, 0x8c, 0x9e,
	0x0c, 0xec, 0x64, 0x84, 0xfb, 0x6d, 0x74, 0x1f, 0x5f, 0x0f, 0xbe, 0x6e, 0x47, 0x9d, 0x6f, 0xdb,
	0x51, 0xe7, 0xc7, 0x76, 0xd4, 0x79, 0xdf, 0xad, 0x67, 0x2b, 0xdf, 0x0e, 0x7c, 0xf1, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0xfb, 0x71, 0xb4, 0xd6, 0x02, 0x00, 0x00,
}
