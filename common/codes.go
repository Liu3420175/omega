package common

var (
	CODES = map[int]string {
		0 : "Success",
		200 : "Sucess",
		10000: "Server error",
		10001: "Parameter error",  // 参数错误
		10002: "Token error",  // token 错误
		10003: "Token expired",  //token过期
		10004: "Please login",
		10005: "Login forbidden",  // 禁止登录
		10006: "Illegal request",
		10007: "You have no permission",
		10008: "Enter a valid URL",
		10009: "Policy is null",
		10010: "Can not support this Format ",
		10012: "Send Email ERROR",
		10013: "Http request method error",
		10014: "Date Format Error,such as yyyy-mm-dd",



		10101: "Email  Format is valid ",
		10102: "This Email had been registered",
		10103: "This Email had been frozen",
		10104: "Verification Code is error",
		10105: "Verification Code had expired",
		10106: "Two cipher inconsistencies",
		10107: "The Account  does not exist",
		10108: "Password error",
		10109: "New password and old password can not be the same",
		10110: "Captcha Code is error",
		10111: "Login forbidden",
		10112: "Email or password error",
		10113: "You can not login backend of shop",
		10114: "Auth code error",
		10115: "Auth code expired",
		10116: "Password format error",
		10117: "Activatin URL is invalid",
		10118: "Activatin URL is disabled",
		10119:"Your account has actived",
		10120: "SQL Error",

	}
)