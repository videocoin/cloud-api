// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profiles/v1/profiles.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProfileId int32

const (
	ProfileIdNone ProfileId = 0
	ProfileIdSD   ProfileId = 1
	ProfileIdHD   ProfileId = 2
	ProfileIdFHD  ProfileId = 3
)

var ProfileId_name = map[int32]string{
	0: "profile_id_none",
	1: "profile_id_sd",
	2: "profile_id_hd",
	3: "profile_id_fhd",
}
var ProfileId_value = map[string]int32{
	"profile_id_none": 0,
	"profile_id_sd":   1,
	"profile_id_hd":   2,
	"profile_id_fhd":  3,
}

func (x ProfileId) String() string {
	return proto.EnumName(ProfileId_name, int32(x))
}
func (ProfileId) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_profiles_5b3e95f08d03fba6, []int{0}
}

type Profiles struct {
	Profiles             []*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Profiles) Reset()         { *m = Profiles{} }
func (m *Profiles) String() string { return proto.CompactTextString(m) }
func (*Profiles) ProtoMessage()    {}
func (*Profiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_profiles_5b3e95f08d03fba6, []int{0}
}
func (m *Profiles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profiles.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Profiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profiles.Merge(dst, src)
}
func (m *Profiles) XXX_Size() int {
	return m.Size()
}
func (m *Profiles) XXX_DiscardUnknown() {
	xxx_messageInfo_Profiles.DiscardUnknown(m)
}

var xxx_messageInfo_Profiles proto.InternalMessageInfo

func (m *Profiles) GetProfiles() []*Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (*Profiles) XXX_MessageName() string {
	return "cloud.api.profiles.v1.Profiles"
}

type Profile struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bitrate              uint32   `protobuf:"varint,3,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	Width                uint32   `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Fps                  float32  `protobuf:"fixed32,5,opt,name=fps,proto3" json:"fps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_profiles_5b3e95f08d03fba6, []int{1}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(dst, src)
}
func (m *Profile) XXX_Size() int {
	return m.Size()
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

func (m *Profile) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Profile) GetFps() float32 {
	if m != nil {
		return m.Fps
	}
	return 0
}

func (*Profile) XXX_MessageName() string {
	return "cloud.api.profiles.v1.Profile"
}
func init() {
	proto.RegisterType((*Profiles)(nil), "cloud.api.profiles.v1.Profiles")
	proto.RegisterType((*Profile)(nil), "cloud.api.profiles.v1.Profile")
	proto.RegisterEnum("cloud.api.profiles.v1.ProfileId", ProfileId_name, ProfileId_value)
}
func (m *Profiles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profiles) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Profiles) > 0 {
		for _, msg := range m.Profiles {
			dAtA[i] = 0xa
			i++
			i = encodeVarintProfiles(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profile) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProfiles(dAtA, i, uint64(m.Id))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProfiles(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Bitrate != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProfiles(dAtA, i, uint64(m.Bitrate))
	}
	if m.Width != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintProfiles(dAtA, i, uint64(m.Width))
	}
	if m.Fps != 0 {
		dAtA[i] = 0x2d
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Fps))))
		i += 4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintProfiles(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Profiles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Profiles) > 0 {
		for _, e := range m.Profiles {
			l = e.Size()
			n += 1 + l + sovProfiles(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProfiles(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProfiles(uint64(l))
	}
	if m.Bitrate != 0 {
		n += 1 + sovProfiles(uint64(m.Bitrate))
	}
	if m.Width != 0 {
		n += 1 + sovProfiles(uint64(m.Width))
	}
	if m.Fps != 0 {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProfiles(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProfiles(x uint64) (n int) {
	return sovProfiles(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Profiles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProfiles
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
			return fmt.Errorf("proto: Profiles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profiles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profiles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiles
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
				return ErrInvalidLengthProfiles
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profiles = append(m.Profiles, &Profile{})
			if err := m.Profiles[len(m.Profiles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProfiles(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProfiles
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
func (m *Profile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProfiles
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
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiles
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
				return ErrInvalidLengthProfiles
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bitrate", wireType)
			}
			m.Bitrate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bitrate |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fps", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Fps = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipProfiles(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProfiles
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
func skipProfiles(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProfiles
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
					return 0, ErrIntOverflowProfiles
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
					return 0, ErrIntOverflowProfiles
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
				return 0, ErrInvalidLengthProfiles
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProfiles
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
				next, err := skipProfiles(dAtA[start:])
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
	ErrInvalidLengthProfiles = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProfiles   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("profiles/v1/profiles.proto", fileDescriptor_profiles_5b3e95f08d03fba6)
}

var fileDescriptor_profiles_5b3e95f08d03fba6 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0x87, 0x9d, 0x44, 0xaf, 0x7a, 0xbc, 0x6a, 0x1c, 0xee, 0x85, 0x21, 0x94, 0x61, 0x90, 0x52,
	0x42, 0xa1, 0x09, 0xb6, 0xbb, 0x2e, 0x8b, 0x88, 0xdd, 0x94, 0x92, 0xee, 0xba, 0x11, 0xe3, 0x24,
	0x66, 0x40, 0x33, 0x41, 0xa3, 0x7d, 0x85, 0xe2, 0x13, 0x74, 0xe3, 0xaa, 0x7d, 0x84, 0xae, 0xba,
	0xea, 0xd2, 0x65, 0x1f, 0xa1, 0xe8, 0x8b, 0x94, 0x8e, 0x31, 0x15, 0xe9, 0xee, 0xf7, 0x9d, 0xf9,
	0xce, 0x1f, 0x18, 0x30, 0xe3, 0x89, 0x0c, 0xc4, 0xc8, 0x9f, 0x3a, 0xf3, 0x96, 0xb3, 0xcb, 0x76,
	0x3c, 0x91, 0x89, 0xc4, 0xff, 0x07, 0x23, 0x39, 0xe3, 0x76, 0x3f, 0x16, 0x76, 0xf6, 0x32, 0x6f,
	0x99, 0x67, 0x43, 0x91, 0x84, 0x33, 0xcf, 0x1e, 0xc8, 0xb1, 0x33, 0x94, 0x43, 0xe9, 0x28, 0xdb,
	0x9b, 0x05, 0x8a, 0x14, 0xa8, 0xb4, 0x9d, 0xd2, 0xec, 0x40, 0xe9, 0x36, 0xed, 0xc6, 0x97, 0x50,
	0xda, 0x4d, 0x22, 0x88, 0xe9, 0x56, 0xe5, 0x9c, 0xda, 0xbf, 0x2e, 0xb1, 0xd3, 0x16, 0x37, 0xf3,
	0x9b, 0x12, 0x8a, 0x69, 0x11, 0xd7, 0x40, 0x13, 0x9c, 0x20, 0x86, 0xac, 0x82, 0xab, 0x09, 0x8e,
	0x31, 0xe4, 0xa3, 0xfe, 0xd8, 0x27, 0x1a, 0x43, 0x56, 0xd9, 0x55, 0x19, 0x13, 0x28, 0x7a, 0x22,
	0x99, 0xf4, 0x13, 0x9f, 0xe8, 0x0c, 0x59, 0x55, 0x77, 0x87, 0xf8, 0x1f, 0x14, 0x1e, 0x04, 0x4f,
	0x42, 0x92, 0x57, 0xf5, 0x2d, 0x60, 0x03, 0xf4, 0x20, 0x9e, 0x92, 0x02, 0x43, 0x96, 0xe6, 0x7e,
	0xc7, 0xd3, 0x57, 0x04, 0xe5, 0x74, 0xe3, 0x35, 0xc7, 0x27, 0x50, 0x4f, 0x4f, 0xe9, 0x09, 0xde,
	0x8b, 0x64, 0xe4, 0x1b, 0x39, 0xb3, 0xb1, 0x58, 0xb2, 0x6a, 0xe6, 0xdc, 0xc8, 0xc8, 0xc7, 0x4d,
	0xa8, 0xee, 0x79, 0x53, 0x6e, 0x20, 0xb3, 0xbe, 0x58, 0xb2, 0x4a, 0x66, 0xdd, 0xb5, 0x0f, 0x9c,
	0x90, 0x1b, 0xda, 0x81, 0xd3, 0x6d, 0xe3, 0x63, 0xa8, 0xed, 0x39, 0x41, 0xc8, 0x0d, 0xdd, 0x34,
	0x16, 0x4b, 0xf6, 0x37, 0x93, 0x3a, 0xdd, 0xb6, 0xd9, 0x78, 0x7c, 0xa6, 0xb9, 0xb7, 0x17, 0xfa,
	0x73, 0xe8, 0xd5, 0xd1, 0x6a, 0x4d, 0xd1, 0xc7, 0x9a, 0xa2, 0xcf, 0x35, 0x45, 0x4f, 0x1b, 0x9a,
	0x7b, 0xdf, 0x50, 0xb4, 0xda, 0x50, 0x74, 0xaf, 0xcd, 0x5b, 0xde, 0x1f, 0xf5, 0x29, 0x17, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0xf0, 0x68, 0xa3, 0xf8, 0x01, 0x00, 0x00,
}
