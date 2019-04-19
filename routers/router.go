package routers

import (
	"free_cms/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin", &admin.MainController{},"get:Admin")
	beego.Router("/main", &admin.MainController{},"get:Main")
	beego.Router("/login", &admin.UserController{},"*:Login")
	beego.Router("/captcha", &admin.UserController{},"*:Captcha")
	beego.Router("/logout", &admin.UserController{},"get:Logout")
	beego.Router("/register", &admin.UserController{},"*:Register")

	beego.Router("/books/index", &admin.BooksController{},"*:Index")
	beego.Router("/books/create", &admin.BooksController{},"*:Create")
	beego.Router("/books/update", &admin.BooksController{},"*:Update")
	beego.Router("/books/delete", &admin.BooksController{},"*:Delete")

	beego.Router("/books-preg/index", &admin.BooksPregController{},"*:Index")
	beego.Router("/books-preg/create", &admin.BooksPregController{},"*:Create")
	beego.Router("/books-preg/update", &admin.BooksPregController{},"*:Update")
	beego.Router("/books-preg/delete", &admin.BooksPregController{},"*:Delete")

	beego.Router("/members/center", &admin.MembersController{},"*:Center")
	beego.Router("/members/center/add", &admin.MembersController{},"*:CenterAdd")
	beego.Router("/members/level", &admin.MembersController{},"*:Level")



	beego.Router("/system/basic", &admin.SystemController{},"*:Basic")
	beego.Router("/system/logs", &admin.SystemController{},"*:Logs")
	beego.Router("/system/links", &admin.SystemController{},"*:Links")
	beego.Router("/system/links/add", &admin.SystemController{},"*:LinksAdd")
	beego.Router("/system/icons", &admin.SystemController{},"*:Icons")

	//beego.Router("/gorm", &admin.GormController{},"*:Find")

	beego.Router("/document/demo1", &admin.DocumentController{},"*:Demo1")
	beego.Router("/document/demo2", &admin.DocumentController{},"*:Demo2")
	beego.Router("/document/demo3", &admin.DocumentController{},"*:Demo3")
}
