package form

import "reflect"

type Form struct {
	HasError         bool
	ErrorMessage    map[string]interface{}
}




func (form *Form)NewCharField(tags string,fieldname string ,dest interface{}) (err error,valid Validator){
	char := new(CharField)
	err = char.ParseTagString(tags,fieldname,dest)
	form.HasError = char.HasError
	form.ErrorMessage[fieldname] = char.ErrorMessage[fieldname]
	valid = char
	return err,valid
}


func (form *Form)NewIntegerField (tags string,fieldname string ,dest interface{}) (err error,valid Validator){
	integer := new(IntegerField)
	err = integer.ParseTagString(tags,fieldname,dest)
	form.HasError = integer.HasError
	form.ErrorMessage[fieldname] = integer.ErrorMessage[fieldname]
	valid = integer
	return err,valid
}



func (form *Form) Valid(object interface{}) (err error,valid Validator) {

	objT := reflect.TypeOf(object)
	objK := objT.Kind()
    objV := reflect.ValueOf(&object)

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
              err,valid = form.NewCharField(tags,fieldname,value)
		case "IntegerField":
			  err,valid = form.NewIntegerField(tags,fieldname,value)
		}

	}


	return err,valid
}