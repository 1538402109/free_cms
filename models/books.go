package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Books struct {
	Model
	BookName       string    `json:"book_name"        form:"book_name"`
	BookType       int       `json:"book_type"        form:"book_type"`
	BookType2      int       `json:"book_type2"        form:"book_type2"`
	BookAuthor     string    `json:"book_author"      form:"book_author"`
	BookNewChapter string    `json:"book_new_chapter" form:"book_new_chapter"`
	BookDescribe   string    `json:"book_describe"    form:"book_describe"`
	BookStatus     int       `json:"book_status"      form:"book_status"`
	BookLastAt     string    `json:"book_last_at"     form:"book_last_at"`
	PregId         int       `json:"preg_id"          form:"preg_id"`
	ListUrl        string    `json:"list_url"         form:"list_url"`
	BookImg        string    `json:"book_img"         form:"book_img"`
	IsTop          int       `json:"is_top"           form:"is_top"`
	SeoTitle       string    `json:"seo_title"        form:"seo_title"`
	SeoKeyword     string    `json:"seo_keyword"      form:"seo_keyword"`
	SeoDescription string    `json:"seo_description"     form:"seo_description"`
	BookTypes      BooksType `gorm:"ForeignKey:Id;AssociationForeignKey:BookType"`

	CreatedAtText int64  `json:"created_at_text"   gorm:"-"`
	BookTypeText  string `json:"book_type_text"    gorm:"-"`
	BookType2Text string `json:"book_type2_text"    gorm:"-"`
}

func NewBooks() (books *Books) {
	books = &Books{}
	books.PregId = 1
	return
}

func (m *Books) AfterFind(scope *gorm.Scope) (err error) {
	m.CreatedAtText = m.CreatedAt.Unix()
	_, booksType := NewBooksType().FindColumn()
	m.BookTypeText = booksType[m.BookType]
	var types = []string{"连载", "完本"}
	m.BookType2Text = types[m.BookType2]
	return
}

func (m *Books) Pagination(offset, limit int, key string) (res []Books, count int) {
	query := Db
	if key != "" {
		query = query.Where("book_name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Books{}).Count(&count)
	return
}

func (m *Books) Create() (err error, newAttr *Books) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *Books) Update() (err error, newAttr Books) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *Books) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Books) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *Books) FindById(id int) (books Books, err error) {
	err = Db.Where("id=?", id).First(&books).Error
	return
}

//使用id查询一条数据
func (m *Books) FindByIdJoinBookTypes(id int) (books Books, err error) {
	err = Db.Where("id=?", id).Preload("BookTypes").First(&books).Error
	return
}

//home控制器使用
func (m *Books) PaginationOfHome(offset, limit int, key string, bookType int) (res []Books, count int) {
	query := Db
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

func (m *Books) FindByBooksType(booksType, status, limit int) (books []Books, err error) {
	query := Db
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
func (m *Books) FindOfBtTable(pid, offset, limit int, key string, bookType2 int) (books []Books, count int, err error) {
	query := Db.Preload("BookTypes")
	if pid > 0 {
		query = query.Where("book_type=?", pid)
	}
	if key != "" {
		query = query.Where("book_name like ?", "%"+key+"%")
	}
	if bookType2 > -1 {
		query = query.Where("book_type2=?", bookType2)
	}

	query.Offset(offset).Limit(limit).Order("id desc").Find(&books)
	query.Model(Books{}).Count(&count)
	return
}
