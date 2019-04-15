package admin

type MembersController struct {
	BaseController
}

func (c *MembersController)Center()  {
	c.TplName = ADMIN_TPL+"members/center.html"
}

func (c *MembersController)CenterAdd(){
	c.TplName = ADMIN_TPL+"members/center_add.html"
}

func (c *MembersController)Level()  {
	c.TplName = ADMIN_TPL+"members/level.html"
}

