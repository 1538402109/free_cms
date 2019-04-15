package admin

import (
	"github.com/astaxie/beego"
)

const ADMIN_TPL = "admin/"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare()  {
	controller,action := c.GetControllerAndAction()
	//if controller!="UserController" && c.GetSession("loginUser") == nil{
	//	c.History("未登录","/login")
	//}

	if controller=="UserController" && action == "Login" && c.GetSession("loginUser") != nil{
		c.History("已登录","/admin")
	}
}

func (c *BaseController) History(msg string, url string) {
	if url == ""{
		c.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1);</script>")
		c.StopRun()
	}else{
		c.Redirect(url,302)
	}
}