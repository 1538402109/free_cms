package home

import (
	"fmt"
	"free_cms/controllers"
	"free_cms/models"
	"free_cms/pkg/d"
	"free_cms/pkg/str"
	"github.com/gocolly/colly"
	"net/url"
	"strconv"
	"strings"
)

type HomeController struct {
	controllers.HomeBaseController
}

func (c *HomeController) Prepare() {
	//获取导航
	var booksType []models.BooksType
	models.Db.Where("pid=? and is_nav=?", 0, 1).Find(&booksType)
	//子导航
	id := c.Ctx.Input.Param(":id")
	id2, _ := strconv.Atoi(id)
	sonType, _ := models.NewBooksType().FindByPid(id2)
	c.Data["sonType"] = sonType
	//友情链接

	//seo
	var seoIndex = make(map[string]string)
	conf, _ := models.NewConfig().FindByName("seo_title")
	seoIndex["seo_title"] = conf.Value
	conf, _ = models.NewConfig().FindByName("seo_keywords")
	seoIndex["seo_keywords"] = conf.Value
	conf, _ = models.NewConfig().FindByName("seo_description")
	seoIndex["seo_description"] = conf.Value

	c.Data["seoIndex"] = seoIndex
	c.Data["booksType"] = booksType
}

//首页
func (c *HomeController) Index() {
	//文章
	books1, _ := models.NewBooks().FindByBooksType(6, 1, 6)
	books2, _ := models.NewBooks().FindByBooksType(7, 1, 6)
	books3, _ := models.NewBooks().FindByBooksType(8, 1, 6)
	books4, _ := models.NewBooks().FindByBooksType(9, 1, 6)
	books5, _ := models.NewBooks().FindByBooksType(10, 1, 6)
	books6, _ := models.NewBooks().FindByBooksType(11, 1, 6)
	books7, _ := models.NewBooks().FindByBooksType(12, 1, 6)
	books8, _ := models.NewBooks().FindByBooksType(-1, 2, 6) //封面
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
		if id2 == 5 {
			//书架
		}
		limit, _ := c.GetInt("limit", 10)
		offset, _ := c.GetInt("offset", 0)
		key := c.GetString("key")
		tabData, count, _ := models.NewBooks().FindOfBtTable(id2, offset, limit, key, bookType2)
		c.Data["json"] = d.TableJson(tabData, offset, limit, count)
		c.ServeJSON()
		return
	}

	category, _ := models.NewBooksType().FindById(id2)
	c.Data["category"] = category

	c.TplName = "home/list.html"
}

//小说章节列表
func (c *HomeController) BooksList() {
	cid := c.Ctx.Input.Param(":id")
	cid2, _ := strconv.Atoi(cid)

	list, _ := GetList(cid, "")

	bookMsg, _ := models.NewBooks().FindById(cid2)
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
	key := c.GetString("key")
	if c.IsAjax() {
		limit, _ := c.GetInt("limit", 10)
		offset, _ := c.GetInt("offset", 0)
		tabData, count, _ := models.NewBooks().FindOfBtTable(-1, offset, limit, key, -1)
		c.Data["json"] = d.TableJson(tabData, offset, limit, count)
		c.ServeJSON()
	}
	c.Data["key"] = key
	c.TplName = "home/search.html"
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

	res, _ := models.NewBooks().FindById(cidI)
	pregOne, _ := models.NewBooksPreg().FindById(res.PregId)

	c.OnHTML("meta[http-equiv='Content-Type']", func(element *colly.HTMLElement) {
		if strings.Index(element.Attr("content"), "gbk") > 0 {
			char = "gbk"
		}
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		e.ForEach(pregOne.ListA, func(i int, element *colly.HTMLElement) {
			title := element.Text
			if char == "gbk" {
				title = str.GbkToUtf8(element.Text)
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
		bookLastTime := e.DOM.Find(pregOne.ListMsgLastTime).Text()
		bookAuthor := e.DOM.Find(pregOne.ListAuthor).Text()
		bookDescribe := e.DOM.Find(pregOne.ListDescribe).Text()
		bookImg, _ := e.DOM.Find(pregOne.ListMsgImg).Attr("src")
		if char == "gbk" {
			bookLastTime = str.GbkToUtf8(bookLastTime)
			bookAuthor = str.GbkToUtf8(bookAuthor)
			bookDescribe = str.GbkToUtf8(bookDescribe)
			bookImg = str.GbkToUtf8(bookImg)
		}
		//过滤
		bookLastTime = strings.Replace(bookLastTime, "最后更新：", "", -1)

		//更新
		if res.BookLastAt != bookLastTime {
			models.Db.Model(&models.Books{}).Where("id=?", res.Id).Update("book_last_at", bookLastTime)
		}
		if res.BookAuthor == "" {
			models.Db.Model(&models.Books{}).Where("id=?", cid).Update("book_author", bookAuthor)
		}
		if res.BookNewChapter != lastTitle {
			models.Db.Model(&models.Books{}).Where("id=?", res.Id).Update("book_new_chapter", lastTitle)
		}

		bookList.BookLastAt = bookLastTime
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
			content = str.GbkToUtf8(content)
			//todo &nbsp;变 聽
			content = strings.Replace(content, "聽", "&nbsp;", -1)
		}
		bookArticle.Content = content
	})
	c.Visit(res.ListUrl)
	return
}
