package models

type BooksPreg struct {
	Model
	Name       string    `json:"name"                           form:"name"`
	ListAuthor string `json:"list_author"                       form:"list_author"`
	ListDescribe string `json:"list_describe"                   form:"list_describe"`
	ListMsgImg string `json:"list_msg_img"                      form:"list_msg_img"`
	ListA string `json:"list_a"                                 form:"list_a"`
	ContentText string `json:"content_text"                     form:"content_text"`
	ListMsgLastTime string `json:"list_msg_last_time"           form:"list_msg_last_time"`
	ListAuthorFilter string `json:"list_author_filter"          form:"list_author_filter"`
	ListDescribeFilter string `json:"list_describe_filter"      form:"list_describe_filter"`
	ContentTextFilter string `json:"content_text_filter"        form:"content_text_filter"`

	CreatedAtText int64 `json:"create_at_text" gorm:"-"`
}

func NewBooksPreg() (books *BooksPreg) {
	return &BooksPreg{}
}
