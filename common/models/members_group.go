package models

type MembersGroup struct {
	Model
	Id                  int       `json:"id"             form:"id"             gorm:"default:''"`
	Name                string    `json:"name"           form:"name"           gorm:"default:''"`
	Rules               string    `json:"rules"          form:"rules"          gorm:"default:''"`
	Createtime          int       `json:"createtime"     form:"createtime"     gorm:"default:''"`
	Updatetime          int       `json:"updatetime"     form:"updatetime"     gorm:"default:''"`
	Status              string    `json:"status"         form:"status"         gorm:"default:''"`
	
}


func NewMembersGroup() (membersGroup *MembersGroup) {
	return &MembersGroup{}
}

func (m *MembersGroup) FindById(id int) (membersGroup MembersGroup, err error) {
	err = Db.Where("id=?", id).First(&membersGroup).Error
	return
}

