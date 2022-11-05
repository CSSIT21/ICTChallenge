package database

type Podium struct {
	Id         uint64  `json:"id"`
	Name       string  `json:"name"`
	Percentile float32 `json:"percentile"`
}
