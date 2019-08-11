package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Post struct {
	Model
	PostTitle       string          `json:"post_title"      form:"post_title"      gorm:"default:''"`
	Author          string          `json:"author"          form:"author"          gorm:"default:''"`
	PostStatus      int             `json:"post_status"     form:"post_status"     gorm:"default:'1'"`
	CommentStatus   int             `json:"comment_status"  form:"comment_status"  gorm:"default:'1'"`
	PostHits        int             `json:"post_hits"       form:"post_hits"       gorm:"default:'0'"`
	PostFavorites   int             `json:"post_favorites"  form:"post_favorites"  gorm:"default:'0'"`
	PostLike        int             `json:"post_like"       form:"post_like"       gorm:"default:'0'"`
	PostComment     int             `json:"post_comment"    form:"post_comment"    gorm:"default:'0'"`
	PostKeywords    string          `json:"post_keywords"   form:"post_keywords"   gorm:"default:''"`
	PostDescription string          `json:"post_description"form:"post_description"gorm:"default:''"`
	Image           string          `json:"image"           form:"image"           gorm:"default:''"`
	PostContent     string          `json:"post_content"    form:"post_content"    gorm:"default:''"`
}


func NewPost() (post *Post) {
	return &Post{}
}

func (m *Post)AfterFind (scope *gorm.Scope) (err error) {
	if m.Image ==""{
		m.Image = "/static/home/blog/images/default-img.png"
	}
	return
}

func (m *Post) Pagination(offset, limit int, key string) (res []Post, count int) {
	query := Db
	if key != "" {
		query = query.Where("post_title like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Post{}).Count(&count)
	return
}

func (m *Post) Create() (newAttr Post, err error) {
	err = Db.Create(m).Error
	newAttr = *m
	return
}

func (m *Post) Update() (newAttr Post, err error) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Post) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Post) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Post) FindById(id int) (post Post, err error) {
	err = Db.Where("id=?", id).First(&post).Error
	return
}

