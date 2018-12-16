package backstage

import (
	"achievement/models"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

//成员管理
//@router /achievement/showadmin [get]
func (this *AdminController)ShowAdmin()  {
	this.TplName = "backstage/setMenager.html"
}

//展示所有管理员
//@router /achievement/showmenager [get]
func (this *IdiomController)Idiom()  {
	user := models.NewUser().GetMenager();
	this.Data["json"] = map[string]interface{}{"name": 1, "message": user}
	this.ServeJSON()
	this.TplName = "backstage/setMenager.html"
}
