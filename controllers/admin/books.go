package admin

import (
	"free_cms/models"
	"frp/utils/log"
	"github.com/astaxie/beego/validation"
)

type BooksController struct {
	BaseController
}

func (c *BooksController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewBooks().Pagination((page-1)*limit, limit, key)
		c.Success(0, "ok", result, count)
	}
	c.TplName = ADMIN_TPL + "books/index.html"
}

func (c *BooksController) Create() {
	if c.Ctx.Input.IsPost() {
		books := models.NewBooks()
		//1.压入数据
		if err := c.ParseForm(books); err != nil {
			c.Error(0, "表单赋值失败")
		}
		//2.验证，在模型中验证不能分场景
		valid := validation.Validation{}
		valid.Required(books.BookName, "book_name").Message("小说名称不能为空")
		valid.Required(books.ListUrl, "list_url").Message("列表页地址不能为空")
		valid.Required(books.PregId, "preg_id").Message("采集规则不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Info(err.Key, err.Message)
			}
			c.Error(0, "验证失败")
		}

		//3.插入数据
		if err, _ := books.Create(); err != nil {
			log.Info("添加数据", err)
			c.Abort("500")
			c.Error(0, "添加失败")
		}
		c.Success(0, "添加成功")
	}

	var booksPregs []models.BooksPreg
	models.Db.Find(&booksPregs)
	c.Data["booksPregs"] = booksPregs
	c.Data["bookType"] = models.BookType

	c.TplName = ADMIN_TPL + "books/create.html"
}

func (c *BooksController) Update() {
	if c.Ctx.Input.IsPost() {
		books := models.NewBooks()
		//1
		if err := c.ParseForm(books); err != nil {
			c.Error(0, "表单赋值失败")
		}
		//2
		valid := validation.Validation{}
		valid.Required(books.Id, "id").Message("id不能为空")
		valid.Required(books.BookName, "book_name").Message("小说名称不能为空")
		valid.Required(books.ListUrl, "list_url").Message("列表页地址不能为空")
		valid.Required(books.PregId, "preg_id").Message("采集规则不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Info(err.Key, err.Message)
			}
			c.Error(0, "验证失败")
		}
		//3
		if err, _ := books.Update(); err != nil {
			c.Error(0, "更新失败")
		}
		c.Success(0, "修改成功")
	}

	var booksPregs []models.BooksPreg
	models.Db.Find(&booksPregs)
	c.Data["booksPregs"] = booksPregs
	c.Data["bookType"] = models.BookType

	c.TplName = ADMIN_TPL + "books/update.html"
}

func (c *BooksController) Delete() {
	books := models.NewBooks()
	id, _ := c.GetInt("id")
	books.Id = id
	if err := books.Delete(); err != nil {
		c.Error(0, "删除失败")
	}
	c.Success(0, "删除成功")
}

func (c *BooksController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.Error(0, "赋值失败")
	}

	books := models.NewBooks()
	if err := books.DelBath(ids); err != nil {
		c.Error(0, "删除失败")
	}
	c.Success(0, "删除成功")
}
