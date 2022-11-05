package repository

import "backend/types/database"

type PodiumRepository interface {
	GetAll() ([]database.Podium, error)
}
