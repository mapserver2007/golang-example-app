// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package proto

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Age

	return nil
}

// GetUserResponseValidationError is the validation error returned by
// GetUserResponse.Validate if the designated constraints aren't met.
type GetUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserResponseValidationError) ErrorName() string { return "GetUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserResponseValidationError{}

// Validate checks the field values on GetUsersResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetUsersResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUsersResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetUsersResponseValidationError is the validation error returned by
// GetUsersResponse.Validate if the designated constraints aren't met.
type GetUsersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersResponseValidationError) ErrorName() string { return "GetUsersResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetUsersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersResponseValidationError{}

// Validate checks the field values on GetUsersAndItemsRespones with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUsersAndItemsRespones) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUsersAndItemsResponesValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUsersAndItemsResponesValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetUsersAndItemsResponesValidationError is the validation error returned by
// GetUsersAndItemsRespones.Validate if the designated constraints aren't met.
type GetUsersAndItemsResponesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersAndItemsResponesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersAndItemsResponesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersAndItemsResponesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersAndItemsResponesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersAndItemsResponesValidationError) ErrorName() string {
	return "GetUsersAndItemsResponesValidationError"
}

// Error satisfies the builtin error interface
func (e GetUsersAndItemsResponesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersAndItemsRespones.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersAndItemsResponesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersAndItemsResponesValidationError{}

// Validate checks the field values on PostUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PostUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 30 {
		return PostUserRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 30 runes, inclusive",
		}
	}

	if !_PostUserRequest_Name_Pattern.MatchString(m.GetName()) {
		return PostUserRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9\\\\s]+$\"",
		}
	}

	if m.GetAge() <= 0 {
		return PostUserRequestValidationError{
			field:  "Age",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// PostUserRequestValidationError is the validation error returned by
// PostUserRequest.Validate if the designated constraints aren't met.
type PostUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostUserRequestValidationError) ErrorName() string { return "PostUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e PostUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPostUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostUserRequestValidationError{}

var _PostUserRequest_Name_Pattern = regexp.MustCompile("^[a-zA-Z0-9\\s]+$")

// Validate checks the field values on PutUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PutUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Age

	return nil
}

// PutUserRequestValidationError is the validation error returned by
// PutUserRequest.Validate if the designated constraints aren't met.
type PutUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutUserRequestValidationError) ErrorName() string { return "PutUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e PutUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutUserRequestValidationError{}

// Validate checks the field values on SimpleApiResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SimpleApiResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// SimpleApiResponseValidationError is the validation error returned by
// SimpleApiResponse.Validate if the designated constraints aren't met.
type SimpleApiResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SimpleApiResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SimpleApiResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SimpleApiResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SimpleApiResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SimpleApiResponseValidationError) ErrorName() string {
	return "SimpleApiResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SimpleApiResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSimpleApiResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SimpleApiResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SimpleApiResponseValidationError{}
