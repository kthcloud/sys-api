package dto

import "time"

type NewsRead struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       []byte    `json:"image,omitempty"`
	PostedAt    time.Time `json:"postedAt"`
}
