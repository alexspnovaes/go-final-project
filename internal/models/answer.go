package model

import "time"

type Answer struct {
	Id        string    `json:"id"`
	User      string    `json:"User"`
	Text      string    `json:"Text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
