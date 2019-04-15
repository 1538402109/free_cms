package admin

import (
	"free_cms/models"
	"free_cms/pkg/d"
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
		c.Data["json"] = d.ReturnJson(0, "", result, count)
		c.ServeJSON()
	}
	c.TplName = ADMIN_TPL + "books/index.html"
}

func (c *BooksController) Create() {
	if c.Ctx.Input.IsPost() {
		books := models.NewBooks()
		//1.压入数据
		if err := c.ParseForm(books); err != nil {
			log.Info("表单赋值", err)
			c.Abort("500")
		}
		//2.验证，在模型中验证不能分场景
		valid := validation.Validation{}
		valid.Required(books.BookName,"book_name").Message("小说名称不能为空")
		valid.Required(books.ListUrl,"list_url").Message("列表页地址不能为空")
		valid.Required(books.PregId,"preg_id").Message("采集规则不能为空")
		if valid.HasErrors(){
			for _, err := range valid.Errors {
				log.Info(err.Key, err.Message)
			}
			c.Abort("500")
		}

		//3.插入数据
		if err, _ := books.Create(); err != nil {
			log.Info("添加数据", err)
			c.Abort("500")
		}
		c.Data["json"] = d.ReturnJson(200, "添加成功")
		c.ServeJSON()
	}
	c.TplName = ADMIN_TPL + "books/update.html"
}

func (c *BooksController) Update() {
	if c.Ctx.Input.IsPost() {
		books := models.NewBooks()
		//1
		if err := c.ParseForm(books); err != nil {
			c.Abort("500")
		}
		//2
		valid := validation.Validation{}
		valid.Required(books.ID,"id").Message("id不能为空")
		valid.Required(books.BookName,"book_name").Message("小说名称不能为空")
		valid.Required(books.ListUrl,"list_url").Message("列表页地址不能为空")
		valid.Required(books.PregId,"preg_id").Message("采集规则不能为空")
		if valid.HasErrors(){
			for _, err := range valid.Errors {
				log.Info(err.Key, err.Message)
			}
			c.Abort("500")
		}
		//3
		if err, _ := books.Update(); err != nil {
			c.Abort("500")
		}
		c.Data["json"] = d.ReturnJson(200, "修改成功")
		c.ServeJSON()
	}
	c.TplName = ADMIN_TPL + "books/update.html"
}

func (c *BooksController) Delete() {
	books := models.NewBooks()
	id, _ := c.GetInt("id")
	books.ID = id
	if err := books.Delete(); err != nil {
		c.Abort("500")
	}
	c.Data["json"] = d.ReturnJson(200, "删除成功")
	c.ServeJSON()
}
