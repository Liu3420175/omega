package routers

import (
	"omega/controllers"
	"github.com/astaxie/beego"
	"omega/backend"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/backend/account/login/",&backend.Requester{},"post:Login")
	beego.Router("/backend/account/add/",&backend.Requester{},"post:AddUser")
    
}
