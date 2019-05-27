package models

import (
	"errors"
)

type MembersGroup struct {
	Model
	Id                  int       `json:"id"             form:"id"             gorm:"default:''"`
	Name                string    `json:"name"           form:"name"           gorm:"default:''"`
	Rules               string    `json:"rules"          form:"rules"          gorm:"default:''"`
	Createtime          int       `json:"createtime"     form:"createtime"     gorm:"default:''"`
	Updatetime          int       `json:"updatetime"     form:"updatetime"     gorm:"default:''"`
	Status              string    `json:"status"         form:"status"         gorm:"default:''"`

}

func NewMembersGroup() (membersGroup *MembersGroup) {
	return &MembersGroup{}
}

func (m *MembersGroup) Pagination(offset, limit int, key string) (res []MembersGroup, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(MembersGroup{}).Count(&count)
	return
}

func (m *MembersGroup) Create() (err error, newAttr *MembersGroup) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *MembersGroup) Update() (err error, newAttr MembersGroup) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *MembersGroup) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *MembersGroup) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *MembersGroup) FindById(id int) (membersGroup MembersGroup, err error) {
	err = Db.Where("id=?", id).First(&membersGroup).Error
	return
}