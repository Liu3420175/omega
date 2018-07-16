package conf

var DATABASES = map[string]map[string]string{
	"default":{
		"HOST":"127.0.0.1",
		"PORT":"3306",
		"USER":"root",
		"PASSWORD":"asdasd",
		"NAME":"shop",
	},
}

var SECRETKEY = "m^a=%2su6%_f7ux8mkm9OSCAR_REQUIRED_ADDRESS_FIELDS+$*^c@&a#8)m@dtd($$!1&2j6)ij^g"


//var PASSWORDHASHERS = "pbkdf2_sha256"

var SESSION_COOKIE_AGE = 7 * 24 * 60 * 60