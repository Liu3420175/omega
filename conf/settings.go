package conf

var DATABASES = map[string]map[string]string{
	"default":{
		"HOST":"localhost",
		"PORT":"3306",
		"USER":"root",
		"PASSWORD":"asdasd",
		"NAME":"omega",
	},
}




const (

	SECRETKEY = "m^a=%2su6%_f7ux8mkm9OSCAR_REQUIRED_ADDRESS_FIELDS+$*^c@&a#8)m@dtd($$!1&2j6)ij^g"
	//var PASSWORDHASHERS = "pbkdf2_sha256"
	SESSION_COOKIE_AGE = 7 * 24 * 60 * 60

	EMAIL_HOST = "m.qq.com"
	EMAIL_USERNAME = "xxxx"
	EMAIL_PASSWORD = "xxxx"

)

