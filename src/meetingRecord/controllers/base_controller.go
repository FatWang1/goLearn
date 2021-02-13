package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

func (c *BaseController) Prepare() {
	loginUser := c.GetSession("loginuser")
	if loginUser != nil {
		c.IsLogin = true
		c.Loginuser = loginUser
	} else {
		c.IsLogin = false
	}
	c.Data["IsLogin"] = c.IsLogin
}
