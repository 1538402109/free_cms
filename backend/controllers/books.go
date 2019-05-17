package controllers

import (
	"free_cms/backend/models"
	model "free_cms/common/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type BooksController struct {
	CommonController
}

func (c *BooksController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewBooks().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "books/index.html"
}

func (c *BooksController) Create() {
	if c.Ctx.Input.IsPost() {
		books := models.NewBooks()
		//1.压入数据
		if err := c.ParseForm(books); err != nil {
			c.JsonResult(0, "表单赋值失败")
		}
		//2.验证，在模型中验证不能分场景
		valid := validation.Validation{}
		valid.Required(books.BookName, "book_name").Message("小说名称不能为空")
		valid.Required(books.ListUrl, "list_url").Message("列表页地址不能为空")
		valid.Required(books.PregId, "preg_id").Message("采集规则不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(0, "验证失败")
		}

		//3.插入数据
		if err, _ := books.Create(); err != nil {
			c.JsonResult(0, "添加失败")
		}
		c.JsonResult(0, "添加成功")
	}

	var booksPregs []models.BooksPreg
	model.Db.Find(&booksPregs)
	c.Data["booksPregs"] = booksPregs
	_, booksType := models.NewBooksType().FindColumn()
	c.Data["bookType"] = booksType

	c.TplName = c.ADMIN_TPL + "books/create.html"
}

func (c *BooksController) Update() {
	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("id")
		books, _ := models.NewBooks().FindById(id)
		//1
		if err := c.ParseForm(&books); err != nil {
			c.JsonResult(0, "表单赋值失败")
		}
		//2
		valid := validation.Validation{}
		valid.Required(books.Id, "id").Message("id不能为空")
		valid.Required(books.BookName, "book_name").Message("小说名称不能为空")
		valid.Required(books.ListUrl, "list_url").Message("列表页地址不能为空")
		valid.Required(books.PregId, "preg_id").Message("采集规则不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(0, "验证失败")
		}
		//3
		if err, _ := books.Update(); err != nil {
			c.JsonResult(0, "更新失败")
		}
		c.JsonResult(0, "修改成功")
	}

	var booksPregs []models.BooksPreg
	model.Db.Find(&booksPregs)
	c.Data["booksPregs"] = booksPregs
	_, booksType := models.NewBooksType().FindColumn()
	c.Data["bookType"] = booksType

	c.TplName = c.ADMIN_TPL + "books/update.html"
}

func (c *BooksController) Delete() {
	books := models.NewBooks()
	id, _ := c.GetInt("id")
	books.Id = id
	if err := books.Delete(); err != nil {
		c.JsonResult(0, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *BooksController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(0, "赋值失败")
	}

	books := models.NewBooks()
	if err := books.DelBath(ids); err != nil {
		c.JsonResult(0, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}
