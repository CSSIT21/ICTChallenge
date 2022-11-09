package payload

type TeamInfo struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Score int32  `json:"score"`
}

type Podium struct {
	Id         uint64  `json:"id"`
	Name       string  `json:"name"`
	Percentile float32 `json:"percentile"`
}
