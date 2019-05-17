package models

import (
	"errors"
	"free_cms/common/models"
)

type MembersGroup struct {
	models.MembersGroup
}

func NewMembersGroup() (membersGroup *MembersGroup) {
	return &MembersGroup{}
}

func (m *MembersGroup) Pagination(offset, limit int, key string) (res []MembersGroup, count int) {
	query := models.Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(MembersGroup{}).Count(&count)
	return
}

func (m *MembersGroup) Create() (err error, newAttr *MembersGroup) {
	err = models.Db.Create(m).Error
	newAttr = m
	return
}

func (m *MembersGroup) Update() (err error, newAttr MembersGroup) {
	if m.Id > 0 {
		err = models.Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *MembersGroup) Delete() (err error) {
	if m.Id > 0 {
		err = models.Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *MembersGroup) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = models.Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}
