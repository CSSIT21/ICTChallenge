package hub

import (
	"backend/types/database"
)

var Hub *Model

type Model struct {
	Teams  []database.Team  `json:"teams"`
	Topics []database.Topic `json:"topics"`
}
