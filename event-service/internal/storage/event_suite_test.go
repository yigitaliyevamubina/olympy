package storage

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	genprotos "olympy/event-service/genproto/event_service"
	"olympy/event-service/internal/config"

	"github.com/stretchr/testify/suite"
)

type EventRepositoryTestSuite struct {
	suite.Suite
	repo        *Event
	CleanUpFunc func()
}

func (s *EventRepositoryTestSuite) SetupSuite() {

	configs, err := config.New()
	if err != nil {
		fmt.Println("config")
		log.Fatal(err)
	}
	
	eventService, err := NewEventService(configs)
	s.Require().NoError(err)
	s.repo = eventService

	s.CleanUpFunc = func() {
		eventService.db.Close()
		eventService.redisClient.Close()
	}
}

func (s *EventRepositoryTestSuite) TearDownSuite() {
	s.CleanUpFunc()
}

func (s *EventRepositoryTestSuite) TestEventCRUD() {
	ctx := context.Background()

	// Add Event
	event := &genprotos.Event{
		Name:      "Olympic Games",
		SportType: "Multi-sport",
		StartTime: time.Now().Format(time.RFC3339),
		EndTime:   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}
	addReq := &genprotos.AddEventRequest{Event: event}
	addResp, err := s.repo.AddEvent(ctx, addReq)
	s.Require().NoError(err)
	s.Require().NotNil(addResp)
	s.Equal(event.Name, addResp.Event.Name)
	s.Equal(event.SportType, addResp.Event.SportType)

	// Edit Event
	event.Id = addResp.Event.Id
	event.Name = "Olympic Games 2024"
	event.SportType = "Summer Olympics"
	editReq := &genprotos.EditEventRequest{Event: event}
	editResp, err := s.repo.EditEvent(ctx, editReq)
	s.Require().NoError(err)
	s.Require().NotNil(editResp)
	s.Equal(event.Name, editResp.Event.Name)
	s.Equal(event.SportType, editResp.Event.SportType)

	// Get Event
	getReq := &genprotos.GetEventRequest{Id: strconv.Itoa(int(event.Id))}
	getResp, err := s.repo.GetEvent(ctx, getReq)
	s.Require().NoError(err)
	s.Require().NotNil(getResp)
	s.Equal(event.Name, getResp.Event.Name)
	s.Equal(event.SportType, getResp.Event.SportType)

	// List Events
	listReq := &genprotos.GetAllEventsRequest{Page: 1, PageSize: 10}
	listResp, err := s.repo.GetAllEvents(ctx, listReq)
	s.Require().NoError(err)
	s.Require().NotNil(listResp)
	s.NotZero(listResp.TotalCount)
	s.NotEmpty(listResp.Events)

	// Search Events
	searchReq := &genprotos.SearchEventsRequest{Query: "Olympic", Page: 1, PageSize: 10}
	searchResp, err := s.repo.SearchEvents(ctx, searchReq)
	s.Require().NoError(err)
	s.Require().NotNil(searchResp)
	s.NotZero(searchResp.TotalCount)
	s.NotEmpty(searchResp.Events)

	// Delete Event
	delReq := &genprotos.DeleteEventRequest{Id: strconv.Itoa(int(event.Id))}
	delResp, err := s.repo.DeleteEvent(ctx, delReq)
	s.Require().NoError(err)
	s.Require().NotNil(delResp)
	s.Contains(delResp.Message, "deleted successfully")
}

func TestEventRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(EventRepositoryTestSuite))
}
