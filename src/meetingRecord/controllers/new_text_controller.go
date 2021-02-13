package controllers

import (
	"meetingRecord/models"
)

type NewTextController struct {
	BaseController
}

func (c *NewTextController) Get() {
	c.TplName = "write_text.html"
}

func (c *NewTextController) Post() {
	projects := c.GetString("projects")
	achievements := c.GetString("achievements")
	goal := c.GetString("goal")
	user := c.GetSession("loginuser").(string)
	txt := models.Txt{Project: projects, Achievements: achievements, Goal: goal, User: user}
	var text models.Text
	text.Txt = txt

	err := models.InsertText(text)
	var response map[string]interface{}

	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}
	c.Data["json"] = response
	c.ServeJSON()
}
