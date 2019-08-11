package routers

import (
	"free_cms/controllers/admin_log"
	"free_cms/controllers/blog"
	"free_cms/controllers/books"
	"free_cms/controllers/books_preg"
	"free_cms/controllers/books_type"
	"free_cms/controllers/category"
	"free_cms/controllers/config"
	"free_cms/controllers/document"
	"free_cms/controllers/home"
	"free_cms/controllers/index"
	"free_cms/controllers/link"
	"free_cms/controllers/members"
	"free_cms/controllers/system"
	"free_cms/controllers/user"
	"github.com/astaxie/beego"
)

func init() {
	//前台模块
	beego.Router("/book", &home.HomeController{}, "get:Index")
	beego.Router("/book/list/:id", &home.HomeController{}, "get:List")
	beego.Router("/book/books-list/:id", &home.HomeController{}, "get:BooksList")
	beego.Router("/book/article/:cid/:id", &home.HomeController{}, "get:Article")
	beego.Router("/book/search", &home.HomeController{}, "get:Search")
	beego.Router("/book/shujia", &home.HomeController{}, "get:Shujia")

	beego.Router("/", &home.BlogController{}, "*:Index")
	beego.Router("/article", &home.BlogController{}, "*:Article")

	//后台模块
	beego.Router("/hhxxttxs", &index.IndexController{}, "get:Admin")
	beego.Router("/main", &index.IndexController{}, "get:Main")
	beego.Router("/upload", &index.IndexController{}, "*:Upload")
	beego.Router("/ueditor-upload", &index.IndexController{}, "*:UeditorUpload")
	beego.Router("/login", &user.UserController{}, "*:Login")
	beego.Router("/captcha", &user.UserController{}, "*:Captcha")
	beego.Router("/logout", &user.UserController{}, "get:Logout")
	beego.Router("/register", &user.UserController{}, "*:Register")

	beego.Router("/category/index", &category.CategoryController{}, "*:Index")
	beego.Router("/category/create", &category.CategoryController{}, "*:Create")
	beego.Router("/category/update", &category.CategoryController{}, "*:Update")
	beego.Router("/category/delete", &category.CategoryController{}, "*:Delete")
	beego.Router("/category/batch-delete", &category.CategoryController{}, "*:BatchDelete")

	//小说模块
	beego.Router("/books/index", &books.BooksController{}, "*:Index")
	beego.Router("/books/create", &books.BooksController{}, "*:Create")
	beego.Router("/books/update", &books.BooksController{}, "*:Update")
	beego.Router("/books/delete", &books.BooksController{}, "*:Delete")
	beego.Router("/books/batch-delete", &books.BooksController{}, "*:BatchDelete")

	beego.Router("/books-preg/index", &books_preg.BooksPregController{}, "*:Index")
	beego.Router("/books-preg/create", &books_preg.BooksPregController{}, "*:Create")
	beego.Router("/books-preg/update", &books_preg.BooksPregController{}, "*:Update")
	beego.Router("/books-preg/delete", &books_preg.BooksPregController{}, "*:Delete")
	beego.Router("/books-preg/batch-delete", &books_preg.BooksPregController{}, "*:BatchDelete")

	beego.Router("/books-type/index", &books_type.BooksTypeController{}, "*:Index")
	beego.Router("/books-type/create", &books_type.BooksTypeController{}, "*:Create")
	beego.Router("/books-type/update", &books_type.BooksTypeController{}, "*:Update")
	beego.Router("/books-type/delete", &books_type.BooksTypeController{}, "*:Delete")
	beego.Router("/books-type/batch-delete", &books_type.BooksTypeController{}, "*:BatchDelete")

	//会员模块
	beego.Router("/members/center", &members.MembersController{}, "*:Center")
	beego.Router("/members/center/add", &members.MembersController{}, "*:CenterAdd")
	beego.Router("/members/level", &members.MembersController{}, "*:Level")

	//系统模块
	beego.Router("/system/basic", &config.ConfigController{}, "*:Index")
	beego.Router("/system/logs", &admin_log.AdminLogController{}, "*:Index")
	//友情链接
	beego.Router("/links/index", &link.LinkController{}, "*:Index")
	beego.Router("/links/create", &link.LinkController{}, "*:Create")
	beego.Router("/links/update", &link.LinkController{}, "*:Update")
	beego.Router("/links/delete", &link.LinkController{}, "*:Delete")
	beego.Router("/links/batch-delete", &link.LinkController{}, "*:BatchDelete")

	beego.Router("/system/icons", &system.SystemController{}, "*:Icons")

	//beego.Router("/gorm", &admin.GormController{},"*:Find")

	beego.Router("/document/demo1", &document.DocumentController{}, "*:Demo1")
	beego.Router("/document/demo2", &document.DocumentController{}, "*:Demo2")
	beego.Router("/document/demo3", &document.DocumentController{}, "*:Demo3")

	//博客
	beego.Router("/blog/index", &blog.PostController{}, "*:Index")
	beego.Router("/blog/update", &blog.PostController{}, "*:Update")
	beego.Router("/blog/create", &blog.PostController{}, "*:Create")
	beego.Router("/blog/delete", &blog.PostController{}, "*:Delete")
	beego.Router("/blog/batch-delete", &blog.PostController{}, "*:BatchDelete")

}
