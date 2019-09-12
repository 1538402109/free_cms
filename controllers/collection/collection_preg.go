package collection

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type CollectionPregController struct {
	controllers.BaseController
}

func (c *CollectionPregController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewCollectionPreg().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "collection_preg/index.html"
}

func (c *CollectionPregController) Create() {
	if c.Ctx.Input.IsPost() {
		collectionPregModel := models.NewCollectionPreg()
		//1.压入数据
		if err := c.ParseForm(collectionPregModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(collectionPregModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := collectionPregModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["vo"] = models.CollectionPreg{}
	c.TplName = c.ADMIN_TPL + "collection_preg/create.html"
}

func (c *CollectionPregController) Update() {
	id, _ := c.GetInt("id")
	collectionPreg, _ := models.NewCollectionPreg().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&collectionPreg); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(collectionPreg); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := collectionPreg.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	c.Data["row"] = collectionPreg
	c.TplName = c.ADMIN_TPL + "collection_preg/update.html"
}

func (c *CollectionPregController) Delete() {
	collectionPregModel := models.NewCollectionPreg()
	id, _ := c.GetInt("id")
	collectionPregModel.Id = id
	if err := collectionPregModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *CollectionPregController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	collectionPregModel := models.NewCollectionPreg()
	if err := collectionPregModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

