package dtos

import "time"

type CreateTaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
}

type CreateTaskResponse struct {
	Meta `json:"meta"`
}

// SearchStaffResponse represents response information of search staff API.
type GetAllTasksResponse struct {
	Meta PaginationMeta `json:"meta"`
	Data []Task         `json:"data"`
}

type Task struct {
}
