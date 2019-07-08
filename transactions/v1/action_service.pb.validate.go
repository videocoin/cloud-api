// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: transactions/v1/action_service.proto

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

// Validate checks the field values on GetActionsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetActionsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Address

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// GetActionsRequestValidationError is the validation error returned by
// GetActionsRequest.Validate if the designated constraints aren't met.
type GetActionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActionsRequestValidationError) ErrorName() string {
	return "GetActionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetActionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActionsRequestValidationError{}

// Validate checks the field values on GetActionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetActionsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return GetActionsResponseValidationError{
						field:  fmt.Sprintf("Actions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// GetActionsResponseValidationError is the validation error returned by
// GetActionsResponse.Validate if the designated constraints aren't met.
type GetActionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActionsResponseValidationError) ErrorName() string {
	return "GetActionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetActionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActionsResponseValidationError{}

// Validate checks the field values on GetActionRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetActionRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Hash

	return nil
}

// GetActionRequestValidationError is the validation error returned by
// GetActionRequest.Validate if the designated constraints aren't met.
type GetActionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActionRequestValidationError) ErrorName() string { return "GetActionRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetActionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActionRequestValidationError{}

// Validate checks the field values on GetActionResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetActionResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetAction()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return GetActionResponseValidationError{
					field:  "Action",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// GetActionResponseValidationError is the validation error returned by
// GetActionResponse.Validate if the designated constraints aren't met.
type GetActionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActionResponseValidationError) ErrorName() string {
	return "GetActionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetActionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActionResponseValidationError{}
