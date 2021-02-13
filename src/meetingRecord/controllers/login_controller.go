package controllers

import (
	"github.com/astaxie/beego"
	"meetingRecord/models"
	"meetingRecord/utils"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	user := models.QueryUserRowForLogin(username, utils.MD5(password))

	if user.ID > 0 {
		c.SetSession("loginuser", username)
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	c.ServeJSON()
}
