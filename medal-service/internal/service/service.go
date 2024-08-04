package service

import (
	"context"
	"log"
	countryservice "olympy/medal-service/genproto/country_service"
	modelservice "olympy/medal-service/genproto/medal_service"
	"olympy/medal-service/internal/storage"
	"os"
)

type (
	MedalService struct {
		modelservice.UnimplementedMedalServiceServer
		medalStorage   storage.Medal
		countryStorage storage.Country
		logger         *log.Logger
		countryservice.UnimplementedCountryServiceServer
	}
)

func New(modelservice storage.Medal, countryservice storage.Country) *MedalService {
	return &MedalService{
		medalStorage:   modelservice,
		countryStorage: countryservice,
		logger:         log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *MedalService) AddMedal(ctx context.Context, req *modelservice.Medal) (*modelservice.Medal, error) {
	s.logger.Println("Add Medal Request")
	return s.medalStorage.AddMedal(ctx, req)
}

func (s *MedalService) EditMedal(ctx context.Context, req *modelservice.Medal) (*modelservice.Medal, error) {
	s.logger.Println("Edit Medal Request")
	return s.medalStorage.EditMedal(ctx, req)
}

func (s *MedalService) DeleteMedal(ctx context.Context, req *modelservice.GetSingleRequest) (*modelservice.Message, error) {
	s.logger.Println("Delete Medal Request")
	return s.medalStorage.DeleteMedal(ctx, req)
}

func (s *MedalService) GetMedal(ctx context.Context, req *modelservice.GetSingleRequest) (*modelservice.Medal, error) {
	s.logger.Println("Get Medal Request")
	return s.medalStorage.GetMedal(ctx, req)
}

func (s *MedalService) ListMedals(ctx context.Context, req *modelservice.ListRequest) (*modelservice.ListResponse, error) {
	s.logger.Println("List Medals Request")
	return s.medalStorage.ListMedals(ctx, req)
}

func (s *MedalService) GetMedalRanking(ctx context.Context, req *modelservice.Empty) (*modelservice.MedalRankingResponse, error) {
	s.logger.Println("Get Medal Ranking Request")
	return s.medalStorage.GetMedalRanking(ctx, req)
}

// country service
func (s *MedalService) AddCountry(ctx context.Context, req *countryservice.Country) (*countryservice.Country, error) {
	s.logger.Println("Add Country Request")
	return s.countryStorage.AddCountry(ctx, req)
}

func (s *MedalService) EditCountry(ctx context.Context, req *countryservice.Country) (*countryservice.Country, error) {
	s.logger.Println("Edit Country Request")
	return s.countryStorage.EditCountry(ctx, req)
}

func (s *MedalService) DeleteCountry(ctx context.Context, req *countryservice.GetSingleRequest) (*countryservice.Message, error) {
	s.logger.Println("Delete Country Request")
	return s.countryStorage.DeleteCountry(ctx, req)
}

func (s *MedalService) GetCountry(ctx context.Context, req *countryservice.GetSingleRequest) (*countryservice.Country, error) {
	s.logger.Println("Get Country Request")
	return s.countryStorage.GetCountry(ctx, req)
}

func (s *MedalService) ListCountries(ctx context.Context, req *countryservice.ListRequest) (*countryservice.ListResponse, error) {
	s.logger.Println("List Countries Request")
	return s.countryStorage.ListCountries(ctx, req)
}
