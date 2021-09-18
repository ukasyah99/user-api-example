package model

type User struct {
	Id int `db:"id" json:"id" example:"1"`
	DisplayName string `db:"display_name" json:"display_name" example:"Fulan"`
	Name string `db:"name" json:"name" example:"fulan_fulan"`
	Email string `db:"email" json:"email" example:"fulan@gmail.com"`
	Password string `db:"password" json:"password" example:"jangankasihtauoranglain"`
	CreatedAt string `db:"created_at" json:"created_at" example:"2018-09-30 10:39:08"`
	UpdatedAt string `db:"updated_at" json:"updated_at" example:"2018-09-30 10:39:08"`
}
