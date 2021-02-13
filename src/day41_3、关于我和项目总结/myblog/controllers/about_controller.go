package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {

	c.Data["wechat"] = "微信：bailu_001"
	c.Data["qq"] = "QQ：527372384"
	c.Data["tel"] = "Tel：18983846023"

	c.TplName = "aboutme.html"
}
