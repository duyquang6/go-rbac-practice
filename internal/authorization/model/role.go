// Package model is a model abstraction of authorization module.
package model

type Role struct {
	ID          int64
	Name        string
	Description string
	Policies    []*Policy `gorm:"many2many:policy_roles;"`
}
