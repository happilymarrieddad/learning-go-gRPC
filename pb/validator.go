package grpc

import (
	vCop "gopkg.in/go-playground/validator.v9"
)

var validator *vCop.Validate

func init() {
	validator = vCop.New()

	// Users
	validator.RegisterStructValidation(func(sl vCop.StructLevel) {
		r := sl.Current().Interface().(CreateUserRequest)

		if r.GetNewUser() == nil {
			sl.ReportError("NewUser", "newuser", "NewUser", "valid-newuser", "")
		} else {
			if len(r.GetNewUser().GetEmail()) == 0 {
				sl.ReportError("Email", "email", "Email", "valid-email", "")
			}
			if len(r.GetNewUser().GetFirstName()) == 0 {
				sl.ReportError("FirstName", "firstname", "FirstName", "valid-firstname", "")
			}
			if len(r.GetNewUser().GetLastName()) == 0 {
				sl.ReportError("LastName", "lastname", "LastName", "valid-lastname", "")
			}
			if len(r.GetNewUser().GetPassword()) == 0 {
				sl.ReportError("Password", "password", "Password", "valid-password", "")
			}
			if len(r.GetNewUser().GetConfirmPassword()) == 0 {
				sl.ReportError("ConfirmPassword", "confirmpassword", "ConfirmPassword", "valid-confirmpassword", "")
			}
		}

	}, CreateUserRequest{})
}

func Validate(t interface{}) error {
	return validator.Struct(t)
}
