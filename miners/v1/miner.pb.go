// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: miners/v1/miner.proto

package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MinerStatus int32

const (
	MinerStatusNew     MinerStatus = 0
	MinerStatusOffline MinerStatus = 1
	MinerStatusIdle    MinerStatus = 2
	MinerStatusBusy    MinerStatus = 3
)

var MinerStatus_name = map[int32]string{
	0: "NEW",
	1: "OFFLINE",
	2: "IDLE",
	3: "BUSY",
}

var MinerStatus_value = map[string]int32{
	"NEW":     0,
	"OFFLINE": 1,
	"IDLE":    2,
	"BUSY":    3,
}

func (x MinerStatus) String() string {
	return proto.EnumName(MinerStatus_name, int32(x))
}

func (MinerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7d99a35436329cc, []int{0}
}

type SystemInfo struct {
	CpuCores             float64  `protobuf:"fixed64,1,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	CpuFreq              float64  `protobuf:"fixed64,2,opt,name=cpu_freq,json=cpuFreq,proto3" json:"cpu_freq,omitempty"`
	CpuUsage             float64  `protobuf:"fixed64,3,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	MemUsage             float64  `protobuf:"fixed64,10,opt,name=mem_usage,json=memUsage,proto3" json:"mem_usage,omitempty"`
	MemTotal             float64  `protobuf:"fixed64,11,opt,name=mem_total,json=memTotal,proto3" json:"mem_total,omitempty"`
	Latitude             float64  `protobuf:"fixed64,12,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,13,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Hw                   string   `protobuf:"bytes,14,opt,name=hw,proto3" json:"hw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemInfo) Reset()         { *m = SystemInfo{} }
func (m *SystemInfo) String() string { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()    {}
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7d99a35436329cc, []int{0}
}
func (m *SystemInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SystemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SystemInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SystemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemInfo.Merge(m, src)
}
func (m *SystemInfo) XXX_Size() int {
	return m.Size()
}
func (m *SystemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SystemInfo proto.InternalMessageInfo

func (m *SystemInfo) GetCpuCores() float64 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *SystemInfo) GetCpuFreq() float64 {
	if m != nil {
		return m.CpuFreq
	}
	return 0
}

func (m *SystemInfo) GetCpuUsage() float64 {
	if m != nil {
		return m.CpuUsage
	}
	return 0
}

func (m *SystemInfo) GetMemUsage() float64 {
	if m != nil {
		return m.MemUsage
	}
	return 0
}

func (m *SystemInfo) GetMemTotal() float64 {
	if m != nil {
		return m.MemTotal
	}
	return 0
}

func (m *SystemInfo) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *SystemInfo) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *SystemInfo) GetHw() string {
	if m != nil {
		return m.Hw
	}
	return ""
}

func (*SystemInfo) XXX_MessageName() string {
	return "cloud.api.miners.v1.SystemInfo"
}

type CryptoInfo struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance              string   `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	SelfStake            string   `protobuf:"bytes,3,opt,name=self_stake,json=selfStake,proto3" json:"self_stake,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoInfo) Reset()         { *m = CryptoInfo{} }
func (m *CryptoInfo) String() string { return proto.CompactTextString(m) }
func (*CryptoInfo) ProtoMessage()    {}
func (*CryptoInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7d99a35436329cc, []int{1}
}
func (m *CryptoInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CryptoInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CryptoInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CryptoInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoInfo.Merge(m, src)
}
func (m *CryptoInfo) XXX_Size() int {
	return m.Size()
}
func (m *CryptoInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoInfo proto.InternalMessageInfo

func (m *CryptoInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CryptoInfo) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *CryptoInfo) GetSelfStake() string {
	if m != nil {
		return m.SelfStake
	}
	return ""
}

func (*CryptoInfo) XXX_MessageName() string {
	return "cloud.api.miners.v1.CryptoInfo"
}

type CapacityInfo struct {
	Encode               int32    `protobuf:"varint,1,opt,name=encode,proto3" json:"encode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CapacityInfo) Reset()         { *m = CapacityInfo{} }
func (m *CapacityInfo) String() string { return proto.CompactTextString(m) }
func (*CapacityInfo) ProtoMessage()    {}
func (*CapacityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7d99a35436329cc, []int{2}
}
func (m *CapacityInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CapacityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CapacityInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CapacityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CapacityInfo.Merge(m, src)
}
func (m *CapacityInfo) XXX_Size() int {
	return m.Size()
}
func (m *CapacityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CapacityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CapacityInfo proto.InternalMessageInfo

func (m *CapacityInfo) GetEncode() int32 {
	if m != nil {
		return m.Encode
	}
	return 0
}

func (*CapacityInfo) XXX_MessageName() string {
	return "cloud.api.miners.v1.CapacityInfo"
}
func init() {
	proto.RegisterEnum("cloud.api.miners.v1.MinerStatus", MinerStatus_name, MinerStatus_value)
	golang_proto.RegisterEnum("cloud.api.miners.v1.MinerStatus", MinerStatus_name, MinerStatus_value)
	proto.RegisterType((*SystemInfo)(nil), "cloud.api.miners.v1.SystemInfo")
	golang_proto.RegisterType((*SystemInfo)(nil), "cloud.api.miners.v1.SystemInfo")
	proto.RegisterType((*CryptoInfo)(nil), "cloud.api.miners.v1.CryptoInfo")
	golang_proto.RegisterType((*CryptoInfo)(nil), "cloud.api.miners.v1.CryptoInfo")
	proto.RegisterType((*CapacityInfo)(nil), "cloud.api.miners.v1.CapacityInfo")
	golang_proto.RegisterType((*CapacityInfo)(nil), "cloud.api.miners.v1.CapacityInfo")
}

func init() { proto.RegisterFile("miners/v1/miner.proto", fileDescriptor_a7d99a35436329cc) }
func init() { golang_proto.RegisterFile("miners/v1/miner.proto", fileDescriptor_a7d99a35436329cc) }

var fileDescriptor_a7d99a35436329cc = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xd1, 0x8a, 0xd3, 0x4c,
	0x14, 0xc7, 0x77, 0xd2, 0xef, 0xdb, 0x6e, 0xce, 0xae, 0xb5, 0x4c, 0x71, 0x89, 0x59, 0x0d, 0xa1,
	0x82, 0x2c, 0x82, 0x29, 0xc5, 0x37, 0x68, 0x6d, 0xa1, 0xb0, 0x76, 0x21, 0x75, 0x11, 0xbd, 0x29,
	0xd3, 0x64, 0x92, 0x06, 0x93, 0x4c, 0x36, 0x33, 0xd3, 0xd2, 0x37, 0x90, 0xbe, 0x43, 0x6f, 0xd4,
	0xa7, 0xf0, 0xca, 0xcb, 0xbd, 0xf4, 0x11, 0xa4, 0x8b, 0xef, 0x21, 0x33, 0xd9, 0xae, 0x15, 0xef,
	0xce, 0xef, 0xfc, 0xce, 0x64, 0xce, 0x9f, 0x0c, 0x3c, 0xca, 0x92, 0x9c, 0x96, 0xbc, 0xb3, 0xe8,
	0x76, 0x74, 0xe5, 0x15, 0x25, 0x13, 0x0c, 0xb7, 0x82, 0x94, 0xc9, 0xd0, 0x23, 0x45, 0xe2, 0x55,
	0x03, 0xde, 0xa2, 0x6b, 0xbf, 0x8c, 0x13, 0x31, 0x97, 0x33, 0x2f, 0x60, 0x59, 0x27, 0x66, 0x31,
	0xeb, 0xe8, 0xd9, 0x99, 0x8c, 0x34, 0x69, 0xd0, 0x55, 0xf5, 0x8d, 0xf6, 0x2f, 0x04, 0x30, 0x59,
	0x71, 0x41, 0xb3, 0x51, 0x1e, 0x31, 0x7c, 0x06, 0x66, 0x50, 0xc8, 0x69, 0xc0, 0x4a, 0xca, 0x2d,
	0xe4, 0xa2, 0x73, 0xe4, 0x1f, 0x05, 0x85, 0xec, 0x2b, 0xc6, 0x8f, 0x41, 0xd5, 0xd3, 0xa8, 0xa4,
	0xd7, 0x96, 0xa1, 0x5d, 0x3d, 0x28, 0xe4, 0xb0, 0xa4, 0xd7, 0xbb, 0x73, 0x92, 0x93, 0x98, 0x5a,
	0xb5, 0xfb, 0x73, 0x57, 0x8a, 0x95, 0xcc, 0x68, 0x76, 0x27, 0xa1, 0x92, 0x19, 0xcd, 0xfe, 0x92,
	0x82, 0x09, 0x92, 0x5a, 0xc7, 0xf7, 0xf2, 0xad, 0x62, 0x6c, 0xc3, 0x51, 0x4a, 0x44, 0x22, 0x64,
	0x48, 0xad, 0x93, 0xca, 0xed, 0x18, 0x3f, 0x01, 0x33, 0x65, 0x79, 0x5c, 0xc9, 0x07, 0x5a, 0xfe,
	0x69, 0xe0, 0x06, 0x18, 0xf3, 0xa5, 0xd5, 0x70, 0xd1, 0xb9, 0xe9, 0x1b, 0xf3, 0x65, 0x7b, 0x0a,
	0xd0, 0x2f, 0x57, 0x85, 0x60, 0x3a, 0xa6, 0x05, 0x75, 0x12, 0x86, 0x25, 0xe5, 0x55, 0x48, 0xd3,
	0xdf, 0xa1, 0x32, 0x33, 0x92, 0x92, 0x3c, 0xa0, 0x3a, 0xa2, 0xe9, 0xef, 0x10, 0x3f, 0x05, 0xe0,
	0x34, 0x8d, 0xa6, 0x5c, 0x90, 0x8f, 0x55, 0x46, 0xd3, 0x37, 0x55, 0x67, 0xa2, 0x1a, 0xed, 0xe7,
	0x70, 0xd2, 0x27, 0x05, 0x09, 0x12, 0xb1, 0xd2, 0x57, 0x9c, 0xc2, 0x21, 0xcd, 0x03, 0x16, 0x52,
	0x7d, 0xc3, 0xff, 0xfe, 0x1d, 0xbd, 0xf8, 0x8c, 0xe0, 0xf8, 0x8d, 0xfa, 0x5b, 0x13, 0x41, 0x84,
	0xe4, 0xf8, 0x0c, 0x6a, 0xe3, 0xc1, 0xbb, 0xe6, 0x81, 0x8d, 0xd7, 0x1b, 0xb7, 0xb1, 0x67, 0xc6,
	0x74, 0x89, 0x9f, 0x41, 0xfd, 0x72, 0x38, 0xbc, 0x18, 0x8d, 0x07, 0x4d, 0x64, 0x9f, 0xae, 0x37,
	0x2e, 0xde, 0x1b, 0xb8, 0x8c, 0xa2, 0x34, 0xc9, 0xd5, 0x62, 0xff, 0x8d, 0x5e, 0x5f, 0x0c, 0x9a,
	0x86, 0xdd, 0x5a, 0x6f, 0xdc, 0x87, 0x7b, 0x13, 0xa3, 0x30, 0xd5, 0xba, 0x77, 0x35, 0x79, 0xdf,
	0xac, 0xfd, 0xa3, 0x7b, 0x92, 0xaf, 0xec, 0xd6, 0xa7, 0x2f, 0xce, 0xc1, 0xb7, 0xaf, 0xce, 0xfe,
	0x52, 0x3d, 0xeb, 0x66, 0xeb, 0xa0, 0x1f, 0x5b, 0x07, 0xfd, 0xdc, 0x3a, 0xe8, 0xfb, 0xad, 0x83,
	0x6e, 0x6e, 0x1d, 0xf4, 0xc1, 0x58, 0x74, 0x67, 0x87, 0xfa, 0xd9, 0xbc, 0xfa, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0x7d, 0x06, 0x57, 0x93, 0x02, 0x00, 0x00,
}

func (m *SystemInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SystemInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SystemInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Hw) > 0 {
		i -= len(m.Hw)
		copy(dAtA[i:], m.Hw)
		i = encodeVarintMiner(dAtA, i, uint64(len(m.Hw)))
		i--
		dAtA[i] = 0x72
	}
	if m.Longitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Longitude))))
		i--
		dAtA[i] = 0x69
	}
	if m.Latitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Latitude))))
		i--
		dAtA[i] = 0x61
	}
	if m.MemTotal != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MemTotal))))
		i--
		dAtA[i] = 0x59
	}
	if m.MemUsage != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MemUsage))))
		i--
		dAtA[i] = 0x51
	}
	if m.CpuUsage != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CpuUsage))))
		i--
		dAtA[i] = 0x19
	}
	if m.CpuFreq != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CpuFreq))))
		i--
		dAtA[i] = 0x11
	}
	if m.CpuCores != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CpuCores))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *CryptoInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CryptoInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CryptoInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SelfStake) > 0 {
		i -= len(m.SelfStake)
		copy(dAtA[i:], m.SelfStake)
		i = encodeVarintMiner(dAtA, i, uint64(len(m.SelfStake)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintMiner(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMiner(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CapacityInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CapacityInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CapacityInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Encode != 0 {
		i = encodeVarintMiner(dAtA, i, uint64(m.Encode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMiner(dAtA []byte, offset int, v uint64) int {
	offset -= sovMiner(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SystemInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CpuCores != 0 {
		n += 9
	}
	if m.CpuFreq != 0 {
		n += 9
	}
	if m.CpuUsage != 0 {
		n += 9
	}
	if m.MemUsage != 0 {
		n += 9
	}
	if m.MemTotal != 0 {
		n += 9
	}
	if m.Latitude != 0 {
		n += 9
	}
	if m.Longitude != 0 {
		n += 9
	}
	l = len(m.Hw)
	if l > 0 {
		n += 1 + l + sovMiner(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CryptoInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMiner(uint64(l))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovMiner(uint64(l))
	}
	l = len(m.SelfStake)
	if l > 0 {
		n += 1 + l + sovMiner(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CapacityInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Encode != 0 {
		n += 1 + sovMiner(uint64(m.Encode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMiner(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMiner(x uint64) (n int) {
	return sovMiner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SystemInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMiner
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
			return fmt.Errorf("proto: SystemInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SystemInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuCores", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.CpuCores = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuFreq", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.CpuFreq = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuUsage", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.CpuUsage = float64(math.Float64frombits(v))
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemUsage", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MemUsage = float64(math.Float64frombits(v))
		case 11:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemTotal", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MemTotal = float64(math.Float64frombits(v))
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Latitude = float64(math.Float64frombits(v))
		case 13:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Longitude = float64(math.Float64frombits(v))
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hw", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiner
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
				return ErrInvalidLengthMiner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMiner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hw = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMiner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMiner
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMiner
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
func (m *CryptoInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMiner
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
			return fmt.Errorf("proto: CryptoInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CryptoInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiner
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
				return ErrInvalidLengthMiner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMiner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiner
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
				return ErrInvalidLengthMiner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMiner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfStake", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiner
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
				return ErrInvalidLengthMiner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMiner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SelfStake = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMiner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMiner
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMiner
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
func (m *CapacityInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMiner
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
			return fmt.Errorf("proto: CapacityInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CapacityInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encode", wireType)
			}
			m.Encode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Encode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMiner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMiner
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMiner
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
func skipMiner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMiner
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
					return 0, ErrIntOverflowMiner
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMiner
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
				return 0, ErrInvalidLengthMiner
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMiner
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMiner
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMiner        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMiner          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMiner = fmt.Errorf("proto: unexpected end of group")
)
