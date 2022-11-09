package services

import (
	"backend/loaders/hub"
	"backend/repository"
	"backend/types/database"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

type teamService struct {
	teamEvent  repository.TeamRepository
	topicEvent repository.TopicRepository
}

func NewTeamService(teamRepository repository.TeamRepository, topicRepository repository.TopicRepository) teamService {
	return teamService{teamEvent: teamRepository, topicEvent: topicRepository}
}

func (s teamService) GetTeamById(id uint64) (*database.Team, error) {
	a, _ := s.teamEvent.GetTeamById(id)
	return a, nil
}

func (s teamService) GetAllTeams() ([]*database.Team, error) {
	teams, _ := s.teamEvent.GetTeams()
	filteredTeams, _ := value.Iterate(teams, func(team *database.Team) (*database.Team, *response.Error) {
		return &database.Team{
			Id:     team.Id,
			Name:   team.Name,
			School: team.School,
			Scores: team.Scores,
		}, nil
	})
	return filteredTeams, nil
}

func (s teamService) ChangeTeamScore(team *database.Team, problem *database.Card, answer bool, correct bool, bonus bool) error {
	return s.teamEvent.ChangeTeamScore(team.Id, problem.Score)
}

func (s teamService) GetPodium() ([]*payload.Podium, error) {
	return nil, nil
}

func (s teamService) GetRanking() ([]*payload.TeamInfo, error) {
	teams, err := s.teamEvent.GetTeams()
	if err != nil {
		return nil, &response.Error{
			Message: "Unable to get teams",
		}
	}

	var rankings []*payload.TeamInfo
	for _, team := range teams {
		var totalScore int32
		if len(team.Scores) == 0 {
			totalScore = 0
		} else {
			totalScore = team.Scores[len(team.Scores)-1].Total
		}

		rankings = append(rankings, &payload.TeamInfo{
			Id:    team.Id,
			Name:  team.Name,
			Score: totalScore,
		})
	}
	return rankings, nil
}

func (s teamService) GetTeamInfos(team *database.Team) (*payload.TeamInfo, error) {
	return nil, nil
}

func (s teamService) UpdateScore(body *payload.UpdateScore) ([]*payload.TeamInfo, error) {
	currentCard, err := s.topicEvent.GetCurrentCard()
	if err != nil {
		return nil, &response.Error{
			Message: "Unable to get current card",
		}
	} else if currentCard == nil {
		return nil, &response.Error{
			Message: "No opening card",
		}
	}

	teams, err := s.teamEvent.GetTeams()
	if err != nil {
		return nil, &response.Error{
			Message: "Unable to get teams",
		}
	}

	for i, update := range body.Update {
		var currentScore int32
		if len(teams[i].Scores) == 0 {
			currentScore = 0
		} else {
			currentScore = teams[i].Scores[len(teams[i].Scores)-1].Total
		}

		switch update {
		case -1:
			teams[i].Scores = append(teams[i].Scores, &database.Score{
				Change: -currentCard.Score,
				Total:  currentScore - currentCard.Score,
			})
		case 0:
			teams[i].Scores = append(teams[i].Scores, &database.Score{Change: 0, Total: currentScore})
		case 1:
			teams[i].Scores = append(teams[i].Scores, &database.Score{
				Change: currentCard.Score,
				Total:  currentScore + currentCard.Score,
			})
		}
	}

	currentCard = nil
	hub.Snapshot()

	return s.GetRanking()
}
