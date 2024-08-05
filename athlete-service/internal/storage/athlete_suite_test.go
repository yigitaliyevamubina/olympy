package storage

import (
	"context"
	"log"
	"testing"

	athleteservice "olympy/athlete-service/genproto/athlete_service"
	"olympy/athlete-service/internal/config"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"
)

type AthleteRepositoryTestSuite struct {
	suite.Suite
	repo        *Athlete
	CleanUpFunc func()
}

func (s *AthleteRepositoryTestSuite) SetupSuite() {
	
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := ConnectDB(*configs)
	s.Require().NoError(err)
	s.repo = &Athlete{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	s.CleanUpFunc = func() {
		db.Close()
	}
}

func (s *AthleteRepositoryTestSuite) TearDownSuite() {
	s.CleanUpFunc()
}

func (s *AthleteRepositoryTestSuite) TestAthleteCRUD() {
	ctx := context.Background()

	// Create athlete
	athlete := &athleteservice.Athlete{
		Id:        20,
		Name:      "John Doe",
		CountryId: 1,
		SportType: "Running",
	}
	createdAthlete, err := s.repo.AddAthlete(ctx, athlete)
	s.Require().NoError(err)
	s.Require().NotNil(createdAthlete)
	s.Equal(athlete.Name, createdAthlete.Name)
	s.Equal(athlete.CountryId, createdAthlete.CountryId)
	s.Equal(athlete.SportType, createdAthlete.SportType)

	// Update athlete
	athlete.Name = "John Updated"
	athlete.SportType = "Swimming"
	updatedAthlete, err := s.repo.EditAthlete(ctx, athlete)
	s.Require().NoError(err)
	s.Require().NotNil(updatedAthlete)
	s.Equal(athlete.Name, updatedAthlete.Name)
	s.Equal(athlete.SportType, updatedAthlete.SportType)

	// Get athlete
	fetchedAthlete, err := s.repo.GetAthlete(ctx, &athleteservice.GetSingleRequest{Id: athlete.Id})
	s.Require().NoError(err)
	s.Require().NotNil(fetchedAthlete)
	s.Equal(athlete.Name, fetchedAthlete.Name)
	s.Equal(athlete.CountryId, fetchedAthlete.CountryId)

	// List athletes
	listReq := &athleteservice.ListRequest{
		Page:      1,
		Limit:     10,
		CountryId: athlete.CountryId,
	}
	listResp, err := s.repo.ListAthletes(ctx, listReq)
	s.Require().NoError(err)
	s.Require().NotNil(listResp)
	s.NotZero(listResp.Count)
	s.NotEmpty(listResp.Athletes)

	// Delete athlete
	msg, err := s.repo.DeleteAthlete(ctx, &athleteservice.GetSingleRequest{Id: athlete.Id})
	s.Require().NoError(err)
	s.Require().NotNil(msg)
	s.Contains(msg.Message, "deleted successfully")
}

func TestAthleteRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AthleteRepositoryTestSuite))
}
