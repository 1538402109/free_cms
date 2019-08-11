package home

import (
	"bytes"
	"free_cms/controllers"
	"free_cms/models"
	"github.com/jinzhu/gorm"
	"html/template"
	"math"
	"strings"
)

type BlogController struct {
	controllers.HomeBaseController
}

func (b *BlogController) Index() {
	blogModel := models.NewPost()
	page, _ := b.GetInt("page", 1)
	keyword := b.GetString("keyword")
	blogList, total := blogModel.Pagination((page-1)*10, 10, keyword)

	var getOption = make(map[string]string)
	getOption["keyword"] = keyword
	display := Page(page, 10, int64(total), getOption)

	var clickRanking []models.Post
	models.Db.Offset(0).Limit(6).Order("post_hits desc").Find(&clickRanking)

	b.Data["clickRanking"] = clickRanking
	b.Data["blogList"] = blogList
	b.Data["page"] = display
	b.TplName = "home/blog/index.html"
}

func (b *BlogController) Article() {
	id, _ := b.GetInt("id")
	article, _ := models.NewPost().FindById(id)

	var clickRanking []models.Post
	models.Db.Offset(0).Limit(6).Order("post_hits desc").Find(&clickRanking)
	b.Data["clickRanking"] = clickRanking
	//自增1
	models.Db.Model(&models.Post{}).Where("id=?", id).UpdateColumn("post_hits", gorm.Expr("post_hits+?", 1))

	b.Data["article"] = article
	b.TplName = "home/blog/article.html"
}

func Page(page, prepage int, nums int64, options map[string]string) (pageStr string) {
	paginator := Paginator(page, prepage, nums)
	var getStr []string
	var s string
	if len(options) > 0 {
		for k, v := range options {
			getStr = append(getStr, k+"="+v)
		}
		s = strings.Join(getStr, "&") + "&"
	}
	pageTpl := `<div class="am-cf"> 共{{.total}}条记录 共记{{.totalpages}} 页 当前页 {{.currpage}}
    <div class="am-fr">
        <ul class="am-pagination">
            <li class=""><a href="?{URL}page={{.firstpage}}">«</a></li>
            {{range $index,$page := .pages}}
                <li {{if eq $.currpage $page }}class="am-active"{{end}}>
                    <a href="?{URL}page={{$page}}">{{$page}}</a></li>
            {{end}}
            <li><a href="?{URL}page={{.lastpage}}">»</a></li>
        </ul>
    </div>
</div>`
	pageTpl = strings.Replace(pageTpl, "{URL}", s, -1)
	tpl, _ := template.New("page").Parse(pageTpl)
	buf := new(bytes.Buffer)
	tpl.Execute(buf, paginator)
	return buf.String()
}

//@page当前页 @prepage 每页记录条数 @nums总条数
func Paginator(page, prepage int, nums int64) map[string]interface{} {
	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
	}

	firstpage = int(math.Max(float64(1), float64(page-1)))
	lastpage = int(math.Min(float64(totalpages), float64(page+1)))

	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	paginatorMap["total"] = nums
	return paginatorMap
}
