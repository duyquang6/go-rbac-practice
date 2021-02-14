package controller

import (
	"context"

	"github.com/duyquang6/go-rbac-practice/internal/user/model"
	"github.com/duyquang6/go-rbac-practice/pkg/rbac"
	"github.com/gorilla/sessions"
)

type contextKey string

const (
	contextRequestIDKey         contextKey = "requestID"
	contextKeySession           contextKey = "session"
	contextKeyUser              contextKey = "user"
	contextKeyPermissionMapping contextKey = "permissionMapping"
)

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

// WithUser stores the current user on the context.
func WithUser(ctx context.Context, u model.User) context.Context {
	return context.WithValue(ctx, contextKeyUser, u)
}

// UserFromContext retrieves the user from the context. If no value exists, it
// returns nil.
func UserFromContext(ctx context.Context) *model.User {
	v := ctx.Value(contextKeyUser)
	if v == nil {
		return nil
	}

	t, ok := v.(*model.User)
	if !ok {
		return nil
	}
	return t
}

// WithPermissionMapping stores the user's available memberships on the context.
func WithPermissionMapping(ctx context.Context, p rbac.PermissionMapping) context.Context {
	return context.WithValue(ctx, contextKeyPermissionMapping, p)
}

func PermissionMappingFromContext(ctx context.Context) rbac.PermissionMapping {
	v := ctx.Value(contextKeyPermissionMapping)
	if v == nil {
		return nil
	}

	t, ok := v.(rbac.PermissionMapping)
	if !ok {
		return nil
	}
	return t
}
