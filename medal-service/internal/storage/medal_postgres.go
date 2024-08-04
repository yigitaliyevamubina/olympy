package storage

import (
	"context"
	"database/sql"
	"fmt"
	"olympy/medal-service/internal/config"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type Medal struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
}

func New(config *config.Config) (*Medal, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	return &Medal{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (m *Medal) AddMedal(ctx context.Context, req *genprotos.Medal) (*genprotos.Medal, error) {
	data := map[string]interface{}{
		"id":         uuid.New().ID(),
		"country_id": req.CountryId,
		"type":       req.Type,
		"event_id":   req.EventId,
		"athlete_id": req.AthleteId,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	query, args, err := m.queryBuilder.Insert("medals").
		SetMap(data).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := m.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return &genprotos.Medal{
		Id:        data["id"].(int64),
		CountryId: req.CountryId,
		Type:      req.Type,
		EventId:   req.EventId,
		AthleteId: req.AthleteId,
		CreatedAt: data["created_at"].(time.Time).String(),
		UpdatedAt: data["updated_at"].(time.Time).String(),
	}, nil
}

func (m *Medal) EditMedal(ctx context.Context, req *genprotos.Medal) (*genprotos.Medal, error) {
	data := map[string]interface{}{
		"country_id": req.CountryId,
		"type":       req.Type,
		"event_id":   req.EventId,
		"athlete_id": req.AthleteId,
		"updated_at": time.Now(),
	}

	query, args, err := m.queryBuilder.Update("medals").
		SetMap(data).
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := m.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	var updatedMedal genprotos.Medal
	err = m.db.QueryRowContext(ctx, "SELECT id, country_id, type, event_id, athlete_id, created_at, updated_at FROM medals WHERE id = $1", req.Id).
		Scan(&updatedMedal.Id, &updatedMedal.CountryId, &updatedMedal.Type, &updatedMedal.EventId, &updatedMedal.AthleteId, &updatedMedal.CreatedAt, &updatedMedal.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated medal: %v", err)
	}

	return &updatedMedal, nil
}

func (m *Medal) DeleteMedal(ctx context.Context, req *genprotos.GetSingleRequest) (*genprotos.Message, error) {
	query, args, err := m.queryBuilder.Delete("medals").
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("medal with ID %d not found", req.Id)
	}

	return &genprotos.Message{Message: fmt.Sprintf("Medal with ID %d deleted successfully", req.Id)}, nil
}

func (m *Medal) GetMedal(ctx context.Context, req *genprotos.GetSingleRequest) (*genprotos.Medal, error) {
	var medal genprotos.Medal
	err := m.db.QueryRowContext(ctx, "SELECT id, country_id, type, event_id, athlete_id, created_at, updated_at FROM medals WHERE id = $1", req.Id).
		Scan(&medal.Id, &medal.CountryId, &medal.Type, &medal.EventId, &medal.AthleteId, &medal.CreatedAt, &medal.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch medal: %v", err)
	}

	return &medal, nil
}

func (m *Medal) GetAllMedals(ctx context.Context, req *genprotos.ListRequest) (*genprotos.ListResponse, error) {
	var medals []*genprotos.Medal
	var total int64

	query := squirrel.Select("id", "country_id", "type", "event_id", "athlete_id", "created_at", "updated_at").
		From("medals").
		Limit(uint64(req.Limit)).
		Offset(uint64((req.Page - 1) * req.Limit))

	if req.Country != 0 {
		query = query.Where(squirrel.Eq{"country_id": req.Country})
	}
	if req.EventId != 0 {
		query = query.Where(squirrel.Eq{"event_id": req.EventId})
	}
	if req.AthleteId != "" {
		query = query.Where(squirrel.Eq{"athlete_id": req.AthleteId})
	}

	rows, err := query.RunWith(m.db).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch medals: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var medal genprotos.Medal
		err := rows.Scan(&medal.Id, &medal.CountryId, &medal.Type, &medal.EventId, &medal.AthleteId, &medal.CreatedAt, &medal.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan medal row: %v", err)
		}
		medals = append(medals, &medal)
	}

	err = m.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM medals").Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total number of medals: %v", err)
	}

	return &genprotos.ListResponse{
		Count:  total,
		Medals: medals,
	}, nil
}
