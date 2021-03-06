// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/buffer/v2/buffer.proto

package v2

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

// Validate checks the field values on Buffer with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Buffer) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetMaxRequestBytes(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return BufferValidationError{
				field:  "MaxRequestBytes",
				reason: "value must be greater than 0",
			}
		}

	}

	if m.GetMaxRequestTime() == nil {
		return BufferValidationError{
			field:  "MaxRequestTime",
			reason: "value is required",
		}
	}

	if d := m.GetMaxRequestTime(); d != nil {
		dur := *d

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return BufferValidationError{
				field:  "MaxRequestTime",
				reason: "value must be greater than 0s",
			}
		}

	}

	return nil
}

// BufferValidationError is the validation error returned by Buffer.Validate if
// the designated constraints aren't met.
type BufferValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferValidationError) ErrorName() string { return "BufferValidationError" }

// Error satisfies the builtin error interface
func (e BufferValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuffer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferValidationError{}

// Validate checks the field values on BufferPerRoute with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BufferPerRoute) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Override.(type) {

	case *BufferPerRoute_Disabled:

		if m.GetDisabled() != true {
			return BufferPerRouteValidationError{
				field:  "Disabled",
				reason: "value must equal true",
			}
		}

	case *BufferPerRoute_Buffer:

		if m.GetBuffer() == nil {
			return BufferPerRouteValidationError{
				field:  "Buffer",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetBuffer()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BufferPerRouteValidationError{
					field:  "Buffer",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return BufferPerRouteValidationError{
			field:  "Override",
			reason: "value is required",
		}

	}

	return nil
}

// BufferPerRouteValidationError is the validation error returned by
// BufferPerRoute.Validate if the designated constraints aren't met.
type BufferPerRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferPerRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferPerRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferPerRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferPerRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferPerRouteValidationError) ErrorName() string { return "BufferPerRouteValidationError" }

// Error satisfies the builtin error interface
func (e BufferPerRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferPerRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferPerRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferPerRouteValidationError{}
