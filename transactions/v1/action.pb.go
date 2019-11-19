// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transactions/v1/action.proto

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
	ActionType_UNKNOWN                              ActionType = 0
	ActionType_DEPOSIT                              ActionType = 1
	ActionType_STREAM_CREATED                       ActionType = 2
	ActionType_STREAM_ENDED                         ActionType = 3
	ActionType_INPUT_CHUNK_ADDED                    ActionType = 4
	ActionType_CHUNK_PROOF_SUBMITTED                ActionType = 5
	ActionType_CHUNK_PROOF_VALIDATED                ActionType = 6
	ActionType_CHUNK_PROOF_SCRAPPED                 ActionType = 14
	ActionType_ACCOUNT_FUNDED                       ActionType = 7
	ActionType_STREAM_REQUESTED                     ActionType = 8
	ActionType_STREAM_REFUNDED                      ActionType = 15
	ActionType_STREAM_APPROVED                      ActionType = 9
	ActionType_VALIDATOR_ADDED                      ActionType = 10
	ActionType_VALIDATOR_REMOVED                    ActionType = 11
	ActionType_REFUND_ALLOWED                       ActionType = 12
	ActionType_REFUND_REVOKED                       ActionType = 13
	ActionType_TRANSFER_EXECUTED                    ActionType = 16
	ActionType_OUT_OF_FUNDS                         ActionType = 17
	ActionType_STREAM_DEPOSITED                     ActionType = 18
	ActionType_STREAM_MANAGER_OWNERSHIP_TRANSFERRED ActionType = 19
	ActionType_BRIDGE_FUNDED                        ActionType = 20
	ActionType_BRIDGE_TRANSFER_WITHIN_LIMITS        ActionType = 21
	ActionType_BRIDGE_TRANSFER_EXCEEDED_LIMITS      ActionType = 22
	ActionType_BRIDGE_TRANSFER_REJECTED             ActionType = 23
	ActionType_BRIDGE_OWNERSHIP_TRANSFERRED         ActionType = 24
	ActionType_DEPOSIT_ERC20                        ActionType = 25
)

var ActionType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "DEPOSIT",
	2:  "STREAM_CREATED",
	3:  "STREAM_ENDED",
	4:  "INPUT_CHUNK_ADDED",
	5:  "CHUNK_PROOF_SUBMITTED",
	6:  "CHUNK_PROOF_VALIDATED",
	14: "CHUNK_PROOF_SCRAPPED",
	7:  "ACCOUNT_FUNDED",
	8:  "STREAM_REQUESTED",
	15: "STREAM_REFUNDED",
	9:  "STREAM_APPROVED",
	10: "VALIDATOR_ADDED",
	11: "VALIDATOR_REMOVED",
	12: "REFUND_ALLOWED",
	13: "REFUND_REVOKED",
	16: "TRANSFER_EXECUTED",
	17: "OUT_OF_FUNDS",
	18: "STREAM_DEPOSITED",
	19: "STREAM_MANAGER_OWNERSHIP_TRANSFERRED",
	20: "BRIDGE_FUNDED",
	21: "BRIDGE_TRANSFER_WITHIN_LIMITS",
	22: "BRIDGE_TRANSFER_EXCEEDED_LIMITS",
	23: "BRIDGE_TRANSFER_REJECTED",
	24: "BRIDGE_OWNERSHIP_TRANSFERRED",
	25: "DEPOSIT_ERC20",
}
var ActionType_value = map[string]int32{
	"UNKNOWN":                              0,
	"DEPOSIT":                              1,
	"STREAM_CREATED":                       2,
	"STREAM_ENDED":                         3,
	"INPUT_CHUNK_ADDED":                    4,
	"CHUNK_PROOF_SUBMITTED":                5,
	"CHUNK_PROOF_VALIDATED":                6,
	"CHUNK_PROOF_SCRAPPED":                 14,
	"ACCOUNT_FUNDED":                       7,
	"STREAM_REQUESTED":                     8,
	"STREAM_REFUNDED":                      15,
	"STREAM_APPROVED":                      9,
	"VALIDATOR_ADDED":                      10,
	"VALIDATOR_REMOVED":                    11,
	"REFUND_ALLOWED":                       12,
	"REFUND_REVOKED":                       13,
	"TRANSFER_EXECUTED":                    16,
	"OUT_OF_FUNDS":                         17,
	"STREAM_DEPOSITED":                     18,
	"STREAM_MANAGER_OWNERSHIP_TRANSFERRED": 19,
	"BRIDGE_FUNDED":                        20,
	"BRIDGE_TRANSFER_WITHIN_LIMITS":        21,
	"BRIDGE_TRANSFER_EXCEEDED_LIMITS":      22,
	"BRIDGE_TRANSFER_REJECTED":             23,
	"BRIDGE_OWNERSHIP_TRANSFERRED":         24,
	"DEPOSIT_ERC20":                        25,
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
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x93, 0x10, 0x92, 0xcb, 0x01, 0xc2, 0x30, 0x84, 0x7b, 0x0d, 0xe2, 0x86, 0xdc, 0x3f,
	0x55, 0x23, 0xa4, 0x3a, 0x40, 0x57, 0xac, 0x2a, 0xe3, 0x39, 0x21, 0x2e, 0x89, 0xed, 0x8e, 0xed,
	0x04, 0x75, 0x63, 0x99, 0x60, 0x42, 0x24, 0x82, 0xa3, 0xc4, 0x89, 0xc4, 0x43, 0x74, 0xdf, 0x47,
	0xea, 0xb2, 0x8f, 0x50, 0xd1, 0x37, 0xe8, 0x13, 0x54, 0x33, 0x76, 0x20, 0x45, 0xad, 0xba, 0x3b,
	0xe7, 0x9b, 0xef, 0x7c, 0xf3, 0x3b, 0x23, 0x1b, 0xf6, 0xe2, 0x71, 0x70, 0x37, 0x09, 0x7a, 0xf1,
	0x20, 0xba, 0x9b, 0xd4, 0x67, 0x47, 0xf5, 0xa4, 0x54, 0x47, 0xe3, 0x28, 0x8e, 0xe8, 0x4e, 0xef,
	0x36, 0x9a, 0x5e, 0xa9, 0xc1, 0x68, 0xa0, 0x2e, 0xfa, 0xd4, 0xd9, 0xd1, 0xee, 0xab, 0xfe, 0x20,
	0xbe, 0x99, 0x5e, 0xaa, 0xbd, 0x68, 0x58, 0xef, 0x47, 0xfd, 0xa8, 0x2e, 0x27, 0x2e, 0xa7, 0xd7,
	0xb2, 0x93, 0x8d, 0xac, 0x92, 0xa4, 0xdd, 0xfd, 0x7e, 0x14, 0xf5, 0x6f, 0xc3, 0x27, 0x57, 0x3c,
	0x18, 0x86, 0x93, 0x38, 0x18, 0x8e, 0x12, 0xc3, 0xbf, 0x1f, 0x72, 0x50, 0xd0, 0x64, 0x3c, 0x7d,
	0x03, 0x85, 0x49, 0x34, 0x1d, 0xf7, 0x42, 0x25, 0x5b, 0xcd, 0xd6, 0x4a, 0xc7, 0x2f, 0xd5, 0x5f,
	0x62, 0xa8, 0xc9, 0x88, 0x23, 0xed, 0x3c, 0x1d, 0xa3, 0x14, 0xf2, 0x37, 0xc1, 0xe4, 0x46, 0xc9,
	0x55, 0xb3, 0xb5, 0x15, 0x2e, 0x6b, 0xa1, 0x5d, 0x8f, 0xa3, 0xa1, 0xb2, 0x94, 0x68, 0xa2, 0xa6,
	0x25, 0xc8, 0xc5, 0x91, 0x92, 0x97, 0x4a, 0x2e, 0x8e, 0xe8, 0x09, 0x40, 0x6f, 0x1c, 0x06, 0x71,
	0x78, 0xe5, 0x07, 0xb1, 0xb2, 0x5c, 0xcd, 0xd6, 0x56, 0x8f, 0x77, 0xd5, 0x84, 0x5c, 0x9d, 0x93,
	0xab, 0xee, 0x9c, 0x9c, 0xaf, 0xa4, 0x6e, 0x2d, 0xa6, 0x65, 0x58, 0x9e, 0x05, 0xb7, 0xd3, 0x50,
	0x29, 0xc8, 0xb4, 0xa4, 0xa1, 0x27, 0x90, 0x8f, 0xef, 0x47, 0xa1, 0x52, 0x94, 0x7b, 0xbc, 0xf8,
	0xed, 0x1e, 0xee, 0xfd, 0x28, 0xe4, 0x72, 0xe4, 0xe0, 0x10, 0xd6, 0x16, 0x77, 0xa3, 0x25, 0x00,
	0xcf, 0x64, 0xe8, 0xa2, 0xee, 0x22, 0x23, 0x19, 0x5a, 0x84, 0x25, 0x74, 0x9b, 0x24, 0x2b, 0x8a,
	0x8e, 0xc1, 0x48, 0xee, 0xe0, 0x5b, 0x1e, 0xe0, 0x29, 0x86, 0xae, 0x42, 0xd1, 0x33, 0xcf, 0x4d,
	0xab, 0x6b, 0x92, 0x8c, 0x68, 0x18, 0xda, 0x96, 0x63, 0xb8, 0x24, 0x4b, 0x29, 0x94, 0x1c, 0x97,
	0xa3, 0xd6, 0xf6, 0x75, 0x8e, 0x9a, 0x88, 0xcb, 0x51, 0x02, 0x6b, 0xa9, 0x86, 0x26, 0x43, 0x46,
	0x96, 0xe8, 0x36, 0x6c, 0x1a, 0xa6, 0xed, 0xb9, 0xbe, 0xde, 0xf4, 0xcc, 0x73, 0x5f, 0x63, 0x42,
	0xce, 0xd3, 0x1d, 0xd8, 0x4e, 0x04, 0x9b, 0x5b, 0x56, 0xc3, 0x77, 0xbc, 0xd3, 0xb6, 0xe1, 0x8a,
	0x8c, 0xe5, 0xe7, 0x47, 0x1d, 0xad, 0x65, 0x30, 0x19, 0x5f, 0xa0, 0x0a, 0x94, 0x7f, 0x98, 0xd2,
	0xb9, 0x66, 0xdb, 0xc8, 0x48, 0x49, 0xc0, 0x68, 0xba, 0x6e, 0x79, 0xa6, 0xeb, 0x37, 0x3c, 0x79,
	0x75, 0x91, 0x96, 0x81, 0xa4, 0x30, 0x1c, 0xdf, 0x79, 0xe8, 0x88, 0x8c, 0x3f, 0xe8, 0x16, 0x6c,
	0x3c, 0xaa, 0xa9, 0x75, 0x63, 0x41, 0xd4, 0x6c, 0x9b, 0x5b, 0x1d, 0x64, 0x64, 0x45, 0x88, 0xe9,
	0xe5, 0x16, 0x4f, 0xc1, 0x41, 0xec, 0xf3, 0x24, 0x72, 0x6c, 0x4b, 0xef, 0xaa, 0xb8, 0x3f, 0x89,
	0xf3, 0xb5, 0x56, 0xcb, 0xea, 0x22, 0x23, 0x6b, 0x0b, 0x1a, 0xc7, 0x8e, 0x75, 0x8e, 0x8c, 0xac,
	0x8b, 0x71, 0x97, 0x6b, 0xa6, 0xd3, 0x40, 0xee, 0xe3, 0x05, 0xea, 0x9e, 0x80, 0x22, 0xe2, 0xdd,
	0x2c, 0xcf, 0xf5, 0xad, 0x86, 0xa4, 0x77, 0xc8, 0xe6, 0x02, 0x7c, 0xfa, 0xe2, 0xc8, 0x08, 0xa5,
	0x35, 0xf8, 0x3f, 0x55, 0xdb, 0x9a, 0xa9, 0x9d, 0x21, 0xf7, 0xad, 0xae, 0x89, 0xdc, 0x69, 0x1a,
	0xb6, 0x3f, 0xcf, 0xe5, 0xc8, 0xc8, 0x16, 0xdd, 0x84, 0xf5, 0x53, 0x6e, 0xb0, 0x33, 0x9c, 0xbf,
	0x47, 0x99, 0xfe, 0x03, 0x7f, 0xa7, 0xd2, 0x23, 0x42, 0xd7, 0x70, 0x9b, 0x86, 0xe9, 0xb7, 0x8c,
	0xb6, 0xe1, 0x3a, 0x64, 0x9b, 0xfe, 0x07, 0xfb, 0xcf, 0x2d, 0x78, 0xa1, 0x23, 0x32, 0x64, 0x73,
	0xd3, 0x9f, 0x74, 0x0f, 0x94, 0xe7, 0x26, 0x8e, 0x6f, 0x93, 0x2f, 0xea, 0x2f, 0x5a, 0x85, 0xbd,
	0xf4, 0xf4, 0xe7, 0x68, 0x8a, 0x40, 0x4b, 0x77, 0xf2, 0x91, 0xeb, 0xc7, 0x87, 0x64, 0xe7, 0xb4,
	0xfc, 0xe9, 0xa1, 0x92, 0xf9, 0xfc, 0x50, 0xc9, 0x7c, 0x79, 0xa8, 0x64, 0x3e, 0x7e, 0xad, 0x64,
	0xde, 0xe7, 0x66, 0x47, 0x97, 0x05, 0xf9, 0xb3, 0xbc, 0xfe, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xe4,
	0x90, 0xa3, 0xf7, 0x5e, 0x04, 0x00, 0x00,
}
