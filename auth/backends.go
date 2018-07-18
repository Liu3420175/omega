package auth

import (
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"

)


const (
	SESSION_KEY = "_auth_user_id"
	BACKEND_SESSION_KEY = "_auth_user_backend"
	HASH_SESSION_KEY = "_auth_user_hash"
)


type Requester struct {
	beego.Controller
	user    User
	session map[string]string
}


func user_can_authenticate(user User) bool{
	return user.IsActive == true
}

func Authenticate(username string ,password string) *User {
	/*
	If the given credentials are valid, return a User object.
	 */
	 o := orm.NewOrm()
	 user1 := User{Username:username}
	 err1 := o.Read(&user1)
	 if err1 == nil{
	 	if user1.CheckPassword(password) && user_can_authenticate(user1){
	 		return &user1
		}
	 }

	 user2 := User{Email:username}
	 err2 := o.Read(&user2)
	 if err2 == nil{
	 	if user2.CheckPassword(password) && user_can_authenticate(user2){
	 		return  &user2
		}
	 }
	 return nil
}





func get_user_permissions(user User) []*Permission {
	o := orm.NewOrm()
	var p []*Permission
	_,ok := o.QueryTable("auth_permission").Filter("Users__User__Id",user.Id).All(&p)
	if ok == nil{
		return p
	}
	return nil
}

func get_user_session

func Login(request *Requester,user *User){
    /*
    Login
     */

	session_auth_hash := ""
	if user == nil{
		user = &request.user
	}
	if user != nil {
		session_auth_hash = user.GetSessionAuthHash()
	}
	value , ok := request.session[SESSION_KEY]
	if ok{

	}
}

//func get_group_permissions()