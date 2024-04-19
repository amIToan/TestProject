package util

import (
	"mime/multipart"

	"github.com/go-playground/validator/v10"
)

var ValidateRole validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if role, ok := fieldLevel.Field().Interface().(string); ok {
		return IsValidRoleOrNot(role)
	}
	return false
}

var AtLeastOneFile validator.Func = func(fl validator.FieldLevel) bool {
	if files, ok := fl.Field().Interface().([]*multipart.FileHeader); ok {
		return len(files) > 0
	}
	return true
}
