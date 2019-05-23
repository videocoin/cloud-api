// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: notifications/v1/notifications_service.proto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.notifications.v1.NotificationService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	Health(context.Context, *empty.Empty) (*rpc.HealthStatus, error)
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.notifications.v1.NotificationService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.notifications.v1.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _NotificationService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications/v1/notifications_service.proto",
}

func init() {
	proto.RegisterFile("notifications/v1/notifications_service.proto", fileDescriptor_notifications_service_2f9ff8736dca8ae1)
}

var fileDescriptor_notifications_service_2f9ff8736dca8ae1 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x71, 0x87, 0x0a, 0x65, 0x42, 0x41, 0x62, 0x48, 0x2b, 0x0f, 0xcc, 0xf4, 0x59, 0x81,
	0x1b, 0x80, 0x90, 0x98, 0x60, 0xa8, 0xc4, 0xc0, 0x82, 0x1c, 0xd7, 0x71, 0x2c, 0xa5, 0x7e, 0x56,
	0xe2, 0x44, 0x82, 0x91, 0x2b, 0xb0, 0x70, 0x1c, 0xc6, 0x8e, 0x48, 0x5c, 0x00, 0x25, 0x1c, 0x04,
	0xc5, 0x0e, 0x6a, 0xb3, 0xf9, 0xf9, 0x7d, 0xff, 0xfb, 0xff, 0x3f, 0xba, 0x30, 0xe8, 0x74, 0xae,
	0x05, 0x77, 0x1a, 0x4d, 0xcd, 0xda, 0x94, 0x4d, 0x3e, 0x9e, 0x6b, 0x59, 0xb5, 0x5a, 0x48, 0xb0,
	0x15, 0x3a, 0x8c, 0x13, 0x51, 0x62, 0xb3, 0x01, 0x6e, 0x35, 0x4c, 0x30, 0x68, 0xd3, 0x64, 0xa1,
	0x10, 0x55, 0x29, 0x99, 0x27, 0xb3, 0x26, 0x67, 0x72, 0x6b, 0xdd, 0x4b, 0x10, 0x26, 0xcb, 0x71,
	0xc9, 0xad, 0x66, 0xdc, 0x18, 0x74, 0xa3, 0x2e, 0x6c, 0x57, 0x4a, 0xbb, 0xa2, 0xc9, 0x40, 0xe0,
	0x96, 0x29, 0x54, 0xb8, 0xbf, 0x31, 0x4c, 0x7e, 0xf0, 0xaf, 0x11, 0x67, 0x07, 0xf8, 0xa3, 0xde,
	0x48, 0xbc, 0x41, 0x6d, 0x98, 0x8f, 0xb6, 0x1a, 0x0c, 0x2a, 0x2b, 0x58, 0x21, 0x79, 0xe9, 0x8a,
	0x20, 0xb8, 0xcc, 0xa3, 0xd3, 0xfb, 0x83, 0xb8, 0xeb, 0xd0, 0x29, 0x7e, 0x88, 0xe6, 0x77, 0x1e,
	0x8b, 0xcf, 0x20, 0xe4, 0x83, 0x7f, 0x63, 0xb8, 0x1d, 0xc2, 0x27, 0x0b, 0xd8, 0x17, 0xae, 0xac,
	0x80, 0x80, 0xaf, 0x1d, 0x77, 0x4d, 0x7d, 0x7e, 0xf2, 0xf6, 0xfd, 0xfb, 0x3e, 0x8b, 0xe2, 0xe3,
	0xd1, 0xec, 0xf5, 0x7a, 0xb9, 0xeb, 0x28, 0xf9, 0xea, 0x28, 0xf9, 0xe9, 0x28, 0xf9, 0xe8, 0xe9,
	0xd1, 0x67, 0x4f, 0xc9, 0xae, 0xa7, 0xe4, 0x69, 0xd6, 0xa6, 0xd9, 0xdc, 0x1f, 0xbf, 0xfa, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x2e, 0x0c, 0x01, 0x41, 0x73, 0x01, 0x00, 0x00,
}
