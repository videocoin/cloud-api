// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transactions/v1/transaction.proto

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		transactions/v1/transaction.proto
		transactions/v1/transaction_service.proto

	It has these top-level messages:
		Transaction
		GetTransactionsRequest
		GetTransactionsResponse
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransactionType int32

const (
	TransactionType_DEPOSIT    TransactionType = 0
	TransactionType_PAYMENT    TransactionType = 1
	TransactionType_WITHDRAWAL TransactionType = 2
)

var TransactionType_name = map[int32]string{
	0: "DEPOSIT",
	1: "PAYMENT",
	2: "WITHDRAWAL",
}
var TransactionType_value = map[string]int32{
	"DEPOSIT":    0,
	"PAYMENT":    1,
	"WITHDRAWAL": 2,
}

func (x TransactionType) String() string {
	return proto.EnumName(TransactionType_name, int32(x))
}
func (TransactionType) EnumDescriptor() ([]byte, []int) { return fileDescriptorTransaction, []int{0} }

type Transaction struct {
	Type      TransactionType             `protobuf:"varint,1,opt,name=type,proto3,enum=cloud.api.transactions.v1.TransactionType" json:"type,omitempty"`
	Hash      string                      `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	From      string                      `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To        string                      `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	CreatedAt *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Value     uint64                      `protobuf:"varint,6,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorTransaction, []int{0} }

func (m *Transaction) GetType() TransactionType {
	if m != nil {
		return m.Type
	}
	return TransactionType_DEPOSIT
}

func (m *Transaction) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Transaction) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Transaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Transaction) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Transaction) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Transaction)(nil), "cloud.api.transactions.v1.Transaction")
	proto.RegisterEnum("cloud.api.transactions.v1.TransactionType", TransactionType_name, TransactionType_value)
}

func init() { proto.RegisterFile("transactions/v1/transaction.proto", fileDescriptorTransaction) }

var fileDescriptorTransaction = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x9b, 0xda, 0x4d, 0x96, 0xc1, 0x1c, 0x61, 0x87, 0xba, 0x43, 0xad, 0x9e, 0xca, 0xc0,
	0x94, 0xcd, 0x93, 0x08, 0x42, 0x65, 0x03, 0x07, 0xfe, 0x19, 0x35, 0x30, 0xf4, 0x22, 0x59, 0x97,
	0x75, 0x85, 0x75, 0x29, 0xed, 0xdb, 0xc2, 0xbe, 0x89, 0x1f, 0xc9, 0xa3, 0x27, 0xcf, 0x32, 0xbf,
	0x88, 0x2c, 0x75, 0x6c, 0x08, 0xde, 0xde, 0xdf, 0xc3, 0xf3, 0x24, 0x3f, 0x7c, 0x0a, 0x29, 0x5f,
	0x66, 0x3c, 0x80, 0x48, 0x2e, 0x33, 0xb7, 0xe8, 0xba, 0x7b, 0x4c, 0x93, 0x54, 0x82, 0x24, 0xc7,
	0xc1, 0x42, 0xe6, 0x53, 0xca, 0x93, 0x88, 0xee, 0x97, 0x69, 0xd1, 0x6d, 0x9f, 0x87, 0x11, 0xcc,
	0xf3, 0x09, 0x0d, 0x64, 0xec, 0x86, 0x32, 0x94, 0xae, 0x5a, 0x4c, 0xf2, 0x99, 0x22, 0x05, 0xea,
	0x2a, 0x5f, 0x6a, 0x9f, 0x84, 0x52, 0x86, 0x0b, 0xb1, 0x6b, 0x41, 0x14, 0x8b, 0x0c, 0x78, 0x9c,
	0x94, 0x85, 0xb3, 0x4f, 0x84, 0xeb, 0x6c, 0xf7, 0x07, 0xb9, 0xc6, 0x06, 0xac, 0x12, 0x61, 0x22,
	0x1b, 0x39, 0x8d, 0x5e, 0x87, 0xfe, 0x6b, 0x42, 0xf7, 0x56, 0x6c, 0x95, 0x08, 0x5f, 0xed, 0x08,
	0xc1, 0xc6, 0x9c, 0x67, 0x73, 0x53, 0xb7, 0x91, 0x53, 0xf3, 0xd5, 0xbd, 0xc9, 0x66, 0xa9, 0x8c,
	0xcd, 0x83, 0x32, 0xdb, 0xdc, 0xa4, 0x81, 0x75, 0x90, 0xa6, 0xa1, 0x12, 0x1d, 0x24, 0xb9, 0xc4,
	0x38, 0x48, 0x05, 0x07, 0x31, 0x7d, 0xe5, 0x60, 0x56, 0x6c, 0xe4, 0xd4, 0x7b, 0x6d, 0x5a, 0xda,
	0xd3, 0xad, 0x3d, 0x65, 0x5b, 0x7b, 0xbf, 0xf6, 0xdb, 0xf6, 0x80, 0xb4, 0x70, 0xa5, 0xe0, 0x8b,
	0x5c, 0x98, 0x55, 0x1b, 0x39, 0x86, 0x5f, 0x42, 0xe7, 0x0a, 0x1f, 0xfd, 0x31, 0x24, 0x75, 0x7c,
	0xd8, 0x1f, 0x8c, 0x1e, 0x9f, 0x86, 0xac, 0xa9, 0x6d, 0x60, 0xe4, 0x3d, 0xdf, 0x0f, 0x1e, 0x58,
	0x13, 0x91, 0x06, 0xc6, 0xe3, 0x21, 0xbb, 0xed, 0xfb, 0xde, 0xd8, 0xbb, 0x6b, 0xea, 0x37, 0xad,
	0xf7, 0xb5, 0xa5, 0x7d, 0xac, 0x2d, 0xed, 0x6b, 0x6d, 0x69, 0x6f, 0xdf, 0x96, 0xf6, 0xa2, 0x17,
	0xdd, 0x49, 0x55, 0x79, 0x5c, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xc4, 0xf6, 0xe4, 0xc2,
	0x01, 0x00, 0x00,
}
