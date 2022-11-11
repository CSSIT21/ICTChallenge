package payload

type UpdateScore struct {
	Update [6]int `json:"update"`
}

type LeaderboardMode struct {
	Mode string `json:"mode"`
}
