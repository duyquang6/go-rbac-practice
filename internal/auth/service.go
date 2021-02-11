package auth

import "context"

type AuthService interface {
	Login(ctx context.Context, loginRequest *dto.LoginRequest)
}
