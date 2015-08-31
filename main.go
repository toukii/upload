package main

import (
	"github.com/astaxie/beego"
	_ "github.com/shaalx/upload/routers"
)

func main() {
	beego.Run()
}
