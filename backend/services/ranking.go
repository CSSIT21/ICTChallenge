package services

import "backend/repository"

type RankingService interface {
	GetAllRankings() (repository.Rankings, error)
}
