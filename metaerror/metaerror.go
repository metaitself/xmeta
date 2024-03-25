package metaerror

import (
	"errors"
	"fmt"
	"github.com/metaitself/xmeta/conv"
	"github.com/metaitself/xmeta/encoding/json"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 500
	// UnknownReason is unknown reason for error info.
	UnknownReason = ""
	// StatusClientClosed is non-standard http status code,
	// which defined by nginx.
	// https://httpstatus.in/499/
	StatusClientClosed = 499
	// SupportPackageIsVersion1 this constant should not be referenced by any other code.
	SupportPackageIsVersion1 = true
)

type MetaError struct {
	MetaStatus
	cause error
}

var (
	_isDebug  = false
	_encoding = "json"
)

func (e *MetaError) Error() string {
	if _encoding == "" {
		return e.Message
	} else {
		return json.MarshalToString(e)
	}
}

func (e *MetaError) Unwrap() error { return e.cause }

// Is matches each error in the chain with the target value.
func (e *MetaError) Is(err error) bool {
	if se := new(MetaError); errors.As(err, &se) {
		return se.Code == e.Code && se.Reason == e.Reason
	}
	return false
}

func (e *MetaError) StatusCode() int32 {
	return e.Status
}

func (e *MetaError) ErrMessage() string { return e.Message }

// WithCause with the underlying cause of the error.
func (e *MetaError) WithCause(cause error) *MetaError {
	err := Clone(e)
	err.cause = cause
	if _isDebug {
		if err.Metadata == nil {
			err.Metadata = map[string]string{}
		}
		err.Metadata["cause"] = cause.Error()
	}
	return err
}

// WithReason with the underlying reason of the error.
func (e *MetaError) WithReason(reason string) *MetaError {
	err := Clone(e)
	err.Reason = reason
	return err
}

// WithMetadata with an MD formed by the mapping of key, value.
func (e *MetaError) WithMetadata(md map[string]string) *MetaError {
	err := Clone(e)
	err.Metadata = md
	return err
}

// GRPCStatus returns the Status represented by se.
func (e *MetaError) GRPCStatus() *status.Status {
	s, _ := status.New(httpStatusToGRPCCode(int(e.Status)), e.Message).
		WithDetails(&errdetails.ErrorInfo{
			Domain:   conv.ToString(e.Code),
			Reason:   e.Reason,
			Metadata: e.Metadata,
		})
	return s
}

// New returns an error object for the code, message.
func New(code, status int, reason, message string) *MetaError {
	return &MetaError{
		MetaStatus: MetaStatus{
			Code:    int32(code),
			Status:  int32(status),
			Message: message,
			Reason:  reason,
		},
	}
}

// Base returns an error object for the code, message and error info.
func Base(code int, format string, a ...interface{}) *MetaError {
	return New(code, http.StatusBadRequest, "", fmt.Sprintf(format, a...))
}

// Code returns the code for an error.
// It supports wrapped errors.
func Code(err error) int {
	if err == nil {
		return 0 //nolint:gomnd
	}
	return int(FromError(err).Code)
}

// StatusCode returns the http code for an error.
// It supports wrapped errors.
func StatusCode(err error) int {
	if err == nil {
		return 200 //nolint:gomnd
	}
	return int(FromError(err).Status)
}

// Reason returns the reason for a particular error.
// It supports wrapped errors.
func Reason(err error) string {
	if err == nil {
		return UnknownReason
	}
	return FromError(err).Reason
}

// Clone deep clone error to a new error.
func Clone(err *MetaError) *MetaError {
	if err == nil {
		return nil
	}
	metadata := make(map[string]string, len(err.Metadata))
	for k, v := range err.Metadata {
		metadata[k] = v
	}
	return &MetaError{
		cause: err.cause,
		MetaStatus: MetaStatus{
			Code:     err.Code,
			Reason:   err.Reason,
			Message:  err.Message,
			Metadata: metadata,
		},
	}
}

// FromError try to convert an error to *MetaError.
// It supports wrapped errors.
func FromError(err error) *MetaError {
	if err == nil {
		return nil
	}
	if se := new(MetaError); errors.As(err, &se) {
		return se
	}
	gs, ok := status.FromError(err)
	if !ok {
		return New(UnknownCode, UnknownCode, UnknownReason, err.Error())
	}
	ret := New(
		UnknownCode,
		httpStatusFromGRPCCode(gs.Code()),
		UnknownReason,
		gs.Message(),
	)
	for _, detail := range gs.Details() {
		switch d := detail.(type) {
		case *errdetails.ErrorInfo:
			ret.Reason = d.Reason
			ret.Code = conv.ToInt32(d.Domain)
			return ret.WithMetadata(d.Metadata)
		}
	}
	if ret.Metadata != nil {
		v, ok := ret.Metadata["cause"]
		if ok {
			ret.cause = errors.New(v)
		}
	}
	return ret
}

func SetDebugMode(b bool) {
	_isDebug = b
}

func SetErrEncode(v string) {
	_encoding = v
}

// httpStatusToGRPCCode converts an HTTP error code into the corresponding gRPC response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func httpStatusToGRPCCode(code int) codes.Code {
	switch code {
	case http.StatusOK:
		return codes.OK
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.Aborted
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case StatusClientClosed:
		return codes.Canceled
	}
	return codes.Unknown
}

// httpStatusFromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
func httpStatusFromGRPCCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return StatusClientClosed
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
