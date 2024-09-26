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

	"github.com/daronenko/auth/config"
	"github.com/daronenko/auth/pkg/logger"
)

type Server struct {
	mux    *http.ServeMux
	cfg    *config.Config
	db     *sql.DB
	logger logger.Logger
}

func NewServer(cfg *config.Config, db *sql.DB, logger logger.Logger) *Server {
	return &Server{http.NewServeMux(), cfg, db, logger}
}

func (s *Server) Run() error {
	s.BindRoutes()

	server := &http.Server{
		Addr:         s.cfg.Server.Port,
		Handler:      s.mux,
		ReadTimeout:  s.cfg.Server.ReadTimeout * time.Second,
		WriteTimeout: s.cfg.Server.WriteTimeout * time.Second,
		IdleTimeout:  s.cfg.Server.IdleTimeout * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), s.cfg.Server.ContextTimeout*time.Second)
	defer shutdown()

	return server.Shutdown(ctx)
}
