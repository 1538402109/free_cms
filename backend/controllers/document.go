package controllers

type DocumentController struct {
	CommonController
}

func(c *DocumentController)Demo1(){

	c.TplName = c.ADMIN_TPL+"document/demo1.html"
}

func(c *DocumentController)Demo2(){

	c.TplName = c.ADMIN_TPL+"document/demo2.html"
}

func(c *DocumentController)Demo3(){

	c.TplName = c.ADMIN_TPL+"document/demo3.html"
}