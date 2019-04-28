package books

import (
	"free_cms/controllers/admin"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type BooksTypeController struct {
	admin.BaseController
}

func (c *BooksTypeController) Index() {
	if c.Ctx.Input.IsAjax() {
		key := c.GetString("key", "")

		_, result := models.NewBooksType().FindTree(key)

		c.JsonResult(0, "ok", result)
	}
	c.TplName = c.ADMIN_TPL + "books_type/index.html"
}

func (c *BooksTypeController) Create() {
	if c.Ctx.Input.IsPost() {
		BooksType := models.NewBooksType()
		//1.压入数据
		if err := c.ParseForm(BooksType); err != nil {
			c.JsonResult(0, "赋值失败")
		}
		//2.验证，在模型中验证不能分场景
		valid := validation.Validation{}
		valid.Required(BooksType.Name, "name").Message("规则名称不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(0, "验证失败")
		}

		//3.插入数据
		if err, _ := BooksType.Create(); err != nil {
			c.JsonResult(0, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}
	_, result := models.NewBooksType().FindTree("")
	c.Data["booksType"] = result
	c.TplName = c.ADMIN_TPL + "books_type/create.html"
}

func (c *BooksTypeController) Update() {
	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("id")
		BooksType, _ := models.NewBooksType().FindById(id)
		//1
		if err := c.ParseForm(BooksType); err != nil {
			c.JsonResult(0, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		valid.Required(BooksType.Id, "id").Message("id不能为空")
		valid.Required(BooksType.Name, "name").Message("规则名称不能为空")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(0, "验证失败")
		}
		//3
		if err, _ := BooksType.Update(); err != nil {
			c.JsonResult(0, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}
	_, result := models.NewBooksType().FindTree("")
	c.Data["booksType"] = result
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
	if err := BooksType.DelBath(ids); err != nil {
		c.JsonResult(0, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}
