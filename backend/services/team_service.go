package services

import (
	"math/rand"
	"sort"

	"backend/loaders/hub"
	"backend/mappers"
	"backend/repository"
	"backend/types/database"
	"backend/types/enum"
	"backend/types/extend"
	"backend/types/message"
	"backend/types/payload"
	"backend/types/response"
)

type teamService struct {
	teamEvent  repository.TeamRepository
	topicEvent repository.TopicRepository
}

func NewTeamService(teamRepository repository.TeamRepository, topicRepository repository.TopicRepository) *teamService {
	return &teamService{teamEvent: teamRepository, topicEvent: topicRepository}
}

func (s *teamService) GetAllTeamInfos() ([]*payload.TeamInfo, error) {
	var teamInfos []*payload.TeamInfo
	for _, team := range s.teamEvent.GetTeams() {
		teamInfos = append(teamInfos, &payload.TeamInfo{
			Id:     team.Id,
			Name:   team.Name,
			School: team.School,
			Scores: team.Scores,
		})
	}

	return teamInfos, nil
}

func (s *teamService) GetCurrentScore(team *database.Team) int32 {
	var score int32
	if len(team.Scores) == 0 {
		score = 0
	} else {
		score = team.Scores[len(team.Scores)-1].Total
	}
	return score
}

func (s *teamService) GetPodium() []*payload.Podium {
	teams := s.teamEvent.GetTeams()

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
			Score:      s.GetCurrentScore(team),
			Percentile: float32(s.GetCurrentScore(team)-min) / float32(min+max),
		})
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Percentile > rankings[j].Percentile
	})

	return rankings
}

func (s *teamService) GetRanking() []*payload.TeamScore {
	teams := s.teamEvent.GetTeams()

	var rankings []*payload.TeamScore
	for _, team := range teams {
		var totalScore int32
		if len(team.Scores) == 0 {
			totalScore = 0
		} else {
			totalScore = s.GetCurrentScore(team)
		}

		rankings = append(rankings, &payload.TeamScore{
			Id:    team.Id,
			Name:  team.Name,
			Score: totalScore,
		})
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Score > rankings[j].Score
	})

	return rankings
}

func (s *teamService) UpdateScore(body *payload.UpdateScore) ([]*payload.TeamScore, error) {
	currentCard := s.topicEvent.GetCurrentCard()
	if currentCard == nil {
		return nil, &response.Error{
			Message: "No opening card",
		}
	}

	teams := s.teamEvent.GetTeams()

	turned := s.teamEvent.GetTurned()
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

	return s.GetRanking(), nil
}

func (s *teamService) GetNextTurn() *database.Team {
	turn := s.teamEvent.GetTurned()
	if len(turn) == 6 {
		s.teamEvent.SetTurned([]*database.Team{})
		turn = []*database.Team{}
	}

	var candidates []*database.Team
	for _, team := range s.teamEvent.GetTeams() {
		exist := false
		for _, t := range turn {
			if team.Id == t.Id {
				exist = true
				break
			}
		}
		if !exist {
			candidates = append(candidates, team)
		}
	}

	selected := candidates[rand.Intn(len(candidates))]
	s.teamEvent.SetTurned(append(turn, selected))

	return selected
}

func (s *teamService) GetStudentsTurn(team *database.Team) *payload.StudentTurn {
	student := &payload.StudentTurn{
		Name:    team.Name,
		Current: true,
		Topics:  mappers.DisplayTopic(s.topicEvent.GetTopics()),
	}

	return student
}

func (s *teamService) GetStudentConns() []*extend.ConnModel {
	return s.teamEvent.GetStudentConns()
}

func (s *teamService) GetAdminConn() *extend.ConnModel {
	return s.teamEvent.GetAdminConn()
}

func (s *teamService) GetLeaderboardConn() *extend.ConnModel {
	return s.teamEvent.GetLeaderboardConn()
}

func (s *teamService) SetMode(mode enum.Mode) {
	s.topicEvent.SetMode(mode)
	if mode == enum.ModePreview {
		s.topicEvent.SetPreviewCount(0)
		s.GetLeaderboardConn().Emit(&message.OutboundMessage{
			Event:   message.LeaderboardPreview,
			Payload: map[string]any{},
		})
	}
	if mode == enum.ModeStarted {
		rankings := s.GetRanking()
		s.GetLeaderboardConn().Emit(&message.OutboundMessage{
			Event: message.LeaderboardRanking,
			Payload: map[string]any{
				"rankings": rankings,
			},
		})
	}
	if mode == enum.ModeEnded {
		rankings := s.GetPodium()
		s.GetLeaderboardConn().Emit(&message.OutboundMessage{
			Event: message.LeaderboardPodium,
			Payload: map[string]any{
				"rankings": rankings,
			},
		})
	}
}

func (s *teamService) IncreasePreview() {
	if s.topicEvent.GetPreviewCount() < uint8(len(s.teamEvent.GetTeams())) {
		s.topicEvent.SetPreviewCount(s.topicEvent.GetPreviewCount() + 1)
		s.teamEvent.GetLeaderboardConn().Emit(&message.OutboundMessage{
			Event: message.LeaderboardPreviewAdd,
			Payload: map[string]any{
				"team": s.teamEvent.GetTeams()[s.topicEvent.GetPreviewCount()-1],
			},
		})
	}
}

func (s *teamService) GetLastTurn() *database.Team {
	turn := s.teamEvent.GetTurned()
	if len(turn) == 0 {
		return nil
	}
	return turn[len(turn)-1]
}
