package home

import (
	"free_cms/models"
)

type Books struct {
	models.Books
}

func NewBooks() (books *Books) {
	return &Books{}
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

func (m *Books) FindByBooksType(booksType, status, limit int) (err error, books []Books) {
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

func (m *Books)FindOfBtTable(pid,offset,limit int,key string,bookType2 int)(err error,books []Books,count int){
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

func (m *Books)FindById(id int)(books Books,err error){
	err = models.Db.Where("id=?",id).Preload("BookTypes").First(&books).Error
	return
}