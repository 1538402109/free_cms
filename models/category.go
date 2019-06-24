package models

import (
	"errors"
	"free_cms/pkg/util"
	"strings"
)

type Category struct {
	Model
	Pid         int    `json:"pid"            form:"pid"            gorm:"default:'0'"`
	Type        int    `json:"type"           form:"type"           gorm:"default:'1'"`
	Name        string `json:"name"           form:"name"           gorm:"default:''"`
	Nickname    string `json:"nickname"       form:"nickname"       gorm:"default:''"`
	Flag        int    `json:"flag"           form:"flag"           gorm:"default:'0'"`
	Href        string `json:"href"           form:"href"           gorm:"default:''"`
	IsNav       int    `json:"is_nav"         form:"is_nav"         gorm:"default:'0'"`
	Image       string `json:"image"          form:"image"          gorm:"default:''"`
	Keywords    string `json:"keywords"       form:"keywords"       gorm:"default:''"`
	Description string `json:"description"    form:"description"    gorm:"default:''"`
	Content     string `json:"content"        form:"content"        gorm:"default:''"`
	Weigh       int    `json:"weigh"          form:"weigh"          gorm:"default:'0'"`
	Status      int    `json:"status"         form:"status"         gorm:"default:'1'"`
	Tpl         string `json:"tpl"            form:"tpl"            gorm:"default:'list'"`
	NameText    string `json:"name_text" gorm:"-"`
}

func NewCategory() (category *Category) {
	return &Category{}
}

//树形菜单
func (m *Category) FindTree(key string) (newCategory *[]Category, err error) {
	var categoryData []Category
	Db.Where("name like ?", "%"+key+"%").Find(&categoryData)
	newCategory = new([]Category)
	m.Tree(categoryData, newCategory, 0, -1)
	return
}

// Tree 树形菜单，递归实现
func (m *Category) Tree(treeData []Category, newTreeData *[]Category, pid int, level int) {
	var position int
	level++
	for _, v := range treeData {
		if v.Pid == pid {
			position++
			v.NameText = Prefix(level, position, m.TNum(treeData, pid)) + v.Name
			*newTreeData = append(*newTreeData, v)
			m.Tree(treeData, newTreeData, v.Id, level)
		}
	}
}
// Prefix 数据前缀
func (m *Category) Prefix(level int, position int, lastPostion int) (str string) {
	if level > 0 {
		if position == lastPostion {
			str = strings.Repeat("│ ", level-1) + "└─"
		} else {
			str = strings.Repeat("│ ", level-1) + "├─"
		}
	}
	return
}
// TNum 判断切片数量
func (m *Category) TNum(b []Category, id int) (num int) {
	for _, v := range b {
		if v.Pid == id {
			num++
		}
	}
	return
}

// TreeData Layui 树形插件
func (m *Category) TreeData() (treeData []util.LayuiTreeData) {
	var categoryData []Category
	Db.Order("id desc").Find(&categoryData)

	var treeRes []util.LayuiTreeDataTpl
	for _, v := range categoryData {
		treeRes = append(treeRes, util.LayuiTreeDataTpl{v.Id, v.Pid, v.Name})
	}

	treeData = util.Tree(treeRes, &treeData, 0)
	return
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

func (m *Category) Create() (newAttr Category, err error) {
	err = Db.Create(m).Error
	newAttr = *m
	return
}

func (m *Category) Update() (newAttr Category, err error) {
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
