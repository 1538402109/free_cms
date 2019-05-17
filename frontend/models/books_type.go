package models

import "free_cms/common/models"

type BooksType struct {
	models.BooksType
}

func NewBooksType()(*BooksType){
	return &BooksType{}
}

