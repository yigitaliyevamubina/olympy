package storage

import (
	"context"
	"database/sql"
	"fmt"
	genprotos "olympy/event-service/genproto/event_service"
	"olympy/event-service/internal/config"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type Event struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
}

func NewEventService(config *config.Config) (*Event, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	return &Event{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (e *Event) AddEvent(ctx context.Context, req *genprotos.AddEventRequest) (*genprotos.AddEventResponse, error) {
	data := map[string]interface{}{
		"id":         uuid.NewString(),
		"name":       req.Event.Name,
		"sport_type": req.Event.SportType,
		"start_time": req.Event.StartTime,
		"end_time":   req.Event.EndTime,
	}

	query, args, err := e.queryBuilder.Insert("events").
		SetMap(data).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := e.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return &genprotos.AddEventResponse{
		Event: &genprotos.Event{
			Id:        data["id"].(string),
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

	return &genprotos.Message{Message: fmt.Sprintf("Event with ID %s deleted successfully", req.Id)}, nil
}

func (e *Event) GetEvent(ctx context.Context, req *genprotos.GetEventRequest) (*genprotos.GetEventResponse, error) {
	var event genprotos.GetEventResponse
	err := e.db.QueryRowContext(ctx, "SELECT id, name, sport_type, start_time, end_time FROM events WHERE id = $1", req.Id).
		Scan(&event.Event.Id, &event.Event.Name, &event.Event.SportType, &event.Event.StartTime, &event.Event.EndTime)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch event: %v", err)
	}

	return &event, nil
}

func (e *Event) GetAllEvents(ctx context.Context, req *genprotos.GetAllEventsRequest) (*genprotos.GetAllEventsResponse, error) {
	// Determine pagination offsets
	page := req.Page
	pageSize := req.PageSize
	offset := (page - 1) * pageSize

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

	return &genprotos.GetAllEventsResponse{
		Events:     events,
		TotalCount: totalCount,
	}, nil
}

func (e *Event) SearchEvents(ctx context.Context, req *genprotos.SearchEventsRequest) (*genprotos.GetAllEventsResponse, error) {
	query, args, err := e.queryBuilder.Select("id", "name", "sport_type", "start_time", "end_time").
		From("events").
		Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", req.Query)}).
		Limit(uint64(req.PageSize)).
		Offset(uint64((req.Page - 1) * req.PageSize)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search events: %v", err)
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

	countQuery, args, err := e.queryBuilder.Select("COUNT(*)").
		From("events").
		Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", req.Query)}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build count query: %v", err)
	}

	var total int32
	err = e.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch total count: %v", err)
	}

	return &genprotos.GetAllEventsResponse{Events: events, TotalCount: total}, nil
}
