package model

type Policy struct {
	ID int64
	// TODO: for moment, not reuse policy
	Role        *Role `gorm:"notnull"`
	Permissions []*Permission
}
