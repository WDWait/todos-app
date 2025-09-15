package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
