package routers

import (
	"basic_blog_go/auth"
	"basic_blog_go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/posts",
				beego.NSInclude(&controllers.PostController{}),
			),
			beego.NSNamespace("/users",
				beego.NSInclude(&controllers.UserController{}),
			),
			beego.NSNamespace("/trends",
				beego.NSInclude(&controllers.TrendController{}),
			),
		)
	beego.AddNamespace(ns)

	beego.InsertFilter("/v1/trends", beego.BeforeRouter, auth.ValidateToken)
	beego.InsertFilter("/v1/posts", beego.BeforeRouter, auth.ValidateToken)
}
