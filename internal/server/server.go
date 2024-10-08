package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
)

type Server struct {
	mux    *mux.Router
	cfg    *config.ServiceConfig
	db     *sql.DB
	logger logger.Logger
}

func NewServer(cfg *config.ServiceConfig, db *sql.DB, logger logger.Logger) *Server {
	return &Server{mux.NewRouter(), cfg, db, logger}
}

func (s *Server) Run() error {
	s.BindRoutes()

	corsedMux := middleware.CORSMiddleware(&s.cfg.CORS, s.mux)
	loggedCorsedMux := middleware.LoggingMiddleware(s.cfg, s.logger, corsedMux)

	server := &http.Server{
		Addr:         s.cfg.Port,
		Handler:      loggedCorsedMux,
		ReadTimeout:  s.cfg.ReadTimeout * time.Second,
		WriteTimeout: s.cfg.WriteTimeout * time.Second,
		IdleTimeout:  s.cfg.IdleTimeout * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), s.cfg.ContextTimeout*time.Second)
	defer shutdown()

	return server.Shutdown(ctx)
}
