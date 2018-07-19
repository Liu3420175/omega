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

func ProcessRequest(request *Requester){
	// the first oparea
	heard := request.Ctx.Request.Header
	token ,ok := heard["x-token-id"]
	if ok {
         request.session = &sessions.SessionStore{SessionKey:token[0]}
	}else{
		panic("") // TODO 
	}

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



func PermissionRequired(request *Requester,perm interface{}){
	perms := []string{}
	switch perm.(type) {
	case string:
		perms = append(perms,perm.(string))

	case []string:
		perms = perm.([]string)

	default:
		perms = []string{}
	}
    for _,v := range perms {
    	 if request.user.GetPerm(v){

		 }
	}
	request.CommonResponse(10007,"")
    return
}



func RequireHttpMethods(request *Requester, method_list []string) {
	method := request.Ctx.Request.Method
	for _,v :=range method_list {
		if method == v {

		}
	}
	request.CommonResponse(10013,"")
	return

}







func (request *Requester)Prepare() {

}