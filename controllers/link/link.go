package link

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type LinkController struct {
	controllers.BaseController
}

func (c *LinkController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewLink().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "link/index.html"
}

func (c *LinkController) Create() {
	if c.Ctx.Input.IsPost() {
		link := models.NewLink()
		//1.压入数据
		if err := c.ParseForm(link); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		valid.Required(link.Name, "name").Message("名称不能为空")
		valid.Required(link.Url, "url").Message("地址不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if err, _ := link.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.TplName = c.ADMIN_TPL + "link/create.html"
}

func (c *LinkController) Update() {
	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("id")

		link, _ := models.NewLink().FindById(id)
		//1
		if err := c.ParseForm(&link); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		valid.Required(link.Id, "id").Message("id不能为空")
		valid.Required(link.Name, "name").Message("名称不能为空")
		valid.Required(link.Url, "url").Message("地址不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if err, _ := link.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.TplName = c.ADMIN_TPL + "link/update.html"
}

func (c *LinkController) Delete() {
	link := models.NewLink()
	id, _ := c.GetInt("id")
	link.Id = id
	if err := link.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *LinkController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	link := models.NewLink()
	if err := link.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}