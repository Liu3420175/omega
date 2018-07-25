package form

import (
	"strings"
	"errors"
	//"encoding/json"
	"strconv"
	"regexp"
	"net/url"
)

const TagName = "form"



var (
	 TagFormatError = errors.New("form tag format error")
     OverMaxLengthError = errors.New("Over MaxLength")
     LessMinLengthError = errors.New("Less MinLength")
     RequiredError = errors.New("Field Required")
     OverMaxValueError = errors.New("Over MaxValue")
     LessMinValueError = errors.New("Less MinValue")
     EmailFormatError = errors.New("Email Format Error")
     URLFormatError = errors.New("Enter a valid URL")

)

type BaseField struct {
	Required        bool
	HelpText        string
	HasError        bool
	ErrorMessage    map[string]interface{}
}



type CharField struct {
	BaseField
	MaxLength        int
	MinLength        int
}


type IntegerField struct {
	BaseField
	MaxValue        int64
	MixValue        int64
}


type FloatField struct {
	BaseField
	MaxValue        float64
	MinValue        float64

}


type EmailField struct {
	CharField
}



type URLField struct {
	CharField
}





//type DateField struct{
//	BaseField

//}

type Validator interface {
	ParseTagString(string, string, interface{}) error
	HasErrors()    bool
	Errors()     map[string]interface{}
}



func Interface2Int(value interface{}) int {
	switch value.(type) {
	case int,int8,int16,int32,int64:
		return value.(int)
	default:
		return 0
	}
}




func ParseFormTagString(tag string) ( error,string,string){
	// first remove ")"

	var err error
	tag = strings.TrimSpace(tag)
	tag = strings.TrimRight(tag,")")
	tag_list := strings.Split(tag,"(")

	if len(tag_list) != 2{
		err = TagFormatError
	}

	fieldtype ,tag_value := tag_list[0] , tag_list[1]
	return err,fieldtype,tag_value
}



func ParseString(tag string)  map[string]string{
	fields := map[string]string{}
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
	return fields
}


func (char *CharField)ParseTagString(tag string,fieldname string ,dest interface{}) error{

	/*
	   tag : tag value
	   field:field name of strcuct object
	   dest : field value
	   is_ok:field  Required
	 */
	errormessage := map[string]interface{}{}
	fields := ParseString(tag)
	var err error
    var dest_value string
	switch dest.(type) {
	case string:
		dest_value = dest.(string)
	default:
		panic("CharFiled Must be string")

	}
	dest_length := len(dest_value)
	errormessage[fieldname] = ""
	 if _,ok := fields["MaxLength"];ok{
		 MaxLegth,err1 := strconv.Atoi(fields["MaxLength"])
		 if err1 != nil{
			 errormessage[fieldname] = errormessage[fieldname].(string) + "MaxLength can not be " + fields["MaxLength"]
			 err = err1
			 char.HasError = true
			 char.ErrorMessage = errormessage
			 return err
		 }
		 if MaxLegth > 0 && dest_length > MaxLegth{
			 errormessage[fieldname] = errormessage[fieldname].(string) + "Max-Legth is Over " + fields["MaxLength"]
			 err = OverMaxLengthError
			 char.ErrorMessage = errormessage
			 char.HasError = true
			 return err
		 }
	 }

	 if _,ok := fields["MinLength"];ok{
		 MinLegth,err2 := strconv.Atoi(fields["MinLength"])
		 if err2 != nil{
			 errormessage[fieldname] = fieldname + " error,MinLength can not be " + fields["MinLength"]
			 err = err2
			 char.HasError = true
			 char.ErrorMessage = errormessage
			 return err
		 }

		 if MinLegth > 0 && dest_length  < MinLegth{
			 errormessage[fieldname] = errormessage[fieldname].(string) + ";Min-Legth is Less " + fields["MinLength"]
			 err = LessMinLengthError
			 char.HasError = true
			 char.ErrorMessage = errormessage
			 return err
		 }
	 }

    if fields["Required"] == "true" {
    	char.Required = true
    	if dest_length == 0 {
			err = RequiredError
			errormessage[fieldname] = errormessage[fieldname].(string) + ";field required"
			char.HasError = true
			char.ErrorMessage = errormessage
			return err
		}

	}
	char.ErrorMessage = errormessage
	return err
}



func (char *CharField) HasErrors() bool {
	 return char.HasError != true
}


func (char *CharField) Errors() map[string]interface{}{
	return char.ErrorMessage
}




func (intfield *IntegerField)ParseTagString(tag string,fieldname string ,dest interface{}) error{
	errormessage := map[string]interface{}{}
	fields := ParseString(tag)
	var err error
	var dest_value int64
	MaxValue,err1 := strconv.Atoi(fields["MaxValue"])
	MinValue,err2 := strconv.Atoi(fields["MinValue"])
	errormessage[fieldname] = ""

	switch dest.(type) {
	case int8,int16,int32,int64,int:
		dest_value = dest.(int64)
	default:
		panic("IntegerField must be integer")
	}

	if err1 != nil {
		err = err1
		intfield.HasError = true
		errormessage[fieldname] = errormessage[fieldname].(string) + "MaxValue can not be " + fields["MaxValue"]
		intfield.ErrorMessage = errormessage
		return err
	}

	if dest_value > int64(MaxValue) {
		err = OverMaxValueError
		intfield.HasError = true
		errormessage[fieldname] = errormessage[fieldname].(string) + ";value can not be greater than " + fields["MaxValue"]
		intfield.ErrorMessage = errormessage
		return err
	}

	if err2 != nil {
		err = err2
		intfield.HasError = true
		errormessage[fieldname] = errormessage[fieldname].(string) + "MinValue can not be " + fields["MinValue"]
		intfield.ErrorMessage = errormessage
		return err
	}

	if dest_value < int64(MinValue) {
		err = LessMinValueError
		intfield.HasError = true
		errormessage[fieldname] = errormessage[fieldname].(string) + ";value can not be less than " + fields["MinValue"]
		intfield.ErrorMessage = errormessage
		return err
	}

	if fields["Required"] == "true" {
		intfield.Required = true
		if dest_value == 0 {
			err = RequiredError
			errormessage[fieldname] = errormessage[fieldname].(string) + ";field required"
			intfield.ErrorMessage = errormessage
			intfield.HasError = true
			return err
		}
	}
	intfield.ErrorMessage = errormessage
	return err
	}



func (intfield *IntegerField)  HasErrors() bool {
	return intfield.HasError != true
}



func ( intfield *IntegerField) Errors() map[string]interface{}{
	return intfield.ErrorMessage
}



func (email *EmailField) HasErrors() bool {
	return email.HasError != true
}

func (email *EmailField) Errors() map[string]interface{}{
	return email.ErrorMessage
}


func (email *EmailField) ParseTagString(tag string,fieldname string ,dest interface{}) error {
	errormessage := map[string]interface{}{}
	fields := ParseString(tag)
	var err error
	var dest_value string
	switch dest.(type) {
	case string:
		dest_value = dest.(string)
	default:
		panic("EmailFiled Must be string")

	}
    emailPattern := regexp.MustCompile(`^[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[a-zA-Z0-9](?:[\w-]*[\w])?$`)
    if !emailPattern.MatchString(dest_value){
		errormessage[fieldname] = "Email Format Error"
		err = EmailFormatError
		email.HasError = true
		email.ErrorMessage = errormessage
		return err
	}
	dest_length := len(dest_value)
	errormessage[fieldname] = ""
	if _,ok := fields["MaxLength"];ok{
		MaxLegth,err1 := strconv.Atoi(fields["MaxLength"])
		if err1 != nil{
			errormessage[fieldname] = errormessage[fieldname].(string) + "MaxLength can not be " + fields["MaxLength"]
			err = err1
			email.HasError = true
			email.ErrorMessage = errormessage
			return err
		}
		if MaxLegth > 0 && dest_length > MaxLegth{
			errormessage[fieldname] = errormessage[fieldname].(string) + "Max-Legth is Over " + fields["MaxLength"]
			err = OverMaxLengthError
			email.HasError = true
			email.ErrorMessage = errormessage
			return err
		}
	}

	if _,ok := fields["MinLength"];ok{
		MinLegth,err2 := strconv.Atoi(fields["MinLength"])
		if err2 != nil{
			errormessage[fieldname] = fieldname + " error,MinLength can not be " + fields["MinLength"]
			err = err2
			email.HasError = true
			email.ErrorMessage = errormessage
			return err
		}

		if MinLegth > 0 && dest_length  < MinLegth{
			errormessage[fieldname] = errormessage[fieldname].(string) + ";Min-Legth is Less " + fields["MinLength"]
			err = LessMinLengthError
			email.HasError = true
			email.ErrorMessage = errormessage
			return err
		}
	}

	if fields["Required"] == "true" {
		email.Required = true
		if dest_length == 0 {
			err = RequiredError
			errormessage[fieldname] = errormessage[fieldname].(string) + ";field required"
			email.HasError = true
			email.ErrorMessage = errormessage
			return err
		}

	}
	email.ErrorMessage = errormessage
	return err
}



func (urlfield *URLField) HasErrors() bool {
	return urlfield.HasError != true
}

func (urlfield *URLField) Errors() map[string]interface{}{
	return urlfield.ErrorMessage
}

func (urlfield *URLField) ParseTagString(tag string,fieldname string ,dest interface{}) error {
	errormessage := map[string]interface{}{}
	fields := ParseString(tag)
	var err error
	var dest_value string
	switch dest.(type) {
	case string:
		dest_value = dest.(string)
	default:
		panic("CharFiled Must be string")

	}
	Url ,_ := url.Parse(dest_value)
	if len(Url.Host) == 0{
		errormessage[fieldname] = "Enter a valid URL."
		err = URLFormatError
		urlfield.HasError = true
		urlfield.ErrorMessage = errormessage
		return err

	}
	dest_length := len(dest_value)
	errormessage[fieldname] = ""
	if _,ok := fields["MaxLength"];ok{
		MaxLegth,err1 := strconv.Atoi(fields["MaxLength"])
		if err1 != nil{
			errormessage[fieldname] = errormessage[fieldname].(string) + "MaxLength can not be " + fields["MaxLength"]
			err = err1
			urlfield.HasError = true
			urlfield.ErrorMessage = errormessage
			return err
		}
		if MaxLegth > 0 && dest_length > MaxLegth{
			errormessage[fieldname] = errormessage[fieldname].(string) + "Max-Legth is Over " + fields["MaxLength"]
			err = OverMaxLengthError
			urlfield.ErrorMessage = errormessage
			urlfield.HasError = true
			return err
		}
	}

	if _,ok := fields["MinLength"];ok{
		MinLegth,err2 := strconv.Atoi(fields["MinLength"])
		if err2 != nil{
			errormessage[fieldname] = fieldname + " error,MinLength can not be " + fields["MinLength"]
			err = err2
			urlfield.HasError = true
			urlfield.ErrorMessage = errormessage
			return err
		}

		if MinLegth > 0 && dest_length  < MinLegth{
			errormessage[fieldname] = errormessage[fieldname].(string) + ";Min-Legth is Less " + fields["MinLength"]
			err = LessMinLengthError
			urlfield.HasError = true
			urlfield.ErrorMessage = errormessage
			return err
		}
	}

	if fields["Required"] == "true" {
		urlfield.Required = true
		if dest_length == 0 {
			err = RequiredError
			errormessage[fieldname] = errormessage[fieldname].(string) + ";field required"
			urlfield.HasError = true
			urlfield.ErrorMessage = errormessage
			return err
		}

	}
	urlfield.ErrorMessage = errormessage
	return err


}