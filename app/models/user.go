package models

type User struct {
	BaseModel
	FullName *string `json:"full_name"`
	Email    string  `json:"email" gorm:"unique"`
	Password string  `json:"password"`

	Posts     []*Post     `json:"posts"`
	Bookmarks []*Bookmark `json:"bookmarks"`
	Comment   []*Comment  `json:"comments"`
	Likes     []*Like     `json:"likes"`
	Followers []*User     `json:"followers" gorm:"many2many:user_followers;"`
}
