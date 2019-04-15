package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Books struct {
	Model
	BookName       string    `json:"book_name"        form:"book_name"`
	BookType       int       `json:"book_type"        form:"book_type"`
	BookAuthor     string    `json:"book_author"      form:"book_author"`
	BookNewChapter string    `json:"book_new_chapter" form:"book_new_chapter"`
	BookDescribe   string    `json:"book_describe"    form:"book_describe"`
	BookStatus     int       `json:"book_status"      form:"book_status"`
	BookLastAt     time.Time `json:"book_last_at"     form:"book_last_at"`
	PregId         int       `json:"preg_id"          form:"preg_id"`
	ListUrl        string    `json:"list_url"         form:"list_url"`
	IsTop          int       `json:"is_top"           form:"is_top"`
	LastId         int       `json:"last_id"          form:"last_id"`
	SeoTitle       string    `json:"seo_title"        form:"seo_title"`
	SeoKeyword     string    `json:"seo_keyword"      form:"seo_keyword"`
	SeoDescribe    string    `json:"seo_describe"     form:"seo_describe"`

	CreatedAtText  int64  `json:"created_at_text"   gorm:"-"`
	BookLastAtText int64  `json:"book_last_at_text" gorm:"-"`
	BookTypeText   string `json:"book_type_text"    gorm:"-"`
}

var BookType = map[int]string{1: "玄幻", 2: "修真", 3: "穿越", 4: "网游", 5: "科幻", 6: "武侠", 7: "言情", 8: "都市"}

func NewBooks() (books *Books) {
	return &Books{}
}

func (b *Books) AfterFind(scope *gorm.Scope) (err error) {
	b.CreatedAtText = b.CreatedAt.Unix()
	b.BookLastAtText = b.BookLastAt.Unix()
	b.BookTypeText = BookType[b.BookType]
	return
}

func (b *Books) BeforeCreate(scope *gorm.Scope) (err error) {
	b.BookLastAt = time.Now()
	return
}

func (b *Books) Pagination(offset, limit int, key string) (res []Books, count int) {
	query := Db
	if key != "" {
		query = query.Where("book_name like ?", fmt.Sprintf("%s%%",key))
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Books{}).Count(&count)
	return
}

func (b *Books) Create() (err error, newAttr *Books) {
	err = Db.Create(b).Error
	newAttr = b
	return
}

func (b *Books) Update() (err error, newAttr Books) {
	if b.ID > 0 {
		err = Db.Model(&newAttr).Where("id=?", b.ID).Update(b).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (b *Books) Delete() (err error) {
	if b.ID > 0 {
		err = Db.Delete(b).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}
