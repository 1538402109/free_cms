package models

import (
	"errors"
	"free_cms/common/models"
	"free_cms/pkg/util"
	"github.com/jinzhu/gorm"
)

type AdminLog struct {
	models.AdminLog
}

func NewAdminLog() (adminlog *AdminLog) {
	return &AdminLog{}
}

func (m *AdminLog) AfterFind(scope *gorm.Scope) (err error) {
	m.IpText = util.Long2ip(uint32(m.Ip))
	m.CreatedAtText = m.CreatedAt.Format("2006-05-04 15:02:01")
	return
}

func (m *AdminLog) Pagination(offset, limit int, key string) (res []AdminLog, count int) {
	query := models.Db.Preload("User")
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(AdminLog{}).Count(&count)
	return
}

func (m *AdminLog) Create() (err error, newAttr *AdminLog) {
	err = models.Db.Create(m).Error
	newAttr = m
	return
}

func (m *AdminLog) Update() (err error, newAttr AdminLog) {
	if m.Id > 0 {
		err = models.Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *AdminLog) Delete() (err error) {
	if m.Id > 0 {
		err = models.Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *AdminLog) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = models.Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *AdminLog) FindById(id int) (adminLog AdminLog, err error) {
	err = models.Db.Where("id=?", id).First(&adminLog).Error
	return
}
