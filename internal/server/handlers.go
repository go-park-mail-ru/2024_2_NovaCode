package server

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	userHandlers "github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/delivery/http"
	userRepo "github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/repository/postgres"
	userUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/usecase"
)

func (s *Server) BindRoutes() {
	userRepo := userRepo.NewUserPostgresRepository(s.db)
	userUsecase := userUsecase.NewUserUsecase(s.cfg, userRepo, s.logger)
	userHandleres := userHandlers.NewUserHandlers(s.cfg, userUsecase, s.logger)

	s.mux.HandleFunc("GET /api/v1/health", userHandleres.Health)

	s.mux.HandleFunc("POST /api/v1/auth/register", userHandleres.Register)
	s.mux.HandleFunc("POST /api/v1/auth/login", userHandleres.Login)
	s.mux.HandleFunc("POST /api/v1/auth/logout", userHandleres.Logout)

	// auth middleware usage example
	s.mux.Handle(
		"GET /api/v1/auth/health",
		middleware.AuthMiddleware(s.cfg, s.logger, http.HandlerFunc(userHandleres.Health)),
	)
}
