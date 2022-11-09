package services

import (
	"backend/loaders/hub"
	"backend/repository"
	"backend/types/database"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
	"sort"
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

func (s teamService) GetCurrentScore(team *database.Team) int32 {
	return team.Scores[len(team.Scores)-1].Total
}

func (s teamService) GetPodium() ([]*payload.Podium, error) {
	teams, err := s.teamEvent.GetTeams()
	if err != nil {
		return nil, &response.Error{
			Message: "Unable to get teams",
		}
	}

	var min, max int32
	for _, team := range teams {
		if min > s.GetCurrentScore(team) {
			min = s.GetCurrentScore(team)
		}
		if max < s.GetCurrentScore(team) {
			max = s.GetCurrentScore(team)
		}
	}

	var rankings []*payload.Podium
	for _, team := range teams {
		rankings = append(rankings, &payload.Podium{
			Id:         team.Id,
			Name:       team.Name,
			Percentile: float32(s.GetCurrentScore(team)-min) / float32(min+max),
		})
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Percentile > rankings[j].Percentile
	})

	return rankings, nil
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
			totalScore = s.GetCurrentScore(team)
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
	currentCard := s.topicEvent.GetCurrentCard()
	if currentCard == nil {
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

	turned, err := s.teamEvent.GetTurned()
	for i, update := range body.Update {
		var currentScore int32
		if len(teams[i].Scores) == 0 {
			currentScore = 0
		} else {
			currentScore = s.GetCurrentScore(teams[i])
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
			change := currentCard.Score
			if teams[i] == turned[len(turned)-1] && currentCard.Bonus {
				change *= 2
			}

			teams[i].Scores = append(teams[i].Scores, &database.Score{
				Change: change,
				Total:  currentScore + change,
			})
		}
	}

	s.topicEvent.SetCurrentCard(nil)
	hub.Snapshot()

	return s.GetRanking()
}
