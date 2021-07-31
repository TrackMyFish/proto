// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: trackmyfish/v1alpha1/trackmyfish.proto

package trackmyfishv1alpha1

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on AddFishRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AddFishRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFish()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddFishRequestValidationError{
				field:  "Fish",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Count

	return nil
}

// AddFishRequestValidationError is the validation error returned by
// AddFishRequest.Validate if the designated constraints aren't met.
type AddFishRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddFishRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddFishRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddFishRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddFishRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddFishRequestValidationError) ErrorName() string { return "AddFishRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddFishRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddFishRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddFishRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddFishRequestValidationError{}

// Validate checks the field values on AddFishResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AddFishResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFish()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddFishResponseValidationError{
				field:  "Fish",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddFishResponseValidationError is the validation error returned by
// AddFishResponse.Validate if the designated constraints aren't met.
type AddFishResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddFishResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddFishResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddFishResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddFishResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddFishResponseValidationError) ErrorName() string { return "AddFishResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddFishResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddFishResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddFishResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddFishResponseValidationError{}

// Validate checks the field values on ListFishRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListFishRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListFishRequestValidationError is the validation error returned by
// ListFishRequest.Validate if the designated constraints aren't met.
type ListFishRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFishRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFishRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFishRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFishRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFishRequestValidationError) ErrorName() string { return "ListFishRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListFishRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFishRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFishRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFishRequestValidationError{}

// Validate checks the field values on ListFishResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListFishResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetFish() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListFishResponseValidationError{
					field:  fmt.Sprintf("Fish[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListFishResponseValidationError is the validation error returned by
// ListFishResponse.Validate if the designated constraints aren't met.
type ListFishResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFishResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFishResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFishResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFishResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFishResponseValidationError) ErrorName() string { return "ListFishResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListFishResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFishResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFishResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFishResponseValidationError{}

// Validate checks the field values on DeleteFishRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteFishRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteFishRequestValidationError is the validation error returned by
// DeleteFishRequest.Validate if the designated constraints aren't met.
type DeleteFishRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteFishRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteFishRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteFishRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteFishRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteFishRequestValidationError) ErrorName() string {
	return "DeleteFishRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteFishRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteFishRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteFishRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteFishRequestValidationError{}

// Validate checks the field values on DeleteFishResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteFishResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFish()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteFishResponseValidationError{
				field:  "Fish",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeleteFishResponseValidationError is the validation error returned by
// DeleteFishResponse.Validate if the designated constraints aren't met.
type DeleteFishResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteFishResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteFishResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteFishResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteFishResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteFishResponseValidationError) ErrorName() string {
	return "DeleteFishResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteFishResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteFishResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteFishResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteFishResponseValidationError{}

// Validate checks the field values on HeartbeatRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HeartbeatRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// HeartbeatRequestValidationError is the validation error returned by
// HeartbeatRequest.Validate if the designated constraints aren't met.
type HeartbeatRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeartbeatRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeartbeatRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeartbeatRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeartbeatRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeartbeatRequestValidationError) ErrorName() string { return "HeartbeatRequestValidationError" }

// Error satisfies the builtin error interface
func (e HeartbeatRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeartbeatRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeartbeatRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeartbeatRequestValidationError{}

// Validate checks the field values on HeartbeatResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HeartbeatResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFishbase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeartbeatResponseValidationError{
				field:  "Fishbase",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HeartbeatResponseValidationError is the validation error returned by
// HeartbeatResponse.Validate if the designated constraints aren't met.
type HeartbeatResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeartbeatResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeartbeatResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeartbeatResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeartbeatResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeartbeatResponseValidationError) ErrorName() string {
	return "HeartbeatResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HeartbeatResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeartbeatResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeartbeatResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeartbeatResponseValidationError{}

// Validate checks the field values on HeartbeatStatus with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HeartbeatStatus) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// HeartbeatStatusValidationError is the validation error returned by
// HeartbeatStatus.Validate if the designated constraints aren't met.
type HeartbeatStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeartbeatStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeartbeatStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeartbeatStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeartbeatStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeartbeatStatusValidationError) ErrorName() string { return "HeartbeatStatusValidationError" }

// Error satisfies the builtin error interface
func (e HeartbeatStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeartbeatStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeartbeatStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeartbeatStatusValidationError{}

// Validate checks the field values on Fish with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Fish) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Genus

	// no validation rules for Species

	// no validation rules for CommonName

	// no validation rules for Name

	// no validation rules for Color

	// no validation rules for Gender

	// no validation rules for PurchaseDate

	// no validation rules for EcosystemName

	// no validation rules for EcosystemType

	// no validation rules for EcosystemLocation

	// no validation rules for Salinity

	// no validation rules for Climate

	return nil
}

// FishValidationError is the validation error returned by Fish.Validate if the
// designated constraints aren't met.
type FishValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FishValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FishValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FishValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FishValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FishValidationError) ErrorName() string { return "FishValidationError" }

// Error satisfies the builtin error interface
func (e FishValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFish.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FishValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FishValidationError{}
