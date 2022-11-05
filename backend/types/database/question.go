package database

type Question struct {
	Id     uint64 `json:"id"`
	Score  int    `json:"score"`
	Opened bool   `json:"opened"`
}
