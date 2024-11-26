package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
	grpcServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/grpc"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	trackService "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/delivery/grpc/service"
	trackHttp "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/delivery/http"
	trackRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/repository"
	trackUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/usecase"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/postgres"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	s3Repo "github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3/repository/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	metrics := metrics.New("backend", "playlist")

	conn, err := grpc.NewClient("novamusic-album:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to album service: %v", err)
	}
	defer conn.Close()
	albumClient := albumService.NewTrackServiceClient(conn)

	conn, err := grpc.NewClient("novamusic-artist:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to artist service: %v", err)
	}
	defer conn.Close()
	artistClient := artistService.NewTrackServiceClient(conn)

	s := httpServer.New(cfg, pg, s3, logger, metrics)
	trackHttp.BindRoutes(s, albumClient. artistClient)

	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}
	}()

	trackPGRepo := trackRepo.NewTrackPostgresRepository(pg, logger)
	trackS3Repo := s3Repo.NewS3Repository(s3, logger)
	trackUsecase := trackUsecase.NewTrackUsecase(&cfg.Service.Auth, &cfg.Minio, trackPGRepo, trackS3Repo, logger)
	registerTrackService := trackService.RegisterTrackService(&cfg.Service.Auth, trackUsecase, logger)

	grpcServer := grpcServer.New(cfg, pg, s3, logger, metrics, registerTrackService)
	if err = grpcServer.Run(); err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}
}
