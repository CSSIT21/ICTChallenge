package services

import (
	"backend/types/database"
)

type RankingService interface {
	GetAllRankings() ([]database.Ranking, error)
}
