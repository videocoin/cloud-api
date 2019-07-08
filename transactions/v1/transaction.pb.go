// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transactions/v1/transaction.proto

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

type Transaction struct {
	Hash        string                      `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From        string                      `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To          string                      `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Timestamp   *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Value       string                      `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	BlockNumber string                      `protobuf:"bytes,7,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorTransaction, []int{0} }

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

func (m *Transaction) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Transaction) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Transaction) GetBlockNumber() string {
	if m != nil {
		return m.BlockNumber
	}
	return ""
}

func init() {
	proto.RegisterType((*Transaction)(nil), "cloud.api.transactions.v1.Transaction")
}

func init() { proto.RegisterFile("transactions/v1/transaction.proto", fileDescriptorTransaction) }

var fileDescriptorTransaction = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8d, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x93, 0x50, 0x8a, 0xea, 0x20, 0x06, 0xab, 0x83, 0xc9, 0x10, 0x5a, 0xa6, 0x2e, 0x38,
	0x0a, 0x2c, 0xcc, 0x1c, 0x80, 0xa1, 0xea, 0xc4, 0x82, 0xec, 0x90, 0x3a, 0x11, 0x71, 0x5e, 0x94,
	0xd8, 0x3e, 0x0b, 0x47, 0xe1, 0x08, 0x8c, 0x1c, 0x01, 0x85, 0x8b, 0xa0, 0x3e, 0xab, 0xa4, 0xdb,
	0xfb, 0x3f, 0x7f, 0xfe, 0x7f, 0xb2, 0x36, 0xbd, 0x68, 0x07, 0x51, 0x98, 0x1a, 0xda, 0x21, 0x73,
	0x79, 0x76, 0x92, 0x79, 0xd7, 0x83, 0x01, 0x7a, 0x5d, 0x34, 0x60, 0xdf, 0xb8, 0xe8, 0x6a, 0x7e,
	0x2a, 0x73, 0x97, 0x27, 0x77, 0xaa, 0x36, 0x95, 0x95, 0xbc, 0x00, 0x9d, 0x29, 0x50, 0x90, 0xe1,
	0x0f, 0x69, 0xf7, 0x98, 0x30, 0xe0, 0xe5, 0x9b, 0x92, 0x1b, 0x05, 0xa0, 0x9a, 0x72, 0xb2, 0x4c,
	0xad, 0xcb, 0xc1, 0x08, 0xdd, 0x79, 0xe1, 0xf6, 0x33, 0x24, 0xf1, 0x6e, 0xda, 0xa0, 0x94, 0xcc,
	0x2a, 0x31, 0x54, 0x2c, 0x5c, 0x85, 0x9b, 0xc5, 0x16, 0xef, 0x03, 0xdb, 0xf7, 0xa0, 0x59, 0xe4,
	0xd9, 0xe1, 0xa6, 0x57, 0x24, 0x32, 0xc0, 0xce, 0x90, 0x44, 0x06, 0xe8, 0x23, 0x59, 0xfc, 0x57,
	0xb3, 0xd9, 0x2a, 0xdc, 0xc4, 0xf7, 0x09, 0xf7, 0xe3, 0xfc, 0x38, 0xce, 0x77, 0x47, 0x63, 0x3b,
	0xc9, 0x74, 0x49, 0xce, 0x9d, 0x68, 0x6c, 0xc9, 0xe6, 0x58, 0xe6, 0x03, 0x5d, 0x93, 0x4b, 0xd9,
	0x40, 0xf1, 0xfe, 0xda, 0x5a, 0x2d, 0xcb, 0x9e, 0x5d, 0xe0, 0x63, 0x8c, 0xec, 0x19, 0xd1, 0xd3,
	0xf2, 0x6b, 0x4c, 0x83, 0xef, 0x31, 0x0d, 0x7e, 0xc6, 0x34, 0xf8, 0xf8, 0x4d, 0x83, 0x97, 0xc8,
	0xe5, 0x72, 0x8e, 0x6b, 0x0f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xb4, 0x33, 0x37, 0x67,
	0x01, 0x00, 0x00,
}
