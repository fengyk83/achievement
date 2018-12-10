package reception

import (
	"github.com/astaxie/beego"
)

type InformationController struct {
	beego.Controller
}

func (c *InformationController) Index() {
	c.TplName = "backstage/index.html"
	//c.TplName = "reception/index.html"
}
