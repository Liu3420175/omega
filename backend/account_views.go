package backend

import (
	"omega/auth"
	"encoding/json"
	"fmt"
	"omega/common"
)
type Requester struct {
	auth.Requester
}

func (request *Requester) Login(){

	form := UserLoginForm{}
	//fmt.Println(request.Ctx.Input.RequestBody)
    json.Unmarshal(request.Ctx.Input.RequestBody,&form)
	username := form.UserName
	password := form.Password
    fmt.Println(username,"wwwww")
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

	user,err := auth.CreateUser(fields)
	if err == nil{
		request.CommonResponse(0,user)
		return
	}else{
		request.CommonResponse(10120,"")
		return
	}
}