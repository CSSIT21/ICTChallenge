package services

import (
	"backend/repository"
	"backend/types/database"
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
	a, _ := s.teamEvent.GetTeams()
	return a, nil
}

func (s teamService) ChangeTeamScore(team *database.Team, problem *database.Question, correct bool, bonus bool) error {
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
