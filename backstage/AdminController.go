package backstage

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

//成员管理
//@router /achievement/showadmin
func (this *AdminController)ShowAdmin()  {
	this.TplName = "backstage/setMenager.html"
}
