package models

import (
	"fmt"
	"errors"
	"github.com/jinzhu/gorm"
)

type BooksPreg struct {
	Model
	Name       string    `json:"name"                           form:"name"`
	ListMsgBlock string `json:"list_msg_block"                  form:"list_msg_block"`
	ListTitle string `json:"list_title"                         form:"list_title"`
	ListAuthor string `json:"list_author"                       form:"list_author"`
	ListAuthorFilter string `json:"list_author_filter"          form:"list_author_filter"`
	ListNewChapter string `json:"list_new_chapter"              form:"list_new_chapter"`
	ListNewChapterFilter string `json:"list_new_chapter_filter" form:"list_new_chapter_filter"`
	ListDescribe string `json:"list_describe"                   form:"list_describe"`
	ListDescribeFilter string `json:"list_describe_filter"      form:"list_describe_filter"`
	ListABlock string `json:"list_a_block"                      form:"list_a_block"`
	ContentBlock string `json:"content_block"                   form:"content_block"`
	ContentTitleFilter string `json:"content_title_filter"      form:"content_title_filter"`
	ContentText string `json:"content_text"                     form:"content_text"`
	ContentTextFilter string `json:"content_text_filter"        form:"content_text_filter"`

	CreatedAtText int64 `json:"create_at_text" gorm:"-"`
}


func NewBooksPreg() (books *BooksPreg) {
	return &BooksPreg{}
}

func (m *BooksPreg) AfterFind(scope *gorm.Scope) (err error) {
	m.CreatedAtText = m.CreatedAt.Unix()
	return
}

func (m *BooksPreg) Pagination(offset, limit int, key string) (res []BooksPreg, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", fmt.Sprintf("%s%%",key))
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(BooksPreg{}).Count(&count)
	return
}

func (m *BooksPreg) Create() (err error, newAttr *BooksPreg) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *BooksPreg) Update() (err error, newAttr BooksPreg) {
	if m.Id > 0 {
		err = Db.Model(&newAttr).Where("id=?", m.Id).Update(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksPreg) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksPreg)DelBath(ids []int)(err error){
	if len(ids)>0{
		err = Db.Where("id in (?)",ids).Delete(m).Error
	}else {
		err = errors.New("id参数错误")
	}
	return
}