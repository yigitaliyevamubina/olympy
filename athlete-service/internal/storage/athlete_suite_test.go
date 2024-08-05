package storage

import (
	"context"
	athleteservice "olympy/athlete-service/genproto/athlete_service"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"
)

type AthleteRepositoryTestSuite struct {
	suite.Suite
	repo    *Athlete
	mock    sqlmock.Sqlmock
	cleanup func()
}

func (s *AthleteRepositoryTestSuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	if err != nil {
		s.FailNow("failed to create sqlmock", err)
	}
	s.mock = mock
	s.repo = &Athlete{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	s.cleanup = func() {
		db.Close()
	}
}

func (s *AthleteRepositoryTestSuite) TearDownSuite() {
	s.cleanup()
}

func (s *AthleteRepositoryTestSuite) TestAthleteCRUD() {
    ctx := context.Background()

    athlete := &athleteservice.Athlete{
        Id:        23,
        Name:      "John Doe",
        CountryId: 1,
        SportType: "Running",
        CreatedAt: time.Now().String(),
        UpdatedAt: time.Now().String(),
    }

    // Log the athlete object before executing the query
    s.T().Logf("Athlete object before INSERT: %+v", athlete)

    s.mock.ExpectExec(`INSERT INTO athletes`).
        WithArgs(athlete.CountryId, sqlmock.AnyArg(), athlete.Id, athlete.Name, athlete.SportType, sqlmock.AnyArg()).
        WillReturnResult(sqlmock.NewResult(1, 1))

    // Test AddAthlete
    createdAthlete, err := s.repo.AddAthlete(ctx, athlete)
    s.NoError(err)
    s.NotNil(createdAthlete)
    s.Equal(athlete.Name, createdAthlete.Name)

    // Log the created athlete object
    s.T().Logf("Created athlete object: %+v", createdAthlete)

    // Mock data for GetAthlete
    s.mock.ExpectQuery(`SELECT id, name, country_id, sport_type, created_at, updated_at FROM athletes WHERE id = \$1`).
        WithArgs(athlete.Id).
        WillReturnRows(sqlmock.NewRows([]string{"id", "name", "country_id", "sport_type", "created_at", "updated_at"}).
            AddRow(athlete.Id, athlete.Name, athlete.CountryId, athlete.SportType, athlete.CreatedAt, athlete.UpdatedAt))

    // Test GetAthlete
    fetchedAthlete, err := s.repo.GetAthlete(ctx, &athleteservice.GetSingleRequest{Id: athlete.Id})
    s.NoError(err)
    s.NotNil(fetchedAthlete)
    s.Equal(athlete.Name, fetchedAthlete.Name)

    // Log the fetched athlete object
    s.T().Logf("Fetched athlete object: %+v", fetchedAthlete)

    // Mock data for EditAthlete
    updatedAthlete := &athleteservice.Athlete{
        Id:        athlete.Id,
        Name:      "Jane Doe",
        CountryId: 2,
        SportType: "Swimming",
        CreatedAt: athlete.CreatedAt,
        UpdatedAt: time.Now().String(),
    }

    s.mock.ExpectExec(`UPDATE athletes`).
        WithArgs(updatedAthlete.Name, updatedAthlete.CountryId, updatedAthlete.SportType, sqlmock.AnyArg(), updatedAthlete.Id).
        WillReturnResult(sqlmock.NewResult(1, 1))

    s.mock.ExpectQuery(`SELECT id, name, country_id, sport_type, created_at, updated_at FROM athletes WHERE id = \$1`).
        WithArgs(updatedAthlete.Id).
        WillReturnRows(sqlmock.NewRows([]string{"id", "name", "country_id", "sport_type", "created_at", "updated_at"}).
            AddRow(updatedAthlete.Id, updatedAthlete.Name, updatedAthlete.CountryId, updatedAthlete.SportType, updatedAthlete.CreatedAt, updatedAthlete.UpdatedAt))

    // Test EditAthlete
    editedAthlete, err := s.repo.EditAthlete(ctx, updatedAthlete)
    s.NoError(err)
    s.NotNil(editedAthlete)
    s.Equal(updatedAthlete.Name, editedAthlete.Name)

    // Log the edited athlete object
    s.T().Logf("Edited athlete object: %+v", editedAthlete)

    // Mock data for DeleteAthlete
    s.mock.ExpectExec(`DELETE FROM athletes WHERE id = \$1`).
        WithArgs(athlete.Id).
        WillReturnResult(sqlmock.NewResult(1, 1))

    // Test DeleteAthlete
    deleteMsg, err := s.repo.DeleteAthlete(ctx, &athleteservice.GetSingleRequest{Id: athlete.Id})
    s.NoError(err)
    s.NotNil(deleteMsg)

    // Mock data for ListAthletes
    s.mock.ExpectQuery(`SELECT id, name, country_id, sport_type, created_at, updated_at FROM athletes LIMIT \$1 OFFSET \$2`).
        WithArgs(int64(5), int64(0)).
        WillReturnRows(sqlmock.NewRows([]string{"id", "name", "country_id", "sport_type", "created_at", "updated_at"}).
            AddRow(athlete.Id, athlete.Name, athlete.CountryId, athlete.SportType, athlete.CreatedAt, athlete.UpdatedAt))

    s.mock.ExpectQuery(`SELECT COUNT\(\*\) FROM athletes`).
        WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

    // Test ListAthletes
    listResponse, err := s.repo.ListAthletes(ctx, &athleteservice.ListRequest{Page: 1, Limit: 5})
    s.NoError(err)
    s.NotNil(listResponse)
    s.Equal(int64(1), listResponse.Count)
    s.Equal(athlete.Name, listResponse.Athletes[0].Name)
}


func TestAthleteRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AthleteRepositoryTestSuite))
}
