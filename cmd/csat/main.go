package main

import (
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	csatHttp "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/csat/delivery/http"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/postgres"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger.New(&cfg.Service.Logger)

	pg, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to create postgres client: %v", err)
	}

	s3, err := s3.New(&cfg.Minio)
	if err != nil {
		log.Fatalf("failed to create s3 client: %v", err)
	}

	metrics := metrics.New("backend", "csat")

	s := httpServer.New(cfg, pg, s3, logger, metrics)
	csatHttp.BindRoutes(s)
	fmt.Println("csat http binded")

	if err = s.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}