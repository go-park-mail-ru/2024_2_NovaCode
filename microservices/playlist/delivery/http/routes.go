package http

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/server"
	playlistRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/repository"
	playlistUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/usecase"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func BindRoutes(s *server.Server) {
	s.MUX.Handle("/metrics", promhttp.Handler())

	playlistRepo := playlistRepo.NewPlaylistRepository(s.PG)
	// trackRepo := trackRepo.NewTrackPGRepository(s.PG)
	// albumRepo := albumRepo.NewAlbumPGRepository(s.PG)
	// artistRepo := artistRepo.NewArtis.PGRepository(s.PG)
	// userRepo := userRepo.NewUserPostgresRepository(s.PG, s.Logger)
	// trackUsecase := trackUsecase.NewTrackUsecase(trackRepo, albumRepo, artistRepo, s.Logger)
	// playlistUsecase := playlistUsecase.NewPlaylistUsecase(trackUsecase, playlistRepo, trackRepo, userRepo, s.Logger)
	playlistUsecase := playlistUsecase.NewPlaylistUsecase(playlistRepo, s.Logger)
	playlistHandleres := NewPlaylistHandlers(playlistUsecase, s.Logger)

	// s.MUX.HandleFunc("/api/v1/playlists", playlistHandleres.GetAllPlaylists).Methods("GET")
	// s.MUX.Handle(
	// 	"/api/v1/playlists",
	// 	middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.CreatePlaylist)),
	// ).Methods("POST")
	// s.MUX.HandleFunc("/api/v1/playlists/{playlistId:[0-9]+}", playlistHandleres.GetPlaylist).Methods("GET")
	// s.MUX.HandleFunc("/api/v1/playlists/{playlistId:[0-9]+}/tracks", playlistHandleres.GetTracksFromPlaylist).Methods("GET")
	// s.MUX.HandleFunc("/api/v1/users/{userId:[0-9a-fA-F-]+}/playlists", playlistHandleres.GetUserPlaylists).Methods("GET")
	s.MUX.Handle(
		"/api/v1/playlists/{playlistId:[0-9]+}/tracks",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.AddToPlaylist)),
	).Methods("POST")
	s.MUX.Handle(
		"/api/v1/playlists/{playlistId:[0-9]+}/tracks",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.RemoveFromPlaylist)),
	).Methods("DELETE")
	s.MUX.Handle(
		"/api/v1/playlists/{playlistId:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.DeletePlaylist)),
	).Methods("DELETE")
}
