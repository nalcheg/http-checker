package repository

//go:generate mockgen -destination=../mocks/mock_repository.go -package=mocks github.com/nalcheg/http-checker/repository RepositoryInterface

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/nalcheg/http-checker/types"
)

type RepositoryInterface interface {
	GetHosts() ([]string, error)
	SaveCheck(result types.Result) error
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(dbDSN string) (*Repository, error) {
	db, err := pgxpool.Connect(context.Background(), dbDSN)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) GetHosts() ([]string, error) {
	var hosts []string
	rows, err := r.db.Query(context.Background(), `SELECT id AS host FROM hosts`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		val, err := rows.Values()
		if err != nil {
			return nil, err
		}
		hosts = append(hosts, val[0].(string))
	}

	return hosts, nil
}

func (r *Repository) SaveCheck(result types.Result) error {
	if _, err := r.db.Exec(
		context.Background(),
		`INSERT INTO checks (created_at, host, state, response_time) VALUES ($1, $2, $3, $4)`,
		result.Time,
		result.Host,
		result.ResponseCode,
		result.ResponseTime,
	); err != nil {
		return err
	}

	return nil
}
