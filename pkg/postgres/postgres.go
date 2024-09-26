package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	_ "github.com/lib/pq"
)

func New(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.DBName,
		cfg.Postgres.Password,
	)

	db, err := sql.Open(cfg.Postgres.Driver, connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.Postgres.ConnMaxIdleLifetime) * time.Second)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(cfg.Postgres.ConnMaxIdleTime) * time.Second)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping postgres instance: %v", err)
	}

	return db, nil
}
