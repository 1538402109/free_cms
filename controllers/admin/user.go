package admin

import (
	"fmt"
	"free_cms/models"
	"free_cms/pkg/util"
)

type UserController struct {
	BaseController
}

func (c *UserController) Index() {

}

func (c *UserController) Register() {
	c.TplName = ADMIN_TPL + "user/register.html"
}

func (c *UserController) Login() {
	if c.Ctx.Input.IsPost() {
		captcha := c.GetString("captcha")
		idkey := c.GetString("idkey")
		user := models.NewUser()
		if err := c.ParseForm(user); err != nil {
			fmt.Println(err)
			c.Abort("500")
		}

		if login := user.Login(); util.VerfiyCaptcha(idkey, captcha) && login !=nil {
			c.SetSession("loginUser", login)
			c.Redirect("admin", 302)
		}
	}
	c.TplName = ADMIN_TPL + "user/login.html"
}

func (c *UserController) Captcha() {
	c.Data["json"] = util.CreateCaptcha()
	c.ServeJSON()
}

func (c *UserController) Logout() {

}
