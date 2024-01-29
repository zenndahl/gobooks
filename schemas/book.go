package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string
	Author string
	Genre  string
	Year   string
}
