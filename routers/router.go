package routers

import (
	"achievement/reception"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/login", &reception.LoginContorller{},"Get:Login")
	beego.Router("/login/judge", &reception.LoginContorller{})
	//beego.Router("/login/verilfication", &reception.VerificationController{},"*:Index")

}
