package home

import (
	"fmt"
	"free_cms/models"
	"free_cms/models/home"
	"free_cms/pkg/d"
	"github.com/gocolly/colly"
	"net/url"
	"strconv"
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
	_, sonType := home.NewBooksType().FindByPid(id2)
	c.Data["sonType"] = sonType
	//友情链接

	c.Data["booksType"] = booksType
}

func (c *HomeController) Index() {
	//文章
	_, books1 := home.NewBooks().FindByBooksType(6, 1, 6)
	_, books2 := home.NewBooks().FindByBooksType(7, 1, 6)
	_, books3 := home.NewBooks().FindByBooksType(8, 1, 6)
	_, books4 := home.NewBooks().FindByBooksType(9, 1, 6)
	_, books5 := home.NewBooks().FindByBooksType(10, 1, 6)
	_, books6 := home.NewBooks().FindByBooksType(11, 1, 6)
	_, books7 := home.NewBooks().FindByBooksType(12, 1, 6)
	_, books8 := home.NewBooks().FindByBooksType(-1, 2, 6) //封面
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
		limit, _ := c.GetInt("limit", 10)
		offset, _ := c.GetInt("offset", 0)
		key := c.GetString("key")
		_, tabData, count := home.NewBooks().FindOfBtTable(id2, offset, limit, key, bookType2)
		c.Data["json"] = d.TableJson(tabData, offset, limit, count)
		c.ServeJSON()
		return
	}

	category, _ := home.NewBooksType().FindById(id2)
	c.Data["category"] = category

	c.TplName = "home/list.html"
}

func (c *HomeController) BooksList() {
	cid := c.Ctx.Input.Param(":id")
	cid2, _ := strconv.Atoi(cid)

	list, _ := GetList(cid, "")

	bookMsg, _ := home.NewBooks().FindById(cid2)
	c.Data["bookMsg"] = bookMsg

	c.Data["list"] = list
	c.TplName = "home/book_list.html"
}

func (c *HomeController) Article() {
	id := c.Ctx.Input.Param(":id")
	cid := c.Ctx.Input.Param(":cid")

	_,article := GetList(cid,id)

	c.Data["article"] = article
	c.TplName = "home/article.html"
}

func (c *HomeController) Search() {

}

type BookList struct {
	Id int //列表页id
	Title   string //标题
	Author string//
	Describe string//
	BookLastAt string//跟新时间
	BookNewChapter string //最新章节
	Href    string //链接
	OldHref string //旧链接
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

func GetList(cid, id string) (bookList []BookList, bookArticle BookArticle) {
	c := colly.NewCollector()
	var lastTitle string
	var lastId int

	cidI, _ := strconv.Atoi(cid)
	idI, _ := strconv.Atoi(id)
	res, _ := home.NewBooks().FindById(cidI)

	pregOne, _ := home.NewBooksPreg().FindById(res.PregId)

	c.OnHTML(pregOne.ListABlock, func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			i++
			i2 := strconv.Itoa(i)
			bookList = append(bookList, BookList{Href: "/article/" + cid + "/" + i2, Title: element.Text, OldHref: element.Attr("href")})
			lastTitle = element.Text
			lastId = i

			if id != "" && i == idI {
				c.Visit(element.Request.AbsoluteURL(element.Attr("href"))+"?book_name="+lastTitle)
			}

			bookArticle.ListHref = "/books-list/"+cid
			//组装内容页的url
			if idI > i-1 { //已到最后一章
				next := strconv.Itoa((idI))
				prev := strconv.Itoa((idI - 1))
				bookArticle.PrevHref = "/article/" + cid + "/" + prev
				bookArticle.NextHref = "/article/" + cid + "/" + next
			} else if idI <= 1 { //已到第一张
				next := strconv.Itoa((idI + 1))
				prev := strconv.Itoa((idI))
				bookArticle.PrevHref = "/article/" + cid + "/" + prev
				bookArticle.NextHref = "/article/" + cid + "/" + next
			} else {
				next := strconv.Itoa((idI + 1))
				prev := strconv.Itoa((idI - 1))
				bookArticle.PrevHref = "/article/" + cid + "/" + prev
				bookArticle.NextHref = "/article/" + cid + "/" + next
			}
		})
		//如果标题和数据库中最新章节对不上，跟新数据库，最新章节，最近跟新时间
		if (res.BookNewChapter != lastTitle) {

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML(pregOne.ContentBlock, func(element *colly.HTMLElement) {
		bookArticle.Id = idI
		bookArticle.Cid = cidI

		v,_:=url.ParseQuery(element.Request.URL.RawQuery)
		bookArticle.Title = v["book_name"][0]

		element.DOM.Find(pregOne.ContentTextFilter).Remove()
		content, _ := element.DOM.Find(pregOne.ContentText).Html()
		bookArticle.Content = content

	})
	c.Visit(res.ListUrl)
	return
}
