// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/ext_authz/v2alpha/ext_authz.proto

package v2alpha

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

// Validate checks the field values on ExtAuthz with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ExtAuthz) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FailureModeAllow

	// no validation rules for SendRequestData

	switch m.Services.(type) {

	case *ExtAuthz_GrpcService:

		if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtAuthzValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ExtAuthz_HttpService:

		if v, ok := interface{}(m.GetHttpService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtAuthzValidationError{
					field:  "HttpService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ExtAuthzValidationError is the validation error returned by
// ExtAuthz.Validate if the designated constraints aren't met.
type ExtAuthzValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtAuthzValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtAuthzValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtAuthzValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtAuthzValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtAuthzValidationError) ErrorName() string { return "ExtAuthzValidationError" }

// Error satisfies the builtin error interface
func (e ExtAuthzValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtAuthz.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtAuthzValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtAuthzValidationError{}

// Validate checks the field values on HttpService with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpService) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetServerUri()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpServiceValidationError{
				field:  "ServerUri",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PathPrefix

	for idx, item := range m.GetAuthorizationHeadersToAdd() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpServiceValidationError{
					field:  fmt.Sprintf("AuthorizationHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpServiceValidationError is the validation error returned by
// HttpService.Validate if the designated constraints aren't met.
type HttpServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpServiceValidationError) ErrorName() string { return "HttpServiceValidationError" }

// Error satisfies the builtin error interface
func (e HttpServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpServiceValidationError{}

// Validate checks the field values on ExtAuthzPerRoute with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ExtAuthzPerRoute) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Override.(type) {

	case *ExtAuthzPerRoute_Disabled:

		if m.GetDisabled() != true {
			return ExtAuthzPerRouteValidationError{
				field:  "Disabled",
				reason: "value must equal true",
			}
		}

	case *ExtAuthzPerRoute_CheckSettings:

		if m.GetCheckSettings() == nil {
			return ExtAuthzPerRouteValidationError{
				field:  "CheckSettings",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetCheckSettings()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtAuthzPerRouteValidationError{
					field:  "CheckSettings",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ExtAuthzPerRouteValidationError{
			field:  "Override",
			reason: "value is required",
		}

	}

	return nil
}

// ExtAuthzPerRouteValidationError is the validation error returned by
// ExtAuthzPerRoute.Validate if the designated constraints aren't met.
type ExtAuthzPerRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtAuthzPerRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtAuthzPerRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtAuthzPerRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtAuthzPerRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtAuthzPerRouteValidationError) ErrorName() string { return "ExtAuthzPerRouteValidationError" }

// Error satisfies the builtin error interface
func (e ExtAuthzPerRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtAuthzPerRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtAuthzPerRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtAuthzPerRouteValidationError{}

// Validate checks the field values on CheckSettings with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckSettings) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ContextExtensions

	return nil
}

// CheckSettingsValidationError is the validation error returned by
// CheckSettings.Validate if the designated constraints aren't met.
type CheckSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckSettingsValidationError) ErrorName() string { return "CheckSettingsValidationError" }

// Error satisfies the builtin error interface
func (e CheckSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckSettingsValidationError{}