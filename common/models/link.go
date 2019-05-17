package models

type Link struct {
	Model
	Name    string    `json:"name" form:"name"`
	Description    string    `json:"description" form:"description"`
	Status    int    `json:"status" form:"status"`
	Url    string    `json:"url" form:"url"`
	Image    string    `json:"image" form:"image"`
	Target    string    `json:"target" form:"target"`
	ListOrder    int    `json:"list_order" form:"list_order" gorm:"default:'1000'"`
	
}


func NewLink() (link *Link) {
	return &Link{}
}


