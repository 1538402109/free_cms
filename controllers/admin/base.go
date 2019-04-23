package admin

import (
	"free_cms/pkg/d"
	"github.com/astaxie/beego"
)

const ADMIN_TPL = "admin/"

type BaseController struct {
	beego.Controller
	ADMIN_TPL string
}

func (c *BaseController) Prepare() {
	c.ADMIN_TPL="admin/"

	controller, action := c.GetControllerAndAction()
	//if controller!="UserController" && c.GetSession("loginUser") == nil{
	//	c.History("未登录","/login")
	//}

	if controller == "UserController" && action == "Login" && c.GetSession("loginUser") != nil {
		c.History("已登录", "/admin")
	}
}

func (c *BaseController) History(msg string, url string) {
	if url == "" {
		c.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		c.StopRun()
	} else {
		c.Redirect(url, 302)
	}
}

func (c *BaseController) Success(code int, msg string, data ...interface{}) {
	if len(data) > 1 {
		c.Data["json"] = d.LayuiJson(code, msg, data[0], data[1])
	} else if len(data) > 0 {
		c.Data["json"] = d.LayuiJson(code, msg, data[0])
	} else {
		c.Data["json"] = d.LayuiJson(code, msg)
	}
	c.ServeJSON()
}

func (c *BaseController) Error(code int, msg string, data ...interface{}) {
	if len(data) > 1 {
		c.Data["json"] = d.LayuiJson(code, msg, data[0])
	} else if len(data) > 0 {
		c.Data["json"] = d.LayuiJson(code, msg, data[0], data[1])
	} else {
		c.Data["json"] = d.LayuiJson(code, msg)
	}
	c.ServeJSON()
}
