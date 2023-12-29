package models

type Post struct {
	BaseModel
	Title   string  `json:"title"`
	Caption *string `json:"caption"`
	Content string  `json:"content"`
	UserID  uint    `json:"user_id"`

	User      User        `json:"user"`
	Comments  []*Comment  `json:"comments"`
	Bookmarks []*Bookmark `json:"bookmarks"`
	Likes     []*Like     `json:"likes" gorm:"polymorphic:Likable"`
}
