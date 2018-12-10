package reception

import "github.com/astaxie/beego"

type ForgetController struct {
	beego.Controller
}

//@router /forget [get]
func (c *ForgetController) Index() {
	c.TplName = "reception/forget.html"
	//c.TplName = "reception/index.html"
}