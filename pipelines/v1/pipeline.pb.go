// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pipelines/v1/pipeline.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "github.com/VideoCoin/cloud-api/profiles/v1"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PipelineStatus int32

const (
	PipelineStatusIdle           PipelineStatus = 0
	PipelineStatusRequestPending PipelineStatus = 1
	PipelineStatusApprovePending PipelineStatus = 2
	PipelineStatusCreatePending  PipelineStatus = 3
	PipelineStatusJobPending     PipelineStatus = 4
	PipelineStatusRunning        PipelineStatus = 5
	PipelineStatusFailed         PipelineStatus = 6
	PipelineStatusCompleted      PipelineStatus = 7
)

var PipelineStatus_name = map[int32]string{
	0: "IDLE",
	1: "PENDING_REQUEST",
	2: "PENDING_APPROVE",
	3: "PENDING_CREATE",
	4: "PENDING_JOB",
	5: "RUNNING",
	6: "FAILED",
	7: "COMPLETED",
}
var PipelineStatus_value = map[string]int32{
	"IDLE":            0,
	"PENDING_REQUEST": 1,
	"PENDING_APPROVE": 2,
	"PENDING_CREATE":  3,
	"PENDING_JOB":     4,
	"RUNNING":         5,
	"FAILED":          6,
	"COMPLETED":       7,
}

func (x PipelineStatus) String() string {
	return proto.EnumName(PipelineStatus_name, int32(x))
}
func (PipelineStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_15f8bb83c61a63a4, []int{0}
}

type Pipeline struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"type:varchar(36);primary_key"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"type:varchar(100)"`
	UserId               string         `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" gorm:"type:varchar(36)"`
	ProfileId            v1.ProfileId   `protobuf:"varint,4,opt,name=profile_id,json=profileId,proto3,enum=cloud.api.profiles.v1.ProfileId" json:"profile_id,omitempty"`
	Status               PipelineStatus `protobuf:"varint,5,opt,name=status,proto3,enum=cloud.api.pipelines.v1.PipelineStatus" json:"status,omitempty" gorm:"type:varchar(100)"`
	ClientAddress        string         `protobuf:"bytes,6,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
	StreamId             uint64         `protobuf:"varint,7,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	StreamAddress        string         `protobuf:"bytes,8,opt,name=stream_address,json=streamAddress,proto3" json:"stream_address,omitempty"`
	CreatedAt            *time.Time     `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	PendingAt            *time.Time     `protobuf:"bytes,11,opt,name=pending_at,json=pendingAt,proto3,stdtime" json:"pending_at,omitempty"`
	RunningAt            *time.Time     `protobuf:"bytes,12,opt,name=running_at,json=runningAt,proto3,stdtime" json:"running_at,omitempty"`
	StoppedAt            *time.Time     `protobuf:"bytes,13,opt,name=stopped_at,json=stoppedAt,proto3,stdtime" json:"stopped_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}
func (*Pipeline) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_15f8bb83c61a63a4, []int{0}
}
func (m *Pipeline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipeline.Unmarshal(m, b)
}
func (m *Pipeline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipeline.Marshal(b, m, deterministic)
}
func (dst *Pipeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipeline.Merge(dst, src)
}
func (m *Pipeline) XXX_Size() int {
	return xxx_messageInfo_Pipeline.Size(m)
}
func (m *Pipeline) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipeline.DiscardUnknown(m)
}

var xxx_messageInfo_Pipeline proto.InternalMessageInfo

func (m *Pipeline) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pipeline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pipeline) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Pipeline) GetProfileId() v1.ProfileId {
	if m != nil {
		return m.ProfileId
	}
	return v1.ProfileIdNone
}

func (m *Pipeline) GetStatus() PipelineStatus {
	if m != nil {
		return m.Status
	}
	return PipelineStatusIdle
}

func (m *Pipeline) GetClientAddress() string {
	if m != nil {
		return m.ClientAddress
	}
	return ""
}

func (m *Pipeline) GetStreamId() uint64 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

func (m *Pipeline) GetStreamAddress() string {
	if m != nil {
		return m.StreamAddress
	}
	return ""
}

func (m *Pipeline) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Pipeline) GetPendingAt() *time.Time {
	if m != nil {
		return m.PendingAt
	}
	return nil
}

func (m *Pipeline) GetRunningAt() *time.Time {
	if m != nil {
		return m.RunningAt
	}
	return nil
}

func (m *Pipeline) GetStoppedAt() *time.Time {
	if m != nil {
		return m.StoppedAt
	}
	return nil
}

type PipelineProfile struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               PipelineStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cloud.api.pipelines.v1.PipelineStatus" json:"status,omitempty"`
	ProfileId            v1.ProfileId   `protobuf:"varint,4,opt,name=profile_id,json=profileId,proto3,enum=cloud.api.profiles.v1.ProfileId" json:"profile_id,omitempty"`
	CreatedAt            *time.Time     `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	PendingAt            *time.Time     `protobuf:"bytes,11,opt,name=pending_at,json=pendingAt,proto3,stdtime" json:"pending_at,omitempty"`
	RunningAt            *time.Time     `protobuf:"bytes,12,opt,name=running_at,json=runningAt,proto3,stdtime" json:"running_at,omitempty"`
	StoppedAt            *time.Time     `protobuf:"bytes,13,opt,name=stopped_at,json=stoppedAt,proto3,stdtime" json:"stopped_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PipelineProfile) Reset()         { *m = PipelineProfile{} }
func (m *PipelineProfile) String() string { return proto.CompactTextString(m) }
func (*PipelineProfile) ProtoMessage()    {}
func (*PipelineProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_15f8bb83c61a63a4, []int{1}
}
func (m *PipelineProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipelineProfile.Unmarshal(m, b)
}
func (m *PipelineProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipelineProfile.Marshal(b, m, deterministic)
}
func (dst *PipelineProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineProfile.Merge(dst, src)
}
func (m *PipelineProfile) XXX_Size() int {
	return xxx_messageInfo_PipelineProfile.Size(m)
}
func (m *PipelineProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_PipelineProfile.DiscardUnknown(m)
}

var xxx_messageInfo_PipelineProfile proto.InternalMessageInfo

func (m *PipelineProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PipelineProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineProfile) GetStatus() PipelineStatus {
	if m != nil {
		return m.Status
	}
	return PipelineStatusIdle
}

func (m *PipelineProfile) GetProfileId() v1.ProfileId {
	if m != nil {
		return m.ProfileId
	}
	return v1.ProfileIdNone
}

func (m *PipelineProfile) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PipelineProfile) GetPendingAt() *time.Time {
	if m != nil {
		return m.PendingAt
	}
	return nil
}

func (m *PipelineProfile) GetRunningAt() *time.Time {
	if m != nil {
		return m.RunningAt
	}
	return nil
}

func (m *PipelineProfile) GetStoppedAt() *time.Time {
	if m != nil {
		return m.StoppedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Pipeline)(nil), "cloud.api.pipelines.v1.Pipeline")
	proto.RegisterType((*PipelineProfile)(nil), "cloud.api.pipelines.v1.PipelineProfile")
	proto.RegisterEnum("cloud.api.pipelines.v1.PipelineStatus", PipelineStatus_name, PipelineStatus_value)
}

func init() {
	proto.RegisterFile("pipelines/v1/pipeline.proto", fileDescriptor_pipeline_15f8bb83c61a63a4)
}

var fileDescriptor_pipeline_15f8bb83c61a63a4 = []byte{
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x92, 0x4f, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0xe3, 0xfc, 0x6d, 0xa6, 0x6c, 0xd6, 0xb2, 0x96, 0xae, 0x71, 0xaa, 0xc4, 0x0a, 0xbb,
	0xab, 0xec, 0x4a, 0x71, 0x9a, 0xad, 0x40, 0x62, 0x91, 0x58, 0x39, 0x89, 0x77, 0xe5, 0xaa, 0xa4,
	0xc1, 0x4d, 0x7b, 0x40, 0x48, 0x91, 0x93, 0x99, 0xba, 0x23, 0x6c, 0xcf, 0x60, 0x4f, 0x2c, 0xf2,
	0x0d, 0x50, 0x4e, 0x9c, 0x50, 0x2f, 0x39, 0xc1, 0xa7, 0xe8, 0x89, 0x23, 0x17, 0x24, 0x3e, 0x41,
	0x41, 0xe5, 0xc4, 0xb5, 0x9f, 0x00, 0xc5, 0x7f, 0xda, 0x24, 0x44, 0xa2, 0x88, 0xeb, 0xde, 0xde,
	0x99, 0x79, 0x9f, 0x9f, 0xfd, 0xbe, 0xcf, 0x03, 0xca, 0x14, 0x53, 0x64, 0x63, 0x17, 0xf9, 0xcd,
	0xa0, 0xd5, 0x4c, 0x0e, 0x0a, 0xf5, 0x08, 0x23, 0xc2, 0xce, 0xd8, 0x26, 0x13, 0xa8, 0x98, 0x14,
	0x2b, 0xb7, 0x6d, 0x4a, 0xd0, 0x92, 0x3e, 0xb1, 0x30, 0x3b, 0x9f, 0x8c, 0x94, 0x31, 0x71, 0x9a,
	0xa7, 0x18, 0x22, 0xd2, 0x21, 0xd8, 0x6d, 0x86, 0xcd, 0x0d, 0x93, 0xe2, 0x26, 0xf5, 0xc8, 0x19,
	0xb6, 0x63, 0x64, 0x5c, 0x47, 0x48, 0x49, 0x5d, 0x92, 0x22, 0x37, 0x20, 0x53, 0xea, 0x91, 0x6f,
	0xa7, 0xcd, 0xf0, 0x71, 0xdc, 0xb0, 0x90, 0xdb, 0x08, 0x4c, 0x1b, 0x43, 0x93, 0xa1, 0xe6, 0x3f,
	0x8a, 0x18, 0xd1, 0x58, 0x42, 0x58, 0xc4, 0x22, 0x91, 0x78, 0x34, 0x39, 0x0b, 0x4f, 0xe1, 0x21,
	0xac, 0xe2, 0xf6, 0xaa, 0x45, 0x88, 0x65, 0xa3, 0xbb, 0x2e, 0x86, 0x1d, 0xe4, 0x33, 0xd3, 0xa1,
	0x51, 0x43, 0xed, 0xd7, 0x1c, 0xd8, 0xea, 0xc7, 0xe3, 0x09, 0x2a, 0x48, 0x63, 0x28, 0x72, 0x32,
	0x57, 0x2f, 0xb6, 0x5b, 0x37, 0x57, 0xd5, 0x0f, 0x2d, 0xe2, 0x39, 0xaf, 0x6a, 0x6c, 0x4a, 0xd1,
	0xab, 0xc0, 0xf4, 0xc6, 0xe7, 0xa6, 0x57, 0xdf, 0xff, 0xf8, 0xf9, 0xa7, 0xd4, 0xc3, 0x8e, 0xe9,
	0x4d, 0x87, 0x5f, 0xa3, 0x69, 0xed, 0xf2, 0xaf, 0x9f, 0x33, 0x39, 0x2f, 0x73, 0xc1, 0x3d, 0x31,
	0xd2, 0x18, 0x0a, 0x7b, 0x20, 0xeb, 0x9a, 0x0e, 0x12, 0xd3, 0x21, 0x64, 0xf7, 0xe6, 0xaa, 0x2a,
	0x6e, 0x80, 0xb4, 0xf6, 0xf6, 0x9e, 0xd7, 0x8c, 0xb0, 0x53, 0xf8, 0x0c, 0x14, 0x26, 0x3e, 0xf2,
	0x86, 0x18, 0x8a, 0x99, 0x50, 0xf4, 0xf4, 0xe6, 0xaa, 0xfa, 0x78, 0xf3, 0x97, 0x97, 0xbf, 0x96,
	0x5f, 0xa8, 0x74, 0x28, 0xbc, 0x06, 0x20, 0x5e, 0xf3, 0x02, 0x91, 0x95, 0xb9, 0x7a, 0xe9, 0xa5,
	0xac, 0x2c, 0x99, 0x97, 0x78, 0x10, 0xb4, 0x94, 0x7e, 0x54, 0xeb, 0xd0, 0x28, 0xd2, 0xa4, 0x14,
	0xbe, 0x02, 0x79, 0x9f, 0x99, 0x6c, 0xe2, 0x8b, 0xb9, 0x50, 0xfc, 0x4c, 0xd9, 0xec, 0xbc, 0x92,
	0xec, 0xe9, 0x38, 0xec, 0xfe, 0x97, 0xe1, 0x62, 0xa6, 0xf0, 0x14, 0x94, 0xc6, 0x36, 0x46, 0x2e,
	0x1b, 0x9a, 0x10, 0x7a, 0xc8, 0xf7, 0xc5, 0xfc, 0x62, 0x4a, 0xe3, 0x41, 0x74, 0xab, 0x46, 0x97,
	0x42, 0x19, 0x14, 0x7d, 0xe6, 0x21, 0xd3, 0x59, 0x0c, 0x51, 0x90, 0xb9, 0x7a, 0xd6, 0xd8, 0x8a,
	0x2e, 0x74, 0xb8, 0x60, 0xc4, 0x8f, 0x09, 0x63, 0x2b, 0x62, 0x44, 0xb7, 0x09, 0xe3, 0x35, 0x00,
	0x63, 0x0f, 0x99, 0x0c, 0xc1, 0xa1, 0xc9, 0x44, 0x20, 0x73, 0xf5, 0xed, 0x97, 0x92, 0x12, 0x25,
	0x40, 0x49, 0x12, 0xa0, 0x0c, 0x92, 0x04, 0xb4, 0xb3, 0xdf, 0xff, 0x5e, 0xe5, 0x8c, 0x62, 0xac,
	0x51, 0x59, 0xb8, 0x4a, 0xe4, 0x42, 0xec, 0x5a, 0x0b, 0xc0, 0xf6, 0x7d, 0x01, 0xb1, 0x26, 0x02,
	0x78, 0x13, 0xd7, 0x8d, 0x01, 0xef, 0xdd, 0x17, 0x10, 0x6b, 0x22, 0x80, 0xcf, 0x08, 0xa5, 0xd1,
	0x08, 0x0f, 0xee, 0x0b, 0x88, 0x35, 0x2a, 0xab, 0x5d, 0x66, 0xc0, 0xc3, 0xc4, 0xa7, 0xd8, 0x6d,
	0xa1, 0x74, 0x17, 0xeb, 0x30, 0xa3, 0xc2, 0x72, 0x46, 0x6f, 0x53, 0x98, 0x84, 0x20, 0xf3, 0x5f,
	0x42, 0x70, 0x6b, 0xf3, 0xff, 0x4e, 0xe1, 0x3b, 0xf3, 0xd8, 0x8b, 0x1f, 0x32, 0xa0, 0xb4, 0xba,
	0x5f, 0x41, 0x06, 0x59, 0xbd, 0x7b, 0xa8, 0xf1, 0x29, 0x69, 0x67, 0x36, 0x97, 0x85, 0xd5, 0x57,
	0x1d, 0xda, 0x48, 0xf8, 0x08, 0x3c, 0xec, 0x6b, 0xbd, 0xae, 0xde, 0x7b, 0x3b, 0x34, 0xb4, 0x2f,
	0x4e, 0xb4, 0xe3, 0x01, 0xcf, 0x49, 0xf2, 0x6c, 0x2e, 0xef, 0xae, 0x59, 0x85, 0xbe, 0x99, 0x20,
	0x9f, 0xf5, 0xa3, 0x79, 0x97, 0x65, 0x6a, 0xbf, 0x6f, 0x1c, 0x9d, 0x6a, 0x7c, 0x7a, 0x93, 0x4c,
	0xa5, 0xd4, 0x23, 0x01, 0x4a, 0x64, 0xfb, 0xa0, 0x94, 0xc8, 0x3a, 0x86, 0xa6, 0x0e, 0x34, 0x3e,
	0x23, 0x55, 0x67, 0x73, 0xb9, 0xbc, 0xaa, 0xea, 0x84, 0xb6, 0x24, 0xa2, 0x06, 0xd8, 0x4e, 0x44,
	0x07, 0x47, 0x6d, 0x3e, 0x2b, 0xed, 0xce, 0xe6, 0xb2, 0xb8, 0xaa, 0x38, 0x20, 0xa3, 0xa4, 0xfd,
	0x19, 0x28, 0x18, 0x27, 0xbd, 0x9e, 0xde, 0x7b, 0xcb, 0xe7, 0xa4, 0x0f, 0x66, 0x73, 0xf9, 0xfd,
	0xb5, 0x49, 0xa2, 0x8d, 0x0b, 0x4f, 0x40, 0xfe, 0x8d, 0xaa, 0x1f, 0x6a, 0x5d, 0x3e, 0x2f, 0x89,
	0xb3, 0xb9, 0xfc, 0x68, 0xb5, 0xed, 0x8d, 0x89, 0x6d, 0x04, 0x85, 0x17, 0xa0, 0xd8, 0x39, 0xfa,
	0xbc, 0x7f, 0xa8, 0x0d, 0xb4, 0x2e, 0x5f, 0x90, 0xca, 0xb3, 0xb9, 0xfc, 0x78, 0xed, 0x67, 0x89,
	0x43, 0x6d, 0xc4, 0x10, 0x94, 0x76, 0xbe, 0xfb, 0xb1, 0x92, 0xba, 0xfc, 0xa9, 0xb2, 0xe6, 0x42,
	0xfb, 0xd1, 0x2f, 0xd7, 0x95, 0xd4, 0x6f, 0xd7, 0x95, 0xd4, 0x1f, 0xd7, 0x95, 0xd4, 0xc5, 0x9f,
	0x95, 0xd4, 0x97, 0xe9, 0xa0, 0x35, 0xca, 0x87, 0x9e, 0xee, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xa8, 0xa8, 0x68, 0xd1, 0x47, 0x07, 0x00, 0x00,
}
