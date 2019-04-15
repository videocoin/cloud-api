// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: users/v1/user.proto

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

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) != 36 {
		return UserValidationError{
			field:  "Id",
			reason: "value length must be 36 runes",
		}

	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return UserValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if utf8.RuneCountInString(m.GetPassword()) < 5 {
		return UserValidationError{
			field:  "Password",
			reason: "value length must be at least 5 runes",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 3 {
		return UserValidationError{
			field:  "Name",
			reason: "value length must be at least 3 runes",
		}
	}

	if m.GetCreatedAt() == nil {
		return UserValidationError{
			field:  "CreatedAt",
			reason: "value is required",
		}
	}

	// no validation rules for Token

	return nil
}

func (m *User) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *User) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on UserProfile with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserProfile) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Email

	// no validation rules for Name

	{
		tmp := m.GetAccount()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return UserProfileValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// UserProfileValidationError is the validation error returned by
// UserProfile.Validate if the designated constraints aren't met.
type UserProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserProfileValidationError) ErrorName() string { return "UserProfileValidationError" }

// Error satisfies the builtin error interface
func (e UserProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserProfileValidationError{}
