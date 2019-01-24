package reception

import (
	"achievement/models"
	"crypto/md5"
	"encoding/hex"
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
// @router /login [get]
func (c *LoginContorller)Login() {
	c.XSRFExpire = 7200
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "reception/index.html"
}

/*登录验证*/
//@router /login/judge [post]
func (c *LoginContorller)Post() {
	valid := validation.Validation{};
	valid.MaxSize(c.GetString("school"),11,"school").Message("您输入的学号不正确")
	valid.AlphaNumeric(c.GetString("school"),"number").Message("请您输入的学号必须是数字字符")
	valid.Match(c.GetString("password"),regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`),"password").Message("您输入的密码格式不正确")
	valid.Required(c.GetString("password"),"password").Message("请您输入密码")
	valid.Required(c.GetString("captcha"),"captcha").Message("请您输入验证码")
	captchaBoolean := cpt.Verify(c.GetString("captchaId"),c.GetString("captcha"));

	if valid.HasErrors() {
		for _,err := range valid.Errors {
			beego.Alert(err.Key+"+++++++"+err.Message)
		}
		return
	}
	h := md5.New()
	h.Write([]byte(c.GetString("password")))
	user := models.NewUser().ReceptionLoginJudge(c.GetString("school"),hex.EncodeToString(h.Sum(nil)))
	if !captchaBoolean || len(user) == 0 {
		if !captchaBoolean {
			c.Data["json"] = map[string]interface{}{"name": -1, "message": "你输入的验证码不正确"}
		}else {
			c.Data["json"] = map[string]interface{}{"name": 0, "message": "你输入的账号或密码不正确"}
		}
	}else {
		c.SetSession("account",c.GetString("school"))
		c.Data["json"] = map[string]interface{}{"name": 1, "message": "登陆成功"}
	}
	c.ServeJSON()
	c.TplName = "reception/index.html"
}

//@router /github [get]
func (c *LoginContorller)GithubCallback()  {
	//c.TplName = "reception/index.html"
}