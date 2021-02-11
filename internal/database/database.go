package database

import (
	"gorm.io/gorm"
)

type DB struct {
	Pool *gorm.DB
}
