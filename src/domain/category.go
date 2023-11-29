package domain

import "time"

type Category struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Slag      string    `json:"slag"`
	IconURL   string    `json:"icon_url"`
}
