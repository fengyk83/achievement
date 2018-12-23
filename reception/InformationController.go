package reception

import (
	"achievement/models"
	"github.com/astaxie/beego"
)

type InformationController struct {
	beego.Controller
}

/*展示学生个人信息,学生的成绩*/
//@router /admin/information [get]
func (this *InformationController) Index() {
	// account,ok := this.GetSession("account").(string)
	//if !ok {
	//	return
	//}
	account := "201301002"
	userInformation := 	models.NewStudent().GetInformation(account)
	exam := models.NewExam().GetExam(userInformation.Gradeid,userInformation.Clazzid)
	this.Data["exam"] = exam
	this.Data["user"] = userInformation
	this.TplName = "reception/information.html";
}
