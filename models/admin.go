package models

type Admin struct {
	ID   int    `gorm:"primary_key" form:"id"`
	User string `form:"username"`
	Pass string `gorm:"type:text" form:"password"`
}
