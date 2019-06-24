package books_type

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type BooksTypeController struct {
	controllers.BaseController
}

func (c *BooksTypeController) Index() {
	if c.Ctx.Input.IsAjax() {
		key := c.GetString("key", "")

		result, _ := models.NewBooksType().FindTree(key)

		c.JsonResult(0, "ok", result)
	}
	c.TplName = c.ADMIN_TPL + "books_type/index.html"
}

func (c *BooksTypeController) Create() {
	if c.Ctx.Input.IsPost() {
		booksType := models.NewBooksType()
		//1.压入数据
		if err := c.ParseForm(booksType); err != nil {
			c.JsonResult(0, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(booksType); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1000, "验证失败")
		}

		//3.插入数据
		if err, _ := booksType.Create(); err != nil {
			c.JsonResult(0, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}
	result, _ := models.NewBooksType().FindTree("")
	c.Data["booksType"] = result
	c.Data["vo"] = models.NewBooksType()
	c.TplName = c.ADMIN_TPL + "books_type/create.html"
}

func (c *BooksTypeController) Update() {
	id, _ := c.GetInt("id")
	booksType, _ := models.NewBooksType().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&booksType); err != nil {
			c.JsonResult(0, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(booksType); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1000, "验证失败")
		}
		//3
		if err, _ := booksType.Update(); err != nil {
			c.JsonResult(0, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}
	result, _ := models.NewBooksType().FindTree("")
	c.Data["booksType"] = result
	c.Data["vo"] = booksType
	c.TplName = c.ADMIN_TPL + "books_type/update.html"
}

func (c *BooksTypeController) Delete() {
	BooksType := models.NewBooksType()
	id, _ := c.GetInt("id")
	BooksType.Id = id
	if err := BooksType.Delete(); err != nil {
		c.JsonResult(0, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *BooksTypeController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(0, "赋值失败")
	}

	BooksType := models.NewBooksType()
	if err := BooksType.DelBatch(ids); err != nil {
		c.JsonResult(0, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}
