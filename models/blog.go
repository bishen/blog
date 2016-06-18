package models

import (
	"time"
)

type Blog struct {
	ID        int    `gorm:"primary_key" form:"id"`
	Cid       int    `form:"cid"`
	Title     string `form:"title"`
	Content   string `gorm:"type:text" form:"content"`
	CreatedAt time.Time
}
