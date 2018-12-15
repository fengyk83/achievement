package backstage

import "github.com/astaxie/beego"

type ScroeController struct {
	beego.Controller
}

//展示成绩
//@router /achievement/showscore [get]
func (this *ScroeController)ShowScore()  {
	this.TplName ="backstage/score.html"
}
