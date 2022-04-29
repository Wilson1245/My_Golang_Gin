package middlewares

import (
	"golangAPI/pojo"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func UserPasd(field validator.FieldLevel) bool {
	if match, _ := regexp.MatchString(`^[A-Z][a-z]\d`, field.Field().String()); match {
		return true
	}
	return false
}

func UserList(field validator.StructLevel) {
	users := field.Current().Interface().(pojo.Users)
	if users.UserListSize == len(users.UserList) {

	} else {
		field.ReportError(users.UserListSize, "Size of user list", "UserListSize", "UserListSizeMustEuqalsUserList", "")
	}
}
