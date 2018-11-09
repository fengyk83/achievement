package main

import (
	_ "achievement/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/beego.log"}`)
	beego.Run();
}

