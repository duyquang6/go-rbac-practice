package dto

// CreateRoleRequest
type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateRoleResponse
type CreateRoleResponse struct {
}

// BindingPolicyRoleRequest
type BindingPolicyRoleRequest struct {
	RoleID    int64   `json:"role_id"`
	PolicyIDs []int64 `json:"policy_ids"`
}
