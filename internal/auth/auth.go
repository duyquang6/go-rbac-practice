package auth

import (
	"context"
	"encoding/json"
	"errors"

	authorizationDB "github.com/duyquang6/go-rbac-practice/internal/authorization/database"
	userDB "github.com/duyquang6/go-rbac-practice/internal/user/database"
	"github.com/gorilla/sessions"

	"github.com/duyquang6/go-rbac-practice/pkg/bcrypt"
	"github.com/duyquang6/go-rbac-practice/pkg/dto"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
	"github.com/duyquang6/go-rbac-practice/pkg/rbac"
)

type AuthService interface {
	Login(ctx context.Context, session *sessions.Session, req dto.LoginRequest) error
	ClearSession(ctx context.Context, session *sessions.Session)
	GetSessionData(ctx context.Context, session *sessions.Session) (SessionData, error)
}

type authService struct {
	userRepo          userDB.UserRepository
	authorizationRepo authorizationDB.AuthorizationRepository
}

func New(userRepo userDB.UserRepository, authorizationRepo authorizationDB.AuthorizationRepository) *authService {
	return &authService{
		userRepo:          userRepo,
		authorizationRepo: authorizationRepo,
	}
}

func (s *authService) Login(ctx context.Context, session *sessions.Session, req dto.LoginRequest) error {
	logger := logging.FromContext(ctx).Named("auth.Login")
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		s.ClearSession(ctx, session)
		logger.Error("get user failed: ", err)
		return err
	}

	if !bcrypt.CheckPasswordHash(req.Password, user.Password) {
		return errors.New("invalid password")
	}

	// Get permission
	permissionData := make(rbac.PermissionMapping)
	permissions, err := s.authorizationRepo.GetPermissionByUser(ctx, int64(user.ID))
	if err != nil {
		s.ClearSession(ctx, session)
		logger.Error("get permission failed: ", err)
		return err
	}

	for _, perm := range permissions {
		permissionData[perm.Object] = rbac.AddImplied(perm.Object, permissionData[perm.Object], perm.Code)
	}

	// store session
	sessionData, err := json.Marshal(&SessionData{
		Username:   user.Username,
		Permission: permissionData,
	})
	if err != nil {
		s.ClearSession(ctx, session)
		logger.Error("marshal sesssion data failed: ", err)
		return err
	}
	if err = sessionSet(session, sessionDataKey, sessionData); err != nil {
		s.ClearSession(ctx, session)
		logger.Error("store session failed: ", err)
		return err
	}

	return nil
}

func (s *authService) GetSessionData(ctx context.Context, session *sessions.Session) (SessionData, error) {
	logger := logging.FromContext(ctx).Named("auth.GetSessionData")
	data := sessionGet(session, sessionDataKey)
	if data == nil {
		logger.Error("session data empty")
		return SessionData{}, errors.New("session data empty")
	}
	var sessionData SessionData

	if err := json.Unmarshal(data.([]byte), &sessionData); err != nil {
		logger.Error("data parse failed:", err)
		return SessionData{}, err
	}

	return sessionData, nil
}

// ClearSession removes any session information for this auth.
func (a *authService) ClearSession(ctx context.Context, session *sessions.Session) {
	sessionClear(session, sessionDataKey)
}
