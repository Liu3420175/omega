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