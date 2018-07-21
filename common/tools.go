package common

import (
	"encoding/json"
	"reflect"
	"fmt"
)

func JsonToObject(data string) (result interface{}){
	json.Unmarshal([]byte(data),&result)
	return result
}


//func StructToMap(object interface{}) map[string]interface{}{
//	v := reflect.ValueOf(object)
//	t := v.Elem().Type()

//	if v.Kind() == reflect.Struct {
//     TODO 这种方法可以获取结构体的字段名及其值
//		var result map[string]interface{}
//		for i := 0;i < v.NumField();i++{
//			key := t.Field(i).Name
//			value := v.Field(i)
//			result[key] = value
//		}
//		return result
//	}else{
//		panic("Not Support to Map")
//	}

//}


func StructToMap(object interface{}) map[string]string{
	v := reflect.ValueOf(object)

	fmt.Println(v.Kind().String(),reflect.Struct)
	if v.Kind() == reflect.Struct {
		var result map[string]string
		json_str,_ := json.Marshal(object)
		json.Unmarshal(json_str,&result)
		return result
	} else{
		panic("Not Support to Map")
	}
}