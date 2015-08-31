package routers

import (
	"github.com/astaxie/beego"
	"github.com/shaalx/upload/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.MainController{})
}
