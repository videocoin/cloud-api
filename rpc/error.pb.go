// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/error.proto

package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ValidationError struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationError) Reset()         { *m = ValidationError{} }
func (m *ValidationError) String() string { return proto.CompactTextString(m) }
func (*ValidationError) ProtoMessage()    {}
func (*ValidationError) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_d237920ec1399604, []int{0}
}
func (m *ValidationError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidationError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidationError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ValidationError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationError.Merge(dst, src)
}
func (m *ValidationError) XXX_Size() int {
	return m.Size()
}
func (m *ValidationError) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationError.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationError proto.InternalMessageInfo

func (m *ValidationError) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ValidationError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MultiValidationError struct {
	Errors               []*ValidationError `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MultiValidationError) Reset()         { *m = MultiValidationError{} }
func (m *MultiValidationError) String() string { return proto.CompactTextString(m) }
func (*MultiValidationError) ProtoMessage()    {}
func (*MultiValidationError) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_d237920ec1399604, []int{1}
}
func (m *MultiValidationError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiValidationError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiValidationError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *MultiValidationError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiValidationError.Merge(dst, src)
}
func (m *MultiValidationError) XXX_Size() int {
	return m.Size()
}
func (m *MultiValidationError) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiValidationError.DiscardUnknown(m)
}

var xxx_messageInfo_MultiValidationError proto.InternalMessageInfo

func (m *MultiValidationError) GetErrors() []*ValidationError {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidationError)(nil), "cloud.api.rpc.ValidationError")
	proto.RegisterType((*MultiValidationError)(nil), "cloud.api.rpc.MultiValidationError")
}
func (m *ValidationError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidationError) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Field) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Field)))
		i += copy(dAtA[i:], m.Field)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *MultiValidationError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiValidationError) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for _, msg := range m.Errors {
			dAtA[i] = 0xa
			i++
			i = encodeVarintError(dAtA, i, uint64(msg.Size()))
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

func encodeVarintError(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ValidationError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MultiValidationError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for _, e := range m.Errors {
			l = e.Size()
			n += 1 + l + sovError(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovError(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozError(x uint64) (n int) {
	return sovError(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidationError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: ValidationError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidationError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthError
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
func (m *MultiValidationError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: MultiValidationError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiValidationError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, &ValidationError{})
			if err := m.Errors[len(m.Errors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthError
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
func skipError(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
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
				return 0, ErrInvalidLengthError
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowError
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
				next, err := skipError(dAtA[start:])
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
	ErrInvalidLengthError = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowError   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpc/error.proto", fileDescriptor_error_d237920ec1399604) }

var fileDescriptor_error_d237920ec1399604 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2a, 0x48, 0xd6,
	0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0xce, 0xc9,
	0x2f, 0x4d, 0xd1, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x2a, 0x48, 0x56, 0x72, 0xe4, 0xe2, 0x0f, 0x4b,
	0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x73, 0x05, 0xa9, 0x13, 0x12, 0xe1, 0x62, 0x4d,
	0xcb, 0x4c, 0xcd, 0x49, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x24, 0xb8,
	0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x98, 0xc0, 0xe2, 0x30, 0xae, 0x92, 0x1f,
	0x97, 0x88, 0x6f, 0x69, 0x4e, 0x49, 0x26, 0xba, 0x39, 0x66, 0x5c, 0x6c, 0x60, 0x8b, 0x8b, 0x25,
	0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xe4, 0xf4, 0x50, 0xac, 0xd6, 0x43, 0x53, 0x1f, 0x04, 0x55,
	0xed, 0x24, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46,
	0x31, 0x17, 0x15, 0x24, 0x27, 0xb1, 0x81, 0xdd, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x80,
	0xeb, 0x8b, 0x07, 0xce, 0x00, 0x00, 0x00,
}
