package routers

import (
	"omega/controllers"
	"github.com/astaxie/beego"
	"omega/backend"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // account
    beego.Router("/backend/account/login/",&backend.Requester{},"*:Login")
	beego.Router("/backend/account/add/",&backend.Requester{},"*:AddUser")
    beego.Router("/backend/account/logout/",&backend.Requester{},"*:Logout")

    // customer
    beego.Router("/backend/customer/list/",&backend.Requester{},"*:CustomerAccountList")
	beego.Router("/backend/authcode/",&backend.Requester{},"*:AuthCode")
    
}
