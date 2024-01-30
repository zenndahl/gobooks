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

type BookResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Genre     string    `json:"genre"`
	Year      string    `json:"year"`
}
