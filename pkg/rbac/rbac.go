package rbac

type Object string

const (
	User Object = "user"
	Task Object = "task"
)

type GeneralPermission int64

const lengthGeneralPermission = 4

const (
	_ GeneralPermission = 1 << iota

	Read
	Create
	Update
	Delete
)

// Custom permission module

type TaskPermission int64

const (
	_ TaskPermission = 1 << (iota + lengthGeneralPermission)

	BulkInsert
	BulkDelete
)
