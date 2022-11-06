package repository

type rankingEvent struct {
	// event
}

func NewRankingEvent() rankingEvent {
	return rankingEvent{}
}

func (r rankingEvent) GetAll() (Rankings, error) {
	// TODO: implement
	return Rankings{
		{
			Id:    1,
			Name:  "Mixko",
			Score: 234,
		},
		{
			Id:    2,
			Name:  "Also Mixko",
			Score: 235,
		},
		{
			Id:    3,
			Name:  "Also Mixko",
			Score: 236,
		},
	}, nil
}
