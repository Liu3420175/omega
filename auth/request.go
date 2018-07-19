package auth

import (
	"github.com/astaxie/beego"
	"../sessions"
	"../common"
)

type Requester struct {
	// TODO 我们也可以不用他的，自己jiyu http.Request封装一个
	beego.Controller
	user    User
	session *sessions.SessionStore
}




func (request *Requester)CommonResponse(code int,r interface{}){
	result := map[string]interface{}{
		"code":code,
		"msg":common.CODES[code],
		"result":r,
	}
	//data ,_ := json.Marshal(result)
	data := result
	request.Data["json"] = data
	request.ServeJSON()
	return
}



func LoginRequired(request *Requester) {
	if request.session.SessionKey == ""{
		request.CommonResponse(10004,"")
		return
	}
	user,err := GetUser(request)
	if err == nil{
		if CompareUser(user,&User{}) {
			request.CommonResponse(10003,"")
			return
		}

		if !user.IsActive{
			request.CommonResponse(10103,"")
			return
		}

		if !user.IsStaff{
			request.CommonResponse(10005,"")
		}
	}

}


func AdminLoginRequired(request *Requester) {
	if request.session.SessionKey == ""{
		request.CommonResponse(10004,"")
		return
	}
	user,err := GetUser(request)
	if err == nil{
		if CompareUser(user,&User{}) {
			request.CommonResponse(10003,"")
			return
		}

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
	}

}



func RequireHttpMethods(method_list [...]string) {
	
}




func (request *Requester)Prepare() {

}