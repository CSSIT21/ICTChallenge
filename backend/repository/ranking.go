package repository

type Rankings []ranking
type Ranking ranking

type ranking struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
}

type RankingRepository interface {
	GetAll() (Rankings, error)
}
