package form

import (
	"reflect"
)

type Form struct {
	HasError         bool
	ErrorMessage    map[string]interface{}
}




func (form *Form)NewCharField(tags string,fieldname string ,dest interface{}) (err error,valid Validator){
	char := new(CharField)
	err,_ = char.ParseTagString(tags,fieldname,dest)
	form.HasError = form.HasError || char.HasError
	if form.ErrorMessage == nil {
		form.ErrorMessage =  map[string]interface{}{}
	}
	form.ErrorMessage[fieldname] = char.ErrorMessage[fieldname]
	valid = char
	return err,valid
}


func (form *Form)NewIntegerField (tags string,fieldname string ,dest interface{}) (err error,valid Validator){
	integer := new(IntegerField)
	err,_ = integer.ParseTagString(tags,fieldname,dest)
	form.HasError = form.HasError || integer.HasError
	if form.ErrorMessage == nil {
		form.ErrorMessage =  map[string]interface{}{}
	}
	form.ErrorMessage[fieldname] = integer.ErrorMessage[fieldname]
	valid = integer
	return err,valid
}


func (form *Form) NewEmailField (tags string,fieldname string ,dest interface{}) (err error,valid Validator) {
	emailfield := new(EmailField)
	err,_ = emailfield.ParseTagString(tags,fieldname,dest)
	form.HasError = form.HasError || emailfield.HasError
	if form.ErrorMessage == nil {
		form.ErrorMessage =  map[string]interface{}{}
	}
	form.ErrorMessage[fieldname] = emailfield.ErrorMessage[fieldname]
	valid = emailfield
	return err,valid

}


func (form *Form)NewURLField(tags string,fieldname string ,dest interface{}) (err error,valid Validator){
	urlfield := new(URLField)
	err,_ = urlfield.ParseTagString(tags,fieldname,dest)
	form.HasError = form.HasError || urlfield.HasError
	if form.ErrorMessage == nil {
		form.ErrorMessage =  map[string]interface{}{}
	}
	form.ErrorMessage[fieldname] = urlfield.ErrorMessage[fieldname]
	valid = urlfield
	return err,valid
}



func (form *Form)NewChoiceField(tags string,fieldname string ,dest interface{}) (err error,valid Validator) {
	choicefield := new(ChoiceField)
	err,_ = choicefield.ParseTagString(tags,fieldname,dest)
	form.HasError = form.HasError || choicefield.HasError
	if form.ErrorMessage == nil {
		form.ErrorMessage =  map[string]interface{}{}
	}
	form.ErrorMessage[fieldname] = choicefield.ErrorMessage[fieldname]
	valid = choicefield
	return err,valid
}



func (form *Form) Valid(object interface{}) (err error,valid Validator) {

	objT := reflect.TypeOf(object)
	objK := objT.Kind()
    objV := reflect.ValueOf(object)
    var errT error
	if objK != reflect.Struct{
		panic("Must be Struct")
	}
	for i := 0; i < objT.NumField();i++ {
		field := objT.Field(i)
		value := objV.Field(i).Interface()
		fieldname := field.Name
		tag := field.Tag.Get(TagName)
		err_,filetype,tags := ParseFormTagString(tag)
		if err_ != nil {
			err = err_
			form.ErrorMessage["error"] = err.Error()
			form.HasError = true
			return err,nil

		}
		switch filetype {
		case "CharField" :
              errT,valid = form.NewCharField(tags,fieldname,value)
              if errT != nil{
              	err = errT
			  }
		case "IntegerField":
			  errT,valid = form.NewIntegerField(tags,fieldname,value)
			  if errT != nil{
				err = errT
			  }
		case "EmailField":
			errT,valid = form.NewEmailField(tags,fieldname,value)
			if errT != nil{
				err = errT
			}
		case "URLField":
			errT,valid = form.NewURLField(tags,fieldname,value)
			if errT != nil{
				err = errT
			}
		case "ChoiceField":
			errT,valid = form.NewURLField(tags,fieldname,value)
			objE := objV.Elem()
			if errT != nil{
				err = errT
			}
			if objE.Field(i).CanSet(){
				// TODO
				//objE.Field(i).SetInt()
			}
		}

	}


	return err,valid
}

func init()  {

}