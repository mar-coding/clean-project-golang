// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: blog.proto

package blog

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on Permission with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Permission) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Permission with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PermissionMultiError, or
// nil if none found.
func (m *Permission) ValidateAll() error {
	return m.validate(true)
}

func (m *Permission) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Optional

	// no validation rules for ValidatePermissions

	// no validation rules for Captcha

	if len(errors) > 0 {
		return PermissionMultiError(errors)
	}

	return nil
}

// PermissionMultiError is an error wrapping multiple validation errors
// returned by Permission.ValidateAll() if the designated constraints aren't met.
type PermissionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PermissionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PermissionMultiError) AllErrors() []error { return m }

// PermissionValidationError is the validation error returned by
// Permission.Validate if the designated constraints aren't met.
type PermissionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PermissionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PermissionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PermissionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PermissionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PermissionValidationError) ErrorName() string { return "PermissionValidationError" }

// Error satisfies the builtin error interface
func (e PermissionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPermission.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PermissionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PermissionValidationError{}

// Validate checks the field values on CreateArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticleRequestMultiError, or nil if none found.
func (m *CreateArticleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTitle()); l < 5 || l > 500 {
		err := CreateArticleRequestValidationError{
			field:  "Title",
			reason: "value length must be between 5 and 500 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Content

	if len(errors) > 0 {
		return CreateArticleRequestMultiError(errors)
	}

	return nil
}

// CreateArticleRequestMultiError is an error wrapping multiple validation
// errors returned by CreateArticleRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateArticleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticleRequestMultiError) AllErrors() []error { return m }

// CreateArticleRequestValidationError is the validation error returned by
// CreateArticleRequest.Validate if the designated constraints aren't met.
type CreateArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleRequestValidationError) ErrorName() string {
	return "CreateArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleRequestValidationError{}

// Validate checks the field values on Article with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Article) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Article with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ArticleMultiError, or nil if none found.
func (m *Article) ValidateAll() error {
	return m.validate(true)
}

func (m *Article) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Content

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ArticleValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ArticleValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArticleValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ArticleValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ArticleValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArticleValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ArticleMultiError(errors)
	}

	return nil
}

// ArticleMultiError is an error wrapping multiple validation errors returned
// by Article.ValidateAll() if the designated constraints aren't met.
type ArticleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleMultiError) AllErrors() []error { return m }

// ArticleValidationError is the validation error returned by Article.Validate
// if the designated constraints aren't met.
type ArticleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleValidationError) ErrorName() string { return "ArticleValidationError" }

// Error satisfies the builtin error interface
func (e ArticleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleValidationError{}

// Validate checks the field values on GetArticleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetArticleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetArticleRequestMultiError, or nil if none found.
func (m *GetArticleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetArticleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetArticleRequestMultiError(errors)
	}

	return nil
}

// GetArticleRequestMultiError is an error wrapping multiple validation errors
// returned by GetArticleRequest.ValidateAll() if the designated constraints
// aren't met.
type GetArticleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetArticleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetArticleRequestMultiError) AllErrors() []error { return m }

// GetArticleRequestValidationError is the validation error returned by
// GetArticleRequest.Validate if the designated constraints aren't met.
type GetArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleRequestValidationError) ErrorName() string {
	return "GetArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleRequestValidationError{}