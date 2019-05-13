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
	Value     float64                     `protobuf:"fixed64,6,opt,name=value,proto3" json:"value,omitempty"`
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

func (m *Transaction) GetValue() float64 {
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
	0x18, 0xc6, 0x9b, 0xba, 0x4d, 0x96, 0xc1, 0x1c, 0x61, 0x87, 0xba, 0x43, 0xad, 0x9e, 0xca, 0xc0,
	0x94, 0xcd, 0x93, 0x08, 0x42, 0x65, 0x03, 0x07, 0xfe, 0x19, 0x35, 0x30, 0xf4, 0x22, 0x59, 0x97,
	0x75, 0x85, 0x75, 0x29, 0xed, 0xdb, 0xc2, 0xbe, 0x89, 0x1f, 0xc9, 0xa3, 0x27, 0xcf, 0x32, 0xbf,
	0x88, 0x2c, 0x75, 0xac, 0x08, 0xde, 0xde, 0xdf, 0xc3, 0xf3, 0x24, 0x3f, 0x7c, 0x0a, 0x09, 0x5f,
	0xa5, 0xdc, 0x87, 0x50, 0xae, 0x52, 0x27, 0xef, 0x39, 0x25, 0xa6, 0x71, 0x22, 0x41, 0x92, 0x63,
	0x7f, 0x29, 0xb3, 0x19, 0xe5, 0x71, 0x48, 0xcb, 0x65, 0x9a, 0xf7, 0x3a, 0xe7, 0x41, 0x08, 0x8b,
	0x6c, 0x4a, 0x7d, 0x19, 0x39, 0x81, 0x0c, 0xa4, 0xa3, 0x16, 0xd3, 0x6c, 0xae, 0x48, 0x81, 0xba,
	0x8a, 0x97, 0x3a, 0x27, 0x81, 0x94, 0xc1, 0x52, 0xec, 0x5b, 0x10, 0x46, 0x22, 0x05, 0x1e, 0xc5,
	0x45, 0xe1, 0xec, 0x13, 0xe1, 0x06, 0xdb, 0xff, 0x41, 0xae, 0x71, 0x05, 0xd6, 0xb1, 0x30, 0x90,
	0x85, 0xec, 0x66, 0xbf, 0x4b, 0xff, 0x35, 0xa1, 0xa5, 0x15, 0x5b, 0xc7, 0xc2, 0x53, 0x3b, 0x42,
	0x70, 0x65, 0xc1, 0xd3, 0x85, 0xa1, 0x5b, 0xc8, 0xae, 0x7b, 0xea, 0xde, 0x66, 0xf3, 0x44, 0x46,
	0xc6, 0x41, 0x91, 0x6d, 0x6f, 0xd2, 0xc4, 0x3a, 0x48, 0xa3, 0xa2, 0x12, 0x1d, 0x24, 0xb9, 0xc4,
	0xd8, 0x4f, 0x04, 0x07, 0x31, 0x7b, 0xe5, 0x60, 0x54, 0x2d, 0x64, 0x37, 0xfa, 0x1d, 0x5a, 0xd8,
	0xd3, 0x9d, 0x3d, 0x65, 0x3b, 0x7b, 0xaf, 0xfe, 0xdb, 0x76, 0x81, 0xb4, 0x71, 0x35, 0xe7, 0xcb,
	0x4c, 0x18, 0x35, 0x0b, 0xd9, 0xc8, 0x2b, 0xa0, 0x7b, 0x85, 0x8f, 0xfe, 0x18, 0x92, 0x06, 0x3e,
	0x1c, 0x0c, 0xc7, 0x8f, 0x4f, 0x23, 0xd6, 0xd2, 0xb6, 0x30, 0x76, 0x9f, 0xef, 0x87, 0x0f, 0xac,
	0x85, 0x48, 0x13, 0xe3, 0xc9, 0x88, 0xdd, 0x0e, 0x3c, 0x77, 0xe2, 0xde, 0xb5, 0xf4, 0x9b, 0xf6,
	0xfb, 0xc6, 0xd4, 0x3e, 0x36, 0xa6, 0xf6, 0xb5, 0x31, 0xb5, 0xb7, 0x6f, 0x53, 0x7b, 0xd1, 0xf3,
	0xde, 0xb4, 0xa6, 0x3c, 0x2e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x05, 0xce, 0xe5, 0x72, 0xc2,
	0x01, 0x00, 0x00,
}
