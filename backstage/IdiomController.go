package backstage

import (
	"achievement/models"
	"github.com/astaxie/beego"
)

type IdiomController struct {
	beego.Controller
}

//展示成员
//@router /achievement/showidiom [get]
func (this *IdiomController)ShowIdiom()  {
	this.Data["grade"] = models.NewGrade().GetGrade()
	this.Data["class"] = models.NewClazz().GetClazz()
	this.TplName = "backstage/people.html"
}

//添加成员
//@router /achievement/addidiom [post]
func (this *IdiomController)AddIdiom()  {
	this.TplName = "backstage/people.html"
}
