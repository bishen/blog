package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Activity struct {
	gorm.Model         // this will include attributes defined in model.go
	Title       string `form:"title" binding:"required"`
	Description string `form:"description"`
	City        string `form:"city" binding:"required"`
	Address     string `form:"address"`
	Contact     string `form:"contact"`
}
