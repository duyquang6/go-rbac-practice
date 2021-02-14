// user package TBU
package user

import (
	"context"

	authorizationModel "github.com/duyquang6/go-rbac-practice/internal/authorization/model"
	"github.com/duyquang6/go-rbac-practice/internal/user/database"
	"github.com/duyquang6/go-rbac-practice/internal/user/model"
	"github.com/duyquang6/go-rbac-practice/pkg/bcrypt"
	"github.com/duyquang6/go-rbac-practice/pkg/dto"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
)

type UserService interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) error
	BindingRoleUser(ctx context.Context, req dto.BindingRoleUserRequest) error
	GetByUsername(ctx context.Context, username string) (model.User, error)
}

type userSvc struct {
	userRepo database.UserRepository
}

func New(userRepo database.UserRepository) *userSvc {
	return &userSvc{
		userRepo: userRepo,
	}
}

func (u *userSvc) CreateUser(ctx context.Context, req dto.CreateUserRequest) error {
	logger := logging.FromContext(ctx).Named("CreateUser")

	hassPw, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		logger.Error("hass password failed: %v", err)
		return err
	}
	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hassPw,
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		logger.Errorf("create user failed: %v", err)
		return err
	}

	return nil
}

func (u *userSvc) BindingRoleUser(ctx context.Context, req dto.BindingRoleUserRequest) error {
	logger := logging.FromContext(ctx).Named("BindingRoleUser")
	var policies []authorizationModel.Role

	for _, id := range req.RoleIDs {
		policies = append(policies, authorizationModel.Role{ID: id})
	}

	if err := u.userRepo.BindingRoleUser(ctx, req.UserID, policies); err != nil {
		logger.Errorf("binding role user failed: %v", err)
		return err
	}
	return nil
}

func (u *userSvc) GetByUsername(ctx context.Context, username string) (model.User, error) {
	logger := logging.FromContext(ctx).Named("BindingRoleUser")
	user, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		logger.Errorf("get user failed: %v", err)
		return user, err
	}
	return user, err
}
