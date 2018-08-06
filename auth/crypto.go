package auth

import (
	"time"
	"bytes"
	"encoding/binary"
	"math/rand"
	"crypto/sha256"
	"strconv"
	"hash"
	"crypto/hmac"

	"omega/conf"
	"crypto/sha1"
	"encoding/hex"
)

var (
	AllowedChars =  []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
)

func RandomChoice(choicechars []byte) byte{
	length := len(choicechars)
	random := rand.Intn(length)
	return choicechars[random]
}


func GetRandomString(length int) string {

	var seek int32
	t := time.Now()
	timesamp := t.UnixNano()
	buf := bytes.Buffer{}
	buf.Write(AllowedChars)
	buf.WriteString(strconv.FormatInt(timesamp,10))
	buf.WriteString(conf.SECRETKEY)
	s := buf.Bytes()// ping jie string
	haser := sha256.New()
	haser.Write(s)
	shash := haser.Sum(nil)
	bytesBuffer := bytes.NewBuffer(shash)

	binary.Read(bytesBuffer, binary.BigEndian, &seek)
    rand.Seed(int64(seek))

	buf2 := bytes.Buffer{}
	for i:=0;i<length;i++{
		tmp := RandomChoice(AllowedChars)
		buf2.WriteByte(tmp)
	}
	return buf2.String()
}



/*
 Returns the HMAC-SHA1 of 'value', using a key generated from key_salt and a
    secret (which defaults to settings.SECRET_KEY).

    A different key_salt should be passed in for every application of HMAC.
 */
func SaltedHhmac(key_salt, value, secret string) string{
	if len(secret) == 0 {
		secret = conf.SECRETKEY
	}
	keysalt := ForceBytes(key_salt)
	secret_ := ForceBytes(secret)
	mac := hmac.New(sha1.New,keysalt)
	mac.Write(secret_)
	mac.Write(ForceBytes(value))
	hash_ := mac.Sum(nil)

	return hex.EncodeToString(hash_) // 转换成16进制字符
}


func Pbkdf2(password string, salt string, iterations int, dklen int,digest func() hash.Hash) []byte {
	/*
	digest 是一个伪随机函数，例如HASH_HMAC函数，它会输出长度为hLen的结果。
    password是用来生成密钥的原文密码。
    salt是一个加密用的盐值。
    iterations是进行重复计算的次数。
    dklen是期望得到的密钥的长度。
	基本原理参见网上资料
	 */
    if iterations == 0 {
    	iterations = 36000
	}
    Password := []byte(password)
    Salt := []byte(salt)
    prf := hmac.New(digest,Password)
    haslen := prf.Size() // returns the number of bytes Sum will return
    numBolcks := (dklen + haslen - 1) / haslen

    var buf [4]byte
    dk := make([]byte,0,numBolcks * haslen)
    U := make([]byte,haslen)

    for block := 1;block <= numBolcks; block++ {
    	prf.Reset()
    	prf.Write(Salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-haslen:]
		copy(U, T)
		for n := 2; n <= iterations; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:dklen]
}