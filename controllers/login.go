package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/sluu99/uuid"
	"github.com/zhangqin/sso-server/models"
)

var (
	p = fmt.Println
)

type LoginController struct {
	beego.Controller
}

func getRedis() (redis.Conn, error) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	//defer c.Close()
	return c, err
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

		c, _ := getRedis()
		c.Do("SET", ticket, user.Id)

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
	ticket := this.GetSession("ticket")
	if ticket == nil {
		ticket = ""
	}
	this.Ctx.WriteString("ticket_callback(\"" + ticket.(string) + "\")")
}

func (this *LoginController) GetUserInfoByTicket() {
	ticket := this.GetString("ticket")
	p(ticket)
	//user_id := this.GetSession(ticket).(int)
	c, _ := getRedis()
	user_id, _ := redis.Int(c.Do("GET", ticket))
	p(user_id)
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
	c, _ := getRedis()
	c.Do("DEL", this.GetSession("ticket"))
	this.DelSession("ticket")
	next := this.GetString("next")
	if next != "" {
		this.Redirect(next, 302)
	} else {
		this.Redirect("/sso/login", 302)
	}
}

func (this *LoginController) Index() {
	this.Ctx.WriteString("My ticket is: " + this.GetSession("ticket").(string))
}
