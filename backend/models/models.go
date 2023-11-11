package models

import "time"

type Post struct {
	Id        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
