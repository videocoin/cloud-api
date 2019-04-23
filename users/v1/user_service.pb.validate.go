// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: users/v1/user_service.proto

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

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Email

	// no validation rules for Password

	// no validation rules for Name

	// no validation rules for ConfirmPassword

	return nil
}

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on LoginUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LoginUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Email

	// no validation rules for Password

	return nil
}

// LoginUserRequestValidationError is the validation error returned by
// LoginUserRequest.Validate if the designated constraints aren't met.
type LoginUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUserRequestValidationError) ErrorName() string { return "LoginUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUserRequestValidationError{}

// Validate checks the field values on LoginUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LoginUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// LoginUserResponseValidationError is the validation error returned by
// LoginUserResponse.Validate if the designated constraints aren't met.
type LoginUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUserResponseValidationError) ErrorName() string {
	return "LoginUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LoginUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUserResponseValidationError{}

// Validate checks the field values on UserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// UserRequestValidationError is the validation error returned by
// UserRequest.Validate if the designated constraints aren't met.
type UserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRequestValidationError) ErrorName() string { return "UserRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRequestValidationError{}

// Validate checks the field values on ResetPasswordUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResetPasswordUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Password

	// no validation rules for ConfirmPassword

	return nil
}

// ResetPasswordUserRequestValidationError is the validation error returned by
// ResetPasswordUserRequest.Validate if the designated constraints aren't met.
type ResetPasswordUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordUserRequestValidationError) ErrorName() string {
	return "ResetPasswordUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ResetPasswordUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPasswordUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordUserRequestValidationError{}

// Validate checks the field values on StartRecoveryUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StartRecoveryUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Email

	return nil
}

// StartRecoveryUserRequestValidationError is the validation error returned by
// StartRecoveryUserRequest.Validate if the designated constraints aren't met.
type StartRecoveryUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartRecoveryUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartRecoveryUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartRecoveryUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartRecoveryUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartRecoveryUserRequestValidationError) ErrorName() string {
	return "StartRecoveryUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StartRecoveryUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartRecoveryUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartRecoveryUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartRecoveryUserRequestValidationError{}

// Validate checks the field values on RecoverUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RecoverUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	// no validation rules for Password

	// no validation rules for ConfirmPassword

	return nil
}

// RecoverUserRequestValidationError is the validation error returned by
// RecoverUserRequest.Validate if the designated constraints aren't met.
type RecoverUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecoverUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecoverUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecoverUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecoverUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecoverUserRequestValidationError) ErrorName() string {
	return "RecoverUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RecoverUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecoverUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecoverUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecoverUserRequestValidationError{}

// Validate checks the field values on WhitelistResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *WhitelistResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// WhitelistResponseValidationError is the validation error returned by
// WhitelistResponse.Validate if the designated constraints aren't met.
type WhitelistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WhitelistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WhitelistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WhitelistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WhitelistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WhitelistResponseValidationError) ErrorName() string {
	return "WhitelistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e WhitelistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWhitelistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WhitelistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WhitelistResponseValidationError{}

// Validate checks the field values on LookupByAddressRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LookupByAddressRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Address

	return nil
}

// LookupByAddressRequestValidationError is the validation error returned by
// LookupByAddressRequest.Validate if the designated constraints aren't met.
type LookupByAddressRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupByAddressRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupByAddressRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupByAddressRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupByAddressRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupByAddressRequestValidationError) ErrorName() string {
	return "LookupByAddressRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LookupByAddressRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupByAddressRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupByAddressRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupByAddressRequestValidationError{}
