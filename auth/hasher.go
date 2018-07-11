package auth

import (
	"strings"
	"encoding/base64"
	"strconv"
	"hash"
)

var UNUSABLE_PASSWORD_PREFIX = '!'  // This will never be a valid encoded hash
var UNUSABLE_PASSWORD_SUFFIX_LENGTH = 40  // number of random chars to add after UNUSABLE_PASSWORD_PREFIX




type PBKDF2PasswordHasher struct {
	Algorithm     string    `default:"pbkdf2_sha256"`
	Iterations    int16     `default:"30000"`
	Digest        func() hash.Hash
}


func (pbkdf *PBKDF2PasswordHasher) Salt() string{
     return GetRandomString(12)
}


func (pbkdf *PBKDF2PasswordHasher) Encode(password string, salt string,iterations int) string {
    if len(password) ==0 {
    	panic("password is null")
	}
	if len(salt) > 0 && strings.Contains(salt,"$"){
        panic("Error")
	}
	if iterations == 0{
		iterations = 36000
	}
	hash_sha256 := Pbkdf2(password,salt,iterations,32,pbkdf.Digest)
	hash := base64.StdEncoding.EncodeToString(hash_sha256)
	return strings.Join([]string{"pbkdf2_sha256",strconv.FormatInt(int64(iterations),10),salt,hash},"$")
}