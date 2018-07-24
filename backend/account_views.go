package backend

import (
	"omega/auth"
	"encoding/json"
	"omega/common"
	"github.com/astaxie/beego/orm"
	"fmt"
	"net/smtp"
	"omega/conf"
	"time"
	"strconv"
)



func (request *Requester) AuthCode() {
	idkey,stream := common.CodeCaptchaCreate()
	request.Ctx.ResponseWriter.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	request.Ctx.ResponseWriter.Header().Set("X-Code-Session",idkey)
	//body := map[string]interface{}{"code": 1, "data": stream,  "msg": "success"}
	//json.NewEncoder(request.Ctx.ResponseWriter).Encode(stream)
	fmt.Println(request.Ctx.ResponseWriter.Header())
	request.Ctx.WriteString(stream)
}


func (request *Requester) Login(){

	form := UserLoginForm{}
	//fmt.Println(request.Ctx.Input.RequestBody)
    json.Unmarshal(request.Ctx.Input.RequestBody,&form)
	username := form.UserName
	password := form.Password

	user := auth.Authenticate(username,password)
	if user != nil {
		if !user.IsActive{

			request.CommonResponse(10103,"")
			return
		}

		if !user.IsStaff{
			request.CommonResponse(10111,"")
			return
		}

		if !user.IsAdmin{
			request.CommonResponse(10117,"")
			return
		}
		//request.Login()
		request.AuthLogin(user)
		data := map[string]string{
			"token": request.Session.SessionKey,
		}
		//fmt.Println(request.Session)
		//fmt.Println(request.User)
		request.CommonResponse(0,data)
		return
		}else{
			request.CommonResponse(10112,"")
			return
	}
	}



func (request *Requester) Logout(){
	request.Session.Flush()
	request.Ctx.ResponseWriter.Header().Del("X-Token-Id")
	request.CommonResponse(0,"")
}


func (request *Requester) SendEmail(){
	//  send code to backend-manager
	form := EmailForm{}
	json.Unmarshal(request.Ctx.Input.RequestBody,&form)
	email := form.Email
	user := auth.User{Email:email}
	err := orm.NewOrm().Read(&user,"Email")
	if err == nil{
		if !user.IsActive{
			request.CommonResponse(10103,"")
			return
		}
		if !user.IsStaff{
			request.CommonResponse(10005,"")
			return
		}
		if !user.IsAdmin{
			request.CommonResponse(10007,"")
			return
		}
        code := auth.GetRandomString(6)
        text := code
		Auth := smtp.PlainAuth("", conf.EMAIL_USERNAME, conf.EMAIL_PASSWORD, conf.EMAIL_HOST)
		smtp.SendMail(conf.EMAIL_HOST,Auth,
			conf.EMAIL_USERNAME,[]string{"xxxx"},[]byte(text))
		request.Session.CycleKey()
		request.Session.SessionCache["email"] = email
		request.Session.SessionCache["email_code"] = code
		request.Session.SessionCache["email_expires"] = strconv.FormatInt(time.Now().Unix() + 30 * 60,10)
	}else{
		request.CommonResponse(10107,"")
		return
	}

}




func (request *Requester) EmailVerification(){
    form := EmailVerifyForm{}
    json.Unmarshal(request.Ctx.Input.RequestBody,&form)
    email , code := form.Email ,form.Code
	if email != request.Session.SessionCache["email"]{
		request.Session.SessionCache["email_code"] = ""
		request.CommonResponse(10001,"")
		return
	}

	if code != request.Session.SessionCache["email_code"]{
		request.Session.SessionCache["email_code"] = ""
		request.CommonResponse(10114,"")
		return
	}else{
		now := time.Now().Unix()
		expires,_ := strconv.Atoi(request.Session.SessionCache["email_expires"])
		if now > int64(expires){
			request.Session.SessionCache["email_code"] = ""
			request.CommonResponse(10114,"")
			return
		}
	}


	request.Session.SessionCache["email_verification"] = "1"
	request.Session.SessionCache["email_code"] = ""
	request.CommonResponse(0,"")
	return
}



func (request *Requester) PasswordRecover() {
	// forget password and recover
	if request.Session.SessionCache["email_verification"] != "1"{
		request.CommonResponse(10002,"")
		return
	}
	now := time.Now().Unix()
	expires,_ := strconv.Atoi(request.Session.SessionCache["email_expires"])
	if now > int64(expires){
		request.CommonResponse(10003,"")
		return
	}

	email := request.Session.SessionCache["email"]
	form := PasswordRecoverForm{}
	json.Unmarshal(request.Ctx.Input.RequestBody,&form)
	password1 , password2 := form.Password1 , form.Password2

	if password1 != password2{
		request.CommonResponse(10109,"")
		return
	}
	if len(password2) < 8{
		request.CommonResponse(10116,"")
		return
	}
	o := orm.NewOrm()
	user := auth.User{Email:email}
	err := o.Read(&user,"Email")
	if err == nil{
		user.SetPassword(password1)
		_,err1 := o.Update(&user)
		if err1 == nil{
			request.CommonResponse(0,"")
			return
		}else{
			request.CommonResponse(10120,"")
			return
		}
	}else{
		request.CommonResponse(10107,"")
		return
	}

}




func (request *Requester) UserPasswordChange() {
	// change password
	form := UserPasswordChangeForm{}
	json.Unmarshal(request.Ctx.Input.RequestBody,&form)
	new_password ,old_password := form.NewPassword,form.OldPassword

	if new_password == old_password{
		request.CommonResponse(10109,"")
		return
	}
	if len(new_password) < 8{
		request.CommonResponse(10116,"")
		return
	}

    user := request.User
    if !user.CheckPassword(old_password){
		request.CommonResponse(10107,"")
		return
	}
	o := orm.NewOrm()
	user.SetPassword(new_password)
	_,err := o.Update(&user)
	if err == nil{
		request.CommonResponse(0,"")
		return
	}else{
		request.CommonResponse(10120,"")
		return
	}
}


func (request *Requester)UserInfo(){
	user := request.User
	if request.Ctx.Request.Method == "GET"{

		result := map[string]interface{}{
			"email":user.Email,
			"username":user.UserName,
			"firstName":user.FirstName,
			"lastname":user.LastName,
			"phone":user.Phone,
		}
		request.CommonResponse(0,result)
		return
	}else{
        form := UserInfoCHangeForm{}
        err := json.Unmarshal(request.Ctx.Input.RequestBody,&form)
        if err == nil{
			user.Phone = form.Phone
			user.FirstName = form.FirstName
			user.LastName = form.LastName
			orm.NewOrm().Update(&user)
			request.CommonResponse(0,"")
			return
		}else{
			request.CommonResponse(10001,"")
			return
		}

	}
}




func (request *Requester) AddUser(){

	form := UserAddForm{}
	json.Unmarshal(request.Ctx.Input.RequestBody,&form)
	// TODO re-test
    fields := common.StructToMap(form)
    username := form.UserName
    email := form.Email
    table_name := new(auth.User)
    cond := orm.NewCondition().Or("user_name",username).Or("email",email)
    num,_ := orm.NewOrm().QueryTable(table_name).SetCond(cond).Count()

    if num > 0{
		request.CommonResponse(10102,"")
		return
	}
	user,err := auth.CreateUser(fields)
	if err == nil{
		result := map[string]interface{}{
			"Id":user.Id,
		}
		request.CommonResponse(0,result)
		return
	}else{
		request.CommonResponse(10120,"")
		return
	}
}