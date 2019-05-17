package models

import (
	"free_cms/common/models"
	"github.com/jinzhu/gorm"
)

type Books struct {
	models.Books
}

func NewBooks() (books *Books) {
	return &Books{}
}

func (m *Books) AfterFind(scope *gorm.Scope) (err error) {
	var types = []string{"连载","完本"}
	m.BookType2Text = types[m.BookType2]
	return
}

func (m *Books) Pagination(offset, limit int, key string, bookType int) (res []Books, count int) {
	query := models.Db
	if key != "" {
		query = query.Where("book_name like ?", "%"+key+"%")
	}
	if bookType > 0 {
		query = query.Where("book_type=?", bookType)
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Books{}).Count(&count)
	return
}

func (m *Books) FindByBooksType(booksType, status, limit int) (books []Books,err error) {
	query := models.Db
	if booksType > 0 {
		query = query.Where("book_type=?", booksType)
	}
	if status > 0 {
		query = query.Where("book_status=?", status)
	}
	err = query.Limit(limit).Order("id desc").Find(&books).Error
	return
}

//根据不同参数查询分页数据
func (m *Books)FindOfBtTable(pid,offset,limit int,key string,bookType2 int)(books []Books,count int,err error){
	query := models.Db.Preload("BookTypes")
	if pid >0{
		query = query.Where("book_type=?",pid)
	}
	if key != "" {
		query = query.Where("book_name like ?", "%"+key+"%")
	}
	if bookType2>-1{
		query = query.Where("book_type2=?",bookType2)
	}

	query.Offset(offset).Limit(limit).Order("id desc").Find(&books)
	query.Model(Books{}).Count(&count)
	return
}

//使用id查询一条数据
func (m *Books)FindById(id int)(books Books,err error){
	err = models.Db.Where("id=?",id).Preload("BookTypes").First(&books).Error
	return
}