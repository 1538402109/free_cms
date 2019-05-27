package models

import (
	"errors"
)

type Category struct {
	Model
	Id                  int       `json:"id"             form:"id"             gorm:"default:''"`
	Pid                 int       `json:"pid"            form:"pid"            gorm:"default:'0'"`
	Type                int       `json:"type"           form:"type"           gorm:"default:'1'"`
	Name                string    `json:"name"           form:"name"           gorm:"default:''"`
	Nickname            string    `json:"nickname"       form:"nickname"       gorm:"default:''"`
	Flag                int       `json:"flag"           form:"flag"           gorm:"default:'0'"`
	Href                string    `json:"href"           form:"href"           gorm:"default:''"`
	IsNav               int       `json:"is_nav"         form:"is_nav"         gorm:"default:'0'"`
	Image               string    `json:"image"          form:"image"          gorm:"default:''"`
	Keywords            string    `json:"keywords"       form:"keywords"       gorm:"default:''"`
	Description         string    `json:"description"    form:"description"    gorm:"default:''"`
	Content             string    `json:"content"        form:"content"        gorm:"default:''"`
	CreatedAt           int       `json:"created_at"     form:"created_at"     gorm:"default:'0'"`
	UpdatedAt           int       `json:"updated_at"     form:"updated_at"     gorm:"default:'0'"`
	Weigh               int       `json:"weigh"          form:"weigh"          gorm:"default:'0'"`
	Status              int       `json:"status"         form:"status"         gorm:"default:'1'"`
	Tpl                 string    `json:"tpl"            form:"tpl"            gorm:"default:'list'"`

}

func NewCategory() (category *Category) {
	return &Category{}
}

func (m *Category) Pagination(offset, limit int, key string) (res []Category, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Category{}).Count(&count)
	return
}

func (m *Category) Create() (err error, newAttr *Category) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *Category) Update() (err error, newAttr Category) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Category) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Category) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Category) FindById(id int) (category Category, err error) {
	err = Db.Where("id=?", id).First(&category).Error
	return
}