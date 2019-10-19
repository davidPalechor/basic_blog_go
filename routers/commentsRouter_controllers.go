package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["basic_blog_go/controllers:PostController"] = append(beego.GlobalControllerRouter["basic_blog_go/controllers:PostController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["basic_blog_go/controllers:UserController"] = append(beego.GlobalControllerRouter["basic_blog_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "CreateUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["basic_blog_go/controllers:UserController"] = append(beego.GlobalControllerRouter["basic_blog_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
