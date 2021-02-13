package rbac

type Object string

const (
	User Object = "user"
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

func IsPermit(source int64, target interface{}) bool {
	if !validatePermissionType(target) {
		panic("target permission isn't compliance")
	}
	return source&target.(int64) != 0
}

func AddImplied(source int64, target interface{}) int64 {
	if !validatePermissionType(target) {
		panic("target permission isn't compliance")
	}

	return source | target.(int64)
}

func validatePermissionType(target interface{}) bool {
	switch target.(type) {
	case GeneralPermission, TaskPermission:
		return true
	default:
		return false
	}
}
