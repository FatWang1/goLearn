package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"meetingRecord/models"
	"meetingRecord/utils"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username, "\t", password)
	exist := models.QueryUserExist(username)
	if exist {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "该用户已被注册。"}
		c.ServeJSON()
		return
	}
	password = utils.MD5(password)
	user := models.Usr{username, password, 1}
	var usr models.User
	usr.Usr = user
	models.InsertUser(usr)
	exist = models.QueryUserExist(username)
	if exist {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败！！"}
	}
	c.ServeJSON()

}
