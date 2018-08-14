package auth

import (
	"strings"
	"bytes"
	"regexp"
	"omega/conf"
)

var (
	REASON_NO_REFERER = "Referer checking failed - no Referer."
	REASON_BAD_REFERER = "Referer checking failed - %s does not match any trusted origins."
	REASON_NO_CSRF_COOKIE = "CSRF cookie not set."
	REASON_BAD_TOKEN = "CSRF token missing or incorrect."
	REASON_MALFORMED_REFERER = "Referer checking failed - Referer is malformed."
	REASON_INSECURE_REFERER = "Referer checking failed - Referer is insecure while host is secure."
	CSRF_ALLOWED_CHARS = string(AllowedChars)
	CSRF_SECRET_LENGTH = 32
	CSRF_TOKEN_LENGTH = 2 * CSRF_SECRET_LENGTH
	CSRF_COOKIE_NAME = "csrftoken"

)


func get_new_csrf_string() string{
	return GetRandomString(CSRF_SECRET_LENGTH)
}




func salt_cipher_secret(secret string) string {
	/*
	Given a secret (assumed to be a string of CSRF_ALLOWED_CHARS), generate a
    token by adding a salt and using it to encrypt the secret.
	 */

	salt := get_new_csrf_string()
	chars := CSRF_ALLOWED_CHARS
	chars_len := len(chars)
    key_value := [][2]int{}
    for i:= 0;i < CSRF_SECRET_LENGTH;i++{
    	k := strings.IndexByte(chars,secret[i])
    	v := strings.IndexByte(chars,salt[i])
    	key_value = append(key_value,[2]int{k,v})
	}
	paris := bytes.Buffer{}
	paris.WriteString(salt)
	for _,v := range key_value {
		index := ( v[0] +v[1]) % chars_len
		paris.WriteByte(chars[index])
	}
    return paris.String()
}



func unsalt_cipher_token(token string) string{
	/*
	Given a token (assumed to be a string of CSRF_ALLOWED_CHARS, of length
    CSRF_TOKEN_LENGTH, and that its first half is a salt), use it to decrypt
    the second half to produce the original secret.
	 */

    salt := token[:CSRF_SECRET_LENGTH]
    token_ := token[CSRF_SECRET_LENGTH:]
	chars := CSRF_ALLOWED_CHARS
	chars_len := len(chars)
	key_value := [][2]int{}
	for i:= 0;i < CSRF_SECRET_LENGTH;i++{
		k := strings.IndexByte(chars,token_[i])
		v := strings.IndexByte(chars,salt[i])
		key_value = append(key_value,[2]int{k,v})
	}
	paris := bytes.Buffer{}
	for _,v := range key_value{
		index := v[0] - v[1]
		if index < 0{
			index = chars_len + index
		}
		paris.WriteByte(chars[index])
	}
	return paris.String()
}


func get_new_csrf_token() string{
	return salt_cipher_secret(get_new_csrf_string())
}


func search(pattern , expr string) bool{
	reg,err := regexp.Compile(pattern)
	if err != nil{
		return false
	}
	s := reg.FindString(expr)
	return s == expr
}

func sanitize_token(token string) string {
	/*

	 */

	 if !search("[0-9A-Za-z]+",token){
	 	return get_new_csrf_token()
	 }else if len(token) == CSRF_TOKEN_LENGTH{
	 	return token
	 }else if len(token) == CSRF_SECRET_LENGTH{
	 	return salt_cipher_secret(token)
	 }else{}
	 return get_new_csrf_token()
}



func (request *Requester) get_token() string{
    /*

     */
     _,ok := request.CSRFMeta["CSRF_COOKIE"]
	csrf_secret := ""
     if ok{
	 csrf_secret = unsalt_cipher_token(request.CSRFMeta["CSRF_COOKIE"].(string))
	 }else{
		 csrf_secret = get_new_csrf_string()
		 request.CSRFMeta["CSRF_COOKIE"] = salt_cipher_secret(csrf_secret)
	 }
	request.CSRFMeta["CSRF_COOKIE_USED"] = true
	return salt_cipher_secret(csrf_secret)
}



func (request *Requester) rotate_token() {
    /*
    Changes the CSRF token in use for a request - should be done on login
    for security purposes.
     */
	request.CSRFMeta["CSRF_COOKIE_USED"] = true
	request.CSRFMeta["CSRF_COOKIE"] = get_new_csrf_token()
    request.CSRFMeta["CSRF_COOKIE_NEEDS_RESET"] = true

}



func (request *Requester) gettoken() string {
	cookie_token := request.Ctx.GetCookie(CSRF_COOKIE_NAME)
	if cookie_token == ""{
		return ""
	}
	csrf_token := sanitize_token(cookie_token)
	if csrf_token != cookie_token {
		request.CSRFMeta["CSRF_COOKIE_NEEDS_RESET"] = true
	}
	return csrf_token
}


func (request *Requester) settoken() {
     request.Ctx.SetCookie(CSRF_COOKIE_NAME,request.CSRFMeta["CSRF_COOKIE"].(string),
     	conf.SESSION_COOKIE_AGE,conf.CSRF_COOKIE_PATH,conf.CSRF_COOKIE_DOMAIN,conf.CSRF_COOKIE_SECURE,conf.CSRF_COOKIE_HTTPONLY)
}




func (request *Requester) CSRFProcess(){

	csrf_token := request.gettoken()
	if csrf_token != ""{
		request.CSRFMeta["CSRF_COOKIE"] = csrf_token
	}
	
}