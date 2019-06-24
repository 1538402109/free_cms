package books_preg

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type BooksPregController struct {
	controllers.BaseController
}

func (c *BooksPregController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewBooksPreg().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "books_preg/index.html"
}

func (c *BooksPregController) Create() {
	if c.Ctx.Input.IsPost() {
		booksPreg := models.NewBooksPreg()
		//1.压入数据
		if err := c.ParseForm(booksPreg); err != nil {
			c.JsonResult(1000, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(booksPreg); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1000, "验证失败")
		}
		//3.插入数据
		if err, _ := booksPreg.Create(); err != nil {
			c.JsonResult(1000, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}
	c.Data["vo"] = models.NewBooksPreg()
	c.TplName = c.ADMIN_TPL + "books_preg/create.html"
}

func (c *BooksPregController) Update() {
	id, _ := c.GetInt("id")
	booksPreg, _ := models.NewBooksPreg().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&booksPreg); err != nil {
			c.JsonResult(0, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(booksPreg); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1000, "验证失败")
		}
		//3
		if err, _ := booksPreg.Update(); err != nil {
			c.JsonResult(1000, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}
	c.Data["vo"] = booksPreg
	c.TplName = c.ADMIN_TPL + "books_preg/update.html"
}

func (c *BooksPregController) Delete() {
	booksPreg := models.NewBooksPreg()
	id, _ := c.GetInt("id")
	booksPreg.Id = id
	if err := booksPreg.Delete(); err != nil {
		c.JsonResult(1000, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *BooksPregController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1000, "赋值失败")
	}

	booksPreg := models.NewBooksPreg()
	if err := booksPreg.DelBatch(ids); err != nil {
		c.JsonResult(1000, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}
