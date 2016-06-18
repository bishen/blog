package models

type Classify struct {
	ID    int    `gorm:"primary_key" form:"id"`
	Title string `form:"title"`
	Url   string `form:"url"`
}
