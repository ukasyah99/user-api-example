package model

type RefreshToken struct {
	Id string `db:"id" json:"id" example:"mDFI2lkHU7jdvuUiue4S6oOq2"`
	UserId int `db:"user_id" json:"user_id" example:"1"`
	CreatedAt string `db:"created_at" json:"created_at" example:"2018-09-30 10:39:08"`
	ValidUntil string `db:"valid_until" json:"valid_until" example:"2018-12-03 10:39:08"`
}
