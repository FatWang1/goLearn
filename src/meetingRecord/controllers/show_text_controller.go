package controllers

import (
	"fmt"
	"meetingRecord/models"
	"strconv"
)

type ShowTextController struct {
	BaseController
}

func (c *ShowTextController) Get() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	txt := models.QueryTextById(id)

	c.Data["Project"] = txt.Project
	c.Data["Achievements"] = txt.Achievements
	c.Data["Goal"] = txt.Goal
	c.TplName = "show_text.html"
}
