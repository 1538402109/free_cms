package home

import (
	"fmt"
	"free_cms/models"
	"free_cms/models/home"
	"free_cms/pkg/d"
	"free_cms/pkg/util"
	"github.com/gocolly/colly"
	"net/url"
	"strconv"
	"strings"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Prepare() {
	//获取导航
	var booksType []models.BooksType
	models.Db.Where("pid=? and is_nav=?", 0, 1).Find(&booksType)
	//子导航
	id := c.Ctx.Input.Param(":id")
	id2, _ := strconv.Atoi(id)
	sonType, _ := home.NewBooksType().FindByPid(id2)
	c.Data["sonType"] = sonType
	//友情链接

	c.Data["booksType"] = booksType
}

//首页
func (c *HomeController) Index() {
	//文章
	books1, _ := home.NewBooks().FindByBooksType(6, 1, 6)
	books2, _ := home.NewBooks().FindByBooksType(7, 1, 6)
	books3, _ := home.NewBooks().FindByBooksType(8, 1, 6)
	books4, _ := home.NewBooks().FindByBooksType(9, 1, 6)
	books5, _ := home.NewBooks().FindByBooksType(10, 1, 6)
	books6, _ := home.NewBooks().FindByBooksType(11, 1, 6)
	books7, _ := home.NewBooks().FindByBooksType(12, 1, 6)
	books8, _ := home.NewBooks().FindByBooksType(-1, 2, 6) //封面
	c.Data["books1"] = books1
	c.Data["books2"] = books2
	c.Data["books3"] = books3
	c.Data["books4"] = books4
	c.Data["books5"] = books5
	c.Data["books6"] = books6
	c.Data["books7"] = books7
	c.Data["books8"] = books8

	c.TplName = "home/index.html"
}

//小说列表
func (c *HomeController) List() {
	id := c.Ctx.Input.Param(":id")
	id2, _ := strconv.Atoi(id)

	if c.IsAjax() {
		if id2 == 1 {
			id2 = 6
		}
		var bookType2 int
		if id2 == 3 {
			bookType2 = 1
			id2 = -1
		}
		if id2 == 4 {
			bookType2 = 0
			id2 = -1
		}
		if id2 == 5{
			//书架
		}
		limit, _ := c.GetInt("limit", 10)
		offset, _ := c.GetInt("offset", 0)
		key := c.GetString("key")
		tabData, count, _ := home.NewBooks().FindOfBtTable(id2, offset, limit, key, bookType2)
		c.Data["json"] = d.TableJson(tabData, offset, limit, count)
		c.ServeJSON()
		return
	}

	category, _ := home.NewBooksType().FindById(id2)
	c.Data["category"] = category

	c.TplName = "home/list.html"
}

//小说章节列表
func (c *HomeController) BooksList() {
	cid := c.Ctx.Input.Param(":id")
	cid2, _ := strconv.Atoi(cid)

	list, _ := GetList(cid, "")

	bookMsg, _ := home.NewBooks().FindById(cid2)
	c.Data["bookMsg"] = bookMsg

	c.Data["list"] = list
	c.TplName = "home/book_list.html"
}

//内容
func (c *HomeController) Article() {
	id := c.Ctx.Input.Param(":id")
	cid := c.Ctx.Input.Param(":cid")

	_, article := GetList(cid, id)

	c.Data["article"] = article
	c.TplName = "home/article.html"
}

func (c *HomeController) Search() {

}

type BookListMsg struct {
	Title   string //标题
	Href    string //链接
	OldHref string //旧链接，没有转换的链接，会链接到采集地址
}
type BookList struct {
	Id             int    //列表页id
	Author         string //
	Describe       string //
	BookLastAt     string //跟新时间
	BookNewChapter string //最新章节
	BookAuthor     string
	BookDescribe   string
	BookImg        string
	BookListMsgs   []BookListMsg
}
type BookArticle struct {
	Id       int
	Cid      int
	Title    string
	Content  string
	PrevHref string
	NextHref string
	ListHref string
}

func GetList(cid, id string) (bookList BookList, bookArticle BookArticle) {
	c := colly.NewCollector()
	c2 := colly.NewCollector()
	var lastTitle string
	var lastId int
	var char string

	cidI, _ := strconv.Atoi(cid)
	idI, _ := strconv.Atoi(id)

	res, _ := home.NewBooks().FindById(cidI)
	pregOne, _ := home.NewBooksPreg().FindById(res.PregId)

	c.OnHTML("meta[http-equiv='Content-Type']", func(element *colly.HTMLElement) {
		if strings.Index(element.Attr("content"), "gbk") > 0 {
			char = "gbk"
		}
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		e.ForEach("#list a[href]", func(i int, element *colly.HTMLElement) {
			title := element.Text
			if char == "gbk" {
				title = util.GbkToUtf8(element.Text)
			}

			i2 := strconv.Itoa(i + 1)
			bookList.BookListMsgs = append(bookList.BookListMsgs, BookListMsg{Href: "/article/" + cid + "/" + i2, Title: title, OldHref: element.Request.AbsoluteURL(element.Attr("href"))})

			lastTitle = title
			lastId = i

			bookArticle.ListHref = "/books-list/" + cid
			//组装内容页的url
			if idI > i { //已到最后一章
				prev := strconv.Itoa((idI - 1))
				bookArticle.PrevHref = "/article/" + cid + "/" + prev
				bookArticle.NextHref = "/books-list/" + cid
			} else if idI <= 1 { //已到第一张
				next := strconv.Itoa((idI + 1))
				bookArticle.PrevHref = "/books-list/" + cid
				bookArticle.NextHref = "/article/" + cid + "/" + next
			} else {
				next := strconv.Itoa((idI + 1))
				prev := strconv.Itoa((idI - 1))
				bookArticle.PrevHref = "/article/" + cid + "/" + prev
				bookArticle.NextHref = "/article/" + cid + "/" + next
			}

			if i+1 == idI {
				c2.Visit(element.Request.AbsoluteURL(element.Attr("href")) + "?book_name=" + lastTitle)
			}
		})

		//跟新最新章节，最近跟新时间,作者，描述，封面图片
		bookList.BookNewChapter = lastTitle

		bookLastTime := e.DOM.Find("#maininfo p:nth-of-type(3)").Text()
		bookAuthor := e.DOM.Find("#maininfo p:nth-of-type(1) a").Text()
		bookDescribe := e.DOM.Find("#intro p").Text()
		bookImg, _ := e.DOM.Find("#fmimg img").Attr("src")
		if char == "gbk" {
			bookLastTime = util.GbkToUtf8(bookLastTime)
			bookAuthor = util.GbkToUtf8(bookAuthor)
			bookDescribe = util.GbkToUtf8(bookDescribe)
			bookImg = util.GbkToUtf8(bookImg)
		}
		//过滤
		bookLastTime = strings.Replace(bookLastTime, "最后更新：", "", -1)

		bookList.BookLastAt = bookLastTime
		if res.BookAuthor == "" {
			models.Db.Model(&models.Books{}).Where("id=?", cid).Update("book_author", bookAuthor)
		}
		bookList.BookAuthor = bookAuthor
		bookList.BookDescribe = bookDescribe
		bookList.BookImg = bookImg
		fmt.Println(bookList.BookAuthor, bookList.BookImg)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c2.OnHTML("html", func(element *colly.HTMLElement) {
		bookArticle.Id = idI
		bookArticle.Cid = cidI

		v, _ := url.ParseQuery(element.Request.URL.RawQuery)
		bookArticle.Title = v["book_name"][0]

		element.DOM.Find(pregOne.ContentTextFilter).Remove()
		content, _ := element.DOM.Find(pregOne.ContentText).Html()
		if char == "gbk" {
			content = util.GbkToUtf8(content)
			//todo &nbsp;变 聽
			content = strings.Replace(content, "聽", "&nbsp;", -1)
		}
		bookArticle.Content = content
	})
	c.Visit(res.ListUrl)
	return
}
