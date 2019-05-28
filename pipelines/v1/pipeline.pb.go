// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pipelines/v1/pipeline.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "github.com/VideoCoin/cloud-api/profiles/v1"
import v11 "github.com/VideoCoin/cloud-api/workorder/v1"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

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
	return fileDescriptor_pipeline_703afdca4440d8e4, []int{0}
}

type Pipeline struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"type:varchar(36);primary_key"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"type:varchar(100)"`
	UserId               string         `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" gorm:"type:varchar(36)"`
	ProfileId            v1.ProfileId   `protobuf:"varint,4,opt,name=profile_id,json=profileId,proto3,enum=cloud.api.profiles.v1.ProfileId" json:"profile_id,omitempty"`
	Status               PipelineStatus `protobuf:"varint,5,opt,name=status,proto3,enum=cloud.api.pipelines.v1.PipelineStatus" json:"status,omitempty" gorm:"type:varchar(100)"`
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
	return fileDescriptor_pipeline_703afdca4440d8e4, []int{0}
}
func (m *Pipeline) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pipeline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pipeline.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Pipeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipeline.Merge(dst, src)
}
func (m *Pipeline) XXX_Size() int {
	return m.Size()
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

func (*Pipeline) XXX_MessageName() string {
	return "cloud.api.pipelines.v1.Pipeline"
}

type PipelineProfile struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               PipelineStatus  `protobuf:"varint,3,opt,name=status,proto3,enum=cloud.api.pipelines.v1.PipelineStatus" json:"status,omitempty"`
	ProfileId            v1.ProfileId    `protobuf:"varint,4,opt,name=profile_id,json=profileId,proto3,enum=cloud.api.profiles.v1.ProfileId" json:"profile_id,omitempty"`
	JobProfile           *v11.JobProfile `protobuf:"bytes,5,opt,name=job_profile,json=jobProfile,proto3" json:"job_profile,omitempty"`
	CreatedAt            *time.Time      `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	PendingAt            *time.Time      `protobuf:"bytes,11,opt,name=pending_at,json=pendingAt,proto3,stdtime" json:"pending_at,omitempty"`
	RunningAt            *time.Time      `protobuf:"bytes,12,opt,name=running_at,json=runningAt,proto3,stdtime" json:"running_at,omitempty"`
	StoppedAt            *time.Time      `protobuf:"bytes,13,opt,name=stopped_at,json=stoppedAt,proto3,stdtime" json:"stopped_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PipelineProfile) Reset()         { *m = PipelineProfile{} }
func (m *PipelineProfile) String() string { return proto.CompactTextString(m) }
func (*PipelineProfile) ProtoMessage()    {}
func (*PipelineProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_703afdca4440d8e4, []int{1}
}
func (m *PipelineProfile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PipelineProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PipelineProfile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PipelineProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineProfile.Merge(dst, src)
}
func (m *PipelineProfile) XXX_Size() int {
	return m.Size()
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

func (m *PipelineProfile) GetJobProfile() *v11.JobProfile {
	if m != nil {
		return m.JobProfile
	}
	return nil
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

func (*PipelineProfile) XXX_MessageName() string {
	return "cloud.api.pipelines.v1.PipelineProfile"
}
func init() {
	proto.RegisterType((*Pipeline)(nil), "cloud.api.pipelines.v1.Pipeline")
	golang_proto.RegisterType((*Pipeline)(nil), "cloud.api.pipelines.v1.Pipeline")
	proto.RegisterType((*PipelineProfile)(nil), "cloud.api.pipelines.v1.PipelineProfile")
	golang_proto.RegisterType((*PipelineProfile)(nil), "cloud.api.pipelines.v1.PipelineProfile")
	proto.RegisterEnum("cloud.api.pipelines.v1.PipelineStatus", PipelineStatus_name, PipelineStatus_value)
	golang_proto.RegisterEnum("cloud.api.pipelines.v1.PipelineStatus", PipelineStatus_name, PipelineStatus_value)
}
func (m *Pipeline) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pipeline) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if m.ProfileId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(m.ProfileId))
	}
	if m.Status != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(m.Status))
	}
	if m.CreatedAt != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)))
		n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PendingAt != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.PendingAt)))
		n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.PendingAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.RunningAt != nil {
		dAtA[i] = 0x62
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.RunningAt)))
		n3, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.RunningAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.StoppedAt != nil {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.StoppedAt)))
		n4, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.StoppedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PipelineProfile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PipelineProfile) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Status != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(m.Status))
	}
	if m.ProfileId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(m.ProfileId))
	}
	if m.JobProfile != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(m.JobProfile.Size()))
		n5, err := m.JobProfile.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.CreatedAt != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)))
		n6, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.PendingAt != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.PendingAt)))
		n7, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.PendingAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.RunningAt != nil {
		dAtA[i] = 0x62
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.RunningAt)))
		n8, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.RunningAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.StoppedAt != nil {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.StoppedAt)))
		n9, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.StoppedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPipeline(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Pipeline) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.ProfileId != 0 {
		n += 1 + sovPipeline(uint64(m.ProfileId))
	}
	if m.Status != 0 {
		n += 1 + sovPipeline(uint64(m.Status))
	}
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.PendingAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.PendingAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.RunningAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.RunningAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.StoppedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.StoppedAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PipelineProfile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovPipeline(uint64(m.Status))
	}
	if m.ProfileId != 0 {
		n += 1 + sovPipeline(uint64(m.ProfileId))
	}
	if m.JobProfile != nil {
		l = m.JobProfile.Size()
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.PendingAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.PendingAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.RunningAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.RunningAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.StoppedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.StoppedAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPipeline(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPipeline(x uint64) (n int) {
	return sovPipeline(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pipeline) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPipeline
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pipeline: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pipeline: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			m.ProfileId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProfileId |= (v1.ProfileId(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (PipelineStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PendingAt == nil {
				m.PendingAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.PendingAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunningAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RunningAt == nil {
				m.RunningAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.RunningAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoppedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StoppedAt == nil {
				m.StoppedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.StoppedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPipeline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPipeline
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
func (m *PipelineProfile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPipeline
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PipelineProfile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PipelineProfile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (PipelineStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			m.ProfileId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProfileId |= (v1.ProfileId(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobProfile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.JobProfile == nil {
				m.JobProfile = &v11.JobProfile{}
			}
			if err := m.JobProfile.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PendingAt == nil {
				m.PendingAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.PendingAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunningAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RunningAt == nil {
				m.RunningAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.RunningAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoppedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPipeline
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StoppedAt == nil {
				m.StoppedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.StoppedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPipeline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPipeline
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
func skipPipeline(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPipeline
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
					return 0, ErrIntOverflowPipeline
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
					return 0, ErrIntOverflowPipeline
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPipeline
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPipeline
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
				next, err := skipPipeline(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthPipeline = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPipeline   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("pipelines/v1/pipeline.proto", fileDescriptor_pipeline_703afdca4440d8e4)
}
func init() {
	golang_proto.RegisterFile("pipelines/v1/pipeline.proto", fileDescriptor_pipeline_703afdca4440d8e4)
}

var fileDescriptor_pipeline_703afdca4440d8e4 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x93, 0xcf, 0x4e, 0x1a, 0x51,
	0x14, 0xc6, 0x1d, 0x40, 0x2c, 0x97, 0x16, 0x27, 0x93, 0x56, 0xa7, 0x83, 0x81, 0x09, 0x6d, 0x2c,
	0x9a, 0x38, 0x88, 0xf6, 0x4f, 0xaa, 0x49, 0x1b, 0xfe, 0x8c, 0x06, 0x63, 0x91, 0x8e, 0xe8, 0xa2,
	0x69, 0x42, 0x06, 0xe6, 0x8a, 0x57, 0x81, 0x7b, 0x3b, 0x5c, 0x68, 0x7c, 0x83, 0x86, 0x55, 0x57,
	0xdd, 0xb1, 0x6a, 0x9f, 0xa2, 0xab, 0x2e, 0x5d, 0x35, 0x7d, 0x02, 0xdb, 0xe8, 0x1b, 0xf8, 0x02,
	0x6d, 0xb8, 0x73, 0x07, 0x81, 0x90, 0x68, 0xd3, 0x6d, 0x77, 0xe7, 0x0c, 0xdf, 0xef, 0x63, 0xce,
	0xf9, 0xce, 0x80, 0x30, 0x41, 0x04, 0xd6, 0x50, 0x03, 0x36, 0x13, 0xed, 0x64, 0xc2, 0x6d, 0x34,
	0x62, 0x63, 0x8a, 0xa5, 0x99, 0x4a, 0x0d, 0xb7, 0x2c, 0xcd, 0x24, 0x48, 0xeb, 0xcb, 0xb4, 0x76,
	0x52, 0x79, 0x5e, 0x45, 0xf4, 0xb0, 0x55, 0xd6, 0x2a, 0xb8, 0x9e, 0xd8, 0x47, 0x16, 0xc4, 0x19,
	0x8c, 0x1a, 0x09, 0x26, 0x5e, 0x32, 0x09, 0x4a, 0x10, 0x1b, 0x1f, 0xa0, 0x1a, 0xb7, 0xe4, 0xb5,
	0x63, 0xa9, 0xac, 0x5f, 0x83, 0xbe, 0xc7, 0xf6, 0x31, 0xb6, 0x2d, 0x68, 0xf7, 0xd8, 0x7e, 0xc3,
	0xe1, 0xa5, 0x01, 0xb8, 0x8a, 0xab, 0x38, 0xc1, 0x1e, 0x97, 0x5b, 0x07, 0xac, 0x63, 0x0d, 0xab,
	0xb8, 0x3c, 0x5a, 0xc5, 0xb8, 0x5a, 0x83, 0x57, 0x2a, 0x8a, 0xea, 0xb0, 0x49, 0xcd, 0x3a, 0x71,
	0x04, 0xb1, 0xef, 0x3e, 0x70, 0xab, 0xc0, 0x07, 0x93, 0x9e, 0x01, 0x0f, 0xb2, 0x64, 0x41, 0x15,
	0xe2, 0x81, 0xf4, 0xa3, 0xcb, 0xb3, 0xe8, 0x83, 0x2a, 0xb6, 0xeb, 0x6b, 0x31, 0x7a, 0x42, 0xe0,
	0x5a, 0xdb, 0xb4, 0x2b, 0x87, 0xa6, 0x1d, 0x5f, 0x7d, 0xba, 0xb0, 0x4e, 0x6c, 0x54, 0x37, 0xed,
	0x93, 0xd2, 0x31, 0x3c, 0x89, 0x19, 0x1e, 0x64, 0x49, 0xcb, 0xc0, 0xd7, 0x30, 0xeb, 0x50, 0xf6,
	0x30, 0x74, 0xee, 0xf2, 0x2c, 0x2a, 0x8f, 0x41, 0x93, 0xcb, 0xcb, 0x0b, 0x31, 0x83, 0x29, 0xa5,
	0xc7, 0x60, 0xaa, 0xd5, 0x84, 0x76, 0x09, 0x59, 0xb2, 0x97, 0x41, 0xe1, 0xcb, 0xb3, 0xe8, 0xec,
	0xf8, 0xff, 0x8b, 0x19, 0xfe, 0x9e, 0x36, 0x67, 0x49, 0x2f, 0x01, 0xe0, 0xcb, 0xec, 0x81, 0x3e,
	0x55, 0x88, 0x87, 0x56, 0x54, 0x6d, 0x20, 0x22, 0x77, 0xd3, 0xed, 0xa4, 0x56, 0x70, 0xea, 0x9c,
	0x65, 0x04, 0x88, 0x5b, 0x4a, 0x6f, 0x81, 0xbf, 0x49, 0x4d, 0xda, 0x6a, 0xca, 0x93, 0x0c, 0x9e,
	0xd7, 0xc6, 0xe7, 0xab, 0xb9, 0x3b, 0xd9, 0x65, 0xea, 0x6b, 0x46, 0xe2, 0x9e, 0xbd, 0xd7, 0xab,
	0xd8, 0xd0, 0xa4, 0xd0, 0x2a, 0x99, 0x54, 0x06, 0xaa, 0x10, 0x0f, 0xae, 0x28, 0x9a, 0x13, 0x81,
	0xe6, 0x46, 0xa0, 0x15, 0xdd, 0x08, 0xd2, 0xbe, 0x8f, 0x3f, 0xa3, 0x82, 0x11, 0xe0, 0x4c, 0x8a,
	0xb2, 0xf9, 0x60, 0xc3, 0x42, 0x8d, 0x6a, 0xcf, 0x20, 0x78, 0x53, 0x03, 0xce, 0x38, 0x06, 0x76,
	0xab, 0xd1, 0xe0, 0x06, 0xb7, 0x6f, 0x6a, 0xc0, 0x19, 0xc7, 0xa0, 0x49, 0x31, 0x21, 0xce, 0x08,
	0x77, 0x6e, 0x6a, 0xc0, 0x99, 0x14, 0x8d, 0xfd, 0xf6, 0x82, 0x69, 0x77, 0x79, 0x3c, 0x02, 0x29,
	0x74, 0x75, 0x57, 0xec, 0x5c, 0xa4, 0xc1, 0x73, 0xe1, 0x07, 0xf1, 0xa2, 0x9f, 0x8c, 0xf7, 0x6f,
	0x92, 0x19, 0xdc, 0xfd, 0xbf, 0x9d, 0x46, 0x06, 0x04, 0x8f, 0x70, 0xb9, 0xc4, 0x1f, 0xb0, 0xfb,
	0x08, 0xae, 0xc4, 0x06, 0x1c, 0xae, 0x3e, 0xc5, 0x76, 0x52, 0xdb, 0xc2, 0x65, 0xee, 0x62, 0x80,
	0xa3, 0x7e, 0xfd, 0xff, 0x02, 0x52, 0x74, 0xf1, 0x93, 0x17, 0x84, 0x86, 0x43, 0x92, 0x54, 0xe0,
	0xcb, 0x65, 0xb7, 0x75, 0x71, 0x42, 0x99, 0xe9, 0x74, 0x55, 0x69, 0xf8, 0xd7, 0x9c, 0x55, 0x83,
	0xd2, 0x13, 0x30, 0x5d, 0xd0, 0xf3, 0xd9, 0x5c, 0x7e, 0xb3, 0x64, 0xe8, 0xaf, 0xf7, 0xf4, 0xdd,
	0xa2, 0x28, 0x28, 0x6a, 0xa7, 0xab, 0xce, 0x8d, 0xe4, 0x0d, 0xdf, 0xb5, 0x60, 0x93, 0x16, 0x9c,
	0x79, 0x07, 0xb1, 0x54, 0xa1, 0x60, 0xec, 0xec, 0xeb, 0xa2, 0x67, 0x1c, 0x96, 0x22, 0xc4, 0xc6,
	0x6d, 0xe8, 0x62, 0xab, 0x20, 0xe4, 0x62, 0x19, 0x43, 0x4f, 0x15, 0x75, 0xd1, 0xab, 0x44, 0x3b,
	0x5d, 0x35, 0x3c, 0x4c, 0x65, 0x58, 0x2c, 0x2e, 0xb4, 0x04, 0x82, 0x2e, 0xb4, 0xb5, 0x93, 0x16,
	0x7d, 0xca, 0x5c, 0xa7, 0xab, 0xca, 0xc3, 0x44, 0xef, 0x2c, 0xb8, 0x7c, 0x1e, 0x4c, 0x19, 0x7b,
	0xf9, 0x7c, 0x2e, 0xbf, 0x29, 0x4e, 0x2a, 0xf7, 0x3b, 0x5d, 0xf5, 0xde, 0xc8, 0x24, 0xce, 0xc6,
	0xa5, 0x87, 0xc0, 0xbf, 0x91, 0xca, 0x6d, 0xeb, 0x59, 0xd1, 0xaf, 0xc8, 0x9d, 0xae, 0x7a, 0x77,
	0x58, 0xb6, 0x61, 0xa2, 0x1a, 0xb4, 0xa4, 0x45, 0x10, 0xc8, 0xec, 0xbc, 0x2a, 0x6c, 0xeb, 0x45,
	0x3d, 0x2b, 0x4e, 0x29, 0xe1, 0x4e, 0x57, 0x9d, 0x1d, 0x79, 0x59, 0x5c, 0x27, 0x35, 0x48, 0xa1,
	0xa5, 0xcc, 0x7c, 0xf8, 0x1c, 0x99, 0xf8, 0xfa, 0x25, 0x32, 0x92, 0x42, 0x5a, 0x3e, 0x3d, 0x8f,
	0x08, 0x3f, 0xce, 0x23, 0xc2, 0xaf, 0xf3, 0x88, 0xf0, 0xed, 0x22, 0x22, 0x9c, 0x5e, 0x44, 0x84,
	0x37, 0x9e, 0x76, 0xb2, 0xec, 0x67, 0xb9, 0xae, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x92,
	0x68, 0x9b, 0x0b, 0x07, 0x00, 0x00,
}
