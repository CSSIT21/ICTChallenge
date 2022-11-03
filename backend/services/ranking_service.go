package services

import "backend/repository"

type rankingService struct {
	rankingEvent repository.RankingRepository
}

func NewRankingService(rankingRepository repository.RankingRepository) rankingService {
	return rankingService{rankingEvent: rankingRepository}
}

func (r rankingService) GetAllRankings() (repository.Rankings, error) {
	a, _ := r.rankingEvent.GetAll()
	return a, nil
}
