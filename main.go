package main

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/zhangqin/sso-server/models"
	_ "github.com/zhangqin/sso-server/routers"
)

var loginFilter = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("ticket").(string)
	if !ok && !strings.HasPrefix(ctx.Request.RequestURI, "/sso/login") && !strings.HasPrefix(ctx.Request.RequestURI, "/sso/userinfo") && !strings.HasPrefix(ctx.Request.RequestURI, "/sso/ticket") {
		ctx.Redirect(302, "/sso/login")
	}
}

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, loginFilter)
	beego.Run()
}
