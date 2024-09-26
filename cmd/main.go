package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/server"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/postgres"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger.New(cfg)

	pg, err := postgres.New(cfg)
	if err != nil {
		log.Fatalf("failed to create postgres client: %v", err)
	}

	s := server.NewServer(cfg, pg, logger)
	if err = s.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
