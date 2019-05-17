package controllers

type MembersController struct {
	CommonController
}

func (c *MembersController)Center()  {
	c.TplName = c.ADMIN_TPL+"members/center.html"
}

func (c *MembersController)CenterAdd(){
	c.TplName = c.ADMIN_TPL+"members/center_add.html"
}

func (c *MembersController)Level()  {
	c.TplName = c.ADMIN_TPL+"members/level.html"
}

