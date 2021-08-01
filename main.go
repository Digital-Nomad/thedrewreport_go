package main

import (
	_ "thedrewreport/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

