package reception

import (
	"github.com/astaxie/beego"
)

type InformationController struct {
	beego.Controller
}

/*展示学生个人信息*/
//@router /admin/information [get]
func (c *InformationController) Index() {
	c.TplName = "reception/information.html"
	//c.TplName = "reception/index.html"
}
