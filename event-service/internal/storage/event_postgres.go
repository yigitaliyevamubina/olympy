package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	genprotos "olympy/event-service/genproto/event_service"
	"olympy/event-service/internal/config"

	"github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v8"
)

type Event struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
	redisClient  *redis.Client
}

func NewEventService(config *config.Config) (*Event, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	return &Event{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		redisClient:  redisClient,
	}, nil
}

func (e *Event) AddEvent(ctx context.Context, req *genprotos.AddEventRequest) (*genprotos.AddEventResponse, error) {
	data := map[string]interface{}{
		"name":       req.Event.Name,
		"sport_type": req.Event.SportType,
		"start_time": req.Event.StartTime,
		"end_time":   req.Event.EndTime,
	}

	query, args, err := e.queryBuilder.Insert("events").
		SetMap(data).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	var id int64

	if err := e.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &genprotos.AddEventResponse{
		Event: &genprotos.Event{
			Id:        id,
			Name:      req.Event.Name,
			SportType: req.Event.SportType,
			StartTime: req.Event.StartTime,
			EndTime:   req.Event.EndTime,
		},
	}, nil
}

func (e *Event) EditEvent(ctx context.Context, req *genprotos.EditEventRequest) (*genprotos.EditEventResponse, error) {
	data := map[string]interface{}{
		"name":       req.Event.Name,
		"sport_type": req.Event.SportType,
		"start_time": req.Event.StartTime,
		"end_time":   req.Event.EndTime,
	}

	query, args, err := e.queryBuilder.Update("events").
		SetMap(data).
		Where(squirrel.Eq{"id": req.Event.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := e.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	idStr := strconv.Itoa(int(req.Event.Id))
	// Invalidate cache for this event
	e.redisClient.Del(ctx, idStr)

	return &genprotos.EditEventResponse{
		Event: &genprotos.Event{
			Id:        req.Event.Id,
			Name:      req.Event.Name,
			SportType: req.Event.SportType,
			StartTime: req.Event.StartTime,
			EndTime:   req.Event.EndTime,
		},
	}, nil
}

func (e *Event) DeleteEvent(ctx context.Context, req *genprotos.DeleteEventRequest) (*genprotos.Message, error) {
	query, args, err := e.queryBuilder.Delete("events").
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	result, err := e.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("event with ID %s not found", req.Id)
	}

	// Invalidate cache for this event
	e.redisClient.Del(ctx, req.Id)

	return &genprotos.Message{Message: fmt.Sprintf("Event with ID %s deleted successfully", req.Id)}, nil
}

func (e *Event) GetEvent(ctx context.Context, req *genprotos.GetEventRequest) (*genprotos.GetEventResponse, error) {
	// Check cache first
	cachedEvent, err := e.redisClient.Get(ctx, req.Id).Result()
	if err == nil {
		var event genprotos.Event
		if err := json.Unmarshal([]byte(cachedEvent), &event); err == nil {
			return &genprotos.GetEventResponse{Event: &event}, nil
		}
	}

	var event genprotos.Event
	err = e.db.QueryRowContext(ctx, "SELECT id, name, sport_type, start_time, end_time FROM events WHERE id = $1", req.Id).
		Scan(&event.Id, &event.Name, &event.SportType, &event.StartTime, &event.EndTime)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch event: %v", err)
	}

	// Cache the result
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event: %v", err)
	}
	e.redisClient.Set(ctx, req.Id, eventJSON, 10*time.Minute)

	return &genprotos.GetEventResponse{Event: &event}, nil
}

func (e *Event) GetAllEvents(ctx context.Context, req *genprotos.GetAllEventsRequest) (*genprotos.GetAllEventsResponse, error) {
	// Determine pagination offsets
	page := req.Page
	pageSize := req.PageSize
	offset := (page - 1) * pageSize

	// Check cache first
	cacheKey := fmt.Sprintf("events_page_%d_size_%d", page, pageSize)
	cachedEvents, err := e.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var eventsResponse genprotos.GetAllEventsResponse
		if err := json.Unmarshal([]byte(cachedEvents), &eventsResponse); err == nil {
			return &eventsResponse, nil
		}
	}

	// Get all events with pagination
	query := "SELECT id, name, sport_type, start_time, end_time FROM events LIMIT $1 OFFSET $2"
	rows, err := e.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %v", err)
	}
	defer rows.Close()

	var events []*genprotos.Event
	for rows.Next() {
		var event genprotos.Event
		if err := rows.Scan(&event.Id, &event.Name, &event.SportType, &event.StartTime, &event.EndTime); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %v", err)
	}

	countQuery := "SELECT COUNT(*) FROM events"
	var totalCount int32
	err = e.db.QueryRowContext(ctx, countQuery).Scan(&totalCount)
	if err != nil {
		return nil, fmt.Errorf("failed to count events: %v", err)
	}

	eventsResponse := &genprotos.GetAllEventsResponse{
		Events:     events,
		TotalCount: totalCount,
	}

	// Cache the result
	eventsJSON, err := json.Marshal(eventsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal events response: %v", err)
	}
	e.redisClient.Set(ctx, cacheKey, eventsJSON, 10*time.Minute)

	return eventsResponse, nil
}

func (e *Event) SearchEvents(ctx context.Context, req *genprotos.SearchEventsRequest) (*genprotos.GetAllEventsResponse, error) {
	// Determine pagination offsets
	page := req.Page
	pageSize := req.PageSize
	offset := (page - 1) * pageSize

	// Check cache first
	cacheKey := fmt.Sprintf("events_search_%s_page_%d_size_%d", req.Query, page, pageSize)
	cachedEvents, err := e.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var eventsResponse genprotos.GetAllEventsResponse
		if err := json.Unmarshal([]byte(cachedEvents), &eventsResponse); err == nil {
			return &eventsResponse, nil
		}
	}

	// Get events with search query and pagination
	query := `SELECT id, name, sport_type, start_time, end_time 
			  FROM events 
			  WHERE name ILIKE '%' || $1 || '%' 
			  OR sport_type ILIKE '%' || $1 || '%' 
			  LIMIT $2 OFFSET $3`
	rows, err := e.db.QueryContext(ctx, query, req.Query, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %v", err)
	}
	defer rows.Close()

	var events []*genprotos.Event
	for rows.Next() {
		var event genprotos.Event
		if err := rows.Scan(&event.Id, &event.Name, &event.SportType, &event.StartTime, &event.EndTime); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %v", err)
	}

	countQuery := `SELECT COUNT(*) 
				   FROM events 
				   WHERE name ILIKE '%' || $1 || '%' 
				   OR sport_type ILIKE '%' || $1 || '%'`
	var totalCount int32
	err = e.db.QueryRowContext(ctx, countQuery, req.Query).Scan(&totalCount)
	if err != nil {
		return nil, fmt.Errorf("failed to count events: %v", err)
	}

	eventsResponse := &genprotos.GetAllEventsResponse{
		Events:     events,
		TotalCount: totalCount,
	}

	// Cache the result
	eventsJSON, err := json.Marshal(eventsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal events response: %v", err)
	}
	e.redisClient.Set(ctx, cacheKey, eventsJSON, 10*time.Minute)

	return eventsResponse, nil
}
