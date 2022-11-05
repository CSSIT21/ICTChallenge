package database

type Ranking struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
}
