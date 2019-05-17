package models

type Category struct {
	Model
	Id                  int       `json:"id"             form:"id"             gorm:"default:''"`
	Pid                 int       `json:"pid"            form:"pid"            gorm:"default:'0'"`
	Type                int       `json:"type"           form:"type"           gorm:"default:'1'"`
	Name                string    `json:"name"           form:"name"           gorm:"default:''"`
	Nickname            string    `json:"nickname"       form:"nickname"       gorm:"default:''"`
	Flag                int       `json:"flag"           form:"flag"           gorm:"default:'0'"`
	Href                string    `json:"href"           form:"href"           gorm:"default:''"`
	IsNav               int       `json:"is_nav"         form:"is_nav"         gorm:"default:'0'"`
	Image               string    `json:"image"          form:"image"          gorm:"default:''"`
	Keywords            string    `json:"keywords"       form:"keywords"       gorm:"default:''"`
	Description         string    `json:"description"    form:"description"    gorm:"default:''"`
	Content             string    `json:"content"        form:"content"        gorm:"default:''"`
	CreatedAt           int       `json:"created_at"     form:"created_at"     gorm:"default:'0'"`
	UpdatedAt           int       `json:"updated_at"     form:"updated_at"     gorm:"default:'0'"`
	Weigh               int       `json:"weigh"          form:"weigh"          gorm:"default:'0'"`
	Status              int       `json:"status"         form:"status"         gorm:"default:'1'"`
	Tpl                 string    `json:"tpl"            form:"tpl"            gorm:"default:'list'"`
	
}


func NewCategory() (category *Category) {
	return &Category{}
}

func (m *Category) FindById(id int) (category Category, err error) {
	err = Db.Where("id=?", id).First(&category).Error
	return
}

