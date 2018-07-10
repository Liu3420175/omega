package auth



var UNUSABLE_PASSWORD_PREFIX = '!'  // This will never be a valid encoded hash
var UNUSABLE_PASSWORD_SUFFIX_LENGTH = 40  // number of random chars to add after UNUSABLE_PASSWORD_PREFIX




type PBKDF2PasswordHasher struct {
	Algorithm     string    `default:"pbkdf2_sha256"`
	Iterations    int16     `default:"30000"`
	Digest        func()
}


func (pbkdf *PBKDF2PasswordHasher) Salt() string{
     return GetRandomString(12)
}


func (pbkdf *PBKDF2PasswordHasher) Encode(password string, salt string,iterations int32) string {
    
}