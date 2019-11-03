package dtos

import "time"

type CreateTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	StartAt time.Time `json:"startAt"`
	EndAt time.Time `json:"endAt"`
}

type CreateTaskResponse struct {
	Meta `json:"meta"`
}