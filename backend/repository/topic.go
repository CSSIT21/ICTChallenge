package repository

import "backend/types/database"

type TopicRepository interface {
	GetAll() ([]*database.Topic, error)
}
