package fc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var controllerTpl = `package category

import (
	"free_cms/controllers"
	"free_cms/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type CategoryController struct {
	controllers.BaseController
}

func (c *CategoryController) PrePare() {

}

func (c *CategoryController) Index() {
	if c.Ctx.Input.IsAjax() {
		key := c.GetString("key", "")
		result, _ := models.NewCategory().FindTree(key)
		c.JsonResult(0, "ok", result)
	}
	c.TplName = c.ADMIN_TPL + "category/index.html"
}

func (c *CategoryController) Create() {
	if c.Ctx.Input.IsPost() {
		categoryModel := models.NewCategory()
		//1.压入数据
		if err := c.ParseForm(categoryModel); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2.验证
		valid := validation.Validation{}
		if b, _ := valid.Valid(categoryModel); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3.插入数据
		if _, err := categoryModel.Create(); err != nil {
			c.JsonResult(1001, "创建失败")
		}
		c.JsonResult(0, "添加成功")
	}

	result, _ := models.NewCategory().FindTree("")
	c.Data["category"] = result
	c.Data["vo"] = models.Category{}
	c.TplName = c.ADMIN_TPL + "category/create.html"
}

func (c *CategoryController) Update() {
	id, _ := c.GetInt("id")
	category, _ := models.NewCategory().FindById(id)

	if c.Ctx.Input.IsPost() {
		//1
		if err := c.ParseForm(&category); err != nil {
			c.JsonResult(1001, "赋值失败")
		}
		//2
		valid := validation.Validation{}
		if b, _ := valid.Valid(category); !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.JsonResult(1001, "验证失败")
		}
		//3
		if _, err := category.Update(); err != nil {
			c.JsonResult(1001, "修改失败")
		}
		c.JsonResult(0, "修改成功")
	}

	result, _ := models.NewCategory().FindTree("")
	c.Data["category"] = result
	c.Data["vo"] = category
	c.TplName = c.ADMIN_TPL + "category/update.html"
}

func (c *CategoryController) Delete() {
	categoryModel := models.NewCategory()
	id, _ := c.GetInt("id")
	categoryModel.Id = id
	if err := categoryModel.Delete(); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

func (c *CategoryController) BatchDelete() {
	var ids []int
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		c.JsonResult(1001, "赋值失败")
	}

	categoryModel := models.NewCategory()
	if err := categoryModel.DelBatch(ids); err != nil {
		c.JsonResult(1001, "删除失败")
	}
	c.JsonResult(0, "删除成功")
}

`

func CreateController(controllerPath, tableName string) {
	controllerData := createModelBase(tableName)

	if err := os.MkdirAll(path.Clean(controllerPath), 777); err != nil {
		log.Println("控制器文件创建失败")
	}

	if err := ioutil.WriteFile(path.Join(controllerPath, fmt.Sprintf("%s.go", tableName)), []byte(controllerData), os.ModeType); err != nil {
		log.Println(err)
	}
}

func createModelBase(tableName string) (controllerData string) {
	controllerData = strings.Replace(controllerTpl, "Category", Hump(tableName, "max"), -1)
	controllerData = strings.Replace(controllerData, "category", Hump(tableName, "min"), -1)
	return
}
