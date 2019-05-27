package admin_log

import (
	"free_cms/controllers"
	"free_cms/models"
)

type AdminLogController struct {
	controllers.BaseController
}

func (c *AdminLogController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewAdminLog().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "admin_log/index.html"
}
