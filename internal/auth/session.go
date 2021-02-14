package auth

import (
	"encoding/gob"
	"errors"

	"github.com/duyquang6/go-rbac-practice/pkg/rbac"
	"github.com/gorilla/sessions"
)

type sessionKey string

const (
	sessionDataKey = "sessionData"
)

type SessionData struct {
	Username   string                 `json:"username"`
	Permission rbac.PermissionMapping `json:"permission_code_mapping"`
}

func init() {
	gob.Register(*new(sessionKey))
	gob.Register(*new(SessionData))
}

func sessionGet(session *sessions.Session, key sessionKey) interface{} {
	if session == nil || session.Values == nil {
		return nil
	}
	return session.Values[key]
}

// sessionSet sets the value in the session.
func sessionSet(session *sessions.Session, key sessionKey, data interface{}) error {
	if session == nil {
		return errors.New("session missing")
	}

	if session.Values == nil {
		session.Values = make(map[interface{}]interface{})
	}

	session.Values[key] = data
	return nil
}
func sessionClear(session *sessions.Session, key sessionKey) {
	if session == nil {
		return
	}
	delete(session.Values, key)
}
