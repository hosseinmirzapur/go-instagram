package models

type Comment struct {
	BaseModel
	Text   string `json:"text"`
	UserID uint   `json:"user_id"`

	User    User       `json:"user"`
	Replies []*Comment `json:"replies"`
	Likes   []*Like    `json:"likes" gorm:"polymorphic:Likable"`
}
