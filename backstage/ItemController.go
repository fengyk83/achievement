package backstage

import (
	"achievement/models"
	"github.com/astaxie/beego"
)

type ItemController struct {
	beego.Controller
}

//展示学期
//@router /achievement/showitem [get]
func (this *ItemController)ShowItem()  {
	item := models.NewExam().GetAllExam();
	this.Data["item"] = item
	this.Data["number"] = len(item)
	this.TplName = "backstage/term.html"
}

//添加考试信息
//@router /achievement/addexam [post]
func (this *IndexController)AddExam()  {
	error := models.NewExam().AddExam(this.GetString("type"),this.GetString("time"))
	this.Data["error"] =error
	this.TplName = "backstage/term.html"
}

