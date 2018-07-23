package backend

import (
	"omega/auth"
	"encoding/json"
	"omega/common"
	"github.com/astaxie/beego/orm"
)



func (request *Requester) AuthCode() {
	stream := common.CodeCaptchaCreate()
	request.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 1, "data": stream,  "msg": "success"}
	json.NewEncoder(request.Ctx.ResponseWriter).Encode(body)
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