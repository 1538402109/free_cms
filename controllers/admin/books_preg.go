package admin

import (
	"free_cms/models"
	"free_cms/pkg/d"
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
		c.Data["json"] = d.LayuiJson(0, "", result, count)
		c.ServeJSON()
	}
	c.TplName = ADMIN_TPL + "books_preg/index.html"
}

func (c *BooksPregController) Create() {
	if c.Ctx.Input.IsPost() {
		booksPreg := models.NewBooksPreg()
		//1.压入数据
		if err := c.ParseForm(booksPreg); err != nil {
			log.Info("表单赋值", err)
			c.Abort("500")
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
			c.Abort("500")
		}

		//3.插入数据
		if err, _ := booksPreg.Create(); err != nil {
			log.Info("添加数据", err)
			c.Abort("500")
		}
		c.Data["json"] = d.ReturnJson(200, "添加成功")
		c.ServeJSON()
	}
	c.TplName = ADMIN_TPL + "books_preg/create.html"
}

func (c *BooksPregController) Update() {
	if c.Ctx.Input.IsPost() {
		booksPreg := models.NewBooksPreg()
		//1
		if err := c.ParseForm(booksPreg); err != nil {
			c.Abort("500")
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
			c.Abort("500")
		}
		//3
		if err, _ := booksPreg.Update(); err != nil {
			c.Abort("500")
		}
		c.Data["json"] = d.ReturnJson(200, "修改成功")
		c.ServeJSON()
	}
	c.TplName = ADMIN_TPL + "books_preg/update.html"
}

func (c *BooksPregController) Delete() {
	booksPreg := models.NewBooksPreg()
	id, _ := c.GetInt("id")
	booksPreg.Id = id
	if err := booksPreg.Delete(); err != nil {
		c.Abort("500")
	}
	c.Data["json"] = d.ReturnJson(200, "删除成功")
	c.ServeJSON()
}

func (c *BooksPregController)BatchDelete(){
	var ids []int
	if err := c.Ctx.Input.Bind(&ids,"id");err !=nil{
		c.Abort("500")
	}

	booksPreg := models.NewBooksPreg()
	if err := booksPreg.DelBath(ids);err != nil{
		c.Abort("500")
	}
	c.Data["json"] = d.ReturnJson(200, "删除成功")
	c.ServeJSON()
}