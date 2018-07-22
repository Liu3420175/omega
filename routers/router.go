package routers

import (
	"omega/controllers"
	"github.com/astaxie/beego"
	"omega/backend"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/backend/account/login/",&backend.Requester{},"*:Login")
	beego.Router("/backend/account/add/",&backend.Requester{},"*:AddUser")


    // customer
    beego.Router("/backend/customer/list/",&backend.Requester{},"*:CustomerAccountList")
    
}
