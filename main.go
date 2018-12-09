package main

import (
	_ "achievements/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/beego.log"}`)
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/ssms?charset=utf8")
	beego.Run();
}

