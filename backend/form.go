package backend

type UserLoginForm struct {
	UserName       string
	Password       string
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