package main

import (
	"log"
	"sync"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	playlistHttp "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/delivery/http"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/postgres"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"
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

	conn, err := grpc.NewClient("novamusic-user:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer conn.Close()
	userClient := userService.NewUserServiceClient(conn)

	s := httpServer.New(cfg, pg, s3, logger, metrics)
	playlistHttp.BindRoutes(s, userClient)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = s.Run(); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
	}()
	wg.Wait()
}
