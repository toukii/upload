package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"Home",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"LHome",
			`/list/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"LoadUpload",
			`/upload`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"LoadUploads",
			`/upload/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"UploadForm",
			`/uploadform`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"DirUploadForm",
			`/uploadform/*`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"Download",
			`/download/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"LoadFile",
			`/loadfile/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"PostDisplay",
			`/display/*`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"Display",
			`/display/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"ListFile",
			`/list/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"DeleteFile",
			`/delfile/*`,
			[]string{"*"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"Upload",
			`/upload/*`,
			[]string{"post", "put"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"GJob",
			`/job`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"GJobs",
			`/job/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"PJob",
			`/job`,
			[]string{"post"},
			nil})
	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"GTopic",
			`/topic`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"GTopics",
			`/topic/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"PTopics",
			`/topic/*`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"PTopic",
			`/topic`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"Bash",
			`/bash`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"GoogleURL",
			`/url`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"GoogleSearchURL",
			`/search`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/toukii/upload/controllers:MainController"],
		beego.ControllerComments{
			"PBash",
			`/bash`,
			[]string{"post"},
			nil})

}
