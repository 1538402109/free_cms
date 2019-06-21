package models

import (
	"errors"
)

type Link struct {
	Model
	Name        string `json:"name" form:"name" valid:"Required"`
	Description string `json:"description" form:"description"`
	Status      int    `json:"status" form:"status"`
	Url         string `json:"url" form:"url"  valid:"Required"`
	Image       string `json:"image" form:"image"`
	Target      string `json:"target" form:"target"`
	ListOrder   int    `json:"list_order" form:"list_order" gorm:"default:'1000'"`
}

func NewLink() (link *Link) {
	return &Link{}
}

func (m *Link) Pagination(offset, limit int, key string) (res []Link, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("list_order desc").Find(&res)
	query.Model(Link{}).Count(&count)
	return
}

func (m *Link) Create() (err error, newAttr *Link) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *Link) Update() (err error, newAttr Link) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Link) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Link) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Link) FindById(id int) (link Link, err error) {
	err = Db.Where("id=?", id).First(&link).Error
	return
}
