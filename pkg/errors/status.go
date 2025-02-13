package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type Status string

const (
	StatusBadRequest          Status = "BadRequest"
	StatusUnauthorized        Status = "Unauthorized"
	StatusForbidden           Status = "Forbidden"
	StatusNotFound            Status = "NotFound"
	StatusTooManyRequests     Status = "TooManyRequests"
	StatusBadGateway          Status = "BadGateway"
	StatusInternalServerError Status = "InternalServerError"
	StatusServiceUnavailable  Status = "ServiceUnavailable"
	StatusGatewayTimeout      Status = "GatewayTimeout"
	StatusAlreadyExists       Status = "AlreadyExists"
	StatusNotImplemented      Status = "NotImplemented"
	StatusConflict            Status = "Conflict"
)

func (s Status) ToHTTPStatus() int {
	switch s {
	case StatusBadRequest:
		return http.StatusBadRequest
	case StatusUnauthorized:
		return http.StatusUnauthorized
	case StatusForbidden:
		return http.StatusForbidden
	case StatusNotFound:
		return http.StatusNotFound
	case StatusTooManyRequests:
		return http.StatusTooManyRequests
	case StatusBadGateway:
		return http.StatusBadGateway
	case StatusInternalServerError:
		return http.StatusInternalServerError
	case StatusServiceUnavailable:
		return http.StatusServiceUnavailable
	case StatusGatewayTimeout:
		return http.StatusGatewayTimeout
	case StatusAlreadyExists:
		return http.StatusConflict
	case StatusNotImplemented:
		return http.StatusNotImplemented
	case StatusConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func (s Status) ToGRPCStatus() codes.Code {
	switch s {
	case StatusBadRequest:
		return codes.InvalidArgument
	case StatusUnauthorized:
		return codes.Unauthenticated
	case StatusForbidden:
		return codes.PermissionDenied
	case StatusNotFound:
		return codes.Unimplemented
	case StatusTooManyRequests:
		return codes.Unavailable
	case StatusBadGateway:
		return codes.Unavailable
	case StatusInternalServerError:
		return codes.Internal
	case StatusServiceUnavailable:
		return codes.Unavailable
	case StatusGatewayTimeout:
		return codes.Unavailable
	case StatusAlreadyExists:
		return codes.AlreadyExists
	case StatusConflict:
		return codes.FailedPrecondition
	default:
		return codes.Internal
	}
}
