package service

import (
	"context"
	"log"
	modelservice "olympy/medal-service/genproto/medal_service"
	"olympy/medal-service/internal/storage"
	"os"
)

type (
	MedalService struct {
		modelservice.UnimplementedMedalServiceServer
		medalStorage storage.Medal
		logger       *log.Logger
	}
)

func New(service storage.Medal) *MedalService {
	return &MedalService{
		medalStorage: service,
		logger:       log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
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
	return s.medalStorage.GetAllMedals(ctx, req)
}
