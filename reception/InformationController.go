package reception

import (
	"achievement/models"
	"fmt"
	"github.com/astaxie/beego"
)

type InformationController struct {
	beego.Controller
}

/*展示学生个人信息,学生的成绩*/
//@router /admins/information [get]
func (this *InformationController) Index() {
	// account,ok := this.GetSession("account").(string)
	//if !ok {
	//	return
	//}
	account := "201301002"
	userInformation := 	models.NewStudent().GetInformation(account)
	fmt.Println(userInformation.ClazzName)
	this.Data["user"] = userInformation
	this.TplName = "reception/information.html"
}
