// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/cluster/v3/cds.proto

package envoy_service_cluster_v3

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

// Validate checks the field values on CdsDummy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CdsDummy) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CdsDummyValidationError is the validation error returned by
// CdsDummy.Validate if the designated constraints aren't met.
type CdsDummyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CdsDummyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CdsDummyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CdsDummyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CdsDummyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CdsDummyValidationError) ErrorName() string { return "CdsDummyValidationError" }

// Error satisfies the builtin error interface
func (e CdsDummyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCdsDummy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CdsDummyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CdsDummyValidationError{}
