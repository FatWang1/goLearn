package controllers

import (
	"fmt"
	"meetingRecord/models"
)

type UpdateTextController struct {
	BaseController
}

func (c *UpdateTextController) Get() {
	id, _ := c.GetInt("id")

	txt := models.QueryTextById(id)

	c.Data["Projects"] = txt.Project
	c.Data["Achievements"] = txt.Achievements
	c.Data["Goal"] = txt.Goal
	c.Data["Id"] = txt.ID

	c.TplName = "write_text.html"
}

func (c *UpdateTextController) Post() {
	id, _ := c.GetInt("id")

	var txt models.Text
	txt.ID = uint(id)
	txt.Achievements = c.GetString("Achievements")
	txt.Goal = c.GetString("Goal")
	txt.Project = c.GetString("Projects")

	err := models.UpdateText(txt)
	if err != nil {
		fmt.Println(err.Error())
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	}
	c.ServeJSON()
}
