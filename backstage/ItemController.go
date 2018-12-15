package backstage

import "github.com/astaxie/beego"

type ItemController struct {
	beego.Controller
}

//展示学期
//@router /achievement/showitem
func (this *ItemController)ShowItem()  {
	this.TplName=""
}


