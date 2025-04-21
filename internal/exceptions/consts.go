package exceptions

import "net/http"

const (
	INTERNAL_ERROR             = "internal_error"
	NOT_FOUND                  = "not_found"
	BAD_REQUEST                = "bad_request"
	UNAUTHORIZED               = "unauthorized"
	FORBIDDEN                  = "forbidden"
	CONFLICT                   = "conflict"
	UNPROCESSABLE_ENTITY       = "unprocessable_entity"
	UNSUPPORTED_MEDIA_TYPE     = "unsupported_media_type"
	TOO_MANY_REQUESTS          = "too_many_requests"
	SERVICE_UNAVAILABLE        = "service_unavailable"
	GATEWAY_TIMEOUT            = "gateway_timeout"
	NOT_IMPLEMENTED            = "not_implemented"
	HTTP_VERSION_NOT_SUPPORTED = "http_version_not_supported"
	METHOD_NOT_ALLOWED         = "method_not_allowed"
)

var ErrorType = map[string]*ErrorInfo{
	INTERNAL_ERROR: {
		StatusCode: http.StatusInternalServerError,
		Title:      http.StatusText(http.StatusInternalServerError),
	},
	NOT_FOUND: {
		StatusCode: http.StatusNotFound,
		Title:      http.StatusText(http.StatusNotFound),
	},
	BAD_REQUEST: {
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
	UNPROCESSABLE_ENTITY: {
		StatusCode: http.StatusUnprocessableEntity,
		Title:      http.StatusText(http.StatusUnprocessableEntity),
	},
	UNSUPPORTED_MEDIA_TYPE: {
		StatusCode: http.StatusUnsupportedMediaType,
		Title:      http.StatusText(http.StatusUnsupportedMediaType),
	},
	TOO_MANY_REQUESTS: {
		StatusCode: http.StatusTooManyRequests,
		Title:      http.StatusText(http.StatusTooManyRequests),
	},
	SERVICE_UNAVAILABLE: {
		StatusCode: http.StatusServiceUnavailable,
		Title:      http.StatusText(http.StatusServiceUnavailable),
	},
	GATEWAY_TIMEOUT: {
		StatusCode: http.StatusGatewayTimeout,
		Title:      http.StatusText(http.StatusGatewayTimeout),
	},
	NOT_IMPLEMENTED: {
		StatusCode: http.StatusNotImplemented,
		Title:      http.StatusText(http.StatusNotImplemented),
	},
	HTTP_VERSION_NOT_SUPPORTED: {
		StatusCode: http.StatusHTTPVersionNotSupported,
		Title:      http.StatusText(http.StatusHTTPVersionNotSupported),
	},
	METHOD_NOT_ALLOWED: {
		StatusCode: http.StatusMethodNotAllowed,
		Title:      http.StatusText(http.StatusMethodNotAllowed),
	},
}
