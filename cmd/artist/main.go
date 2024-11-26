package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
	grpcServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/grpc"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/delivery/grpc/service"
	artistHttp "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/delivery/http"
	artistRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/repository"
	artistUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/usecase"
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

	metrics := metrics.New("backend", "artist")

	httpServer := httpServer.New(cfg, pg, s3, logger, metrics)
	artistHttp.BindRoutes(httpServer)

	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}
	}()

	artistPGRepo := artistRepo.NewArtistPGRepository(pg)
	artistUsecase := artistUsecase.NewArtistUsecase(artistPGRepo, logger)
	registerArtistService := artistService.RegisterArtistService(artistUsecase, logger)

	grpcServer := grpcServer.New(cfg, pg, s3, logger, metrics, registerArtistService)
	if err = grpcServer.Run(); err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}
}
