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

func NewBooks() (books *Books) {
	return &Books{}
}

func (m *Books) AfterFind(scope *gorm.Scope) (err error) {
	m.CreatedAtText = m.CreatedAt.Unix()
	m.BookLastAtText = m.BookLastAt.Unix()
	_,booksType:=NewBooksType().FindColumn()
	m.BookTypeText = booksType[m.BookType]
	return
}

func (m *Books) BeforeCreate(scope *gorm.Scope) (err error) {
	m.BookLastAt = time.Now()
	return
}

func (m *Books) Pagination(offset, limit int, key string) (res []Books, count int) {
	query := Db
	if key != "" {
		query = query.Where("book_name like ?", fmt.Sprintf("%s%%", key))
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Books{}).Count(&count)
	return
}

func (m *Books) Create() (err error, newAttr *Books) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *Books) Update() (err error, newAttr Books) {
	if m.Id > 0 {
		err = Db.Model(&newAttr).Where("id=?", m.Id).Update(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Books) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Books) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}
