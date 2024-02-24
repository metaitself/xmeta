package metaerror

import (
	stderrors "errors"
	"fmt"
	"github.com/metaitself/xmeta/encoding/json"
	"github.com/metaitself/xmeta/metadata"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"strings"
)

// MetaError is a status error.
type MetaError struct {
	Code     int32             `json:"code,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Message  string            `json:"message,omitempty"`
	Metadata metadata.Metadata `json:"metadata,omitempty"`
	cause    error
}

var isDebugMode = false

func (e *MetaError) Error() string {
	toString, err := json.MarshalToString(e)
	if err != nil {
		return e.Message
	}
	return toString
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *MetaError) Unwrap() error { return e.cause }

// Is matches each error in the chain with the target value.
func (e *MetaError) Is(err error) bool {
	if se := new(MetaError); stderrors.As(err, &se) {
		return se.Code == e.Code && se.Reason == e.Reason
	}
	return false
}

func (e *MetaError) Errorf(format string, a ...interface{}) *MetaError {
	err := Clone(e)
	err.Message = fmt.Sprintf(format, a...)
	return err
}

func (e *MetaError) Append(format string, a ...interface{}) *MetaError {
	err := Clone(e)
	msg := strings.Builder{}
	msg.WriteString(err.Message)
	msg.WriteString(" [")
	msg.WriteString(fmt.Sprintf(format, a...))
	msg.WriteString("]")
	err.Message = msg.String()
	return err
}

func (e *MetaError) WithMessage(msg string) *MetaError {
	return &MetaError{
		Code:    e.Code,
		Message: msg,
	}
}

// WithCause with the underlying cause of the error.
func (e *MetaError) WithCause(cause error) *MetaError {
	err := Clone(e)
	err.cause = cause
	if isDebugMode {
		err.Reason = cause.Error()
	}
	return err
}

// WithMetadata with an MD formed by the mapping of key, value.
func (e *MetaError) WithMetadata(md metadata.Metadata) *MetaError {
	err := Clone(e)
	err.Metadata = md
	return err
}

// New returns an error object for the code, message.
func New(code int32, message string) *MetaError {
	return &MetaError{
		Code:    code,
		Message: message,
		Reason:  "",
		cause:   nil,
	}
}

// Errorf returns an error object for the code, message and error info.
func Errorf(code int32, format string, a ...interface{}) error {
	return New(code, fmt.Sprintf(format, a...))
}

func NewFromJson(buf string) *MetaError {
	err := MetaError{}
	if e := json.UnmarshalFromString(buf, &err); e != nil {
		return ErrUnmarshal
	}
	return &err
}

// Clone deep clone error to a new error.
func Clone(err *MetaError) *MetaError {
	if err == nil {
		return nil
	}
	md := make(metadata.Metadata, len(err.Metadata))
	for k, v := range err.Metadata {
		md[k] = v
	}
	return &MetaError{
		Code:     err.Code,
		cause:    err.cause,
		Reason:   err.Reason,
		Message:  err.Message,
		Metadata: md,
	}
}

// FromError try to convert an error to *ErrCode. It supports wrapped errcode.
func FromError(err error) *MetaError {
	if err == nil {
		return nil
	}

	if se := new(MetaError); stderrors.As(err, &se) {
		return se
	}

	gs, ok := status.FromError(err)
	if !ok {
		return ErrUnknown.WithMessage(gs.Message())
	}

	ret := New(int32(gs.Code()), gs.Message())
	for _, detail := range gs.Details() {
		switch d := detail.(type) {
		case *errdetails.ErrorInfo:
			ret.Reason = d.Reason
			fm := metadata.New()
			fm.FromStrMap(d.Metadata)
			return ret.WithMetadata(fm)
		}
	}
	return ret
}

func SetDebugMode(b bool) {
	isDebugMode = b
}

func Is(err, target error) bool {
	ee := FromError(err)
	if ee == nil {
		return false
	}
	te := FromError(target)
	if te == nil {
		return false
	}

	return ee.Code == te.Code
}

func IsCanceled(err error) bool {
	if err == nil {
		return false
	}

	if strings.Contains(err.Error(), "context canceled") {
		return true
	}
	return strings.Contains(err.Error(), "context deadline exceeded")
}
