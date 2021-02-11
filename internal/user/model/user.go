package model

import (
	"fmt"
	"strings"

	"github.com/duyquang6/go-rbac-practice/internal/dberror"
	_validator "github.com/duyquang6/go-rbac-practice/pkg/validator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	dberror.Errorable
	Name  string `validate:"required,max=255"`
	Email string `validate:"required,email"`
}

// BeforeSave is used by callbacks.
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Do validation
	if err := _validator.GetValidate().Struct(u); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			u.AddError("name", "err")
		} else {
			for _, _err := range err.(validator.ValidationErrors) {
				u.AddError(_err.Field(), _err.Error())
				fmt.Println(_err.Namespace())
				fmt.Println(_err.Field())
				fmt.Println(_err.StructNamespace())
				fmt.Println(_err.StructField())
				fmt.Println(_err.Tag())
				fmt.Println(_err.ActualTag())
				fmt.Println(_err.Kind())
				fmt.Println(_err.Type())
				fmt.Println(_err.Value())
				fmt.Println(_err.Param())
				fmt.Println()
			}
		}
	}

	// Custom validation (if any)

	if msgs := u.ErrorMessages(); len(msgs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(msgs, ", "))
	}

	return nil
}
