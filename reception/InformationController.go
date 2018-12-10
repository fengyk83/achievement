package reception

import (
	"github.com/astaxie/beego"
)

type InformationController struct {
	beego.Controller
}

func (c *InformationController) Index() {
	c.TplName = "reception/information.html"
	//c.TplName = "reception/index.html"
}
