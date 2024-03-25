package metaerror

import "net/http"

// BadRequest new BadRequest error that is mapped to a 400 response.
func BadRequest(code int, reason, message string) *MetaError {
	return New(code, http.StatusBadRequest, reason, message)
}

// IsBadRequest determines if err is an error which indicates a BadRequest error.
// It supports wrapped errors.
func IsBadRequest(err error) bool {
	return Code(err) == http.StatusBadRequest
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(code int, reason, message string) *MetaError {
	return New(code, http.StatusUnauthorized, reason, message)
}

// IsUnauthorized determines if err is an error which indicates an Unauthorized error.
// It supports wrapped errors.
func IsUnauthorized(err error) bool {
	return Code(err) == http.StatusUnauthorized
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(code int, reason, message string) *MetaError {
	return New(code, http.StatusForbidden, reason, message)
}

// IsForbidden determines if err is an error which indicates a Forbidden error.
// It supports wrapped errors.
func IsForbidden(err error) bool {
	return Code(err) == http.StatusForbidden
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(code int, reason, message string) *MetaError {
	return New(code, http.StatusNotFound, reason, message)
}

// IsNotFound determines if err is an error which indicates an NotFound error.
// It supports wrapped errors.
func IsNotFound(err error) bool {
	return Code(err) == http.StatusNotFound
}

// Conflict new Conflict error that is mapped to a 409 response.
func Conflict(code int, reason, message string) *MetaError {
	return New(code, http.StatusConflict, reason, message)
}

// IsConflict determines if err is an error which indicates a Conflict error.
// It supports wrapped errors.
func IsConflict(err error) bool {
	return Code(err) == http.StatusConflict
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(code int, reason, message string) *MetaError {
	return New(code, http.StatusInternalServerError, reason, message)
}

// IsInternalServer determines if err is an error which indicates an Internal error.
// It supports wrapped errors.
func IsInternalServer(err error) bool {
	return Code(err) == http.StatusInternalServerError
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to an HTTP 503 response.
func ServiceUnavailable(code int, reason, message string) *MetaError {
	return New(code, http.StatusServiceUnavailable, reason, message)
}

// IsServiceUnavailable determines if err is an error which indicates an Unavailable error.
// It supports wrapped errors.
func IsServiceUnavailable(err error) bool {
	return Code(err) == http.StatusServiceUnavailable
}

// GatewayTimeout new GatewayTimeout error that is mapped to an HTTP 504 response.
func GatewayTimeout(code int, reason, message string) *MetaError {
	return New(code, http.StatusBadGateway, reason, message)
}

// IsGatewayTimeout determines if err is an error which indicates a GatewayTimeout error.
// It supports wrapped errors.
func IsGatewayTimeout(err error) bool {
	return Code(err) == http.StatusBadGateway
}

// ClientClosed new ClientClosed error that is mapped to an HTTP 499 response.
func ClientClosed(code int, reason, message string) *MetaError {
	return New(code, StatusClientClosed, reason, message)
}

// IsClientClosed determines if err is an error which indicates a IsClientClosed error.
// It supports wrapped errors.
func IsClientClosed(err error) bool {
	return Code(err) == StatusClientClosed
}
