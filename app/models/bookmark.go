package models

type Bookmark struct {
	BaseModel
	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`

	Post Post `json:"post"`
	User User `json:"user"`
}
