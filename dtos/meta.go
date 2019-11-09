package dtos

import (
	"todolist-facebook-chatbot/common/httpcode"
)

// Meta is common meta.
type Meta struct {
	Code    int           `json:"code"`
	Message string        `json:"message,omitempty"`
	Params  []interface{} `json:"-"`
}

// PaginationMeta is meta information with pagination info.
type PaginationMeta struct {
	Meta
	Total int64 `json:"total"`
}

// CommonResponse is common structure of response without Data.
type CommonResponse struct {
	Meta `json:"meta"`
}

// Data when export file
type Data struct {
	URL string `json:"url"`
}

// BuildCommonResponse returns a bad request response.
func BuildCommonResponse(code int, msg string) CommonResponse {
	responseMessage := msg
	if msg == "" {
		responseMessage = httpcode.GetHTTPStatusText(code)
	}
	return CommonResponse{
		Meta: Meta{
			Code:    code,
			Message: responseMessage,
		},
	}
}

// NewMeta returns a new meta with message.
func NewMeta(code int, messages ...string) Meta {
	msg := httpcode.GetHTTPStatusText(code)
	if len(messages) > 0 {
		msg = messages[0]
	}
	return Meta{
		Code:    code,
		Message: msg,
	}
}

// NewMetaWithParams returns a new meta with custom message with params
func NewMetaWithParams(code int, params ...interface{}) Meta {
	meta := NewMeta(code)
	meta.Params = params
	return meta
}
