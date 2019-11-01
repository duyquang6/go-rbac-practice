package dtos

import "net/http"

// Meta is metadata of response.
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// MetaCode returns Meta.Code as integer.
func (m *Meta) MetaCode() int {
	return m.Code
}

// NewMeta returns a new instance of Meta.
func NewMeta(code int, msg ...string) Meta {
	m := http.StatusText(code)

	if len(msg) > 0 && msg[0] != "" {
		m = msg[0]
	}

	return Meta{
		Code:    code,
		Message: m,
	}
}
