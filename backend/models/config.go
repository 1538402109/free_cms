package models

import (
	"errors"
	"free_cms/common/models"
)

type Config struct {
	models.Config
}

func NewConfig() (config *Config) {
	return &Config{}
}

func (m *Config) Pagination(offset, limit int, key string) (res []Config, count int) {
	query := models.Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Config{}).Count(&count)
	return
}

func (m *Config) Create() (err error, newAttr *Config) {
	err = models.Db.Create(m).Error
	newAttr = m
	return
}

func (m *Config) Update() (err error, newAttr Config) {
	if m.Id > 0 {
		err = models.Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Config) Delete() (err error) {
	if m.Id > 0 {
		err = models.Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Config) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = models.Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Config) FindById(id int) (config Config, err error) {
	err = models.Db.Where("id=?", id).First(&config).Error
	return
}

func (m *Config) FindByName(name string) (config models.Config, err error) {
	err = models.Db.Where("name=?", name).First(&config).Error
	return
}

func (m *Config) FindAll() (configs []models.Config, err error) {
	err = models.Db.Find(&configs).Error
	return
}
