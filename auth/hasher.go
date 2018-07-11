package auth

import (
	"strings"
	"encoding/base64"
	"strconv"
	"hash"
	"crypto/sha256"
)

var UNUSABLE_PASSWORD_PREFIX = '!'  // This will never be a valid encoded hash
var UNUSABLE_PASSWORD_SUFFIX_LENGTH = 40  // number of random chars to add after UNUSABLE_PASSWORD_PREFIX


//func GetHasher()


func CheckPassword(password,encoded string ) bool {
	/*
	Returns a boolean of whether the raw password matches the three
    part encoded digest.

    If setter is specified, it'll be called when you need to
    regenerate the password.
	 */

	p := PBKDF2PasswordHasher{
		Iterations:36000,
		Algorithm:"pbkdf2_sha256",
		Digest:sha256.New,
	}
	return p.Verify(password,encoded)

}


func MakePassword(password, salt string) string {
    /*
    Turn a plain-text password into a hash for database storage
     */
    //if conf.PASSWORDHASHERS == "pbkdf2_sha256"{
    p := PBKDF2PasswordHasher{
    		Iterations:36000,
    		Algorithm:"pbkdf2_sha256",
    		Digest:sha256.New,
    	}
    if len(salt) == 0{
    	salt = GetRandomString(UNUSABLE_PASSWORD_SUFFIX_LENGTH)
	}
	return p.Encode(password,salt,p.Iterations)
	//}
}



func ConstantTtimeCompare(val1, val2 string) bool {
	/*
	 Returns True if the two strings are equal, False otherwise.
	 */
	if len(val1) != len(val2){
		return false
	}
	return val1 == val2

}

func MaskHash(hash string, show int ) string{
	/*
	Returns the given hash, with only the first ``show`` number shown. The
    rest are masked with ``*`` for security reasons.
	 */
	 masked := hash[:show]
	 masked += strings.Repeat("*" ,len(hash[show:]))
	 return masked
}




type PBKDF2PasswordHasher struct {
	Algorithm     string    `default:"pbkdf2_sha256"`
	Iterations    int       `default:"30000"`
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
	hashs := base64.StdEncoding.EncodeToString(hash_sha256)
	return strings.Join([]string{"pbkdf2_sha256",strconv.FormatInt(int64(iterations),10),salt,hashs},"$")
}


func (pbkdf *PBKDF2PasswordHasher) Verify(password string, encoded string) bool {
	encoded_list := strings.Split(encoded,"$")
	algorithm, iterations, salt, _ := encoded_list[0],encoded_list[1],encoded_list[2],encoded_list[3]
	if algorithm != pbkdf.Algorithm{
		panic("Algorithm of Hash Error")
	}
	i,_ := strconv.Atoi(iterations)
	encoded_2 := pbkdf.Encode(password,salt,i)
	return ConstantTtimeCompare(encoded,encoded_2)
}

func (pbkdf *PBKDF2PasswordHasher) SafeSummary(encoded string) map[string]string{
	encoded_list := strings.Split(encoded,"$")
	algorithm, iterations, salt, hashs := encoded_list[0],encoded_list[1],encoded_list[2],encoded_list[3]
	if algorithm != pbkdf.Algorithm{
		panic("Algorithm of Hash Error")
	}
	return map[string]string{
		"Algorithm" : algorithm,
		"Salt":MaskHash(salt,6),
		"Iterations":iterations,
		"Hash":MaskHash(hashs,6),
	}
}


func (pbkdf *PBKDF2PasswordHasher) MustUpdate(encoded string) bool{
	encoded_list := strings.Split(encoded,"$")
	iterations := encoded_list[1]
    iterations_int,_ := strconv.Atoi(iterations)
	return iterations_int != pbkdf.Iterations
}

func (pbkdf *PBKDF2PasswordHasher) HardenRuntime(password string, encoded string) {
	encoded_list := strings.Split(encoded,"$")
	iterations,salt := encoded_list[1],encoded_list[2]
	iterations_int,_ := strconv.Atoi(iterations)
	extra_iterations := pbkdf.Iterations - iterations_int
	if extra_iterations > 0{
		pbkdf.Encode(password,salt,extra_iterations)
	}

}