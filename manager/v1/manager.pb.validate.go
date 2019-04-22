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

	v1 "github.com/VideoCoin/cloud-api/profiles/v1"
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

// Validate checks the field values on UpdateProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStreamId() < 0 {
		return UpdateProfileRequestValidationError{
			field:  "StreamId",
			reason: "value must be greater than or equal to 0",
		}
	}

	if _, ok := v1.ProfileId_name[int32(m.GetProfileId())]; !ok {
		return UpdateProfileRequestValidationError{
			field:  "ProfileId",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// UpdateProfileRequestValidationError is the validation error returned by
// UpdateProfileRequest.Validate if the designated constraints aren't met.
type UpdateProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProfileRequestValidationError) ErrorName() string {
	return "UpdateProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProfileRequestValidationError{}

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

// Validate checks the field values on ProfileRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := v1.ProfileId_name[int32(m.GetProfileId())]; !ok {
		return ProfileRequestValidationError{
			field:  "ProfileId",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ProfileRequestValidationError is the validation error returned by
// ProfileRequest.Validate if the designated constraints aren't met.
type ProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileRequestValidationError) ErrorName() string { return "ProfileRequestValidationError" }

// Error satisfies the builtin error interface
func (e ProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileRequestValidationError{}

// Validate checks the field values on CheckBalanceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckBalanceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ContractAddress

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

	// no validation rules for StreamId

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

// Validate checks the field values on TranscoderStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TranscoderStatusRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TranscoderId

	// no validation rules for Status

	return nil
}

// TranscoderStatusRequestValidationError is the validation error returned by
// TranscoderStatusRequest.Validate if the designated constraints aren't met.
type TranscoderStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TranscoderStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TranscoderStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TranscoderStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TranscoderStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TranscoderStatusRequestValidationError) ErrorName() string {
	return "TranscoderStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TranscoderStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTranscoderStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TranscoderStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TranscoderStatusRequestValidationError{}

// Validate checks the field values on StreamStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamStatusRequest) Validate() error {
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

// StreamStatusRequestValidationError is the validation error returned by
// StreamStatusRequest.Validate if the designated constraints aren't met.
type StreamStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamStatusRequestValidationError) ErrorName() string {
	return "StreamStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StreamStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamStatusRequestValidationError{}

// Validate checks the field values on JobResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *JobResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RtmpInputUrl

	return nil
}

// JobResponseValidationError is the validation error returned by
// JobResponse.Validate if the designated constraints aren't met.
type JobResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobResponseValidationError) ErrorName() string { return "JobResponseValidationError" }

// Error satisfies the builtin error interface
func (e JobResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobResponseValidationError{}

// Validate checks the field values on JobRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JobRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PipelineId

	// no validation rules for ClientAddress

	if _, ok := v1.ProfileId_name[int32(m.GetProfileId())]; !ok {
		return JobRequestValidationError{
			field:  "ProfileId",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// JobRequestValidationError is the validation error returned by
// JobRequest.Validate if the designated constraints aren't met.
type JobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobRequestValidationError) ErrorName() string { return "JobRequestValidationError" }

// Error satisfies the builtin error interface
func (e JobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobRequestValidationError{}

// Validate checks the field values on StreamRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StreamRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for StreamId

	return nil
}

// StreamRequestValidationError is the validation error returned by
// StreamRequest.Validate if the designated constraints aren't met.
type StreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamRequestValidationError) ErrorName() string { return "StreamRequestValidationError" }

// Error satisfies the builtin error interface
func (e StreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamRequestValidationError{}
