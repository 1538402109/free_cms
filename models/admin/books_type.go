package models

import (
	"errors"
	"free_cms/models"
)

type BooksType struct {
	models.BooksType
}

func NewBooksType() (books *BooksType) {
	return &BooksType{}
}

func (m *BooksType) Pagination(offset, limit int, key string) (res []BooksType, count int) {
	query := models.Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(BooksType{}).Count(&count)
	return
}

func (m *BooksType) Create() (err error, newAttr *BooksType) {
	err = models.Db.Create(m).Error
	newAttr = m
	return
}

func (m *BooksType) Update() (err error, newAttr BooksType) {
	if m.Id > 0 {
		err = models.Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *BooksType) Delete() (err error) {
	if m.Id > 0 {
		err = models.Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksType) DelBath(ids []int) (err error) {
	if len(ids) > 0 {
		err = models.Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksType) FindById(id int) (booksType BooksType, err error) {
	err = models.Db.Where("id=?", id).First(&booksType).Error
	return
}

func (m *BooksType) FindColumn() (err error, booksType map[int]string) {
	booksType = make(map[int]string)
	var booksTypes []BooksType
	err = models.Db.Find(&booksTypes).Error
	for _, v := range booksTypes {
		booksType[v.Id] = v.Name
	}
	return
}

//树形菜单
func (m *BooksType) FindTree(key string) (err error, booksType2 *[]BooksType) {
	var booksType []BooksType
	Db.Where("name like ?", "%"+key+"%").Find(&booksType)

	booksType2 = new([]BooksType)
	for k, v := range booksType {
		var level int = 0
		booksType[k].NameText = v.Name
		v.NameText = v.Name
		if key == "" && v.Pid == 0 {
			*booksType2 = append(*booksType2, v)
			T(booksType, v.Id, booksType2, level)
		}
	}
	if key != "" {
		booksType2 = &booksType
	}
	return
}

func T(b []BooksType, id int, booksType2 *[]BooksType, level int) {
	var position int
	level++
	for _, v := range b {
		if v.Pid == id {
			position++
			v.NameText = Tpl(level, position, TNum(b, id)) + v.Name
			*booksType2 = append(*booksType2, v)
			T(b, v.Id, booksType2, level)
		}
	}
}

func Tpl(level int, position int, lastPostion int) (str string) {
	if level == 1 && position == lastPostion {
		str = "└─"
	}

	if level == 1 && position != lastPostion {
		str = "├─"
	}

	if level == 2 && position == lastPostion {
		str = "│ └─"
	}

	if level == 2 && position != lastPostion {
		str = "│ ├─"
	}

	return
}

func TNum(b []BooksType, id int) (num int) {
	for _, v := range b {
		if v.Pid == id {
			num++
		}
	}
	return
}
