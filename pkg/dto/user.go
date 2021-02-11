package dto

// CreateUserRequest
type CreateUserRequest struct {
	Username string `json:"username"`
	Pasword  string `json:"password"`
	Email    string `json:"email"`
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
