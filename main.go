package main

import (
	"flag"
	"free_cms/fc"
	_ "free_cms/routers"
	"github.com/astaxie/beego"
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
	tableName := flag.String("t", "", "表名")
	modelPath := flag.String("m", "", "模型地址")
	controllerPath := flag.String("c", "", "控制器地址")
	viewPath := flag.String("v", "", "视图地址")
	flag.Parse()

	if *tableName != "" {
		fc.Fc(*tableName, *modelPath, *controllerPath, *viewPath)
		return
	}
	beego.ErrorHandler("404", page_not_found)
	beego.Run()
}
