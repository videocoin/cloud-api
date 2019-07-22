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
		GetActionsByStreamIdRequest
		GetActionsByStreamIdResponse
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
	ActionType_STREAM_REQUESTED      ActionType = 8
	ActionType_STREAM_APPROVED       ActionType = 9
	ActionType_VALIDATOR_ADDED       ActionType = 10
	ActionType_VALIDATOR_REMOVED     ActionType = 11
	ActionType_REFUND_ALLOWED        ActionType = 12
)

var ActionType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "DEPOSIT",
	2:  "ESCROW",
	3:  "STREAM_ENDED",
	4:  "INPUT_CHUNK_ADDED",
	5:  "CHUNK_PROOF_SUBMITTED",
	6:  "CHUNK_PROOF_VALIDATED",
	7:  "ACCOUNT_FUNDED",
	8:  "STREAM_REQUESTED",
	9:  "STREAM_APPROVED",
	10: "VALIDATOR_ADDED",
	11: "VALIDATOR_REMOVED",
	12: "REFUND_ALLOWED",
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
	"STREAM_REQUESTED":      8,
	"STREAM_APPROVED":       9,
	"VALIDATOR_ADDED":       10,
	"VALIDATOR_REMOVED":     11,
	"REFUND_ALLOWED":        12,
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
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xd1, 0x8a, 0xda, 0x40,
	0x14, 0x4d, 0xa2, 0xc6, 0x7a, 0x15, 0x3b, 0x9d, 0xba, 0x90, 0x95, 0x62, 0xa5, 0x50, 0x2a, 0x0b,
	0x8d, 0x75, 0xfb, 0xb4, 0x4f, 0x25, 0x6b, 0x66, 0x59, 0x59, 0x4d, 0xec, 0x64, 0xa2, 0xd0, 0x97,
	0x10, 0xdd, 0xac, 0x0a, 0xba, 0x23, 0x3a, 0x0a, 0xfb, 0x11, 0x85, 0x3e, 0xf6, 0x93, 0xfa, 0xd8,
	0x4f, 0x28, 0xf6, 0x47, 0xca, 0x4c, 0xb2, 0xb8, 0x14, 0x4a, 0xdf, 0xce, 0x3d, 0x73, 0xee, 0x39,
	0xe7, 0x32, 0xf0, 0x4a, 0x6c, 0xe2, 0xfb, 0x6d, 0x3c, 0x15, 0x0b, 0x7e, 0xbf, 0x6d, 0xef, 0x3b,
	0xed, 0x14, 0xda, 0xeb, 0x0d, 0x17, 0x1c, 0x9f, 0x4e, 0x97, 0x7c, 0x77, 0x6b, 0xc7, 0xeb, 0x85,
	0xfd, 0x54, 0x67, 0xef, 0x3b, 0xf5, 0xf7, 0xb3, 0x85, 0x98, 0xef, 0x26, 0xf6, 0x94, 0xaf, 0xda,
	0x33, 0x3e, 0xe3, 0x6d, 0xb5, 0x31, 0xd9, 0xdd, 0xa9, 0x49, 0x0d, 0x0a, 0xa5, 0x4e, 0xf5, 0xd7,
	0x33, 0xce, 0x67, 0xcb, 0xe4, 0xa8, 0x12, 0x8b, 0x55, 0xb2, 0x15, 0xf1, 0x6a, 0x9d, 0x0a, 0xde,
	0x7c, 0x35, 0xc0, 0x74, 0x94, 0x3d, 0xfe, 0x04, 0xe6, 0x96, 0xef, 0x36, 0xd3, 0xc4, 0xd2, 0x9b,
	0x7a, 0xab, 0x7a, 0xfe, 0xce, 0xfe, 0x67, 0x0d, 0x3b, 0x5d, 0x09, 0x94, 0x9c, 0x66, 0x6b, 0x18,
	0x43, 0x7e, 0x1e, 0x6f, 0xe7, 0x96, 0xd1, 0xd4, 0x5b, 0x25, 0xaa, 0xb0, 0xe4, 0xee, 0x36, 0x7c,
	0x65, 0xe5, 0x52, 0x4e, 0x62, 0x5c, 0x05, 0x43, 0x70, 0x2b, 0xaf, 0x18, 0x43, 0x70, 0x7c, 0x01,
	0x30, 0xdd, 0x24, 0xb1, 0x48, 0x6e, 0xa3, 0x58, 0x58, 0x85, 0xa6, 0xde, 0x2a, 0x9f, 0xd7, 0xed,
	0xb4, 0xb9, 0xfd, 0xd8, 0xdc, 0x66, 0x8f, 0xcd, 0x69, 0x29, 0x53, 0x3b, 0x02, 0xd7, 0xa0, 0xb0,
	0x8f, 0x97, 0xbb, 0xc4, 0x32, 0x95, 0x5b, 0x3a, 0xe0, 0x0b, 0xc8, 0x8b, 0x87, 0x75, 0x62, 0x15,
	0xd5, 0x1d, 0x6f, 0xff, 0x7b, 0x07, 0x7b, 0x58, 0x27, 0x54, 0xad, 0x9c, 0x7d, 0x80, 0xca, 0xd3,
	0xdb, 0x70, 0x15, 0x20, 0xf4, 0x5c, 0xc2, 0x48, 0x97, 0x11, 0x17, 0x69, 0xb8, 0x08, 0x39, 0xc2,
	0xae, 0x91, 0x2e, 0xc1, 0xa8, 0xe7, 0x22, 0xe3, 0xec, 0x9b, 0x01, 0x70, 0xb4, 0xc1, 0x65, 0x28,
	0x86, 0xde, 0x8d, 0xe7, 0x8f, 0x3d, 0xa4, 0xc9, 0xc1, 0x25, 0x43, 0x3f, 0xe8, 0x31, 0xa4, 0x63,
	0x00, 0x93, 0x04, 0x5d, 0xea, 0x8f, 0x91, 0x81, 0x11, 0x54, 0x02, 0x46, 0x89, 0x33, 0x88, 0x88,
	0xe7, 0x12, 0x17, 0xe5, 0xf0, 0x09, 0xbc, 0xe8, 0x79, 0xc3, 0x90, 0x45, 0xdd, 0xeb, 0xd0, 0xbb,
	0x89, 0x1c, 0x57, 0xd2, 0x79, 0x7c, 0x0a, 0x27, 0x29, 0x31, 0xa4, 0xbe, 0x7f, 0x15, 0x05, 0xe1,
	0xe5, 0xa0, 0xc7, 0x64, 0x95, 0xc2, 0xdf, 0x4f, 0x23, 0xa7, 0xdf, 0x73, 0x1d, 0xf9, 0x64, 0x62,
	0x0c, 0x55, 0xa7, 0xdb, 0xf5, 0x43, 0x8f, 0x45, 0x57, 0xa1, 0x0a, 0x28, 0xe2, 0x1a, 0xa0, 0x2c,
	0x92, 0x92, 0xcf, 0x21, 0x09, 0xa4, 0xf2, 0x19, 0x7e, 0x09, 0xcf, 0x33, 0xd6, 0x19, 0x0e, 0xa9,
	0x3f, 0x22, 0x2e, 0x2a, 0x49, 0x32, 0x73, 0xf3, 0x69, 0xd6, 0x04, 0x64, 0xc1, 0x23, 0x49, 0xc9,
	0x40, 0x69, 0xcb, 0x32, 0x8a, 0x12, 0x19, 0x12, 0x39, 0xfd, 0xbe, 0x3f, 0x26, 0x2e, 0xaa, 0x5c,
	0xd6, 0x7e, 0x1c, 0x1a, 0xda, 0xcf, 0x43, 0x43, 0xfb, 0x75, 0x68, 0x68, 0xdf, 0x7f, 0x37, 0xb4,
	0x2f, 0xc6, 0xbe, 0x33, 0x31, 0xd5, 0x57, 0x7e, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x81, 0xe4,
	0x02, 0xfa, 0xfc, 0x02, 0x00, 0x00,
}
