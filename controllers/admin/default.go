package admin

import (
	"free_cms/models"
	"free_cms/pkg/d"
	"runtime"
	"strconv"
)

type MainController struct {
	BaseController
}

func (c *MainController) Admin() {
/*	var bookstype []models.BooksType
	models.Db.Find(&bookstype)
	_,v:=models.NewBooksType().FindTree(bookstype)
	c.Data["json"] =v
	c.ServeJSON()*/
	c.TplName = "admin/default/index.html"
}

func (c *MainController) Main() {
	if c.Ctx.Input.IsAjax() {
		var system = make(map[string]string)
		system["cms_version"] = "v1.0"
		system["cms_author"] = "stone"
		system["blog"] = "http://www.yyq6.cn"
		system["server"] = runtime.GOARCH + " " + runtime.GOOS //系统
		system["go_version"] = runtime.Version() //go版本
		system["numcpu"] = strconv.Itoa(runtime.NumCPU()) //逻辑cpu
		system["numgoroutine"] = strconv.Itoa(runtime.NumGoroutine())//当前go携程数
		system["mysql_version"] = models.GetMysqlMsg()["version"]
		c.Data["json"] = d.ReturnJson(200, "ok", system)
		c.ServeJSON()
	}
	c.TplName = "admin/default/main.html"
}
