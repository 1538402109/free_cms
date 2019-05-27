package models

import (
	"errors"
	"free_cms/pkg/util"
	"github.com/jinzhu/gorm"
)

type AdminLog struct {
	Model
	Route       string `json:"route" form:"route"`
	Description string `json:"description" form:"description"`
	UserId      int    `json:"user_id" form:"user_id"`
	Ip          int    `json:"ip" form:"ip"`
	Method      string `json:"method" form:"method"`
	User        User   `gorm:"ForeignKey:Id;AssociationForeignKey:UserId"`

	IpText        string `json:"ip_text" gorm:"-"`
	CreatedAtText string `json:"created_at_text" gorm:"-"`
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
	query := Db.Preload("User")
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(AdminLog{}).Count(&count)
	return
}

func (m *AdminLog) Create() (err error, newAttr *AdminLog) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *AdminLog) Update() (err error, newAttr AdminLog) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *AdminLog) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *AdminLog) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *AdminLog) FindById(id int) (adminLog AdminLog, err error) {
	err = Db.Where("id=?", id).First(&adminLog).Error
	return
}
