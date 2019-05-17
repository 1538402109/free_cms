package models

import (
	"errors"
	"free_cms/common/models"
)

type Link struct {
	models.Link
}

func NewLink() (link *Link) {
	return &Link{}
}

func (m *Link) Pagination(offset, limit int, key string) (res []Link, count int) {
	query := models.Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Link{}).Count(&count)
	return
}

func (m *Link) Create() (err error, newAttr *Link) {
	err = models.Db.Create(m).Error
	newAttr = m
	return
}

func (m *Link) Update() (err error, newAttr Link) {
	if m.Id > 0 {
		err = models.Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Link) Delete() (err error) {
	if m.Id > 0 {
		err = models.Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Link) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = models.Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Link) FindById(id int) (link Link, err error) {
	err = models.Db.Where("id=?", id).First(&link).Error
	return
}