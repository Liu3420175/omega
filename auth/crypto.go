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
	return buf2.String()
}