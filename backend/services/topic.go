package services

import (
	"backend/types/payload"
)

type TopicService interface {
	OpenCard(body *payload.OpenCard) error
}
