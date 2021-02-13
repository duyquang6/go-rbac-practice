package controller

import (
	"context"
	"encoding/gob"
	"time"

	"github.com/gorilla/sessions"
)

type contextKey string

const (
	contextRequestIDKey    contextKey = "requestID"
	contextKeySession      contextKey = "session"
	sessionKeyLastActivity contextKey = "lastActivity"
)

func init() {
	gob.Register(*new(contextKey))
}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextRequestIDKey, id)
}

func RequestIDFromContext(ctx context.Context) string {
	id, ok := ctx.Value(contextRequestIDKey).(string)
	if !ok {
		return ""
	}
	return id
}

// WithSession stores the session on the request's context for retrieval later.
// Use Session(r) to retrieve the session.
func WithSession(ctx context.Context, session *sessions.Session) context.Context {
	return context.WithValue(ctx, contextKeySession, session)
}

// SessionFromContext retrieves the session on the provided context. If no
// session exists, or if the value in the context is not of the correct type, it
// returns nil.
func SessionFromContext(ctx context.Context) *sessions.Session {
	v := ctx.Value(contextKeySession)
	if v == nil {
		return nil
	}

	t, ok := v.(*sessions.Session)
	if !ok {
		return nil
	}
	return t
}

// StoreSessionLastActivity stores the last time the user did something. This is
// used to track idle session timeouts.
func StoreSessionLastActivity(session *sessions.Session, t time.Time) {
	if session == nil {
		return
	}
	session.Values[sessionKeyLastActivity] = t.Unix()
}
