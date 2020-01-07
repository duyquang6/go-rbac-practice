package dtos

import "time"

type CreateTaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
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
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}

// TaskKafkaMsg present kafka message format for Task
type TaskKafkaMsg struct {
	ID          string           `json:"id"`
	RequestID   string           `json:"request_id"`
	RefEventID  string           `json:"ref_event_id"`
	Event       string           `json:"event"`
	ServiceCode string           `json:"service_code"`
	TimeStamp   int64            `json:"timestamp"`
	UserID      string           `json:"user_id"`
	PayloadID   string           `json:"payload_id"`
	Payload     TaskKafkaPayload `json:"payload"`
}

// StaffProfileKafkaPayload present kafka msg payload format for Staff
type TaskKafkaPayload struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}
