package httpcode

import (
	"fmt"
	"net/http"
	"strconv"
)

// HTTP code definition.
const (
	SuccessfullyCode        int = 200
	NoContentCode           int = 204
	BadRequestCode          int = 400
	UnauthorizedCode        int = 401
	ForbiddenCode           int = 403
	NotFoundCode            int = 404
	ConflictCode            int = 409
	ServerErrorCode         int = 500
	ForceChangePasswordCode int = 406
)

// GetHTTPCode returns HTTP status code.
func GetHTTPCode(code int) int {
	httpCode := ParseHTTPStatus(code, SuccessfullyCode)
	switch httpCode {
	case UnauthorizedCode, ForbiddenCode:
		return httpCode
	default:
		return SuccessfullyCode
	}
}

// GetHTTPStatusText returns status text of HTTP status code.
func GetHTTPStatusText(code int) string {
	return http.StatusText(ParseHTTPStatus(code, SuccessfullyCode))
}

// ParseHTTPStatus returns HTTP status code in error code.
func ParseHTTPStatus(code int, defaultCode ...int) int {
	var (
		def     int
		codeStr = fmt.Sprintf("%v", code)
	)
	if len(defaultCode) > 0 {
		def = defaultCode[0]
	}
	if len(codeStr) > 2 && codeStr[0] <= '9' && codeStr[0] >= '0' {
		httpStatus, _ := strconv.Atoi(codeStr[:3])
		return httpStatus
	}
	return def
}
