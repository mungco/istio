// Code generated by protoc-gen-validate
// source: envoy/admin/v2alpha/clusters.proto
// DO NOT EDIT!!!

package envoy_admin_v2alpha

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

// Validate checks the field values on Clusters with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Clusters) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetClusterStatuses() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClustersValidationError{
					Field:  fmt.Sprintf("ClusterStatuses[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// ClustersValidationError is the validation error returned by
// Clusters.Validate if the designated constraints aren't met.
type ClustersValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ClustersValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusters.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ClustersValidationError{}

// Validate checks the field values on ClusterStatus with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ClusterStatus) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for AddedViaApi

	if v, ok := interface{}(m.GetSuccessRateEjectionThreshold()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterStatusValidationError{
				Field:  "SuccessRateEjectionThreshold",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetHostStatuses() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterStatusValidationError{
					Field:  fmt.Sprintf("HostStatuses[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// ClusterStatusValidationError is the validation error returned by
// ClusterStatus.Validate if the designated constraints aren't met.
type ClusterStatusValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ClusterStatusValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterStatus.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ClusterStatusValidationError{}

// Validate checks the field values on HostStatus with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HostStatus) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HostStatusValidationError{
				Field:  "Address",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetStats() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HostStatusValidationError{
					Field:  fmt.Sprintf("Stats[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetHealthStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HostStatusValidationError{
				Field:  "HealthStatus",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSuccessRate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HostStatusValidationError{
				Field:  "SuccessRate",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// HostStatusValidationError is the validation error returned by
// HostStatus.Validate if the designated constraints aren't met.
type HostStatusValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HostStatusValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHostStatus.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HostStatusValidationError{}

// Validate checks the field values on HostHealthStatus with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HostHealthStatus) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FailedActiveHealthCheck

	// no validation rules for FailedOutlierCheck

	// no validation rules for EdsHealthStatus

	return nil
}

// HostHealthStatusValidationError is the validation error returned by
// HostHealthStatus.Validate if the designated constraints aren't met.
type HostHealthStatusValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HostHealthStatusValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHostHealthStatus.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HostHealthStatusValidationError{}
