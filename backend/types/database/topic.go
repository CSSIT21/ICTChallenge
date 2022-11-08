package database

import "time"

type Topic struct {
	Id    uint64  `json:"id"`
	Title string  `json:"title"`
	Cards []*Card `json:"cards"`
}

type Card struct {
	Id       uint64        `json:"id"`
	Score    int32         `json:"score"`
	Opened   bool          `json:"opened"`
	Question string        `json:"question"`
	Duration time.Duration `json:"duration"`
}
