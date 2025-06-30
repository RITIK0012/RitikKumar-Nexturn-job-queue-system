package models

import "time"

type Job struct {
	ID        int64     `json:"id"`
	Payload   string    `json:"payload"`
	Status    string    `json:"status"`
	Result    string    `json:"result"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
