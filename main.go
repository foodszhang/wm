package main

import (
	"fmt"
	_ "wm/docs"
	_ "wm/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://wm:wwww@localhost/wm")
	orm.RunCommand()
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
