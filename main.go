package main

import (
	"book/models"
	_ "free_cms/routers"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

func init(){
	models.Init()
}

func page_not_found(rw http.ResponseWriter, r *http.Request){
	t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/404.html")
	data :=make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func main() {
	beego.ErrorHandler("404",page_not_found)
	beego.Run()
}


