package routers

import (
	"basic_blog_go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/posts",
				beego.NSInclude(&controllers.PostController{}),
			),
		)
	beego.AddNamespace(ns)
}
