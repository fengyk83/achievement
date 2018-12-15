package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["achievement/reception:ForgetController"] = append(beego.GlobalControllerRouter["achievement/reception:ForgetController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/forget`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/reception:InformationController"] = append(beego.GlobalControllerRouter["achievement/reception:InformationController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/admin/information`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/reception:LoginContorller"] = append(beego.GlobalControllerRouter["achievement/reception:LoginContorller"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/reception:LoginContorller"] = append(beego.GlobalControllerRouter["achievement/reception:LoginContorller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/login/judge`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
