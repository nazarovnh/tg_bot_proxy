package database

import (
	"context"
	"time"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	ConnectionURI   string        `yaml:"connectionURI"`
	MaxConn         int32         `yaml:"maxConn"`
	MaxConnLifetime time.Duration `yaml:"maxConnLifetime"`
}

func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.ConnectionURI)
	if err != nil {
		return nil, errors.New("failed parse connectionURI for build pool")
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, errors.New("failed ping to database")
	}

	return pool, nil
}
