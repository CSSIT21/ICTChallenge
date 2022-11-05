package hub

import (
	"backend/types/database"
)

var Hub *Model

type Model struct {
	Teams  map[uint64]*database.Team
	Topics map[uint64]*database.Topic
}
