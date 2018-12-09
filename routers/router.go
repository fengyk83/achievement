package routers

import (
	"achievements/backstage"
	"achievements/reception"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	/*前台路由*/
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok  {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterUser)
    beego.Router("/login", &reception.LoginContorller{},"Get:Login")
	beego.Router("/login/judge", &reception.LoginContorller{})
	beego.Router("/login/information", &reception.InformationController{},"*:Index")

	/*后台路由*/
	beego.Router("/back/login",&backstage.LoginController{},"Get:Login")
	beego.Router("/back/index",&backstage.IndexController{},"Get:Index")
}
