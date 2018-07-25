package backend

import "omega/form"

type Form struct {

	form.Form
}

type UserLoginForm struct {
	UserName       string           `form:"EmailField(Required=true,MaxLength=128)"`
	Password       string           `form:"CharField(Required=true,MinLength=8)"`
}


type UserAddForm struct {
	UserName        string
	Email           string
	Password        string
	FirstName       string
	LastName        string
	Phone           string
}


type EmailForm struct {
	Email           string
}


type EmailVerifyForm struct {
	Email           string
	Code            string
}


type PasswordRecoverForm struct {
	Password1       string
	Password2       string
}

type UserPasswordChangeForm struct {
	NewPassword     string
	OldPassword     string
}


type UserInfoCHangeForm struct {
	FirstName       string
	LastName        string
	Phone           string
}