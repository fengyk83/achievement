package backstage

import (
	"achievement/models"
	"fmt"
	"github.com/astaxie/beego"
)

type IdiomController struct {
	beego.Controller
}

//展示成员
//@router /achievement/showidiom [get]
func (this *IdiomController)ShowIdiom()  {
	this.Data["grade"] = models.NewGrade().GetGrade()
	this.Data["class"] = models.NewClazz().GetClazz()
	this.TplName = "backstage/people.html"
}

//添加成员
//@router /achievement/addidiom [post]
func (this *IdiomController)AddIdiom()  {

	this.TplName = "backstage/people.html"
}

// 上传Excel 表格
//@router /achievement/excel [post]
func (this *IdiomController)AddExcel()  {
	f, h, error := this.GetFile("myfile")                  //获取上传的文件
	fmt.Println(f)
	fmt.Println(h)
	if error == nil {
		
	}
	//path := h.Filename 　　　　　　　　　　　　　 			//文件目录
	//f.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	//this.SaveToFile("myfile", path)                    //存文件
}
