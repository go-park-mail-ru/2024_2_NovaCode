package http

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	httpServer "github.com/go-park-mail-ru/2024_2_NovaCode/internal/server/http"
	albumRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/repository"
	albumUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/usecase"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func BindRoutes(s *httpServer.Server, artistClient artistService.ArtistServiceClient) {
	s.MUX.Handle("/metrics", promhttp.Handler())

	albumRepo := albumRepo.NewAlbumPGRepository(s.PG)
	albumUsecase := albumUsecase.NewAlbumUsecase(albumRepo, artistClient, s.Logger)
	albumHandleres := NewAlbumHandlers(albumUsecase, s.Logger)

	s.MUX.HandleFunc("/api/v1/albums/search", albumHandleres.SearchAlbum).Methods("GET")
	s.MUX.HandleFunc("/api/v1/albums/{id:[0-9]+}", albumHandleres.ViewAlbum).Methods("GET")
	s.MUX.HandleFunc("/api/v1/albums", albumHandleres.GetAll).Methods("GET")
	s.MUX.HandleFunc("/api/v1/albums/byArtistId/{artistId:[0-9]+}", albumHandleres.GetAllByArtistID).Methods("GET")

	s.MUX.Handle(
		"/api/v1/albums/favorite/byUser/{userID}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(albumHandleres.GetFavoriteAlbums)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/albums/favorite/count/{userID:[0-9a-fA-F-]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(albumHandleres.GetFavoriteAlbumsCount)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/albums/favorite/{albumID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(albumHandleres.IsFavoriteAlbum)),
	).Methods("GET")

	s.MUX.Handle(
		"/api/v1/albums/favorite/{albumID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(albumHandleres.AddFavoriteAlbum)),
	).Methods("POST")

	s.MUX.Handle(
		"/api/v1/albums/favorite/{albumID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(albumHandleres.DeleteFavoriteAlbum)),
	).Methods("DELETE")

	s.MUX.Handle(
		"/api/v1/albums/likes/{albumID:[0-9]+}",
		middleware.AuthMiddleware(&s.CFG.Service.Auth, s.Logger, http.HandlerFunc(albumHandleres.GetAlbumLikesCount)),
	).Methods("GET")
}
