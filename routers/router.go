package routers

import (
	"achievement/backstage"
	"achievement/reception"
	"github.com/astaxie/beego"
)

func init() {
	/*前台路由*/
    beego.Router("/login", &reception.LoginContorller{},"Get:Login")
	beego.Router("/login/judge", &reception.LoginContorller{})
	//beego.Router("/login/verilfication", &reception.VerificationController{},"*:Index")


	/*后台路由*/
	beego.Router("/back/login",&backstage.LoginController{},"Get:Login")

}
