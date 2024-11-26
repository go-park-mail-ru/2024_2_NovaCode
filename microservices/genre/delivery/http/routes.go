package http

import (
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	genreRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/genre/repository"
	genreUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/genre/usecase"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func BindRoutes(s *httpServer.Server) {
	s.MUX.Handle("/metrics", promhttp.Handler())

	genreRepo := genreRepo.NewGenrePGRepository(s.PG)
	genreUsecase := genreUsecase.NewGenreUsecase(genreRepo, s.Logger)
	genreHandleres := NewGenreHandlers(genreUsecase, s.Logger)

	s.MUX.HandleFunc("/api/v1/genres", genreHandleres.GetAll).Methods("GET")
	s.MUX.HandleFunc("/api/v1/genres/byArtistId/{artistId:[0-9]+}", genreHandleres.GetAllByArtistID).Methods("GET")
	s.MUX.HandleFunc("/api/v1/genres/byTrackId/{trackId:[0-9]+}", genreHandleres.GetAllByTrackID).Methods("GET")
}
