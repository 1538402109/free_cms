package models

import (
	"errors"
	"free_cms/common/models"
)

type Category struct {
	models.Category
}

func NewCategory() (category *Category) {
	return &Category{}
}

func (m *Category) Pagination(offset, limit int, key string) (res []Category, count int) {
	query := models.Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Category{}).Count(&count)
	return
}

func (m *Category) Create() (err error, newAttr *Category) {
	err = models.Db.Create(m).Error
	newAttr = m
	return
}

func (m *Category) Update() (err error, newAttr Category) {
	if m.Id > 0 {
		err = models.Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Category) Delete() (err error) {
	if m.Id > 0 {
		err = models.Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Category) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = models.Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}
