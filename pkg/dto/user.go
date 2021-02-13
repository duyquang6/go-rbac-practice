package dto

// CreateUserRequest
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// CreateUserRequest
type BindingRoleUserRequest struct {
	UserID  int64   `json:"user_id"`
	RoleIDs []int64 `json:"role_ids"`
}

// CreateUserResponse
type CreateUserResponse struct {
	Meta Meta `json:"meta"`
}

// DeleteUserRequest
type DeleteUserRequest struct {
	UserID int64 `json:"user_id"`
}

// DeleteUserResponse
type DeleteUserResponse struct {
	Meta Meta `json:"meta"`
}
