package auth


var (
	REASON_NO_REFERER = "Referer checking failed - no Referer."
	REASON_BAD_REFERER = "Referer checking failed - %s does not match any trusted origins."
	REASON_NO_CSRF_COOKIE = "CSRF cookie not set."
	REASON_BAD_TOKEN = "CSRF token missing or incorrect."
	REASON_MALFORMED_REFERER = "Referer checking failed - Referer is malformed."
	REASON_INSECURE_REFERER = "Referer checking failed - Referer is insecure while host is secure."

	CSRF_SECRET_LENGTH = 32
	CSRF_TOKEN_LENGTH = 2 * CSRF_SECRET_LENGTH

)


func GetNewCsrfString() string{
	return GetRandomString(CSRF_SECRET_LENGTH)
}


