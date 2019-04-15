package admin

type MainController struct {
	BaseController
}

func (c *MainController) Admin() {
	c.TplName = "admin/default/index.html"
}

func (c *MainController)Main(){
	c.TplName="admin/default/main.html"
}