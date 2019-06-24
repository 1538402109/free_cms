package books

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type BooksController struct {
	controllers.BaseController
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
			c.JsonResult(1000, "表单赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(books); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1000, "验证失败")
		}

		//3.插入数据
		if err, _ := books.Create(); err != nil {
			c.JsonResult(1000, "添加失败")
		}
		c.JsonResult(0, "添加成功")
	}

	var booksPregs []models.BooksPreg
	models.Db.Find(&booksPregs)
	_, booksType := models.NewBooksType().FindColumn()

	c.Data["booksPregs"] = booksPregs
	c.Data["bookType"] = booksType
	c.Data["vo"] = models.NewBooks()

	c.TplName = c.ADMIN_TPL + "books/create.html"
}

func (c *BooksController) Update() {
	id, _ := c.GetInt("id")
	books, _ := models.NewBooks().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&books); err != nil {
			c.JsonResult(1000, "表单赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(books); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1000, "验证失败")
		}
		//3
		if err, _ := books.Update(); err != nil {
			c.JsonResult(1000, "更新失败")
		}
		c.JsonResult(0, "修改成功")
	}

	var booksPregs []models.BooksPreg
	models.Db.Find(&booksPregs)
	_, booksType := models.NewBooksType().FindColumn()

	c.Data["booksPregs"] = booksPregs
	c.Data["bookType"] = booksType
	c.Data["vo"] = books

	c.TplName = c.ADMIN_TPL + "books/update.html"
}

func (c *BooksController) Delete() {
	books := models.NewBooks()
	id, _ := c.GetInt("id")
	books.Id = id
	if err := books.Delete(); err != nil {
		c.JsonResult(1000, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *BooksController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1000, "赋值失败")
	}

	books := models.NewBooks()
	if err := books.DelBatch(ids); err != nil {
		c.JsonResult(1000, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}
