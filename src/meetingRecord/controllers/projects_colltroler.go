package controllers

import "meetingRecord/models"

type ProjectsController struct {
	BaseController
}

func (c *ProjectsController) Get() {
	projects := models.CountTextByPara("project")

	c.Data["Projects"] = models.HandleProjectsListData(projects)

	c.TplName = "projects.html"
}
