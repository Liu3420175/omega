package sessions

import (
	"github.com/astaxie/beego/orm"

	"../auth"
	"encoding/json"
	"encoding/base64"
)

type SessionStore struct {
	SessionKey     string
	Accessed       bool
	Modified       bool
}



func (store *SessionStore) _Session() map[string]string{
    store.Accessed = true
    return nil
}



func (store *SessionStore) Decode(sessiondata string) map[string]string {

    return  nil
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

func (store *SessionStore) Load() map[string] string{
	o := orm.NewOrm()
	session := Session{SessionKey:store.SessionKey}
	err := o.Read(&session)
	if err == nil{

	}
	return nil
}