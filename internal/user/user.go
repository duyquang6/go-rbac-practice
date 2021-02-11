// user package TBU
package user

import (
	"context"

	"github.com/duyquang6/go-rbac-practice/internal/user/model"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user model.User)
	UpdateProfile(ctx, user model.User)
}
