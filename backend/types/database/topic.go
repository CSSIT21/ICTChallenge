package database

type Topic struct {
	Id        uint64      `json:"id"`
	Title     string      `json:"title"`
	Questions []*Question `json:"questions"`
}
