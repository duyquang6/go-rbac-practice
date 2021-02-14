package database

import (
	"context"

	authorizationModel "github.com/duyquang6/go-rbac-practice/internal/authorization/model"
	"github.com/duyquang6/go-rbac-practice/internal/database"
	userModel "github.com/duyquang6/go-rbac-practice/internal/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user userModel.User) error
	BindingRoleUser(ctx context.Context, userID int64, role []authorizationModel.Role) error
	GetByUsername(ctx context.Context, username string) (userModel.User, error)
}

type authorizationDB struct {
	db *database.DB
}

func New(db *database.DB) *authorizationDB {
	return &authorizationDB{
		db: db,
	}
}

func (s *authorizationDB) Create(ctx context.Context, user userModel.User) error {
	res := s.db.Pool.Create(&user)
	return res.Error
}

func (s *authorizationDB) BindingRoleUser(ctx context.Context, userID int64, role []authorizationModel.Role) error {
	return s.db.Pool.Model(&userModel.User{
		Model: gorm.Model{ID: uint(userID)},
	}).Association("Roles").Append(role)
}

func (s *authorizationDB) GetByUsername(ctx context.Context, username string) (userModel.User, error) {
	user := userModel.User{
		Username: username,
	}
	res := s.db.Pool.First(&user)
	return user, res.Error
}
