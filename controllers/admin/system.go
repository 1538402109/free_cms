package admin

type SystemController struct {
	BaseController
}

func(c *SystemController)Basic(){
	c.TplName = ADMIN_TPL+"system/basic.html"
}

func(c *SystemController)Logs(){
	c.TplName = ADMIN_TPL+"system/logs.html"
}

func(c *SystemController)Links(){
	c.TplName = ADMIN_TPL+"system/links.html"
}
func(c *SystemController)LinksAdd(){
	c.TplName = ADMIN_TPL+"system/links_add.html"
}


func(c *SystemController)Icons(){
	c.TplName = ADMIN_TPL+"system/icons.html"
}