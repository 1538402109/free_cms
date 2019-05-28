package nav

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type NavController struct {
	controllers.BaseController
}

func (c *NavController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewNav().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "nav/index.html"
}

func (c *NavController) Create() {
	if c.Ctx.Input.IsPost() {
		nav := models.NewNav()
		//1.压入数据
		if err := c.ParseForm(nav); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if err, _ := nav.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.TplName = c.ADMIN_TPL + "nav/create.html"
}

func (c *NavController) Update() {
	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("id")

		nav, _ := models.NewNav().FindById(id)
		//1
		if err := c.ParseForm(&nav); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		valid.Required(nav.Id, "id").Message("id不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if err, _ := nav.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.TplName = c.ADMIN_TPL + "nav/update.html"
}

func (c *NavController) Delete() {
	nav := models.NewNav()
	id, _ := c.GetInt("id")
	nav.Id = id
	if err := nav.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *NavController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	nav := models.NewNav()
	if err := nav.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}
