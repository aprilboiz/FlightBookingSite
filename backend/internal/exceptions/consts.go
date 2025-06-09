package exceptions

import "net/http"

const (
	INTERNAL                = "internal_error"
	NotFound                = "not_found"
	BadRequest              = "bad_request"
	UNAUTHORIZED            = "unauthorized"
	FORBIDDEN               = "forbidden"
	CONFLICT                = "conflict"
	UnprocessableEntity     = "unprocessable_entity"
	UnsupportedMediaType    = "unsupported_media_type"
	TooManyRequests         = "too_many_requests"
	ServiceUnavailable      = "service_unavailable"
	GatewayTimeout          = "gateway_timeout"
	NotImplemented          = "not_implemented"
	HttpVersionNotSupported = "http_version_not_supported"
	MethodNotAllowed        = "method_not_allowed"
)

var ErrorType = map[string]*ErrorInfo{
	INTERNAL: {
		StatusCode: http.StatusInternalServerError,
		Title:      http.StatusText(http.StatusInternalServerError),
	},
	NotFound: {
		StatusCode: http.StatusNotFound,
		Title:      http.StatusText(http.StatusNotFound),
	},
	BadRequest: {
		StatusCode: http.StatusBadRequest,
		Title:      http.StatusText(http.StatusBadRequest),
	},
	UNAUTHORIZED: {
		StatusCode: http.StatusUnauthorized,
		Title:      http.StatusText(http.StatusUnauthorized),
	},
	FORBIDDEN: {
		StatusCode: http.StatusForbidden,
		Title:      http.StatusText(http.StatusForbidden),
	},
	CONFLICT: {
		StatusCode: http.StatusConflict,
		Title:      http.StatusText(http.StatusConflict),
	},
	UnprocessableEntity: {
		StatusCode: http.StatusUnprocessableEntity,
		Title:      http.StatusText(http.StatusUnprocessableEntity),
	},
	UnsupportedMediaType: {
		StatusCode: http.StatusUnsupportedMediaType,
		Title:      http.StatusText(http.StatusUnsupportedMediaType),
	},
	TooManyRequests: {
		StatusCode: http.StatusTooManyRequests,
		Title:      http.StatusText(http.StatusTooManyRequests),
	},
	ServiceUnavailable: {
		StatusCode: http.StatusServiceUnavailable,
		Title:      http.StatusText(http.StatusServiceUnavailable),
	},
	GatewayTimeout: {
		StatusCode: http.StatusGatewayTimeout,
		Title:      http.StatusText(http.StatusGatewayTimeout),
	},
	NotImplemented: {
		StatusCode: http.StatusNotImplemented,
		Title:      http.StatusText(http.StatusNotImplemented),
	},
	HttpVersionNotSupported: {
		StatusCode: http.StatusHTTPVersionNotSupported,
		Title:      http.StatusText(http.StatusHTTPVersionNotSupported),
	},
	MethodNotAllowed: {
		StatusCode: http.StatusMethodNotAllowed,
		Title:      http.StatusText(http.StatusMethodNotAllowed),
	},
}
