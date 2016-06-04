package routers

import (
	"github.com/astaxie/beego"
	"github.com/shaalx/goutils"
	"github.com/shaalx/upload/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.MainController{})
	beego.AddFuncMap("ToString", func(b []byte) string {
		return goutils.ToString(b)
	})
}
