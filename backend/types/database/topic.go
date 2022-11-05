package database

type Topic struct {
	Id    uint64      `json:"id"`
	Title string      `json:"title"`
	Cards []*Question `json:"cards"`
}

type Question struct {
	Id      uint64  `json:"id"`
	Score   float64 `json:"score"`
	Opened  bool    `json:"opened"`
	Problem string  `json:"problem"`
}
