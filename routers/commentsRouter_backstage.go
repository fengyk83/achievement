package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["achievement/backstage:IndexController"] = append(beego.GlobalControllerRouter["achievement/backstage:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/achievement/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:LoginController"] = append(beego.GlobalControllerRouter["achievement/backstage:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/backlogin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
