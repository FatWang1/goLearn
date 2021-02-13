package controllers

import (
	"meetingRecord/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {

	project := c.GetString("tag")
	page, _ := c.GetInt("page")

	var textList []models.Text

	if len(project) > 0 {
		textList = models.QueryTextByProject(project)
		c.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}

		textList = models.QueryTextByPage(page)
		c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		c.Data["HasFooter"] = true
	}
	c.Data["Content"] = models.MakeHomeBlocks(textList, c.IsLogin)
	c.TplName = "home.html"
}
