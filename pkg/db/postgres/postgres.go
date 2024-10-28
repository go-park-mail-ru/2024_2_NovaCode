package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	_ "github.com/lib/pq"
)

type Client *sql.DB

func New(cfg *config.PostgresConfig) (Client, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.Password,
	)

	db, err := sql.Open(cfg.Driver, connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxIdleLifetime) * time.Second)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping postgres instance: %v", err)
	}

	return db, nil
}
