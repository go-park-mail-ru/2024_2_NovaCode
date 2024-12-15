package http

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	playlistRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/repository"
	playlistUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/usecase"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func BindRoutes(s *httpServer.Server, userClient userService.UserServiceClient) {
	s.MUX.Handle("/metrics", promhttp.Handler())

	playlistRepo := playlistRepo.NewPlaylistRepository(s.PG)
	playlistUsecase := playlistUsecase.NewPlaylistUsecase(playlistRepo, userClient, s.Logger)
	playlistHandleres := NewPlaylistHandlers(playlistUsecase, s.Logger)

	s.MUX.HandleFunc("/api/v1/playlists", playlistHandleres.GetAllPlaylists).Methods("GET")

	s.MUX.Handle(
		"/api/v1/playlists",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.CreatePlaylist)),
	).Methods("POST")

	s.MUX.HandleFunc("/api/v1/playlists/{playlistId:[0-9]+}", playlistHandleres.GetPlaylist).Methods("GET")

	s.MUX.HandleFunc("/api/v1/playlists/{userId:[0-9a-fA-F-]+}/allPlaylists", playlistHandleres.GetUserPlaylists).Methods("GET")

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

	s.MUX.Handle(
		"/api/v1/playlists/favorite",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.GetFavoritePlaylists)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/playlists/favorite/{artistID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.IsFavoritePlaylist)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/playlists/favorite/{artistID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.AddFavoritePlaylist)),
	).Methods("POST")

	s.MUX.Handle(
		"/api/v1/playlists/favorite/{artistID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(playlistHandleres.DeleteFavoritePlaylist)),
	).Methods("DELETE")
}
