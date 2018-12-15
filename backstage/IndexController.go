package backstage

import (
	"achievement/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

//后台页面展示
//@router /achievement/index [get]
func (this *IndexController)Index() {
	this.TplName = "backstage/index.html"
}

//获取所有考试信息
//@router /achievement/exam [get]
func (this *IndexController)GetInformation() {
	exam := models.NewExam().GetAllExam();
	this.Data["exam"] = exam
	this.TplName = "backstage/index.html"
}

//添加考试信息
//@router /achievement/addexam [get]
func (this *IndexController)AddExam()  {
	this.TplName = "backstage/index.html"
}



//添加考试信息
//@router /S [get]
func (this *IndexController)S()  {

	this.TplName = "backstage/welcome.html"
}