package http

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	csatRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/csat/repository"
	csatUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/csat/usecase"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func BindRoutes(s *httpServer.Server) {
	s.MUX.Handle("/metrics", promhttp.Handler())

	csatRepo := csatRepo.NewCSATPGRepository(s.PG, s.Logger)
	csatUsecase := csatUsecase.NewCSATUsecase(csatRepo, s.Logger)
	csatHandlers := NewCSATHandlers(csatUsecase, s.Logger)

	s.MUX.Handle(
		"/api/v1/csat/stat",
		middleware.AuthMiddleware(
			&s.CFG.Service.Auth, s.Logger,
			middleware.AdminMiddleware(
				&s.CFG.Service.Auth, s.Logger,
				http.HandlerFunc(csatHandlers.GetStatistics),
			),
		),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/csat/questions",
		middleware.AuthMiddleware(
			&s.CFG.Service.Auth, s.Logger,
			middleware.CSRFMiddleware(&s.CFG.Service.Auth.CSRF, s.Logger, http.HandlerFunc(csatHandlers.GetQuestionsByTopic)),
		),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/csat/questions/{questionID:[0-9]+}/submit",
		middleware.AuthMiddleware(
			&s.CFG.Service.Auth, s.Logger,
			middleware.CSRFMiddleware(&s.CFG.Service.Auth.CSRF, s.Logger, http.HandlerFunc(csatHandlers.SubmitAnswer)),
		),
	).Methods("POST")

	fmt.Println("routes have binded")
}
