// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transactions/v1/action.proto

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		transactions/v1/action.proto
		transactions/v1/action_service.proto
		transactions/v1/block.proto
		transactions/v1/block_service.proto
		transactions/v1/transaction.proto
		transactions/v1/transaction_service.proto

	It has these top-level messages:
		Action
		GetActionsRequest
		GetActionsResponse
		GetActionRequest
		GetActionResponse
		Block
		BlockDetail
		GetBlocksRequest
		GetBlocksResponse
		GetBlockRequest
		GetBlockResponse
		Transaction
		TransactionDetail
		GetTransactionsRequest
		GetTransactionsResponse
		GetTransactionRequest
		GetTransactionResponse
		GetAllTransactionsRequest
		GetAllTransactionsResponse
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

type ActionSource int32

const (
	ActionSource_UNDETECTED ActionSource = 0
	ActionSource_ETH        ActionSource = 1
	ActionSource_VID        ActionSource = 2
)

var ActionSource_name = map[int32]string{
	0: "UNDETECTED",
	1: "ETH",
	2: "VID",
}
var ActionSource_value = map[string]int32{
	"UNDETECTED": 0,
	"ETH":        1,
	"VID":        2,
}

func (x ActionSource) String() string {
	return proto.EnumName(ActionSource_name, int32(x))
}
func (ActionSource) EnumDescriptor() ([]byte, []int) { return fileDescriptorAction, []int{0} }

type ActionType int32

const (
	ActionType_UNKNOWN               ActionType = 0
	ActionType_DEPOSIT               ActionType = 1
	ActionType_ESCROW                ActionType = 2
	ActionType_STREAM_ENDED          ActionType = 3
	ActionType_INPUT_CHUNK_ADDED     ActionType = 4
	ActionType_CHUNK_PROOF_SUBMITTED ActionType = 5
	ActionType_CHUNK_PROOF_VALIDATED ActionType = 6
	ActionType_ACCOUNT_FUNDED        ActionType = 7
)

var ActionType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DEPOSIT",
	2: "ESCROW",
	3: "STREAM_ENDED",
	4: "INPUT_CHUNK_ADDED",
	5: "CHUNK_PROOF_SUBMITTED",
	6: "CHUNK_PROOF_VALIDATED",
	7: "ACCOUNT_FUNDED",
}
var ActionType_value = map[string]int32{
	"UNKNOWN":               0,
	"DEPOSIT":               1,
	"ESCROW":                2,
	"STREAM_ENDED":          3,
	"INPUT_CHUNK_ADDED":     4,
	"CHUNK_PROOF_SUBMITTED": 5,
	"CHUNK_PROOF_VALIDATED": 6,
	"ACCOUNT_FUNDED":        7,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}
func (ActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptorAction, []int{1} }

type Action struct {
	Source    ActionSource                `protobuf:"varint,1,opt,name=source,proto3,enum=cloud.api.transactions.v1.ActionSource" json:"source,omitempty"`
	Hash      string                      `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	From      string                      `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To        string                      `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	CreatedAt *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Value     string                      `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	Type      ActionType                  `protobuf:"varint,7,opt,name=type,proto3,enum=cloud.api.transactions.v1.ActionType" json:"type,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptorAction, []int{0} }

func (m *Action) GetSource() ActionSource {
	if m != nil {
		return m.Source
	}
	return ActionSource_UNDETECTED
}

func (m *Action) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Action) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Action) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Action) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Action) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Action) GetType() ActionType {
	if m != nil {
		return m.Type
	}
	return ActionType_UNKNOWN
}

func init() {
	proto.RegisterType((*Action)(nil), "cloud.api.transactions.v1.Action")
	proto.RegisterEnum("cloud.api.transactions.v1.ActionSource", ActionSource_name, ActionSource_value)
	proto.RegisterEnum("cloud.api.transactions.v1.ActionType", ActionType_name, ActionType_value)
}

func init() { proto.RegisterFile("transactions/v1/action.proto", fileDescriptorAction) }

var fileDescriptorAction = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xc1, 0x6a, 0xdb, 0x40,
	0x14, 0x94, 0x64, 0x5b, 0x26, 0x2f, 0xc1, 0xa8, 0x8f, 0x04, 0x14, 0x53, 0x5c, 0x53, 0x28, 0x35,
	0x81, 0xae, 0xeb, 0xf4, 0x94, 0x53, 0x51, 0xac, 0x0d, 0x11, 0x69, 0x24, 0x23, 0xad, 0x12, 0xe8,
	0x45, 0xc8, 0x8a, 0x22, 0x1b, 0xec, 0xac, 0x90, 0x56, 0x86, 0x7c, 0x44, 0xef, 0xfd, 0x81, 0xfe,
	0x4b, 0x8f, 0xfd, 0x84, 0xe2, 0xfe, 0x48, 0xd1, 0x2a, 0x21, 0xa1, 0x50, 0x72, 0x9b, 0x99, 0x9d,
	0x19, 0xcd, 0x13, 0xbc, 0x16, 0x45, 0x7c, 0x57, 0xc6, 0x89, 0x58, 0xf2, 0xbb, 0x72, 0xbc, 0x99,
	0x8c, 0x1b, 0x48, 0xf2, 0x82, 0x0b, 0x8e, 0x87, 0xc9, 0x8a, 0x57, 0x37, 0x24, 0xce, 0x97, 0xe4,
	0xb9, 0x8f, 0x6c, 0x26, 0xfd, 0x0f, 0xd9, 0x52, 0x2c, 0xaa, 0x39, 0x49, 0xf8, 0x7a, 0x9c, 0xf1,
	0x8c, 0x8f, 0x65, 0x62, 0x5e, 0xdd, 0x4a, 0x26, 0x89, 0x44, 0x4d, 0x53, 0xff, 0x4d, 0xc6, 0x79,
	0xb6, 0x4a, 0x9f, 0x5c, 0x62, 0xb9, 0x4e, 0x4b, 0x11, 0xaf, 0xf3, 0xc6, 0xf0, 0xf6, 0x9b, 0x06,
	0xba, 0x25, 0xeb, 0xf1, 0x33, 0xe8, 0x25, 0xaf, 0x8a, 0x24, 0x35, 0xd5, 0xa1, 0x3a, 0xea, 0x1d,
	0xbf, 0x27, 0xff, 0x9d, 0x41, 0x9a, 0x48, 0x20, 0xed, 0xfe, 0x43, 0x0c, 0x11, 0xda, 0x8b, 0xb8,
	0x5c, 0x98, 0xda, 0x50, 0x1d, 0xed, 0xf8, 0x12, 0xd7, 0xda, 0x6d, 0xc1, 0xd7, 0x66, 0xab, 0xd1,
	0x6a, 0x8c, 0x3d, 0xd0, 0x04, 0x37, 0xdb, 0x52, 0xd1, 0x04, 0xc7, 0x13, 0x80, 0xa4, 0x48, 0x63,
	0x91, 0xde, 0x44, 0xb1, 0x30, 0x3b, 0x43, 0x75, 0xb4, 0x7b, 0xdc, 0x27, 0xcd, 0x72, 0xf2, 0xb8,
	0x9c, 0xb0, 0xc7, 0xe5, 0xfe, 0xce, 0x83, 0xdb, 0x12, 0xb8, 0x0f, 0x9d, 0x4d, 0xbc, 0xaa, 0x52,
	0x53, 0x97, 0x6d, 0x0d, 0xc1, 0x13, 0x68, 0x8b, 0xfb, 0x3c, 0x35, 0xbb, 0xf2, 0x8e, 0x77, 0x2f,
	0xde, 0xc1, 0xee, 0xf3, 0xd4, 0x97, 0x91, 0xa3, 0x8f, 0xb0, 0xf7, 0xfc, 0x36, 0xec, 0x01, 0x84,
	0xae, 0x4d, 0x19, 0x9d, 0x32, 0x6a, 0x1b, 0x0a, 0x76, 0xa1, 0x45, 0xd9, 0xb9, 0xa1, 0xd6, 0xe0,
	0xca, 0xb1, 0x0d, 0xed, 0xe8, 0x87, 0x0a, 0xf0, 0x54, 0x83, 0xbb, 0xd0, 0x0d, 0xdd, 0x0b, 0xd7,
	0xbb, 0x76, 0x0d, 0xa5, 0x26, 0x36, 0x9d, 0x79, 0x81, 0xc3, 0x0c, 0x15, 0x01, 0x74, 0x1a, 0x4c,
	0x7d, 0xef, 0xda, 0xd0, 0xd0, 0x80, 0xbd, 0x80, 0xf9, 0xd4, 0xba, 0x8c, 0xa8, 0x6b, 0x53, 0xdb,
	0x68, 0xe1, 0x01, 0xbc, 0x72, 0xdc, 0x59, 0xc8, 0xa2, 0xe9, 0x79, 0xe8, 0x5e, 0x44, 0x96, 0x5d,
	0xcb, 0x6d, 0x3c, 0x84, 0x83, 0x46, 0x98, 0xf9, 0x9e, 0x77, 0x16, 0x05, 0xe1, 0xe9, 0xa5, 0xc3,
	0xea, 0x29, 0x9d, 0x7f, 0x9f, 0xae, 0xac, 0x2f, 0x8e, 0x6d, 0xd5, 0x4f, 0x3a, 0x22, 0xf4, 0xac,
	0xe9, 0xd4, 0x0b, 0x5d, 0x16, 0x9d, 0x85, 0xf2, 0x03, 0xdd, 0xd3, 0xfd, 0x9f, 0xdb, 0x81, 0xf2,
	0x6b, 0x3b, 0x50, 0x7e, 0x6f, 0x07, 0xca, 0xf7, 0x3f, 0x03, 0xe5, 0xab, 0xb6, 0x99, 0xcc, 0x75,
	0xf9, 0x7f, 0x3f, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x32, 0xd8, 0x39, 0x10, 0x91, 0x02, 0x00,
	0x00,
}
