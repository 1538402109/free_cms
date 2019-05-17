package models

type Config struct {
	Model
	Name string `json:"name" form:"name"`
	Group string `json:"group" form:"group"`
	Title string `json:"title" form:"title"`
	Tip string `json:"tip" form:"tip"`
	Type string `json:"type" form:"type"`
	Value string `json:"value" form:"value"`
	Content string `json:"content" form:"content"`
	Rule string `json:"rule" form:"rule"`
	Extend string `json:"extend" form:"extend"`
}


func NewConfig() (config *Config) {
	return &Config{}
}

