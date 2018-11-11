package reception

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
	"regexp"
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
	valid.MaxSize(c.GetString("school"),11,"school").Message("您输入的学号不正确")
	valid.AlphaNumeric(c.GetString("school"),"number").Message("请您输入的学号必须是数字字符")
	valid.Match(c.GetString("password"),regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`),"password").Message("您输入的密码格式不正确")
	valid.Required(c.GetString("password"),"password").Message("请您输入密码")
	valid.Required(c.GetString("captcha"),"captcha").Message("请您输入验证码")
	//id,value = c.GetString(c.Input().Get("captcha")), c.GetString("captcha")
	captchaBoolean := cpt.Verify(c.GetString("captchaId"),c.GetString("captcha"));

	if valid.HasErrors() {
		for _,err := range valid.Errors {
			beego.Alert(err.Key+"+++++++"+err.Message)
		}
	}
	if captchaBoolean {

	}
	c.Data["json"] = map[string]interface{}{"name": c.Input().Get("password"), "num": 1111}
	c.ServeJSON()
	c.TplName = "reception/index.html"
}