package controllers


type SystemController struct {
	CommonController
}

func(c *SystemController)Basic(){
	c.TplName = c.ADMIN_TPL+"system/basic.html"
}

func(c *SystemController)Logs(){
	c.TplName = c.ADMIN_TPL+"system/logs.html"
}

func(c *SystemController)Links(){
	c.TplName = c.ADMIN_TPL+"system/links.html"
}
func(c *SystemController)LinksAdd(){
	c.TplName = c.ADMIN_TPL+"system/links_add.html"
}

func(c *SystemController)Icons(){
	c.TplName = c.ADMIN_TPL+"system/icons.html"
}