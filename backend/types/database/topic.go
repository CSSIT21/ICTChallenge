package database

import "time"

type Topic struct {
	Id    uint64  `json:"id"`
	Title string  `json:"title"`
	Cards []*Card `json:"cards"`
}

type Card struct {
	Id       uint64        `json:"id,omitempty"`
	Score    int32         `json:"score,omitempty"`
	Opened   bool          `json:"opened,omitempty"`
	Bonus    bool          `json:"bonus,omitempty"`
	Question string        `json:"question,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
}
