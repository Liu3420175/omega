package auth

import (
	"github.com/astaxie/beego"
	"omega/common"

)

type Requester struct {
	// TODO 我们也可以不用他的，自己jiyu http.Request封装一个
	beego.Controller
	User    User
	Session *SessionStore
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

func (request *Requester)ProcessRequest(){
	// the first oparea
	heard := request.Ctx.Request.Header
	token ,ok := heard["x-token-id"]
	if ok {
         request.Session = &SessionStore{SessionKey:token[0]}
	}else{
		panic("") // TODO 
	}

}

func (request *Requester)LoginRequired() {
	if request.Session.SessionKey == ""{
		request.CommonResponse(10004,"")
		return
	}
	user,err := request.GetUser()
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


func (request *Requester)AdminLoginRequired() {
	if request.Session.SessionKey == ""{
		request.CommonResponse(10004,"")
		return
	}
	user,err := request.GetUser()
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



func (request *Requester)PermissionRequired(perm interface{}){
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
    	 if request.User.GetPerm(v){

		 }
	}
	request.CommonResponse(10007,"")
    return
}



func (request *Requester)RequireHttpMethods( method_list []string) {
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