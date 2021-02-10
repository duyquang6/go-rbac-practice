package model

import "time"

type Task struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
	Title     string
	Body      string
}
