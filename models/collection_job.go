package models

import (
	"errors"
)

type CollectionJob struct {
	Model
	Name              string `json:"name" form:"name" gorm:"default:''"`
	TriggerTime       string `json:"trigger_time"form:"trigger_time"gorm:"default:''"`
	PregIds           string `json:"preg_ids"    form:"preg_ids"    gorm:"default:''"`
	ImagePath         string `json:"image_path"  form:"image_path"  gorm:"default:''"`
	ApiUrl            string `json:"api_url"  form:"api_url"  gorm:"default:''"`
	ArticleCategoryId string `json:"article_category_id"  form:"article_category_id"  gorm:"default:''"`
}

func NewCollectionJob() (collectionJob *CollectionJob) {
	collectionJob = &CollectionJob{}
	collectionJob.ImagePath = `C:\www\thinkcmf_new\public\upload\caiji`
	collectionJob.ApiUrl = "http://thinkcmf3.com/api/portal/articles"
	collectionJob.TriggerTime = "23:00:00"
	return
}

func (m *CollectionJob) Pagination(offset, limit int, key string) (res []CollectionJob, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(CollectionJob{}).Count(&count)
	return
}

func (m *CollectionJob) Create() (newAttr CollectionJob, err error) {
	err = Db.Create(m).Error
	newAttr = *m
	return
}

func (m *CollectionJob) Update() (newAttr CollectionJob, err error) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *CollectionJob) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *CollectionJob) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *CollectionJob) FindById(id int) (collectionJob CollectionJob, err error) {
	err = Db.Where("id=?", id).First(&collectionJob).Error
	return
}
