package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	StartAt     time.Time
	EndAt       time.Time
}
