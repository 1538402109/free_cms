package models

import (
	"errors"
	"free_cms/common/models"
	"github.com/jinzhu/gorm"
)

type Books struct {
	models.Books
}

func NewBooks() (books *Books) {
	return &Books{}
}
func (m *Books) AfterFind(scope *gorm.Scope) (err error) {
	m.CreatedAtText = m.CreatedAt.Unix()
	_, booksType := NewBooksType().FindColumn()
	m.BookTypeText = booksType[m.BookType]
	return
}

func (m *Books) Pagination(offset, limit int, key string) (res []Books, count int) {
	query := models.Db
	if key != "" {
		query = query.Where("book_name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Books{}).Count(&count)
	return
}

func (m *Books) Create() (err error, newAttr *Books) {
	err = models.Db.Create(m).Error
	newAttr = m
	return
}

func (m *Books) Update() (err error, newAttr Books) {
	if m.Id > 0 {
		err = models.Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Books) Delete() (err error) {
	if m.Id > 0 {
		err = models.Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Books) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = models.Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Books) FindById(id int) (books Books, err error) {
	err = models.Db.Where("id=?", id).First(&books).Error
	return
}
