package domain

import "time"

type Page struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CategoryID string    `json:"category_id"`
	Title      string    `json:"title"`
	Text       string    `json:"text"`
}
