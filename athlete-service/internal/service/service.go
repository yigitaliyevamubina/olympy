package service

import (
	"context"
	"log"
	athleteservice "olympy/athlete-service/genproto/athlete_service"
	"olympy/athlete-service/internal/storage"
	"os"
)

type (
	AthleteService struct {
		athleteservice.UnimplementedAthleteServiceServer
		athleteStorage storage.Athlete
		logger         *log.Logger
	}
)

func NewAthleteService(athleteservice storage.Athlete) *AthleteService {
	return &AthleteService{
		athleteStorage: athleteservice,
		logger:         log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *AthleteService) AddAthlete(ctx context.Context, req *athleteservice.Athlete) (*athleteservice.Athlete, error) {
	s.logger.Println("Add Medal Request")
	return s.athleteStorage.AddAthlete(ctx, req)
}

func (s *AthleteService) EditAthlete(ctx context.Context, req *athleteservice.Athlete) (*athleteservice.Athlete, error) {
	s.logger.Println("Edit Medal Request")
	return s.athleteStorage.EditAthlete(ctx, req)
}

func (s *AthleteService) DeleteAthlete(ctx context.Context, req *athleteservice.GetSingleRequest) (*athleteservice.Message, error) {
	s.logger.Println("Delete Medal Request")
	return s.athleteStorage.DeleteAthlete(ctx, req)
}

func (s *AthleteService) GetAthlete(ctx context.Context, req *athleteservice.GetSingleRequest) (*athleteservice.Athlete, error) {
	s.logger.Println("Get Medal Request")
	return s.athleteStorage.GetAthlete(ctx, req)
}

func (s *AthleteService) ListAthletes(ctx context.Context, req *athleteservice.ListRequest) (*athleteservice.ListResponse, error) {
	s.logger.Println("List Medals Request")
	return s.athleteStorage.ListAthletes(ctx, req)
}
