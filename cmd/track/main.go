package main

import (
	"log"
	"sync"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	trackHttp "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/delivery/http"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/postgres"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
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

	metrics := metrics.New("backend", "track")

	connArtist, err := grpc.NewClient("novamusic-artist:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer connArtist.Close()
	artistClient := artistService.NewArtistServiceClient(connArtist)

	connAlbum, err := grpc.NewClient("novamusic-album:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer connAlbum.Close()
	albumClient := albumService.NewAlbumServiceClient(connAlbum)

	httpServer := httpServer.New(cfg, pg, s3, logger, metrics)
	trackHttp.BindRoutes(httpServer, artistClient, albumClient)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}
	}()
	wg.Wait()
}
