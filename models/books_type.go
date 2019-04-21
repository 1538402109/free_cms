package models

import (
	"errors"
	"fmt"
)

type BooksType struct {
	Model
	Name string `json:"name" form:"name"`
}

func NewBooksType() (books *BooksType) {
	return &BooksType{}
}

func (m *BooksType) Pagination(offset, limit int, key string) (res []BooksType, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", fmt.Sprintf("%s%%", key))
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(BooksType{}).Count(&count)
	return
}

func (m *BooksType) Create() (err error, newAttr *BooksType) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *BooksType) Update() (err error, newAttr BooksType) {
	if m.Id > 0 {
		err = Db.Model(&newAttr).Where("id=?", m.Id).Update(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksType) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksType) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksType) FindColumn() (err error, booksType map[int]string) {
	booksType = make(map[int]string)
	var booksTypes []BooksType
	err = Db.Find(&booksTypes).Error
	for _,v:= range booksTypes{
		booksType[v.Id] = v.Name
	}
	return
}
