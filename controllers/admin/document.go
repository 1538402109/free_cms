package admin

type DocumentController struct {
	BaseController
}

func(c *DocumentController)Demo1(){

	c.TplName = ADMIN_TPL+"document/demo1.html"
}

func(c *DocumentController)Demo2(){

	c.TplName = ADMIN_TPL+"document/demo2.html"
}

func(c *DocumentController)Demo3(){

	c.TplName = ADMIN_TPL+"document/demo3.html"
}