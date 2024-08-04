package service

import (
	"context"
	"log"
	"olympy/medal-service/genprotos"
	"olympy/medal-service/internal/storage"
	"os"
)

type (
	MedalService struct {
		genprotos.UnimplementedMedalServiceServer
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

func (s *MedalService) AddMedal(ctx context.Context, req *genprotos.AddMedalRequest) (*genprotos.AddMedalResponse, error) {
	s.logger.Println("Add Medal Request")
	return s.medalStorage.AddMedal(ctx, req)
}

func (s *MedalService) EditMedal(ctx context.Context, req *genprotos.EditMedalRequest) (*genprotos.EditMedalResponse, error) {
	s.logger.Println("Edit Medal Request")
	return s.medalStorage.EditMedal(ctx, req)
}

func (s *MedalService) DeleteMedal(ctx context.Context, req *genprotos.DeleteMedalRequest) (*genprotos.Message, error) {
	s.logger.Println("Delete Medal Request")
	return s.medalStorage.DeleteMedal(ctx, req)
}

func (s *MedalService) GetMedal(ctx context.Context, req *genprotos.GetMedalRequest) (*genprotos.GetMedalResponse, error) {
	s.logger.Println("Get Medal Request")
	return s.medalStorage.GetMedal(ctx, req)
}

func (s *MedalService) ListMedals(ctx context.Context, req *genprotos.ListMedalsRequest) (*genprotos.ListMedalsResponse, error) {
	s.logger.Println("List Medals Request")
	return s.medalStorage.GetAllMedals(ctx, req)
}
