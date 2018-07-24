package form

import (
	"strings"
	"github.com/pkg/errors"
	//"encoding/json"
	"strconv"
)

const TagName = "form"



var TagFormatError = errors.New("form tag format error")
var OverMaxLengthError = errors.New("Over MaxLength")
var LessMixLengthError = errors.New("Less MixLength")


type BaseField struct {
	Required        bool
	HelpText        string
	ErrorMessage    map[string]string
}



type CharField struct {
	BaseField
	MaxLength        int
	MinLength        int
	Default          string
}


type IntegerField struct {
	BaseField
	MaxValue        int64
	MixValue        int64
	Default         int64
}


type FloatField struct {
	BaseField
	MaxValue        float64
	MinValue        float64
	Default         float64
}

//type DateField struct{
//	BaseField

//}


func Interface2Int(value interface{}) int {
	switch value.(type) {
	case int,int8,int16,int32,int64:
		return value.(int)
	default:
		return 0
	}
}




func ParseFormTagString(tag string) (err error,errormessage map[string]string,object interface{}){
	// first remove ")"
	tag = strings.TrimSpace(tag)
	tag = strings.TrimLeft(tag,")")
	tag_list := strings.Split(tag,"(")
	if len(tag_list) != 2{
		err = TagFormatError
	}

	field ,tag_value := tag_list[0] , tag_list[1]

	switch field {
	case "CharField":
		err,errormessage,object = ParseCharField(tag_value)

	default:

	}

	return err,errormessage,object
}





func ParseCharField(tag string,fieldname string ,dest string) (error,map[string]string) {
	/*
	   tag : tag value
	   field:field name of strcuct object
	   dest : field value
	 */
	fields := map[string]string{}
	errormessage := map[string]string{}
	var err error
	if len(tag) == 0{
		// use default value
	}else{
		labels := strings.Split(tag,",")
		for _,values := range labels{
			key_value := strings.Split(values,"=")
			if len(key_value) >= 2{
				key := key_value[0]
				value := key_value[1]
				fields[key] = value
			}

		}
	}
	dest_length := len(dest)
	MaxLegth,err1 := strconv.Atoi(fields["MaxLength"])
	MinLegth,err2 := strconv.Atoi(fields["MinLength"])
	if err1 != nil{
		errormessage[fieldname] = fieldname + " error,MaxLength ca be " + fields["MaxLength"]
		err = err1
	}

	if MaxLegth > 0 && dest_length > MaxLegth{
		errormessage[fieldname] = errormessage[fieldname] + ";MaxLegth is Over " + fields["MaxLength"]
		err = OverMaxLengthError
	}

	if err2 != nil{
		errormessage[fieldname] = fieldname + " error,MinLength ca be " + fields["MinLength"]
		err = err2
	}

	if MinLegth > 0 && dest_length  < MinLegth{
		errormessage[fieldname] = errormessage[fieldname] + ";MinLegth is Less " + fields["MinLength"]
		err = LessMixLengthError
	}

}


func ParseIntegerField(tag string) (err error,errormessage map[string]string,field IntegerField){



	return err,errormessage,IntegerField{}
}