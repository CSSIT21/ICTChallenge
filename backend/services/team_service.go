package services

import (
	"backend/repository"
	"backend/types"
	"backend/types/database"
	"backend/types/response"
	"backend/utils/value"
)

type teamService struct {
	teamEvent repository.TeamRepository
}

func NewTeamService(teamRepository repository.TeamRepository) teamService {
	return teamService{teamEvent: teamRepository}
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

func (s teamService) ChangeTeamScore(team *database.Team, problem *database.Card, correct bool, bonus bool) error {
	return s.teamEvent.ChangeTeamScore(team.Id, problem.Score)
}

func (s teamService) GetPodium() ([]*database.Podium, error) {
	a, _ := s.teamEvent.GetPodium()
	return a, nil
}

func (s teamService) GetRanking() ([]*database.Ranking, error) {
	a, _ := s.teamEvent.GetRanking()
	return a, nil
}

func (s teamService) GetTeamInfo(team *database.Team) (*types.TeamInfo, error) {
	return nil, nil
}
