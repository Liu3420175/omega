package common

import "encoding/json"

func JsonToObject(data string) (result interface{}){
	json.Unmarshal([]byte(data),&result)
	return result
}