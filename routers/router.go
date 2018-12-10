package routers

import (
	"achievement/backstage"
	"achievement/reception"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	/*前台路由*/
	beego.InsertFilter("/admin/*",beego.BeforeRouter,func(ctx *context.Context) {   //前台路由过滤
		_, ok := ctx.Input.Session("account").(string)
		if !ok  {
			ctx.Redirect(302, "/reception/login")
		}
	})
	ns := beego.NewNamespace("reception",
				beego.NSInclude(
					&reception.LoginContorller{},
					&reception.InformationController{},
				),
			)
	beego.AddNamespace(ns)


	/*后台路由*/
	beego.Router("/back/login",&backstage.LoginController{},"Get:Login")
	beego.Router("/back/index",&backstage.IndexController{},"Get:Index")
}
