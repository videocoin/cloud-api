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
		GetTransactionRequest
		GetTransactionResponse
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

type TransactionSource int32

const (
	TransactionSource_UNDETECTED TransactionSource = 0
	TransactionSource_ETH        TransactionSource = 1
	TransactionSource_VID        TransactionSource = 2
)

var TransactionSource_name = map[int32]string{
	0: "UNDETECTED",
	1: "ETH",
	2: "VID",
}
var TransactionSource_value = map[string]int32{
	"UNDETECTED": 0,
	"ETH":        1,
	"VID":        2,
}

func (x TransactionSource) String() string {
	return proto.EnumName(TransactionSource_name, int32(x))
}
func (TransactionSource) EnumDescriptor() ([]byte, []int) { return fileDescriptorTransaction, []int{0} }

type TransactionType int32

const (
	TransactionType_UNKNOWN               TransactionType = 0
	TransactionType_DEPOSIT               TransactionType = 1
	TransactionType_ESCROW                TransactionType = 2
	TransactionType_STREAM_ENDED          TransactionType = 3
	TransactionType_INPUT_CHUNK_ADDED     TransactionType = 4
	TransactionType_CHUNK_PROOF_SUBMITTED TransactionType = 5
	TransactionType_CHUNK_PROOF_VALIDATED TransactionType = 6
	TransactionType_ACCOUNT_FUNDED        TransactionType = 7
)

var TransactionType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DEPOSIT",
	2: "ESCROW",
	3: "STREAM_ENDED",
	4: "INPUT_CHUNK_ADDED",
	5: "CHUNK_PROOF_SUBMITTED",
	6: "CHUNK_PROOF_VALIDATED",
	7: "ACCOUNT_FUNDED",
}
var TransactionType_value = map[string]int32{
	"UNKNOWN":               0,
	"DEPOSIT":               1,
	"ESCROW":                2,
	"STREAM_ENDED":          3,
	"INPUT_CHUNK_ADDED":     4,
	"CHUNK_PROOF_SUBMITTED": 5,
	"CHUNK_PROOF_VALIDATED": 6,
	"ACCOUNT_FUNDED":        7,
}

func (x TransactionType) String() string {
	return proto.EnumName(TransactionType_name, int32(x))
}
func (TransactionType) EnumDescriptor() ([]byte, []int) { return fileDescriptorTransaction, []int{1} }

type Transaction struct {
	Source    TransactionSource           `protobuf:"varint,1,opt,name=source,proto3,enum=cloud.api.transactions.v1.TransactionSource" json:"source,omitempty"`
	Hash      string                      `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	From      string                      `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To        string                      `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	CreatedAt *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Value     string                      `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	Type      TransactionType             `protobuf:"varint,7,opt,name=type,proto3,enum=cloud.api.transactions.v1.TransactionType" json:"type,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorTransaction, []int{0} }

func (m *Transaction) GetSource() TransactionSource {
	if m != nil {
		return m.Source
	}
	return TransactionSource_UNDETECTED
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

func (m *Transaction) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Transaction) GetType() TransactionType {
	if m != nil {
		return m.Type
	}
	return TransactionType_UNKNOWN
}

func init() {
	proto.RegisterType((*Transaction)(nil), "cloud.api.transactions.v1.Transaction")
	proto.RegisterEnum("cloud.api.transactions.v1.TransactionSource", TransactionSource_name, TransactionSource_value)
	proto.RegisterEnum("cloud.api.transactions.v1.TransactionType", TransactionType_name, TransactionType_value)
}

func init() { proto.RegisterFile("transactions/v1/transaction.proto", fileDescriptorTransaction) }

var fileDescriptorTransaction = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0xc1, 0x6a, 0xdb, 0x40,
	0x14, 0x94, 0x14, 0x5b, 0x26, 0xcf, 0xc5, 0x55, 0x96, 0x04, 0x14, 0x1f, 0x5c, 0xb7, 0x27, 0x63,
	0xda, 0x35, 0x4e, 0xe9, 0xa1, 0x97, 0x82, 0xa2, 0xdd, 0x10, 0x91, 0x46, 0x32, 0xd2, 0x2a, 0x81,
	0x5e, 0x84, 0xac, 0x28, 0xb2, 0xc1, 0xce, 0x0a, 0x69, 0x65, 0xc8, 0x9f, 0xf4, 0xd0, 0x2f, 0xe8,
	0x97, 0xf4, 0xd8, 0x4f, 0x28, 0xee, 0x8f, 0x14, 0xad, 0x12, 0x62, 0x52, 0x0a, 0xbd, 0xcd, 0xcc,
	0xce, 0xbc, 0xf7, 0x66, 0xe1, 0xb5, 0x28, 0xe2, 0xbb, 0x32, 0x4e, 0xc4, 0x92, 0xdf, 0x95, 0x93,
	0xcd, 0x74, 0xb2, 0xc3, 0x71, 0x5e, 0x70, 0xc1, 0xd1, 0x71, 0xb2, 0xe2, 0xd5, 0x0d, 0x8e, 0xf3,
	0x25, 0xde, 0x35, 0xe3, 0xcd, 0xb4, 0xff, 0x2e, 0x5b, 0x8a, 0x45, 0x35, 0xc7, 0x09, 0x5f, 0x4f,
	0x32, 0x9e, 0xf1, 0x89, 0x4c, 0xcc, 0xab, 0x5b, 0xc9, 0x24, 0x91, 0xa8, 0x99, 0xd4, 0x7f, 0x95,
	0x71, 0x9e, 0xad, 0xd2, 0x27, 0x97, 0x58, 0xae, 0xd3, 0x52, 0xc4, 0xeb, 0xbc, 0x31, 0xbc, 0xf9,
	0xa6, 0x41, 0x97, 0x3d, 0xed, 0x40, 0x04, 0xf4, 0x92, 0x57, 0x45, 0x92, 0x9a, 0xea, 0x50, 0x1d,
	0xf5, 0x4e, 0xde, 0xe2, 0x7f, 0xde, 0x82, 0x77, 0x72, 0x81, 0xcc, 0xf8, 0x0f, 0x59, 0x84, 0xa0,
	0xb5, 0x88, 0xcb, 0x85, 0xa9, 0x0d, 0xd5, 0xd1, 0xbe, 0x2f, 0x71, 0xad, 0xdd, 0x16, 0x7c, 0x6d,
	0xee, 0x35, 0x5a, 0x8d, 0x51, 0x0f, 0x34, 0xc1, 0xcd, 0x96, 0x54, 0x34, 0xc1, 0xd1, 0x47, 0x80,
	0xa4, 0x48, 0x63, 0x91, 0xde, 0x44, 0xb1, 0x30, 0xdb, 0x43, 0x75, 0xd4, 0x3d, 0xe9, 0xe3, 0xa6,
	0x03, 0x7e, 0xec, 0x80, 0xd9, 0x63, 0x07, 0x7f, 0xff, 0xc1, 0x6d, 0x09, 0x74, 0x08, 0xed, 0x4d,
	0xbc, 0xaa, 0x52, 0x53, 0x97, 0xd3, 0x1a, 0x82, 0x3e, 0x41, 0x4b, 0xdc, 0xe7, 0xa9, 0xd9, 0x91,
	0x65, 0xc6, 0xff, 0x57, 0x86, 0xdd, 0xe7, 0xa9, 0x2f, 0x73, 0xe3, 0x0f, 0x70, 0xf0, 0x57, 0x4b,
	0xd4, 0x03, 0x08, 0x5d, 0x42, 0x19, 0xb5, 0x19, 0x25, 0x86, 0x82, 0x3a, 0xb0, 0x47, 0xd9, 0xb9,
	0xa1, 0xd6, 0xe0, 0xca, 0x21, 0x86, 0x36, 0xfe, 0xae, 0xc2, 0xcb, 0x67, 0x03, 0x51, 0x17, 0x3a,
	0xa1, 0x7b, 0xe1, 0x7a, 0xd7, 0xae, 0xa1, 0xd4, 0x84, 0xd0, 0x99, 0x17, 0x38, 0xcc, 0x50, 0x11,
	0x80, 0x4e, 0x03, 0xdb, 0xf7, 0xae, 0x0d, 0x0d, 0x19, 0xf0, 0x22, 0x60, 0x3e, 0xb5, 0x2e, 0x23,
	0xea, 0x12, 0x4a, 0x8c, 0x3d, 0x74, 0x04, 0x07, 0x8e, 0x3b, 0x0b, 0x59, 0x64, 0x9f, 0x87, 0xee,
	0x45, 0x64, 0x91, 0x5a, 0x6e, 0xa1, 0x63, 0x38, 0x6a, 0x84, 0x99, 0xef, 0x79, 0x67, 0x51, 0x10,
	0x9e, 0x5e, 0x3a, 0xac, 0xbe, 0xa7, 0xfd, 0xfc, 0xe9, 0xca, 0xfa, 0xec, 0x10, 0xab, 0x7e, 0xd2,
	0x11, 0x82, 0x9e, 0x65, 0xdb, 0x5e, 0xe8, 0xb2, 0xe8, 0x2c, 0x94, 0x0b, 0x3a, 0xa7, 0x87, 0x3f,
	0xb6, 0x03, 0xe5, 0xe7, 0x76, 0xa0, 0xfc, 0xda, 0x0e, 0x94, 0xaf, 0xbf, 0x07, 0xca, 0x17, 0x6d,
	0x33, 0x9d, 0xeb, 0xf2, 0xbb, 0xdf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xaa, 0x0b, 0x4f,
	0xaf, 0x02, 0x00, 0x00,
}
