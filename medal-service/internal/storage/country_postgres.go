package storage

import (
	"context"
	"database/sql"
	"fmt"
	"olympy/medal-service/internal/config"
	"time"

	countryproto "olympy/medal-service/genproto/country_service"

	"github.com/Masterminds/squirrel"
)

type Country struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
}

func NewCountryService(config *config.Config) (*Country, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	return &Country{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (c *Country) AddCountry(ctx context.Context, req *countryproto.Country) (*countryproto.Country, error) {
	data := map[string]interface{}{
		"name":       req.Name,
		"flag":       req.Flag,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	query, args, err := c.queryBuilder.Insert("countries").
		SetMap(data).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	var id int64

	if err := c.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to fetch inserted ID: %v", err)
	}

	return &countryproto.Country{
		Id:        id,
		Name:      req.Name,
		Flag:      req.Flag,
		CreatedAt: data["created_at"].(time.Time).String(),
		UpdatedAt: data["updated_at"].(time.Time).String(),
	}, nil
}

func (c *Country) EditCountry(ctx context.Context, req *countryproto.Country) (*countryproto.Country, error) {
	data := map[string]interface{}{
		"name":       req.Name,
		"flag":       req.Flag,
		"updated_at": time.Now(),
	}

	query, args, err := c.queryBuilder.Update("countries").
		SetMap(data).
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := c.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	var updatedCountry countryproto.Country
	err = c.db.QueryRowContext(ctx, "SELECT id, name, flag, created_at, updated_at FROM countries WHERE id = $1", req.Id).
		Scan(&updatedCountry.Id, &updatedCountry.Name, &updatedCountry.Flag, &updatedCountry.CreatedAt, &updatedCountry.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated country: %v", err)
	}

	return &updatedCountry, nil
}

func (c *Country) DeleteCountry(ctx context.Context, req *countryproto.GetSingleRequest) (*countryproto.Message, error) {
	query, args, err := c.queryBuilder.Delete("countries").
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	result, err := c.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("country with ID %d not found", req.Id)
	}

	return &countryproto.Message{Message: fmt.Sprintf("Country with ID %d deleted successfully", req.Id)}, nil
}

func (c *Country) GetCountry(ctx context.Context, req *countryproto.GetSingleRequest) (*countryproto.Country, error) {
	var country countryproto.Country
	err := c.db.QueryRowContext(ctx, "SELECT id, name, flag, created_at, updated_at FROM countries WHERE id = $1", req.Id).
		Scan(&country.Id, &country.Name, &country.Flag, &country.CreatedAt, &country.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch country: %v", err)
	}

	return &country, nil
}

func (c *Country) ListCountries(ctx context.Context, req *countryproto.ListRequest) (*countryproto.ListResponse, error) {
	var countries []*countryproto.Country
	var total int64

	query := squirrel.Select("id", "name", "flag", "created_at", "updated_at").
		From("countries").
		Limit(uint64(req.Limit)).
		Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(c.db).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch countries: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var country countryproto.Country
		err := rows.Scan(&country.Id, &country.Name, &country.Flag, &country.CreatedAt, &country.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan country row: %v", err)
		}
		countries = append(countries, &country)
	}

	err = c.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM countries").Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total number of countries: %v", err)
	}

	return &countryproto.ListResponse{
		Count:     total,
		Countries: countries,
	}, nil
}
