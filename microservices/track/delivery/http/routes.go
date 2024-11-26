package http

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	trackRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/repository"
	trackUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/usecase"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func BindRoutes(s *httpServer.Server, artistClient artistService.ArtistServiceClient, albumClient albumService.AlbumServiceClient) {
	s.MUX.Handle("/metrics", promhttp.Handler())

	trackRepo := trackRepo.NewTrackPGRepository(s.PG)
	trackUsecase := trackUsecase.NewTrackUsecase(trackRepo, artistClient, albumClient, s.Logger)
	trackHandleres := NewTrackHandlers(trackUsecase, s.Logger)

	s.MUX.HandleFunc("/api/v1/tracks/search", trackHandleres.SearchTrack).Methods("GET")
	s.MUX.HandleFunc("/api/v1/tracks/{id:[0-9]+}", trackHandleres.ViewTrack).Methods("GET")
	s.MUX.HandleFunc("/api/v1/tracks", trackHandleres.GetAll).Methods("GET")
	s.MUX.HandleFunc("/api/v1/tracks/byArtistId/{artistId:[0-9]+}", trackHandleres.GetAllByArtistID).Methods("GET")
	s.MUX.HandleFunc("/api/v1/tracks/byAlbumId/{albumId:[0-9]+}", trackHandleres.GetAllByAlbumID).Methods("GET")

	s.MUX.Handle(
		"/api/v1/tracks/favorite",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(trackHandleres.GetFavoriteTracks)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/tracks/favorite/{trackID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(trackHandleres.IsFavoriteTrack)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/tracks/favorite/{trackID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(trackHandleres.AddFavoriteTrack)),
	).Methods("POST")

	s.MUX.Handle(
		"/api/v1/tracks/favorite/{trackID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(trackHandleres.DeleteFavoriteTrack)),
	).Methods("DELETE")
}
