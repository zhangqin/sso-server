package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sluu99/uuid"
	"github.com/zhangqin/sso-server/models"
)

var (
	p = fmt.Println
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	next := this.GetString("next")

	user := models.Users{}
	if r := user.CheckLogin(username, password); r == true {
		this.Data["msg"] = "login success"

		ticket := uuid.Rand().Hex()
		this.SetSession("ticket", ticket)
		this.SetSession(ticket, user.Id)

		if next != "" {
			this.Redirect(next, 302)
		} else {
			this.Redirect("/sso/index", 302)
		}

	} else {
		this.Data["msg"] = "login fail"
	}
	this.TplNames = "login.tpl"
}

func (this *LoginController) Ticket() {
	this.Ctx.WriteString("ticket_callback(\"" + this.GetSession("ticket").(string) + "\")")
}

func (this *LoginController) GetUserInfoByTicket() {
	ticket := this.GetSession("ticket")
	user_id := this.GetSession(ticket).(int)
	user := models.Users{}
	user.GetUserInfoById(user_id)
	info := map[string]interface{}{
		"username": user.Username,
		"phone":    user.Phone,
	}
	this.Data["json"] = &info
	this.ServeJson()
}

func (this *LoginController) GetURL() {
	next := this.GetString("next")
	this.Ctx.WriteString(next)
}
func (this *LoginController) Logout() {
	this.DelSession("ticket")
	this.Redirect("/sso/login", 302)
}

func (this *LoginController) Index() {
	this.Ctx.WriteString("My ticket is: " + this.GetSession("ticket").(string))
}
