package auth

import (
	"github.com/astaxie/beego/orm"

	"strconv"
	"github.com/pkg/errors"
)


var (

	UserDoesNotExist = errors.New("UserDoesNotExist")
)


const (
	SESSION_KEY = "_auth_user_id"
	BACKEND_SESSION_KEY = "_auth_user_backend"
	HASH_SESSION_KEY = "_auth_user_hash"
)



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
	value , ok := request.session.SessionCache[SESSION_KEY]
	if ok{
		user_id,_ := strconv.Atoi(value)
		if int64(user_id) != user.Id || (len(session_auth_hash) > 0 &&
			!ConstantTtimeCompare(request.session.SessionCache[HASH_SESSION_KEY],session_auth_hash)){
				request.session.Flush()
		}

	}else{
		request.session.CycleKey()
	}
	request.session.SessionCache[SESSION_KEY] = string(user.Id)
	request.session.SessionCache[HASH_SESSION_KEY] = session_auth_hash
	request.session.SessionCache[BACKEND_SESSION_KEY] = "omega.sessions.session"
	request.user = *user

	// TODO set csrf
}


func Logout(request *Requester){
	/*
	logout
	 */
	 //user := request.user
     request.session.Flush()
     request.user = User{}

}


func GetUser(request *Requester) (*User,error){

	value ,ok := request.session.SessionCache[SESSION_KEY]
	if ok{
		user_id,_ := strconv.Atoi(value)
		o := orm.NewOrm()
		user := User{Id:int64(user_id)}
		err := o.Read(&user)
		if err == nil{
			return &user,nil
		}
	}
	return &User{},UserDoesNotExist

}




func UpdateSessionAuthHash(request *Requester,user *User ){
	/*
	Updating a user's password logs out all sessions for the user.

    This function takes the current request and the updated user object from
    which the new session hash will be derived and updates the session hash
    appropriately to prevent a password change from logging out the session
    from which the password was changed.

	 */
      request.session.CycleKey()
      if CompareUser(&request.user,user){
          request.session.SessionCache[HASH_SESSION_KEY] = user.GetSessionAuthHash()
	  }

}
//func get_group_permissions()


