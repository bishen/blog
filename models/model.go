package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	ID        uint `gorm:"primary_key" form:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
