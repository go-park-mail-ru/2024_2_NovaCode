package http

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	artistRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/repository"
	artistUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/usecase"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func BindRoutes(s *httpServer.Server) {
	s.MUX.Handle("/metrics", promhttp.Handler())

	artistRepo := artistRepo.NewArtistPGRepository(s.PG)
	artistUsecase := artistUsecase.NewArtistUsecase(artistRepo, s.Logger)
	artistHandlers := NewArtistHandlers(artistUsecase, s.Logger)

	s.MUX.HandleFunc("/api/v1/artists/search", artistHandlers.SearchArtist).Methods("GET")
	s.MUX.HandleFunc("/api/v1/artists/{id:[0-9]+}", artistHandlers.ViewArtist).Methods("GET")
	s.MUX.HandleFunc("/api/v1/artists", artistHandlers.GetAll).Methods("GET")

	s.MUX.Handle(
		"/api/v1/artists/favorite/byUser/{userID}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(artistHandlers.GetFavoriteArtists)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/artists/favorite/{artistID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(artistHandlers.IsFavoriteArtist)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/artists/favorite/{artistID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(artistHandlers.AddFavoriteArtist)),
	).Methods("POST")

	s.MUX.Handle(
		"/api/v1/artists/favorite/{artistID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(artistHandlers.DeleteFavoriteArtist)),
	).Methods("DELETE")
}
