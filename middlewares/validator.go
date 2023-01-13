package middlewares

import (
	"SampleGoGin/pojo"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func UserPwd(field validator.FieldLevel) bool {
	if match, _ := regexp.MatchString(`^[A-Z]\w{4,10}$`, field.Field().String()); match {
		return true
	}
	return false
}

func UserList(field validator.StructLevel) {
	users := field.Current().Interface().(pojo.Users)
	if users.UserListSize == len(users.UserList) {

	} else {
		field.ReportError(users.UserListSize, "UserListSize", "UserListSize", "UserListSizeMustEqualsUserList", "")
	}
}
