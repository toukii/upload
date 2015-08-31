package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"],
		beego.ControllerComments{
			"LoadUpload",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"],
		beego.ControllerComments{
			"UploadForm",
			`/uploadform`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"],
		beego.ControllerComments{
			"Download",
			`/download/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"],
		beego.ControllerComments{
			"LoadFile",
			`/loadfile/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"],
		beego.ControllerComments{
			"ListFile",
			`/list/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"],
		beego.ControllerComments{
			"DeleteFile",
			`/delfile`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/upload/controllers:MainController"],
		beego.ControllerComments{
			"Upload",
			`/upload/*`,
			[]string{"*"},
			nil})

}
