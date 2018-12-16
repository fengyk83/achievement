package reception

import (
	"github.com/astaxie/beego"
	"net/smtp"
	"strings"
)

type ForgetController struct {
	beego.Controller
}

//@router /forget [get]
func (this *ForgetController) Index() {
	//this.TplName = "reception/forget.html"
	//c.TplName = "reception/index.html"
	//err := this.Email("2181550591@qq.com", "xsywrvvxqzwxebbi", "smtp.qq.com:25", "3438196289@qq.com", "大橘猫", "sssss", "html")
	//r := rand.New(rand.NewSource(time.Now().Unix()))

	this.TplName = "reception/forget.html"
}

//err := SendMail("发送的邮箱", "发送的邮箱密码", "smtp.qq.com:25", "目标邮箱", "邮件标题", "邮件内容", "html")
func (this *ForgetController)Email(user, password, host, to, subject, body, mailtype string) error  {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}