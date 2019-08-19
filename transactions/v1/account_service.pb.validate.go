// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: transactions/v1/account_service.proto

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

// Validate checks the field values on GetAccountRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetAccountRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Address

	return nil
}

// GetAccountRequestValidationError is the validation error returned by
// GetAccountRequest.Validate if the designated constraints aren't met.
type GetAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountRequestValidationError) ErrorName() string {
	return "GetAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountRequestValidationError{}

// Validate checks the field values on GetAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAccountResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetAccount()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return GetAccountResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// GetAccountResponseValidationError is the validation error returned by
// GetAccountResponse.Validate if the designated constraints aren't met.
type GetAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountResponseValidationError) ErrorName() string {
	return "GetAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountResponseValidationError{}
