package auth

import "github.com/astaxie/beego/orm"

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