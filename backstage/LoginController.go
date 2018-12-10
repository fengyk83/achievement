package backstage

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

//后台登陆
//@router /backlogin [get]
func (this *LoginController)Login()  {
	this.TplName = "backstage/login.html";
}

