package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	athleteservice "olympy/athlete-service/genproto/athlete_service"
	"olympy/athlete-service/internal/config"

	"github.com/Masterminds/squirrel"
	//"github.com/google/uuid"
)

type Athlete struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
}

func NewAthleteService(config *config.Config) (*Athlete, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	return &Athlete{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (a *Athlete) AddAthlete(ctx context.Context, req *athleteservice.Athlete) (*athleteservice.Athlete, error) {
	data := map[string]interface{}{
		"id":         req.Id,
		"name":       req.Name,
		"country_id": req.CountryId,
		"sport_type": req.SportType,
		"created_at": time.Now(),
	}

	query, args, err := a.queryBuilder.Insert("athletes").
		SetMap(data).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	var id int64

	if err := a.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &athleteservice.Athlete{
		Id:        id,
		Name:      req.Name,
		CountryId: req.CountryId,
		SportType: req.SportType,
		CreatedAt: data["created_at"].(time.Time).String(),
	}, nil
}

func (a *Athlete) EditAthlete(ctx context.Context, req *athleteservice.Athlete) (*athleteservice.Athlete, error) {
	data := map[string]interface{}{
		"name":       req.Name,
		"country_id": req.CountryId,
		"sport_type": req.SportType,
		"updated_at": time.Now(),
	}

	query, args, err := a.queryBuilder.Update("athletes").
		SetMap(data).
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := a.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	var updatedAthlete athleteservice.Athlete
	err = a.db.QueryRowContext(ctx, "SELECT id, name, country_id, sport_type, created_at, updated_at FROM athletes WHERE id = $1", req.Id).
		Scan(&updatedAthlete.Id, &updatedAthlete.Name, &updatedAthlete.CountryId, &updatedAthlete.SportType, &updatedAthlete.CreatedAt, &updatedAthlete.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated athlete: %v", err)
	}

	return &updatedAthlete, nil
}

func (a *Athlete) DeleteAthlete(ctx context.Context, req *athleteservice.GetSingleRequest) (*athleteservice.Message, error) {
	query, args, err := a.queryBuilder.Delete("athletes").
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	result, err := a.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("athlete with ID %d not found", req.Id)
	}

	return &athleteservice.Message{Message: fmt.Sprintf("Athlete with ID %d deleted successfully", req.Id)}, nil
}

func (a *Athlete) GetAthlete(ctx context.Context, req *athleteservice.GetSingleRequest) (*athleteservice.Athlete, error) {
	var athlete athleteservice.Athlete
	err := a.db.QueryRowContext(ctx, "SELECT id, name, country_id, sport_type, created_at, updated_at FROM athletes WHERE id = $1", req.Id).
		Scan(&athlete.Id, &athlete.Name, &athlete.CountryId, &athlete.SportType, &athlete.CreatedAt, &athlete.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch athlete: %v", err)
	}

	return &athlete, nil
}

func (a *Athlete) ListAthletes(ctx context.Context, req *athleteservice.ListRequest) (*athleteservice.ListResponse, error) {
	var athletes []*athleteservice.Athlete
	var total int64

	query := a.queryBuilder.Select("id", "name", "country_id", "sport_type", "created_at", "updated_at").
		From("athletes").
		Limit(uint64(req.Limit)).
		Offset(uint64((req.Page - 1) * req.Limit))

	if req.CountryId != 0 {
		query = query.Where(squirrel.Eq{"country_id": req.CountryId})
	}
	if req.SportType != "" {
		query = query.Where(squirrel.Eq{"sport_type": req.SportType})
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	// fmt.Printf("Generated SQL: %s, Args: %v\n", sql, args)

	rows, err := a.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch athletes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var athlete athleteservice.Athlete
		err := rows.Scan(&athlete.Id, &athlete.Name, &athlete.CountryId, &athlete.SportType, &athlete.CreatedAt, &athlete.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan athlete row: %v", err)
		}
		athletes = append(athletes, &athlete)
	}

	err = a.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM athletes").Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total number of athletes: %v", err)
	}

	return &athleteservice.ListResponse{
		Count:    total,
		Athletes: athletes,
	}, nil
}

