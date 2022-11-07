package database

type Raw struct {
	Topics []*Topic `json:"topics"`
	Teams  []*Team  `json:"teams"`
}
