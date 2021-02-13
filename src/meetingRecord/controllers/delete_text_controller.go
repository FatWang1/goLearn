package controllers

import (
	"fmt"
	"meetingRecord/models"
)

type DeleteTextController struct {
	BaseController
}

func (c *DeleteTextController) Get() {
	id, _ := c.GetInt("id")

	err := models.DeleteTextById(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Redirect("/", 302)
}
