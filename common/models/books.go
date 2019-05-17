package models

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
	return &Books{}
}
