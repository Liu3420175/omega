package common

import (
	"encoding/json"
	"reflect"
	"github.com/astaxie/beego/orm"
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



func  GetSearchResults(query orm.QuerySeter,fields []string,search_term string) orm.QuerySeter {
	/*
    数据库模糊查询
    */

	cond := orm.NewCondition()

    for _,field := range fields {

    	cond = cond.Or(field + "__icontains",search_term)
    	//cond = cond.Or(field + "__istartswith",search_term)
    	//cond = cond.Or(field + "__endswith",search_term)
    	//cond = cond.Or(field + "__iexact",search_term)

	}
	//fmt.Println(cond)
    query.SetCond(cond)
    return query

}

func StructToMap(object interface{}) map[string]string{
	// TODO 有待改进，可以用反射
	v := reflect.ValueOf(object)
	if v.Kind() == reflect.Struct {
		var result map[string]string
		json_str,_ := json.Marshal(object)
		json.Unmarshal(json_str,&result)
		return result
	} else{
		panic("Not Support to Map")
	}
}