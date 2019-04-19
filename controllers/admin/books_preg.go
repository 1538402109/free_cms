package admin

import (
	"free_cms/models"
	"frp/utils/log"
	"github.com/astaxie/beego/validation"
)

type BooksPregController struct {
	BaseController
}

func (c *BooksPregController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewBooksPreg().Pagination((page-1)*limit, limit, key)
		c.Success(0,"ok",result,count)
	}
	c.TplName = ADMIN_TPL + "books_preg/index.html"
}

func (c *BooksPregController) Create() {
	if c.Ctx.Input.IsPost() {
		booksPreg := models.NewBooksPreg()
		//1.压入数据
		if err := c.ParseForm(booksPreg); err != nil {
			c.Error(0,"赋值失败")
		}
		//2.验证，在模型中验证不能分场景
		valid := validation.Validation{}
		valid.Required(booksPreg.Name,"name").Message("规则名称不能为空")
		valid.Required(booksPreg.ListABlock,"list_a_block").Message("列表链接不能为空")
		valid.Required(booksPreg.ListTitle,"list_title").Message("内容标题不能为空")
		valid.Required(booksPreg.ContentBlock,"content_block").Message("内容块不能为空")
		valid.Required(booksPreg.ContentText,"content_text").Message("内容文本不能为空")
		if valid.HasErrors(){
			for _, err := range valid.Errors {
				log.Info(err.Key, err.Message)
			}
			c.Error(0,"验证失败")
		}

		//3.插入数据
		if err, _ := booksPreg.Create(); err != nil {
			c.Error(0,"创建失败")
		}
		c.Success(0,"添加成功")
	}
	c.TplName = ADMIN_TPL + "books_preg/create.html"
}

func (c *BooksPregController) Update() {
	if c.Ctx.Input.IsPost() {
		booksPreg := models.NewBooksPreg()
		//1
		if err := c.ParseForm(booksPreg); err != nil {
			c.Error(0,"赋值失败")
		}
		//2
		valid := validation.Validation{}
		valid.Required(booksPreg.Id,"id").Message("id不能为空")
		valid.Required(booksPreg.Name,"name").Message("规则名称不能为空")
		valid.Required(booksPreg.ListABlock,"list_a_block").Message("列表链接不能为空")
		valid.Required(booksPreg.ListTitle,"list_title").Message("内容标题不能为空")
		valid.Required(booksPreg.ContentBlock,"content_block").Message("内容块不能为空")
		valid.Required(booksPreg.ContentText,"content_text").Message("内容文本不能为空")
		if valid.HasErrors(){
			for _, err := range valid.Errors {
				log.Info(err.Key, err.Message)
			}
			c.Error(0,"验证失败")
		}
		//3
		if err, _ := booksPreg.Update(); err != nil {
			c.Error(0,"修改失败")
		}
		c.Error(0,"修改成功")
	}
	c.TplName = ADMIN_TPL + "books_preg/update.html"
}

func (c *BooksPregController) Delete() {
	booksPreg := models.NewBooksPreg()
	id, _ := c.GetInt("id")
	booksPreg.Id = id
	if err := booksPreg.Delete(); err != nil {
		c.Error(0,"删除失败")
	}
	c.Error(0,"删除成功")
}

func (c *BooksPregController)BatchDelete(){
	var ids []int
	if err := c.Ctx.Input.Bind(&ids,"ids");err !=nil{
		c.Error(0,"赋值失败")
	}

	booksPreg := models.NewBooksPreg()
	if err := booksPreg.DelBath(ids);err != nil{
		c.Error(0,"删除失败")
	}
	c.Success(0,"删除成功")
}