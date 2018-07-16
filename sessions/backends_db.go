package sessions

import (
	"github.com/astaxie/beego/orm"

	"../auth"
	"encoding/json"
	"encoding/base64"
	"strings"
	"time"
)

type SessionStore struct {
	SessionKey     string
	Accessed       bool
	Modified       bool
	SessionCache   map[string]string
}



func (store *SessionStore) _Session() map[string]string{
    store.Accessed = true
    if len(store.SessionCache) != 0{
    	return store.SessionCache
	}

    if len(store.SessionKey) == 0 {
		store.SessionCache =  map[string]string{}
	}else{
		store.SessionCache = store.Load()
	}
	return store.SessionCache
}




func (store *SessionStore) Decode(sessiondata string) map[string]string {
	encoded_data,_ := base64.StdEncoding.DecodeString(sessiondata)
	encoded_data_str := string(encoded_data)
	hash_serialized := strings.Split(encoded_data_str,":")
	haser := hash_serialized[0] + ":"
	serialized_str := strings.Trim(encoded_data_str,haser)
	//TODO 有待优化需要验证
	var result map[string]string
	json.Unmarshal([]byte(serialized_str),&result)
    return  result
}


func (store *SessionStore) Hash(value string) string {
	key_salt := "omega.sessions"
    return auth.SaltedHhmac(key_salt,value,"")
}


//Returns the given session dictionary serialized and encoded as a string.
func (store *SessionStore) Encode(session_map map[string]string) string {
	serialized,_ := json.Marshal(session_map)
	serialized_str := string(serialized)
	serialized_hash := store.Hash(serialized_str)
	s := serialized_hash + ":" + serialized_str
	hashs := base64.StdEncoding.EncodeToString([]byte(s))
	return hashs
}


func (store *SessionStore) Get(key string) string{
	data := store._Session()
	return data[key]
}

func (store *SessionStore) Pop(key string) string{
	data := store._Session()
	value,ok := data[key]
	store.Modified = store.Modified ||  ok
	delete(data,key)
	return value
}


func (store *SessionStore)SetDefault(key,value string) string {
	data := store._Session()
	_,ok := data[key]
	if ok {
		return data[key]
	}else{
		store.Modified = true
		store.SessionCache[key] = value
		return value
	}
}


func (store * SessionStore) Update(dict map[string]string) {
	for key,value := range dict {
		store.SessionCache[key] = value
	}
	store.Modified = true
}


func (store *SessionStore) HasKey(key string) bool {
	data := store._Session()
	_,ok := data[key]
	return  ok
}

func (store *SessionStore) Keys() (result []string) {
	data := store._Session()
	for key,_ := range data {
		result = append(result,key)
	}
	return
}


func (store *SessionStore) Values() (result []string) {
	data := store._Session()
	for _,value := range data {
		result = append(result,value)
	}
	return
}


func (store *SessionStore) Clear(){
	store.Modified = true
	store.Accessed = true
	store.SessionCache = map[string]string{}
}



func (store *SessionStore) IsEmpty() bool {
	return len(store.SessionCache) != 0 && len(store.SessionKey) != 0
}



func (store *SessionStore) Load() map[string] string{
	o := orm.NewOrm()
	session := Session{SessionKey:store.SessionKey}
	err := o.Read(&session)
	if err == nil{
        if session.ExpireDate.Before(time.Now()){
        	// exprired
        	return map[string]string{}
		}else{
			return store.Decode(session.SessionData)
		}
	}else{
		store.SessionKey = ""
		return map[string]string{}
	}
}