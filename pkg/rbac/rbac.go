package rbac

import (
	"errors"
	"log"
)

type Object string
type PermissionMapping map[string]int64

const (
	User Object = "user"
)

type GeneralPermission int64

const lengthGeneralPermission = 4

const (
	Read GeneralPermission = 1 << iota
	Create
	Update
	Delete
)

// Custom permission module

type UserPermission int64

const (
	BulkInsert UserPermission = 1 << (iota + lengthGeneralPermission)
	BulkDelete
)

func IsPermit(permissionMapping PermissionMapping, object Object, target interface{}) bool {
	if permissionMapping == nil {
		return false
	}

	source := permissionMapping[string(object)]
	_target, err := permissionToInt(target)
	if err != nil {
		log.Println("target permission isn't compliance:", err)
		panic("target permission isn't compliance")
	}
	return source&_target != 0
}

func AddImplied(object string, source int64, target int64) int64 {
	_, err := objectPermissionType(object, target)
	if err != nil {
		log.Println("add implied error:", err)
		return 0
	}

	return source | target
}

func permissionToInt(target interface{}) (int64, error) {
	switch v := target.(type) {
	case GeneralPermission:
		return int64(v), nil
	default:
		return 0, errors.New("no permission mapping")
	}
}

func objectPermissionType(object string, target int64) (interface{}, error) {
	if target <= 1<<lengthGeneralPermission {
		return GeneralPermission(target), nil
	}
	switch object {
	case string(User):
		return UserPermission(target), nil
	default:
		return 0, errors.New("no permission mapping")
	}
}

func validatePermissionType(target interface{}) bool {
	switch target.(type) {
	case GeneralPermission, UserPermission:
		return true
	default:
		return false
	}
}
