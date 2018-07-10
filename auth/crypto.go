package auth

import (
	"time"
    "../conf"
	"bytes"
	"encoding/binary"
	"math/rand"
	"crypto/sha256"
	"strconv"
)


func RandomChoice(choicechars string) string{
	length := len(choicechars)
	random := rand.Intn(length)
	return choicechars[random:random+1]
}


func GetRandomString(length int) string {
	AllowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seek int32
	t := time.Now()
	timesamp := t.UnixNano()
	buf := bytes.Buffer{}
	buf.WriteString(AllowedChars)
	buf.WriteString(strconv.FormatInt(timesamp,10))
	buf.WriteString(conf.SECRETKEY)
	s := buf.String()
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
	return buf.String()
}