package main

import (
	"free_cms/fc"
	_ "free_cms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"html/template"
	"net/http"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func main() {
	//gii
	if b, err := beego.AppConfig.Bool("gii"); b {
		if err == nil {
			fc.Run() //开启gii
		}
		//return
	}

	//log
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//输出文件名，行号
	logs.EnableFuncCallDepth(true)
	//异步log
	logs.Async(1e3)
	//404
	beego.ErrorHandler("404", page_not_found)
	//run
	beego.Run()
}
