package models

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type CollectionPreg struct {
	Model
	Name        string `json:"name"        form:"name"        gorm:"default:''"`
	List        string `json:"list"        form:"list"        gorm:"default:''"`
	ListA       string `json:"list_a"      form:"list_a"      gorm:"default:''"`
	ListLi      string `json:"list_li"      form:"list_li"      gorm:"default:''"`
	Time        string `json:"time"      form:"time"      gorm:"default:''"`
	TimeFormat  string `json:"time_format"      form:"time_format"      gorm:"default:''"`
	ContentText string `json:"content_text"form:"content_text"gorm:"default:''"`
	IsChecked   bool   `json:"is_checked" gorm:"-"`
}

func NewCollectionPreg() (collectionPreg *CollectionPreg) {
	return &CollectionPreg{}
}

func (m *CollectionPreg) FindCheck(id string) (collection []CollectionPreg) {
	err := Db.Find(&collection).Error
	if err != nil {
		log.Println(err)
	}
	ids := strings.Split(id, ",")
	for k, v := range collection {
		for _, v2 := range ids {
			v2Id, _ := strconv.Atoi(v2)
			if v.Id == v2Id {
				collection[k].IsChecked = true
			}
		}
	}
	return
}

func (m *CollectionPreg) Pagination(offset, limit int, key string) (res []CollectionPreg, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(CollectionPreg{}).Count(&count)
	return
}

func (m *CollectionPreg) Create() (newAttr CollectionPreg, err error) {
	err = Db.Create(m).Error
	newAttr = *m
	return
}

func (m *CollectionPreg) Update() (newAttr CollectionPreg, err error) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *CollectionPreg) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *CollectionPreg) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *CollectionPreg) FindById(id int) (collectionPreg CollectionPreg, err error) {
	err = Db.Where("id=?", id).First(&collectionPreg).Error
	return
}
