package model

import "time"

type Question struct {
	Id        string    `json:"id"`
	User      string    `json:"user"`
	Question  string    `json:"question"`
	Answer    Answer    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
