package models

type Category struct {
	Model
	Pid         int
	Type        int
	Name        string
	Nickname    string
	Flag        int
	Href        string
	IsNav       int
	Image       string
	Keywords    string
	Description string
	Content     string
	Weight      int
	Status      int
	Tpl         string
}

func NewCategory() (category *Category) {
	return &Category{}
}

func (c *Category) Create() {

}
