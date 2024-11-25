package http

import (
	"fmt"
	"net/http"

	s3Repo "github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3/repository/s3"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	userRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/repository/postgres"
	userUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user/usecase"
)

func BindRoutes(s *httpServer.Server) {
	s.MUX.Handle("/metrics", promhttp.Handler())

	userPGRepo := userRepo.NewUserPostgresRepository(s.PG, s.Logger)
	userS3Repo := s3Repo.NewS3Repository(s.S3, s.Logger)
	userUsecase := userUsecase.NewUserUsecase(&s.CFG.Service.Auth, &s.CFG.Minio, userPGRepo, userS3Repo, s.Logger)
	userHandleres := NewUserHandlers(&s.CFG.Service.Auth, userUsecase, s.Logger)

	s.MUX.HandleFunc("/api/v1/health", userHandleres.Health).Methods("GET")

	s.MUX.HandleFunc("/api/v1/auth/register", userHandleres.Register).Methods("POST")
	s.MUX.HandleFunc("/api/v1/auth/login", userHandleres.Login).Methods("POST")

	s.MUX.Handle(
		"/api/v1/auth/csrf",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(userHandleres.GetCSRFToken)),
	).Methods("GET")

	s.MUX.HandleFunc("/api/v1/auth/logout", userHandleres.Logout).Methods("POST")

	s.MUX.HandleFunc("/api/v1/users/{username:[a-zA-Z0-9-_]+}", userHandleres.GetUserByUsername).Methods("GET")

	s.MUX.Handle(
		"/api/v1/users/me",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(userHandleres.GetMe)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/users/{user_id:[0-9a-fA-F-]+}",
		middleware.AuthMiddleware(
			&s.CFG.Service.Auth, s.Logger,
			middleware.CSRFMiddleware(&s.CFG.Service.Auth.CSRF, s.Logger, http.HandlerFunc(userHandleres.Update)),
		),
	).Methods("PUT")

	s.MUX.Handle(
		"/api/v1/users/{user_id:[0-9a-fA-F-]+}/image",
		middleware.AuthMiddleware(
			&s.CFG.Service.Auth, s.Logger,
			middleware.CSRFMiddleware(&s.CFG.Service.Auth.CSRF, s.Logger, http.HandlerFunc(userHandleres.UploadImage)),
		),
	).Methods("POST")

	s.MUX.Handle(
		"/api/v1/auth/health",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(userHandleres.Health)),
	).Methods("GET")

	fmt.Println("routes have binded")
}
