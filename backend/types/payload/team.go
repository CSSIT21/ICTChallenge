package payload

import "backend/types/database"

type TeamInfo struct {
	Id     uint64            `json:"id"`
	Name   string            `json:"name"`
	School string            `json:"school"`
	Scores []*database.Score `json:"score"`
}

type TeamScore struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Score int32  `json:"score"`
}

type Podium struct {
	Id         uint64  `json:"id"`
	Name       string  `json:"name"`
	Score      int32   `json:"score"`
	Percentile float32 `json:"percentile"`
}

type StudentTurn struct {
	Name    string            `json:"name"`
	Current bool              `json:"current"`
	Topics  []*database.Topic `json:"topics"`
}
