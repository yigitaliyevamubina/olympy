package storage

import (
	"context"
	"fmt"
	"log"
	"testing"

	medalproto "olympy/medal-service/genproto/medal_service"
	"olympy/medal-service/internal/config"

	"github.com/stretchr/testify/suite"
)

type MedalRepositoryTestSuite struct {
	suite.Suite
	repo        *Medal
	CleanUpFunc func()
}

func (s *MedalRepositoryTestSuite) SetupSuite() {
	configs, err := config.New()
	if err != nil {
		fmt.Println("config")
		log.Fatal(err)
	}

	medalService, err := NewMedalService(configs)
	s.Require().NoError(err)
	s.repo = medalService

	s.CleanUpFunc = func() {
		medalService.db.Close()
	}
}

func (s *MedalRepositoryTestSuite) TearDownSuite() {
	s.CleanUpFunc()
}

func (s *MedalRepositoryTestSuite) TestMedalCRUD() {
	ctx := context.Background()

	// Add Medal
	medal := &medalproto.Medal{
		CountryId: 1,
		Type:      "Gold",
		EventId:   1,
		AthleteId: "1",
	}
	addResp, err := s.repo.AddMedal(ctx, medal)
	s.Require().NoError(err)
	s.Require().NotNil(addResp)
	s.Equal(medal.CountryId, addResp.CountryId)
	s.Equal(medal.Type, addResp.Type)

	// Edit Medal
	medal.Id = addResp.Id
	medal.Type = "Silver"
	editResp, err := s.repo.EditMedal(ctx, medal)
	s.Require().NoError(err)
	s.Require().NotNil(editResp)
	s.Equal(medal.Type, editResp.Type)

	// Get Medal
	getResp, err := s.repo.GetMedal(ctx, &medalproto.GetSingleRequest{Id: medal.Id})
	s.Require().NoError(err)
	s.Require().NotNil(getResp)
	s.Equal(medal.Type, getResp.Type)

	// List Medals
	listReq := &medalproto.ListRequest{Page: 1, Limit: 10}
	listResp, err := s.repo.ListMedals(ctx, listReq)
	s.Require().NoError(err)
	s.Require().NotNil(listResp)
	s.NotZero(listResp.Count)
	s.NotEmpty(listResp.Medals)

	// Delete Medal
	delReq := &medalproto.GetSingleRequest{Id: medal.Id}
	delResp, err := s.repo.DeleteMedal(ctx, delReq)
	s.Require().NoError(err)
	s.Require().NotNil(delResp)
	s.Contains(delResp.Message, "deleted successfully")
}

func (s *MedalRepositoryTestSuite) TestMedalRanking() {
	ctx := context.Background()

	// Medal Ranking
	rankingResp, err := s.repo.GetMedalRanking(ctx, &medalproto.Empty{})
	s.Require().NoError(err)
	s.Require().NotNil(rankingResp)
	s.NotEmpty(rankingResp.CountryMedalCounts)
}

func TestMedalRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MedalRepositoryTestSuite))
}
