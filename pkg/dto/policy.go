package dto

type AppendPermissionPolicyRequest struct {
	PolicyID      int64   `json:"policy_id"`
	PermissionIDs []int64 `json:"permission_ids"`
}

// CreatePolicyRequest
type CreatePolicyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
