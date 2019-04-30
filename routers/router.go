package routers

import (
	"free_cms/controllers/admin"
	"free_cms/controllers/admin/books"
	"free_cms/controllers/home"
	"github.com/astaxie/beego"
)

func init() {
	//前台模块
	beego.Router("/", &home.HomeController{},"get:Index")
	beego.Router("/list/:id", &home.HomeController{},"get:List")
	beego.Router("/books-list/:id", &home.HomeController{},"get:BooksList")
	beego.Router("/article/:cid/:id", &home.HomeController{},"get:Article")
	beego.Router("/search", &home.HomeController{},"get:Search")


	//后台模块
	beego.Router("/admin", &admin.MainController{},"get:Admin")
	beego.Router("/main", &admin.MainController{},"get:Main")
	beego.Router("/upload", &admin.MainController{},"*:Upload")
	beego.Router("/login", &admin.UserController{},"*:Login")
	beego.Router("/captcha", &admin.UserController{},"*:Captcha")
	beego.Router("/logout", &admin.UserController{},"get:Logout")
	beego.Router("/register", &admin.UserController{},"*:Register")

	//小说模块
	beego.Router("/books/index", &books.BooksController{},"*:Index")
	beego.Router("/books/create", &books.BooksController{},"*:Create")
	beego.Router("/books/update", &books.BooksController{},"*:Update")
	beego.Router("/books/delete", &books.BooksController{},"*:Delete")
	beego.Router("/books/batch-delete", &books.BooksController{},"*:BatchDelete")

	beego.Router("/books-preg/index", &books.BooksPregController{},"*:Index")
	beego.Router("/books-preg/create", &books.BooksPregController{},"*:Create")
	beego.Router("/books-preg/update", &books.BooksPregController{},"*:Update")
	beego.Router("/books-preg/delete", &books.BooksPregController{},"*:Delete")
	beego.Router("/books-preg/batch-delete", &books.BooksPregController{},"*:BatchDelete")

	beego.Router("/books-type/index", &books.BooksTypeController{},"*:Index")
	beego.Router("/books-type/create", &books.BooksTypeController{},"*:Create")
	beego.Router("/books-type/update", &books.BooksTypeController{},"*:Update")
	beego.Router("/books-type/delete", &books.BooksTypeController{},"*:Delete")
	beego.Router("/books-type/batch-delete", &books.BooksTypeController{},"*:BatchDelete")

	//
	beego.Router("/members/center", &admin.MembersController{},"*:Center")
	beego.Router("/members/center/add", &admin.MembersController{},"*:CenterAdd")
	beego.Router("/members/level", &admin.MembersController{},"*:Level")


	//系统模块
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
