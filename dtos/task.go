package dtos

import "time"

type CreateTaskRequest struct {
	Title string
	Description string
	StartAt time.Time
	EndAt time.Time
}

type CreateTaskResponse struct {
	Meta `json:"meta"`
}