package auth

import (
	"github.com/astaxie/beego"
	"../sessions"
	)

type Requester struct {
	// TODO 我们也可以不用他的，自己jiyu http.Request封装一个
	beego.Controller
	user    User
	session *sessions.SessionStore
}



func (request *Requester)Prepare() {

}