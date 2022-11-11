package database

type Team struct {
	Id     uint64   `json:"id"`
	Name   string   `json:"name"`
	School string   `json:"school"`
	Token  string   `json:"token,omitempty"`
	Scores []*Score `json:"scores"`
}
