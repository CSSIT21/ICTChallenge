package repository

import "backend/types/database"

type RankingRepository interface {
	GetAll() ([]database.Ranking, error)
}
