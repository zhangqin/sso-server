package routers

import (
	"github.com/astaxie/beego"
	"github.com/zhangqin/sso-server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sso/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/sso/ticket", &controllers.LoginController{}, "*:Ticket")
	beego.Router("/sso/userinfo", &controllers.LoginController{}, "*:GetUserInfoByTicket")
	beego.Router("/sso/index", &controllers.LoginController{}, "*:Index")
	beego.Router("/sso/logout", &controllers.LoginController{}, "*:Logout")
	beego.Router("/sso/geturl", &controllers.LoginController{}, "*:GetURL")
}
