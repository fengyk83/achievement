package reception

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
)

type LoginContorller struct {
	beego.Controller
}
var cpt *captcha.Captcha

func init()  {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

/*返回登录首页*/
func (c *LoginContorller)Login() {
	c.TplName = "reception/index.html"

}

/*登录验证*/
func (c *LoginContorller)Post() {
	valid := validation.Validation{};
	valid.MaxSize(c.Input().Get("school"),11,"school").Message("您输入的学号不正确")
	valid.AlphaNumeric(c.Input().Get("school"),"number").Message("请您输入的学号必须是数字字符")
	//valid.Match(c.Input().Get("password"),'',"password").Message("您输入的密码格式不正确")
	valid.Required(c.Input().Get("password"),"password").Message("请您输入密码")
	valid.Required(c.Input().Get("captcha"),"captcha").Message("请您输入验证码")
	if valid.HasErrors() {
		for _,err := range valid.Errors {
			beego.Alert(err.Key+"+++++++"+err.Message)
		}
	}
	c.Data["json"] = map[string]interface{}{"name": c.Input().Get("password"), "num": 1111}
	c.ServeJSON()
	return
}