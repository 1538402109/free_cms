package collection

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
	"strings"
)

type CollectionJobController struct {
	controllers.BaseController
}

func (c *CollectionJobController) Prepare() {
	c.BaseController.Prepare()
}

func (c *CollectionJobController) Index() {
	if c.Ctx.Input.IsAjax() {
		page, _ := c.GetInt("page")
		limit, _ := c.GetInt("limit")
		key := c.GetString("key", "")

		result, count := models.NewCollectionJob().Pagination((page-1)*limit, limit, key)
		c.JsonResult(0, "ok", result, count)
	}
	c.TplName = c.ADMIN_TPL + "collection_job/index.html"
}

func (c *CollectionJobController) Create() {
	if c.Ctx.Input.IsPost() {
		collectionJobModel := models.NewCollectionJob()

		var pregIds []string
		c.Ctx.Input.Bind(&pregIds, "preg_ids")
		var p []string
		for _, v := range pregIds {
			if v != "" {
				p = append(p, v)
			}
		}
		collectionJobModel.PregIds = strings.Join(p, ",")

		//1.压入数据
		if err := c.ParseForm(collectionJobModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(collectionJobModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := collectionJobModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	c.Data["preg"] = models.NewCollectionPreg().FindCheck("")
	c.Data["row"] = models.NewCollectionJob()
	c.TplName = c.ADMIN_TPL + "collection_job/create.html"
}

func (c *CollectionJobController) Update() {
	id, _ := c.GetInt("id")
	collectionJob, _ := models.NewCollectionJob().FindById(id)

	if c.Ctx.Input.IsPost() {
		var pregIds []string
		c.Ctx.Input.Bind(&pregIds, "preg_ids")
		var p []string
		for _, v := range pregIds {
			if v != "" {
				p = append(p, v)
			}
		}
		collectionJob.PregIds = strings.Join(p, ",")
		//1
		if err := c.ParseForm(&collectionJob); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(collectionJob); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := collectionJob.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}
	c.Data["preg"] = models.NewCollectionPreg().FindCheck(collectionJob.PregIds)
	c.Data["row"] = collectionJob
	c.TplName = c.ADMIN_TPL + "collection_job/update.html"
}

func (c *CollectionJobController) Delete() {
	collectionJobModel := models.NewCollectionJob()
	id, _ := c.GetInt("id")
	collectionJobModel.Id = id
	if err := collectionJobModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *CollectionJobController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	collectionJobModel := models.NewCollectionJob()
	if err := collectionJobModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}
