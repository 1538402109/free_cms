package models

type AdminLog struct {
	Model
	Route    string    `json:"route" form:"route"`
	Description    string    `json:"description" form:"description"`
	UserId    int    `json:"user_id" form:"user_id"`
	Ip    int    `json:"ip" form:"ip"`
	Method    string    `json:"method" form:"method"`
	User User `gorm:"ForeignKey:Id;AssociationForeignKey:UserId"`

	IpText string `json:"ip_text" gorm:"-"`
	CreatedAtText string `json:"created_at_text" gorm:"-"`
}


func NewAdminLog() (adminLog *AdminLog) {
	return &AdminLog{}
}
