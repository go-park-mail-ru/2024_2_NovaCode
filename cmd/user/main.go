package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
	grpcServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/grpc"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/delivery/grpc/service"
	userHttp "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/delivery/http"
	userRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/repository/postgres"
	userUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/usecase"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/postgres"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	s3Repo "github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3/repository/s3"
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

	metrics := metrics.New("backend", "user")

	httpServer := httpServer.New(cfg, pg, s3, logger, metrics)
	userHttp.BindRoutes(httpServer)

	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}
	}()

	userPGRepo := userRepo.NewUserPostgresRepository(pg, logger)
	userS3Repo := s3Repo.NewS3Repository(s3, logger)
	userUsecase := userUsecase.NewUserUsecase(&cfg.Service.Auth, &cfg.Minio, userPGRepo, userS3Repo, logger)
	registerUserService := userService.RegisterUserService(&cfg.Service.Auth, userUsecase, logger)

	grpcServer := grpcServer.New(cfg, pg, s3, logger, metrics, registerUserService)
	if err = grpcServer.Run(); err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}
}
