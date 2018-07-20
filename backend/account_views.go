package backend

import (
	"omega/auth"
	"encoding/json"
	"fmt"
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
    fmt.Println(user)
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
    fields := map[string]string{
    	"UserName":form.UserName,
    	"Password":form.Password,
    	"Email":form.Email,
		"FirstName":form.FirstName,
		"LastName" : form.LastName,
		"Phone":   form.Phone,
	}

	user,err := auth.CreateSuperuser(fields)
	if err == nil{
		request.CommonResponse(0,user)
		return
	}else{
		request.CommonResponse(10120,"")
		return
	}
}