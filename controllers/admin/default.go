package admin

import (
	"crypto/md5"
	"fmt"
	"free_cms/models"
	"free_cms/pkg/d"
	"math/rand"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
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

func (c *MainController)Upload(){
	f,h,_ := c.GetFile("file")
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":true,
		".jpeg":true,
		".png":true,
	}
	if _,ok:=AllowExtMap[ext];!ok{
		c.Ctx.WriteString( "后缀名不符合上传要求" )
		return
	}
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("20060102/")
	err := os.MkdirAll( uploadDir , 777)
	if err != nil {
		c.Ctx.WriteString( fmt.Sprintf("%v",err) )
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )

	fileName := fmt.Sprintf("%x",hashName) + ext
	//this.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	//保存文件到指定的位置
	if err := c.SaveToFile("file", fpath);err !=nil{
		c.JsonResult(2000,"error")
	}
	c.JsonResult(200,"success","/"+fpath)
}