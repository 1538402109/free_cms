package models

import (
	"errors"
	"free_cms/models"
	"github.com/jinzhu/gorm"
)

type BooksPreg struct {
	models.BooksPreg
}


func NewBooksPreg() (books *BooksPreg) {
	return &BooksPreg{}
}

func (m *BooksPreg) AfterFind(scope *gorm.Scope) (err error) {
	m.CreatedAtText = m.CreatedAt.Unix()
	return
}

func (m *BooksPreg) Pagination(offset, limit int, key string) (res []BooksPreg, count int) {
	query := models.Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(BooksPreg{}).Count(&count)
	return
}

func (m *BooksPreg) Create() (err error, newAttr *BooksPreg) {
	err = models.Db.Create(m).Error
	newAttr = m
	return
}

func (m *BooksPreg) Update() (err error, newAttr BooksPreg) {
	if m.Id > 0 {
		err = models.Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *BooksPreg) Delete() (err error) {
	if m.Id > 0 {
		err = models.Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksPreg)DelBath(ids []int)(err error){
	if len(ids)>0{
		err = models.Db.Where("id in (?)",ids).Delete(m).Error
	}else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksPreg)FindById(id int)(booksPreg BooksPreg,err error){
	err = models.Db.Where("id=?",id).First(&booksPreg).Error
	return
}