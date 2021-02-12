package database

import (
	"context"

	authorizedModel "github.com/duyquang6/go-rbac-practice/internal/authorization/model"
	"github.com/duyquang6/go-rbac-practice/internal/database"
)

type AuthorizationRepository interface {
	CreateRole(ctx context.Context, role authorizedModel.Role) error
	BindingRole(ctx context.Context, role_id int64) error
	CreatePolicy(ctx context.Context, policy authorizedModel.Policy) error
}

type authorizationDB struct {
	db *database.DB
}

func New(db *database.DB) *authorizationDB {
	return &authorizationDB{
		db: db,
	}
}

func (db *authorizationDB) CreateRole(ctx context.Context, role authorizedModel.Role) error {
	return nil
}

func (db *authorizationDB) BindingRole(ctx context.Context, role_id int64) error {
	return nil
}

func (db *authorizationDB) CreatePolicy(ctx context.Context, policy authorizedModel.Policy) error {
	return nil
}
