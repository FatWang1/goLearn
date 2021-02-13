package main

import (
	"github.com/astaxie/beego"
	_ "meetingRecord/routers"
	"meetingRecord/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
