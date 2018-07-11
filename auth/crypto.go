package auth

import (
	"time"
    "../conf"
	"bytes"
	"encoding/binary"
	"math/rand"
	"crypto/sha256"
	"strconv"
	"strings"
	"hash"
	"crypto/hmac"
)


func RandomChoice(choicechars []string) string{
	length := len(choicechars)
	random := rand.Intn(length)
	return choicechars[random]
}


func GetRandomString(length int) string {
	AllowedChars := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m",
	"n","o","p","q","r","s","t","u","v","w","x","y","z","A","B","C","D","E","F","G","H","I",
	"J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","0","1","2","3",
	"4","5","6","7","8","9"}


	var seek int32
	t := time.Now()
	timesamp := t.UnixNano()
	buf := bytes.Buffer{}
	buf.WriteString(strings.Join(AllowedChars,""))
	buf.WriteString(strconv.FormatInt(timesamp,10))
	buf.WriteString(conf.SECRETKEY)
	s := buf.String() // ping jie string
	haser := sha256.New()
	haser.Write([]byte(s))
	shash := haser.Sum(nil)
	bytesBuffer := bytes.NewBuffer(shash)

	binary.Read(bytesBuffer, binary.BigEndian, &seek)
    rand.Seed(int64(seek))

	buf2 := bytes.Buffer{}
	for i:=0;i<length;i++{
		tmp := RandomChoice(AllowedChars)
		buf2.WriteString(tmp)
	}
	return buf2.String()
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