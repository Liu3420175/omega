package common

import (
	"math/rand"
	"bytes"
	"strings"
	"image"
	"image/color"
)

var InitChars = []string{"a","b","c","d","e","f","g","h","j","k","m",
		"n","p","q","r","s","t","u","v","w","x","y","A","B","C","D","E","F","G","H",
		"J","K","M","N","P","Q","R","S","T","U","V","W","X","Y","1","2","3",
		"4","5","6","7","8","9"}


var point_chance = 4

type AuthCode struct {
	Width          int
	Height         int
	Length         int
	FontSize       int
	DrawPoints     bool
	LineNumber     int  //绘制干扰线
	Model          string
}

// TODO 没有字符集，暂时搁置

func RandomChoice(choicechars []string) string{
	length := len(choicechars)
	random := rand.Intn(length)
	return choicechars[random]
}


func (authcode *AuthCode) GetChars() string{
	buf := bytes.Buffer{}
	for i := 0;i < authcode.Length;i++ {
		str := RandomChoice(InitChars)
		buf.WriteString(str)
	}
	return buf.String()
}


func (authcode *AuthCode)BackGroundColor() color.Color{
	  r := rand.Intn(100) + 155
	  g := rand.Intn(100) + 155
	  b := rand.Intn(100) + 155
	  s := color.RGBA{R:uint8(r),G:uint8(g),B:uint8(b),A:255}
	  return s
}


func (authcode *AuthCode) SetBackGroundColor() *image.RGBA{
	if len(authcode.Model) == 0 || strings.ToUpper(authcode.Model) == "RGBA"{
		rgba := image.NewRGBA(image.Rect(0,0,authcode.Width,authcode.Height))
		Color := authcode.BackGroundColor()
		for i:=0;i<authcode.Width;i++{
			for j:=0;j<authcode.Height;j++{
				rgba.Set(i,j,Color)
			}
		}
    return rgba
	}else{
		return nil
	}
}


func (authcode *AuthCode) DrowLine(){
	//绘制干扰线
	num := 0
	if authcode.LineNumber <= 0{
		num = 0
	}else if authcode.LineNumber > 2{
		num = 2
	}else{
		num = rand.Intn(2) + 1

	}

	for i := 0;i < num;i++ {

	}
}



func (authcode *AuthCode) DrowPoints(rgba *image.RGBA) *image.RGBA{
	//绘制干扰点

	for i:=0;i<authcode.Width;i++{
		for j:=0;j<authcode.Height;j++{
			tmp := rand.Intn(100)
			if tmp > 100 - point_chance {
				rgba.Set(i,j,color.RGBA{R:0,G:0,B:0,A:255})
			}
		}
	}
	return rgba
}




