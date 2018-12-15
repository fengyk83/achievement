package backstage

import (
	"achievement/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"regexp"
)

type LoginController struct {
	beego.Controller
}

//后台登陆
//@router /backlogin [get]
func (this *LoginController)Login()  {
	this.TplName = "backstage/login.html";
}

//@router /backlogin/judge [post]
func (this *LoginController)JudgeLogin()  {
	valid := validation.Validation{};
	valid.MaxSize(this.GetString("school"),11,"school").Message("您输入的学号不正确")
	valid.AlphaNumeric(this.GetString("school"),"number").Message("请您输入的学号必须是数字字符")
	valid.Match(this.GetString("password"),regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`),"password").Message("您输入的密码格式不正确")
	valid.Required(this.GetString("password"),"password").Message("请您输入密码")
	valid.Required(this.GetString("captcha"),"captcha").Message("请您输入验证码")
	if valid.HasErrors() {
		for _,err := range valid.Errors {
			beego.Alert(err.Key+"+++++++"+err.Message)
		}
		return
	}
	user := models.NewUser().LoginJudge(this.GetString("school"),this.GetString("password"))
	if len(user) == 0 {
		this.Data["json"] = map[string]interface{}{"name": 0, "message": "你输入的账号或密码不正确"}
	}else {
		if user[0].Type == 3 {
			this.Data["json"] = map[string]interface{}{"name": 1, "message": "您没有权限登陆"}
		}else {
			this.SetSession("account",this.GetString("school"))
			this.Data["json"] = map[string]interface{}{"name": 1, "message": "登陆成功"}
		}
	}
	this.ServeJSON()
	this.TplName = "backstage/index.html"
}

