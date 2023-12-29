package models

type Like struct {
	UserID      uint   `json:"user_id"`
	LikableID   uint   `json:"likable_id"`
	LikableType string `json:"likable_type"`

	User User `json:"user"`
}
