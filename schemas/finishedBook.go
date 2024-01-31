package schemas

import (
	"time"

	"gorm.io/gorm"
)

type FinishedBook struct {
	gorm.Model
	Name       string
	Author     string
	Genre      string
	Year       string
	Finished   string
	FinishedAt time.Time
	Rating     int
}

type FinishedBookResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	Name       string    `json:"name"`
	Author     string    `json:"author"`
	Genre      string    `json:"genre"`
	Year       string    `json:"year"`
	Finished   string    `json:"finished"`
	FinishedAt time.Time `json:"finishedAt"`
	Rating     int       `json:"rating"`
}
