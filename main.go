package main

import (
	"Data/db_mysql"
	"github.com/astaxie/beego"
	_ "Data/routers"

)

func main() {
	db_mysql.Connect()

	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")
	beego.Run()
}

