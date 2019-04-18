// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: manager/v1/manager.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on AddProfileRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AddProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStreamId() == nil {
		return AddProfileRequestValidationError{
			field:  "StreamId",
			reason: "value is required",
		}
	}

	if a := m.GetStreamId(); a != nil {

	}

	if m.GetProfileId() == nil {
		return AddProfileRequestValidationError{
			field:  "ProfileId",
			reason: "value is required",
		}
	}

	if a := m.GetProfileId(); a != nil {

	}

	return nil
}

// AddProfileRequestValidationError is the validation error returned by
// AddProfileRequest.Validate if the designated constraints aren't met.
type AddProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddProfileRequestValidationError) ErrorName() string {
	return "AddProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddProfileRequestValidationError{}

// Validate checks the field values on Heartbeat with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Heartbeat) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// HeartbeatValidationError is the validation error returned by
// Heartbeat.Validate if the designated constraints aren't met.
type HeartbeatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeartbeatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeartbeatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeartbeatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeartbeatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeartbeatValidationError) ErrorName() string { return "HeartbeatValidationError" }

// Error satisfies the builtin error interface
func (e HeartbeatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeartbeat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeartbeatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeartbeatValidationError{}

// Validate checks the field values on GetProfileRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetProfileId() == nil {
		return GetProfileRequestValidationError{
			field:  "ProfileId",
			reason: "value is required",
		}
	}

	if a := m.GetProfileId(); a != nil {

	}

	return nil
}

// GetProfileRequestValidationError is the validation error returned by
// GetProfileRequest.Validate if the designated constraints aren't met.
type GetProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileRequestValidationError) ErrorName() string {
	return "GetProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileRequestValidationError{}

// Validate checks the field values on CheckBalanceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckBalanceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetContractAddress() == nil {
		return CheckBalanceRequestValidationError{
			field:  "ContractAddress",
			reason: "value is required",
		}
	}

	if a := m.GetContractAddress(); a != nil {

	}

	return nil
}

// CheckBalanceRequestValidationError is the validation error returned by
// CheckBalanceRequest.Validate if the designated constraints aren't met.
type CheckBalanceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckBalanceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckBalanceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckBalanceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckBalanceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckBalanceRequestValidationError) ErrorName() string {
	return "CheckBalanceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CheckBalanceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckBalanceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckBalanceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckBalanceRequestValidationError{}

// Validate checks the field values on CheckBalanceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckBalanceResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Balance

	return nil
}

// CheckBalanceResponseValidationError is the validation error returned by
// CheckBalanceResponse.Validate if the designated constraints aren't met.
type CheckBalanceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckBalanceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckBalanceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckBalanceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckBalanceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckBalanceResponseValidationError) ErrorName() string {
	return "CheckBalanceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CheckBalanceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckBalanceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckBalanceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckBalanceResponseValidationError{}

// Validate checks the field values on ContractAddrRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ContractAddrRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStreamId() == nil {
		return ContractAddrRequestValidationError{
			field:  "StreamId",
			reason: "value is required",
		}
	}

	if a := m.GetStreamId(); a != nil {

	}

	// no validation rules for ContractAddress

	return nil
}

// ContractAddrRequestValidationError is the validation error returned by
// ContractAddrRequest.Validate if the designated constraints aren't met.
type ContractAddrRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContractAddrRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContractAddrRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContractAddrRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContractAddrRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContractAddrRequestValidationError) ErrorName() string {
	return "ContractAddrRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ContractAddrRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContractAddrRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContractAddrRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContractAddrRequestValidationError{}

// Validate checks the field values on VerifyChunkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *VerifyChunkRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStreamId() <= 0 {
		return VerifyChunkRequestValidationError{
			field:  "StreamId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetSourceChunkId() < 0 {
		return VerifyChunkRequestValidationError{
			field:  "SourceChunkId",
			reason: "value must be greater than or equal to 0",
		}
	}

	if m.GetResultChunkId() < 0 {
		return VerifyChunkRequestValidationError{
			field:  "ResultChunkId",
			reason: "value must be greater than or equal to 0",
		}
	}

	// no validation rules for HashDistance

	// no validation rules for Bitrate

	return nil
}

// VerifyChunkRequestValidationError is the validation error returned by
// VerifyChunkRequest.Validate if the designated constraints aren't met.
type VerifyChunkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyChunkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyChunkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyChunkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyChunkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyChunkRequestValidationError) ErrorName() string {
	return "VerifyChunkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyChunkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyChunkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyChunkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyChunkRequestValidationError{}

// Validate checks the field values on ChunkCreatedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ChunkCreatedRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStreamId() <= 0 {
		return ChunkCreatedRequestValidationError{
			field:  "StreamId",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for SourceChunkId

	// no validation rules for ResultChunkId

	// no validation rules for Bitrate

	return nil
}

// ChunkCreatedRequestValidationError is the validation error returned by
// ChunkCreatedRequest.Validate if the designated constraints aren't met.
type ChunkCreatedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChunkCreatedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChunkCreatedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChunkCreatedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChunkCreatedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChunkCreatedRequestValidationError) ErrorName() string {
	return "ChunkCreatedRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChunkCreatedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChunkCreatedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChunkCreatedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChunkCreatedRequestValidationError{}

// Validate checks the field values on StopStreamRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StopStreamRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStreamId() <= 0 {
		return StopStreamRequestValidationError{
			field:  "StreamId",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for ContractAddress

	// no validation rules for WalletAddress

	return nil
}

// StopStreamRequestValidationError is the validation error returned by
// StopStreamRequest.Validate if the designated constraints aren't met.
type StopStreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StopStreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StopStreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StopStreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StopStreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StopStreamRequestValidationError) ErrorName() string {
	return "StopStreamRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StopStreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StopStreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StopStreamRequestValidationError{}

// Validate checks the field values on UpdateTranscoderStatusRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateTranscoderStatusRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TranscoderId

	// no validation rules for Status

	return nil
}

// UpdateTranscoderStatusRequestValidationError is the validation error
// returned by UpdateTranscoderStatusRequest.Validate if the designated
// constraints aren't met.
type UpdateTranscoderStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTranscoderStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTranscoderStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTranscoderStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTranscoderStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTranscoderStatusRequestValidationError) ErrorName() string {
	return "UpdateTranscoderStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTranscoderStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTranscoderStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTranscoderStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTranscoderStatusRequestValidationError{}

// Validate checks the field values on UpdateStreamStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateStreamStatusRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TranscoderId

	// no validation rules for StreamId

	// no validation rules for Status

	// no validation rules for Refunded

	// no validation rules for IngestStatus

	return nil
}

// UpdateStreamStatusRequestValidationError is the validation error returned by
// UpdateStreamStatusRequest.Validate if the designated constraints aren't met.
type UpdateStreamStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStreamStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStreamStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStreamStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStreamStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStreamStatusRequestValidationError) ErrorName() string {
	return "UpdateStreamStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStreamStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStreamStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStreamStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStreamStatusRequestValidationError{}

// Validate checks the field values on AddJobResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AddJobResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RtmpInputUrl

	return nil
}

// AddJobResponseValidationError is the validation error returned by
// AddJobResponse.Validate if the designated constraints aren't met.
type AddJobResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddJobResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddJobResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddJobResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddJobResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddJobResponseValidationError) ErrorName() string { return "AddJobResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddJobResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddJobResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddJobResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddJobResponseValidationError{}

// Validate checks the field values on AddJobRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AddJobRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for StreamId

	// no validation rules for WalletAddress

	// no validation rules for ProfileId

	return nil
}

// AddJobRequestValidationError is the validation error returned by
// AddJobRequest.Validate if the designated constraints aren't met.
type AddJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddJobRequestValidationError) ErrorName() string { return "AddJobRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddJobRequestValidationError{}

// Validate checks the field values on GetJobRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetJobRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetJobRequestValidationError is the validation error returned by
// GetJobRequest.Validate if the designated constraints aren't met.
type GetJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetJobRequestValidationError) ErrorName() string { return "GetJobRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetJobRequestValidationError{}

// Validate checks the field values on GetStreamRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetStreamRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for StreamId

	return nil
}

// GetStreamRequestValidationError is the validation error returned by
// GetStreamRequest.Validate if the designated constraints aren't met.
type GetStreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStreamRequestValidationError) ErrorName() string { return "GetStreamRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetStreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStreamRequestValidationError{}
