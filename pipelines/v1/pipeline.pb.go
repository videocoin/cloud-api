// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pipelines/v1/pipeline.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	v11 "github.com/videocoin/cloud-api/jobs/v1"
	v1 "github.com/videocoin/cloud-api/profiles/v1"
	io "io"
	math "math"
	time "time"
)

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

type Pipeline struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"type:varchar(36);primary_key"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"type:varchar(100)"`
	UserId               string       `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" gorm:"type:varchar(36)"`
	ProfileId            v1.ProfileId `protobuf:"varint,4,opt,name=profile_id,json=profileId,proto3,enum=cloud.api.profiles.v1.ProfileId" json:"profile_id,omitempty"`
	CreatedAt            *time.Time   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	IsActivated          bool         `protobuf:"varint,13,opt,name=is_activated,json=isActivated,proto3" json:"is_activated,omitempty"`
	AccessCode           string       `protobuf:"bytes,14,opt,name=access_code,json=accessCode,proto3" json:"access_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}
func (*Pipeline) Descriptor() ([]byte, []int) {
	return fileDescriptor_8672028def87e47d, []int{0}
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
func (m *Pipeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipeline.Merge(m, src)
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

func (m *Pipeline) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Pipeline) GetIsActivated() bool {
	if m != nil {
		return m.IsActivated
	}
	return false
}

func (m *Pipeline) GetAccessCode() string {
	if m != nil {
		return m.AccessCode
	}
	return ""
}

func (*Pipeline) XXX_MessageName() string {
	return "cloud.api.pipelines.v1.Pipeline"
}

type PipelineProfile struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProfileId            v1.ProfileId     `protobuf:"varint,4,opt,name=profile_id,json=profileId,proto3,enum=cloud.api.profiles.v1.ProfileId" json:"profile_id,omitempty"`
	Jobs                 *v11.JobProfiles `protobuf:"bytes,5,opt,name=jobs,proto3" json:"jobs,omitempty"`
	CreatedAt            *time.Time       `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	IsActivated          bool             `protobuf:"varint,13,opt,name=is_activated,json=isActivated,proto3" json:"is_activated,omitempty"`
	AccessCode           string           `protobuf:"bytes,14,opt,name=access_code,json=accessCode,proto3" json:"access_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PipelineProfile) Reset()         { *m = PipelineProfile{} }
func (m *PipelineProfile) String() string { return proto.CompactTextString(m) }
func (*PipelineProfile) ProtoMessage()    {}
func (*PipelineProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_8672028def87e47d, []int{1}
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
func (m *PipelineProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineProfile.Merge(m, src)
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

func (m *PipelineProfile) GetProfileId() v1.ProfileId {
	if m != nil {
		return m.ProfileId
	}
	return v1.ProfileIdNone
}

func (m *PipelineProfile) GetJobs() *v11.JobProfiles {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *PipelineProfile) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PipelineProfile) GetIsActivated() bool {
	if m != nil {
		return m.IsActivated
	}
	return false
}

func (m *PipelineProfile) GetAccessCode() string {
	if m != nil {
		return m.AccessCode
	}
	return ""
}

func (*PipelineProfile) XXX_MessageName() string {
	return "cloud.api.pipelines.v1.PipelineProfile"
}
func init() {
	proto.RegisterType((*Pipeline)(nil), "cloud.api.pipelines.v1.Pipeline")
	golang_proto.RegisterType((*Pipeline)(nil), "cloud.api.pipelines.v1.Pipeline")
	proto.RegisterType((*PipelineProfile)(nil), "cloud.api.pipelines.v1.PipelineProfile")
	golang_proto.RegisterType((*PipelineProfile)(nil), "cloud.api.pipelines.v1.PipelineProfile")
}

func init() { proto.RegisterFile("pipelines/v1/pipeline.proto", fileDescriptor_8672028def87e47d) }
func init() { golang_proto.RegisterFile("pipelines/v1/pipeline.proto", fileDescriptor_8672028def87e47d) }

var fileDescriptor_8672028def87e47d = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xb1, 0x6e, 0xd3, 0x40,
	0x18, 0xc7, 0x75, 0x26, 0x94, 0xe6, 0x02, 0x41, 0xf2, 0x00, 0x56, 0x8a, 0x6c, 0x63, 0x06, 0xdc,
	0xa1, 0xe7, 0x38, 0x45, 0x20, 0xca, 0x80, 0x1a, 0xa6, 0x32, 0x55, 0x16, 0x13, 0x8b, 0x75, 0xf6,
	0x5d, 0xdd, 0x2b, 0x71, 0xee, 0x74, 0xbe, 0x58, 0xca, 0x5b, 0xf0, 0x26, 0xbc, 0x02, 0x63, 0x47,
	0x9e, 0x20, 0xa0, 0xf4, 0x09, 0xc8, 0x13, 0x20, 0x9f, 0xcf, 0x34, 0x43, 0x24, 0x06, 0x96, 0x6e,
	0xdf, 0xf7, 0xe9, 0xff, 0xfb, 0x72, 0xdf, 0x2f, 0x86, 0x07, 0x82, 0x09, 0x3a, 0x63, 0x73, 0x5a,
	0x45, 0x75, 0x1c, 0x75, 0x0d, 0x12, 0x92, 0x2b, 0x6e, 0x3f, 0xc9, 0x67, 0x7c, 0x41, 0x10, 0x16,
	0x0c, 0xfd, 0x8d, 0xa1, 0x3a, 0x1e, 0xbd, 0x2d, 0x98, 0xba, 0x5c, 0x64, 0x28, 0xe7, 0x65, 0x54,
	0x33, 0x42, 0x79, 0xce, 0xd9, 0x3c, 0xd2, 0xe1, 0x23, 0x2c, 0x58, 0x24, 0x24, 0xbf, 0x60, 0x33,
	0xb3, 0xd2, 0xd4, 0xed, 0xca, 0xd1, 0xf8, 0x1f, 0xe8, 0x15, 0xcf, 0x34, 0x76, 0xc5, 0x33, 0x43,
	0x1c, 0x6d, 0x11, 0x05, 0x2f, 0x78, 0xa4, 0xc7, 0xd9, 0xe2, 0x42, 0x77, 0xba, 0xd1, 0x95, 0x89,
	0x7b, 0x05, 0xe7, 0xc5, 0x8c, 0xde, 0xa6, 0x14, 0x2b, 0x69, 0xa5, 0x70, 0x29, 0xda, 0x40, 0xf0,
	0xdb, 0x82, 0xfb, 0xe7, 0xe6, 0x1a, 0xfb, 0x0d, 0xb4, 0x18, 0x71, 0x80, 0x0f, 0xc2, 0xfe, 0xf4,
	0xe5, 0x66, 0xe5, 0xbd, 0x28, 0xb8, 0x2c, 0x4f, 0x02, 0xb5, 0x14, 0xf4, 0xa4, 0xc6, 0x32, 0xbf,
	0xc4, 0x32, 0x3c, 0x7e, 0x7d, 0xf8, 0x4e, 0x48, 0x56, 0x62, 0xb9, 0x4c, 0xbf, 0xd0, 0x65, 0x90,
	0x58, 0x8c, 0xd8, 0x63, 0xd8, 0x9b, 0xe3, 0x92, 0x3a, 0x96, 0x46, 0x9f, 0x6d, 0x56, 0x9e, 0xb3,
	0x03, 0x8d, 0xc7, 0xe3, 0xc3, 0x20, 0xd1, 0x49, 0xfb, 0x15, 0x7c, 0xb0, 0xa8, 0xa8, 0x4c, 0x19,
	0x71, 0xee, 0x69, 0xe8, 0x60, 0xb3, 0xf2, 0x9e, 0xee, 0xfe, 0xbd, 0x20, 0xd9, 0x6b, 0xb2, 0x67,
	0xc4, 0x7e, 0x0f, 0xa1, 0x31, 0xd8, 0x80, 0x3d, 0x1f, 0x84, 0xc3, 0x89, 0x8f, 0xb6, 0xfe, 0x97,
	0x4e, 0x6f, 0x1d, 0xa3, 0xf3, 0xb6, 0x3e, 0x23, 0x49, 0x5f, 0x74, 0x65, 0xb3, 0x20, 0x97, 0x14,
	0x2b, 0x4a, 0x52, 0xac, 0x1c, 0xe8, 0x83, 0x70, 0x30, 0x19, 0xa1, 0x56, 0x12, 0xea, 0x24, 0xa1,
	0x4f, 0x9d, 0xa4, 0x69, 0xef, 0xeb, 0x4f, 0x0f, 0x24, 0x7d, 0xc3, 0x9c, 0x2a, 0xfb, 0x39, 0x7c,
	0xc8, 0xaa, 0x14, 0xe7, 0x8a, 0xd5, 0xcd, 0xc4, 0x79, 0xe4, 0x83, 0x70, 0x3f, 0x19, 0xb0, 0xea,
	0xb4, 0x1b, 0xd9, 0x1e, 0x1c, 0xe0, 0x3c, 0xa7, 0x55, 0x95, 0xe6, 0x9c, 0x50, 0x67, 0xd8, 0x9c,
	0x97, 0xc0, 0x76, 0xf4, 0x81, 0x13, 0x1a, 0x7c, 0xb3, 0xe0, 0xe3, 0xce, 0xb9, 0x79, 0xa5, 0x3d,
	0xbc, 0x55, 0xaf, 0x8d, 0xda, 0xdb, 0x46, 0x8d, 0xb3, 0xff, 0xbe, 0x7e, 0x02, 0x7b, 0xcd, 0x17,
	0xe5, 0xdc, 0xd7, 0x77, 0xbb, 0x5b, 0x68, 0x33, 0x6e, 0xb0, 0x8f, 0x3c, 0x33, 0x64, 0x95, 0xe8,
	0xec, 0x9d, 0x30, 0x36, 0x75, 0xae, 0xd7, 0x2e, 0xf8, 0xb1, 0x76, 0xc1, 0xaf, 0xb5, 0x0b, 0xbe,
	0xdf, 0xb8, 0xe0, 0xfa, 0xc6, 0x05, 0x9f, 0xad, 0x3a, 0xce, 0xf6, 0xf4, 0x13, 0x8e, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x13, 0x6c, 0x62, 0xcc, 0xba, 0x03, 0x00, 0x00,
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
	if m.IsActivated {
		dAtA[i] = 0x68
		i++
		if m.IsActivated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.AccessCode) > 0 {
		dAtA[i] = 0x72
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.AccessCode)))
		i += copy(dAtA[i:], m.AccessCode)
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
	if m.ProfileId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(m.ProfileId))
	}
	if m.Jobs != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(m.Jobs.Size()))
		n2, err := m.Jobs.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.CreatedAt != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)))
		n3, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.IsActivated {
		dAtA[i] = 0x68
		i++
		if m.IsActivated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.AccessCode) > 0 {
		dAtA[i] = 0x72
		i++
		i = encodeVarintPipeline(dAtA, i, uint64(len(m.AccessCode)))
		i += copy(dAtA[i:], m.AccessCode)
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
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.IsActivated {
		n += 2
	}
	l = len(m.AccessCode)
	if l > 0 {
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
	if m.ProfileId != 0 {
		n += 1 + sovPipeline(uint64(m.ProfileId))
	}
	if m.Jobs != nil {
		l = m.Jobs.Size()
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovPipeline(uint64(l))
	}
	if m.IsActivated {
		n += 2
	}
	l = len(m.AccessCode)
	if l > 0 {
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
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActivated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActivated = bool(v != 0)
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessCode", wireType)
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
			m.AccessCode = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field Jobs", wireType)
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
			if m.Jobs == nil {
				m.Jobs = &v11.JobProfiles{}
			}
			if err := m.Jobs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActivated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPipeline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActivated = bool(v != 0)
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessCode", wireType)
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
			m.AccessCode = string(dAtA[iNdEx:postIndex])
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
