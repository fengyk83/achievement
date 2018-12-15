package backstage

import "github.com/astaxie/beego"

type IdiomController struct {
	beego.Controller
}

//展示成员
//@router /achievement/showidiom [get]
func (this *IdiomController)ShowIdiom()  {
	this.TplName = ""
}
