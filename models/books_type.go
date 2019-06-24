package models

import (
	"errors"
	"strings"
)

type BooksType struct {
	Model
	Name     string `json:"name" form:"name"`
	Pid      int    `json:"pid" form:"pid"`
	IsNav    int    `json:"is_nav" form:"is_nav"`
	NameText string `json:"name_text" gorm:"-"`
}

func NewBooksType() (books *BooksType) {
	return &BooksType{}
}

func (m *BooksType) Pagination(offset, limit int, key string) (res []BooksType, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(BooksType{}).Count(&count)
	return
}

func (m *BooksType) Create() (err error, newAttr *BooksType) {
	err = Db.Create(m).Error
	newAttr = m
	return
}

func (m *BooksType) Update() (err error, newAttr BooksType) {
	if m.Id > 0 {
		err = Db.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	newAttr = *m
	return
}

func (m *BooksType) Delete() (err error) {
	if m.Id > 0 {
		err = Db.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksType) DelBatch(ids []int) (err error) {
	if len(ids) > 0 {
		err = Db.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
	return
}

func (m *BooksType) FindColumn() (err error, booksType map[int]string) {
	booksType = make(map[int]string)
	var booksTypes []BooksType
	err = Db.Find(&booksTypes).Error
	for _, v := range booksTypes {
		booksType[v.Id] = v.Name
	}
	return
}

//树形菜单
func (m *BooksType) FindTree(key string) (booksType2 *[]BooksType, err error) {
	var booksTypeData []BooksType
	Db.Where("name like ?", "%"+key+"%").Find(&booksTypeData)

	booksType2 = new([]BooksType)
	Tree(booksTypeData, booksType2, 0, -1)
	return
}

func Tree(treeData []BooksType, booksType2 *[]BooksType, pid int, level int) {
	var position int
	level++
	for _, v := range treeData {
		if v.Pid == pid {
			position++
			v.NameText = Prefix(level, position, TNum(treeData, pid)) + v.Name
			*booksType2 = append(*booksType2, v)
			Tree(treeData, booksType2, v.Id, level)
		}
	}
}
func Prefix(level int, position int, lastPostion int) (str string) {
	if level > 0 {
		if position == lastPostion {
			str = strings.Repeat("│ ", level-1) + "└─"
		} else {
			str = strings.Repeat("│ ", level-1) + "├─"
		}
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

func (m *BooksType) FindByPid(pid int) (booksType []BooksType, err error) {
	err = Db.Where("pid=?", pid).Find(&booksType).Error
	return
}

func (m *BooksType) FindById(id int) (booksType BooksType, err error) {
	err = Db.Where("id=?", id).Find(&booksType).Error
	return
}
