package models

import "errors"

type Nav struct {
	Model
	Id                  int       `json:"id"             form:"id"             gorm:"default:''"`
	Pid                 int       `json:"pid"            form:"pid"            gorm:"default:'0'"`
	Name                string    `json:"name"           form:"name"           gorm:"default:''"`
	Mca                 string    `json:"mca"            form:"mca"            gorm:"default:''"`
	Ico                 string    `json:"ico"            form:"ico"            gorm:"default:''"`
	OrderNumber         int       `json:"order_number"   form:"order_number"   gorm:"default:''"`
	
}


func NewNav() (nav *Nav) {
	return &Nav{}
}

func (m *Nav) Pagination(offset, limit int, key string) (res []Nav, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Nav{}).Count(&count)
	return
}

func (m *Nav) Create() (err error, newAttr *Nav) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *Nav) Update() (err error, newAttr Nav) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Nav) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Nav) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Nav) FindById(id int) (nav Nav, err error) {
	err = Db.Where("id=?", id).First(&nav).Error
	return
}

