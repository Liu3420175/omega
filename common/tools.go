package common

import (
	"encoding/json"
	"reflect"
)

func JsonToObject(data string) (result interface{}){
	json.Unmarshal([]byte(data),&result)
	return result
}


func StructToMap(object interface{}) map[string]interface{}{
	v := reflect.ValueOf(object)
	t := v.Elem().Type()

	if v.Kind() == reflect.Struct {
		var result map[string]interface{}
		for i := 0;i < v.NumField();i++{
			key := t.Field(i).Name
			value := v.Field(i)
			result[key] = value
		}
		return result
	}else{
		panic("Not Support to Map")
	}

}