package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["wm/controllers:LoginController"] = append(beego.GlobalControllerRouter["wm/controllers:LoginController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["wm/controllers:ObjectController"] = append(beego.GlobalControllerRouter["wm/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["wm/controllers:ObjectController"] = append(beego.GlobalControllerRouter["wm/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["wm/controllers:ObjectController"] = append(beego.GlobalControllerRouter["wm/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["wm/controllers:ObjectController"] = append(beego.GlobalControllerRouter["wm/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["wm/controllers:ObjectController"] = append(beego.GlobalControllerRouter["wm/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["wm/controllers:SigninController"] = append(beego.GlobalControllerRouter["wm/controllers:SigninController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}
