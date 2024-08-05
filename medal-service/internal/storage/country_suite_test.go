package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	countryproto "olympy/medal-service/genproto/country_service"
	"olympy/medal-service/internal/config"

	"github.com/stretchr/testify/suite"
)

type CountryRepositoryTestSuite struct {
	suite.Suite
	repo        *Country
	db          *sql.DB
	CleanUpFunc func()
}

func (s *CountryRepositoryTestSuite) SetupSuite() {
	configs, err := config.New()
	if err != nil {
		fmt.Println("config")
		log.Fatal(err)
	}

	countryService, err := NewCountryService(configs)
	s.Require().NoError(err)
	s.repo = countryService
	s.db = countryService.db

	s.CleanUpFunc = func() {
		s.db.Close()
	}
}

func (s *CountryRepositoryTestSuite) TearDownSuite() {
	s.CleanUpFunc()
}

func (s *CountryRepositoryTestSuite) TestCountryCRUD() {
	ctx := context.Background()

	// Add Country
	country := &countryproto.Country{
		Name: "Uzbekistan",
		Flag: "uz_flag_url",
	}
	addResp, err := s.repo.AddCountry(ctx, country)
	s.Require().NoError(err)
	s.Require().NotNil(addResp)
	s.Equal(country.Name, addResp.Name)
	s.Equal(country.Flag, addResp.Flag)

	// Edit Country
	country.Id = addResp.Id
	country.Name = "Republic of Uzbekistan"
	country.Flag = "new_uz_flag_url"
	editResp, err := s.repo.EditCountry(ctx, country)
	s.Require().NoError(err)
	s.Require().NotNil(editResp)
	s.Equal(country.Name, editResp.Name)
	s.Equal(country.Flag, editResp.Flag)

	// Get Country
	getResp, err := s.repo.GetCountry(ctx, &countryproto.GetSingleRequest{Id: country.Id})
	s.Require().NoError(err)
	s.Require().NotNil(getResp)
	s.Equal(country.Name, getResp.Name)
	s.Equal(country.Flag, getResp.Flag)

	// List Countries
	listReq := &countryproto.ListRequest{Page: 1, Limit: 10}
	listResp, err := s.repo.ListCountries(ctx, listReq)
	s.Require().NoError(err)
	s.Require().NotNil(listResp)
	s.NotZero(listResp.Count)
	s.NotEmpty(listResp.Countries)

	// Delete Country
	delReq := &countryproto.GetSingleRequest{Id: country.Id}
	delResp, err := s.repo.DeleteCountry(ctx, delReq)
	s.Require().NoError(err)
	s.Require().NotNil(delResp)
	s.Contains(delResp.Message, "deleted successfully")
}

func TestCountryRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CountryRepositoryTestSuite))
}
