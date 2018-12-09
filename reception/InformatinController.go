package reception

import (
	"github.com/astaxie/beego"
)

type InformationController struct {
	beego.Controller
}

func init() {
	//store := cache.NewMemoryCache()
	//cpt = captcha.NewWithFilter("/captcha/", store)
}
//
func (c *InformationController) Index() {
	c.TplName = "backstage/index.html"
}
