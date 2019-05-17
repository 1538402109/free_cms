package controllers

import (
	"free_cms/backend/models"
)

type AdminLogController struct {
	CommonController
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
