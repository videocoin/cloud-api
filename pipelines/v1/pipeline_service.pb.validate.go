// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pipelines/v1/pipeline_service.proto

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

	proto "github.com/VideoCoin/common/proto"
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

	_ = proto.ProfileId(0)
)

// Validate checks the field values on CreatePipelineRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatePipelineRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return CreatePipelineRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// CreatePipelineRequestValidationError is the validation error returned by
// CreatePipelineRequest.Validate if the designated constraints aren't met.
type CreatePipelineRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePipelineRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePipelineRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePipelineRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePipelineRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePipelineRequestValidationError) ErrorName() string {
	return "CreatePipelineRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePipelineRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePipelineRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePipelineRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePipelineRequestValidationError{}

// Validate checks the field values on PipelineRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PipelineRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// PipelineRequestValidationError is the validation error returned by
// PipelineRequest.Validate if the designated constraints aren't met.
type PipelineRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PipelineRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PipelineRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PipelineRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PipelineRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PipelineRequestValidationError) ErrorName() string { return "PipelineRequestValidationError" }

// Error satisfies the builtin error interface
func (e PipelineRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPipelineRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PipelineRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PipelineRequestValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}

// Validate checks the field values on UpdatePipelineRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdatePipelineRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) != 36 {
		return UpdatePipelineRequestValidationError{
			field:  "Id",
			reason: "value length must be 36 runes",
		}

	}

	// no validation rules for UserId

	// no validation rules for Name

	// no validation rules for StreamId

	// no validation rules for StreamAddress

	// no validation rules for Status

	// no validation rules for ProfileId

	return nil
}

// UpdatePipelineRequestValidationError is the validation error returned by
// UpdatePipelineRequest.Validate if the designated constraints aren't met.
type UpdatePipelineRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePipelineRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePipelineRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePipelineRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePipelineRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePipelineRequestValidationError) ErrorName() string {
	return "UpdatePipelineRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePipelineRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePipelineRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePipelineRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePipelineRequestValidationError{}
