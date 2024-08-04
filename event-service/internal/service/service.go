package service

import (
	"context"
	"log"
	genprotos "olympy/event-service/genproto/event_service"
	"olympy/event-service/internal/storage"
	"os"
)

type EventService struct {
	genprotos.UnimplementedEventServiceServer
	eventService storage.Event
	logger       *log.Logger
}

func NewEventService(eventStorage storage.Event) *EventService {
	return &EventService{
		eventService: eventStorage,
		logger:       log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *EventService) AddEvent(ctx context.Context, req *genprotos.AddEventRequest) (*genprotos.AddEventResponse, error) {
	s.logger.Println("Add Event Request")
	return s.eventService.AddEvent(ctx, req)
}

func (s *EventService) EditEvent(ctx context.Context, req *genprotos.EditEventRequest) (*genprotos.EditEventResponse, error) {
	s.logger.Println("Edit Event Request")
	return s.eventService.EditEvent(ctx, req)
}

func (s *EventService) DeleteEvent(ctx context.Context, req *genprotos.DeleteEventRequest) (*genprotos.Message, error) {
	s.logger.Println("Delete Event Request")
	return s.eventService.DeleteEvent(ctx, req)
}

func (s *EventService) GetEvent(ctx context.Context, req *genprotos.GetEventRequest) (*genprotos.GetEventResponse, error) {
	s.logger.Println("Get Event Request")
	return s.eventService.GetEvent(ctx, req)
}

func (s *EventService) GetAllEvents(ctx context.Context, req *genprotos.GetAllEventsRequest) (*genprotos.GetAllEventsResponse, error) {
	s.logger.Println("Get All Events Request")
	return s.eventService.GetAllEvents(ctx, req)
}

func (s *EventService) SearchEvents(ctx context.Context, req *genprotos.SearchEventsRequest) (*genprotos.GetAllEventsResponse, error) {
	s.logger.Println("Search Events Request")
	return s.eventService.SearchEvents(ctx, req)
}
