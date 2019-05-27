package models

import (
	"errors"
)

type Config struct {
	Model
	Name string `json:"name" form:"name"`
	Group string `json:"group" form:"group"`
	Title string `json:"title" form:"title"`
	Tip string `json:"tip" form:"tip"`
	Type string `json:"type" form:"type"`
	Value string `json:"value" form:"value"`
	Content string `json:"content" form:"content"`
	Rule string `json:"rule" form:"rule"`
	Extend string `json:"extend" form:"extend"`
}


func NewConfig() (config *Config) {
	return &Config{}
}

func (m *Config) Pagination(offset, limit int, key string) (res []Config, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Config{}).Count(&count)
	return
}

func (m *Config) Create() (err error, newAttr *Config) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *Config) Update() (err error, newAttr Config) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Config) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Config) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Config) FindById(id int) (config Config, err error) {
	err = Db.Where("id=?", id).First(&config).Error
	return
}

func (m *Config) FindByName(name string) (config Config, err error) {
	err = Db.Where("name=?", name).First(&config).Error
	return
}

func (m *Config) FindAll() (configs []Config, err error) {
	err = Db.Find(&configs).Error
	return
}
