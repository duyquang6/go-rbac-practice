package model

type Policy struct {
	ID          int64
	Name        string
	Description string
	Roles       []*Role       `gorm:"many2many:policy_roles;"`
	Permissions []*Permission `gorm:"many2many:permission_policies;"`
}
