package routers

import (
	admin "free_cms/backend/controllers"
	home "free_cms/frontend/controllers"
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
	beego.Router("/books/index", &admin.BooksController{},"*:Index")
	beego.Router("/books/create", &admin.BooksController{},"*:Create")
	beego.Router("/books/update", &admin.BooksController{},"*:Update")
	beego.Router("/books/delete", &admin.BooksController{},"*:Delete")
	beego.Router("/books/batch-delete", &admin.BooksController{},"*:BatchDelete")

	beego.Router("/books-preg/index", &admin.BooksPregController{},"*:Index")
	beego.Router("/books-preg/create", &admin.BooksPregController{},"*:Create")
	beego.Router("/books-preg/update", &admin.BooksPregController{},"*:Update")
	beego.Router("/books-preg/delete", &admin.BooksPregController{},"*:Delete")
	beego.Router("/books-preg/batch-delete", &admin.BooksPregController{},"*:BatchDelete")

	beego.Router("/books-type/index", &admin.BooksTypeController{},"*:Index")
	beego.Router("/books-type/create", &admin.BooksTypeController{},"*:Create")
	beego.Router("/books-type/update", &admin.BooksTypeController{},"*:Update")
	beego.Router("/books-type/delete", &admin.BooksTypeController{},"*:Delete")
	beego.Router("/books-type/batch-delete", &admin.BooksTypeController{},"*:BatchDelete")

	//会员模块
	beego.Router("/members/center", &admin.MembersController{},"*:Center")
	beego.Router("/members/center/add", &admin.MembersController{},"*:CenterAdd")
	beego.Router("/members/level", &admin.MembersController{},"*:Level")


	//系统模块
	beego.Router("/system/basic", &admin.ConfigController{},"*:Index")
	beego.Router("/system/logs", &admin.AdminLogController{},"*:Index")
	//友情链接
	beego.Router("/links/index", &admin.LinkController{},"*:Index")
	beego.Router("/links/create", &admin.LinkController{},"*:Create")
	beego.Router("/links/update", &admin.LinkController{},"*:Update")
	beego.Router("/links/delete", &admin.LinkController{},"*:Delete")
	beego.Router("/links/batch-delete", &admin.LinkController{},"*:BatchDelete")

	beego.Router("/system/icons", &admin.SystemController{},"*:Icons")

	//beego.Router("/gorm", &admin.GormController{},"*:Find")

	beego.Router("/document/demo1", &admin.DocumentController{},"*:Demo1")
	beego.Router("/document/demo2", &admin.DocumentController{},"*:Demo2")
	beego.Router("/document/demo3", &admin.DocumentController{},"*:Demo3")
}
