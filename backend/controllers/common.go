package controllers

import (
	"free_cms/common"
	"free_cms/common/models"
	"free_cms/pkg/d"
)

type CommonController struct {
	common.BaseController
	ADMIN_TPL string
}

func Init() {

}

func (c *CommonController) Prepare() {
	c.ADMIN_TPL = "admin/"

	common.Fc = c.Ctx
	if user := c.GetSession("loginUser"); user != nil {
		common.UserId = user.(*models.User).Id
	}


/*	controller, action := c.GetControllerAndAction()
	if controller!="UserController" && c.GetSession("loginUser") == nil{
		c.History("未登录","/login")
	}

	if controller == "UserController" && action == "Login" && c.GetSession("loginUser") != nil {
		c.History("已登录", "/admin")
	}*/
}

func (c *CommonController) History(msg string, url string) {
	if url == "" {
		c.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		c.StopRun()
	} else {
		c.Redirect(url, 302)
	}
}

func (c *CommonController) JsonResult(code int, msg string, data ...interface{}) {
	if len(data) > 1 {
		c.Data["json"] = d.LayuiJson(code, msg, data[0], data[1])
	} else if len(data) > 0 {
		c.Data["json"] = d.LayuiJson(code, msg, data[0])
	} else {
		c.Data["json"] = d.LayuiJson(code, msg)
	}
	c.ServeJSON()
	c.StopRun()
}
