package home

import "free_cms/models"

type BooksType struct {
	models.BooksType
}

func NewBooksType()(*BooksType){
	return &BooksType{}
}

func (m *BooksType) FindByPid(pid int) (booksType []BooksType,err error) {
	err = models.Db.Where("pid=?", pid).Find(&booksType).Error
	return
}

func (m *BooksType) FindById(id int) (booksType BooksType,err error) {
	err = models.Db.Where("id=?", id).Find(&booksType).Error
	return
}