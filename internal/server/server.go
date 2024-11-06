package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/postgres"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
)

type Server struct {
	mux    *mux.Router
	cfg    *config.Config
	pg     postgres.Client
	s3     *minio.Client
	logger logger.Logger
}

func NewServer(cfg *config.Config, pg postgres.Client, s3 *minio.Client, logger logger.Logger) *Server {
	return &Server{mux.NewRouter(), cfg, pg, s3, logger}
}

func (s *Server) Run() error {
	s.BindRoutes()

	corsedMux := middleware.CORSMiddleware(&s.cfg.Service.CORS, s.mux)
	loggedCorsedMux := middleware.LoggingMiddleware(&s.cfg.Service, s.logger, corsedMux)

	server := &http.Server{
		Addr:         s.cfg.Service.Port,
		Handler:      loggedCorsedMux,
		ReadTimeout:  s.cfg.Service.ReadTimeout * time.Second,
		WriteTimeout: s.cfg.Service.WriteTimeout * time.Second,
		IdleTimeout:  s.cfg.Service.IdleTimeout * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), s.cfg.Service.ContextTimeout*time.Second)
	defer shutdown()

	return server.Shutdown(ctx)
}
