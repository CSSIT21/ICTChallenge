package services

import (
	"backend/repository"
	"backend/types/database"
)

type rankingService struct {
	rankingEvent repository.RankingRepository
}

func NewRankingService(rankingRepository repository.RankingRepository) rankingService {
	return rankingService{rankingEvent: rankingRepository}
}

func (r rankingService) GetAllRankings() ([]database.Ranking, error) {
	a, _ := r.rankingEvent.GetAll()
	return a, nil
}
