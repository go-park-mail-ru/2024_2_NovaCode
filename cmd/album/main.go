package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
	grpcServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/grpc"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/delivery/grpc/service"
	albumHttp "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/delivery/http"
	albumRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/repository"
	albumUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/usecase"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/postgres"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
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

	metrics := metrics.New("backend", "artist")

	conn, err := grpc.NewClient("novamusic-artist:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to artist service: %v", err)
	}
	defer conn.Close()
	artistClient := artistService.NewArtistServiceClient(conn)

	s := httpServer.New(cfg, pg, s3, logger, metrics)
	albumHttp.BindRoutes(s, artistClient)

	if err = s.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}
	}()

	albumPGRepo := albumRepo.NewAlbumPostgresRepository(pg, logger)
	albumUsecase := albumUsecase.NewAlbumUsecase(albumPGRepo, artistClient, logger)
	registerAlbumService := albumService.RegisterAlbumService(&cfg.Service.Auth, albumUsecase, logger)

	grpcServer := grpcServer.New(cfg, pg, s3, logger, metrics, registerAlbumService)
	if err = grpcServer.Run(); err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}
}
