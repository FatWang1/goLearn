package controllers

type ExitController struct {
	BaseController
}

func (c *ExitController) Get() {
	c.DelSession("loginuser")
	c.Redirect("/", 302)
}
