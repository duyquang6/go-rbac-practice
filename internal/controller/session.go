package controller

import (
	"encoding/gob"
	"time"

	"github.com/gorilla/sessions"
)

type sessionKey string

const (
	sessionKeyLastActivity = sessionKey("lastActivity")
)

func init() {
	gob.Register(*new(sessionKey))
}

// StoreSessionLastActivity stores the last time the user did something. This is
// used to track idle session timeouts.
func StoreSessionLastActivity(session *sessions.Session, t time.Time) {
	if session == nil {
		return
	}
	session.Values[sessionKeyLastActivity] = t.Unix()
}

// ClearLastActivity clears the session last activity time.
func ClearLastActivity(session *sessions.Session) {
	sessionClear(session, sessionKeyLastActivity)
}

// LastActivityFromSession extracts the last time the user did something.
func LastActivityFromSession(session *sessions.Session) time.Time {
	v := sessionGet(session, sessionKeyLastActivity)
	if v == nil {
		return time.Time{}
	}

	i, ok := v.(int64)
	if !ok || i == 0 {
		delete(session.Values, sessionKeyLastActivity)
		return time.Time{}
	}

	return time.Unix(i, 0)
}

func sessionGet(session *sessions.Session, key sessionKey) interface{} {
	if session == nil || session.Values == nil {
		return nil
	}
	return session.Values[key]
}

func sessionClear(session *sessions.Session, key sessionKey) {
	if session == nil {
		return
	}
	delete(session.Values, key)
}
