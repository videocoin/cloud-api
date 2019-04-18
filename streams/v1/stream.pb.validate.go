// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: streams/v1/stream.proto

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

// Validate checks the field values on Stream with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Stream) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) != 36 {
		return StreamValidationError{
			field:  "Id",
			reason: "value length must be 36 runes",
		}

	}

	// no validation rules for UserId

	// no validation rules for Name

	// no validation rules for Status

	{
		tmp := m.GetCreatedAt()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return StreamValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for ProfileId

	// no validation rules for StreamId

	// no validation rules for StreamAddress

	return nil
}

// StreamValidationError is the validation error returned by Stream.Validate if
// the designated constraints aren't met.
type StreamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamValidationError) ErrorName() string { return "StreamValidationError" }

// Error satisfies the builtin error interface
func (e StreamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStream.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamValidationError{}

// Validate checks the field values on StreamProfile with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StreamProfile) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Status

	// no validation rules for ProfileId

	{
		tmp := m.GetCreatedAt()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return StreamProfileValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// StreamProfileValidationError is the validation error returned by
// StreamProfile.Validate if the designated constraints aren't met.
type StreamProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamProfileValidationError) ErrorName() string { return "StreamProfileValidationError" }

// Error satisfies the builtin error interface
func (e StreamProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamProfileValidationError{}
