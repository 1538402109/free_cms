package models

type BooksType struct {
	Model
	Name string `json:"name" form:"name"`
	Pid  int    `json:"pid" form:"pid"`
	IsNav int `json:"is_nav" form:"is_nav"`
	NameText string `json:"name_text" gorm:"-"`
}

func NewBooksType() (books *BooksType) {
	return &BooksType{}
}

func (m *BooksType) FindByPid(pid int) (booksType []BooksType,err error) {
	err = Db.Where("pid=?", pid).Find(&booksType).Error
	return
}

func (m *BooksType) FindById(id int) (booksType BooksType,err error) {
	err = Db.Where("id=?", id).Find(&booksType).Error
	return
}
