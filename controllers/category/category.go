package category

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type CategoryController struct {
	controllers.BaseController
}

func (c *CategoryController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewCategory().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "category/index.html"
}

func (c *CategoryController) Create() {
	if c.Ctx.Input.IsPost() {
		categoryModel := models.NewCategory()
		//1.压入数据
		if err := c.ParseForm(categoryModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		//valid.Required(category.Name, "name").Message("名称不能为空")
		//valid.Required(category.Url, "url").Message("地址不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if err, _ := categoryModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.TplName = c.ADMIN_TPL + "category/create.html"
}

func (c *CategoryController) Update() {
	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("id")

		categoryModel := models.NewCategory()
		category, _ := categoryModel.FindById(id)
		//1
		if err := c.ParseForm(&category); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		//valid.Required(category.Id, "id").Message("id不能为空")
		//valid.Required(category.Name, "name").Message("名称不能为空")
		//valid.Required(category.Url, "url").Message("地址不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if err, _ := categoryModel.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.TplName = c.ADMIN_TPL + "category/update.html"
}

func (c *CategoryController) Delete() {
	categoryModel := models.NewCategory()
	id, _ := c.GetInt("id")
	categoryModel.Id = id
	if err := categoryModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *CategoryController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	categoryModel := models.NewCategory()
	if err := categoryModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}
