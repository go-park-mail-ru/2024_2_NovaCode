package http

import (
	// "net/http"

	// "github.com/prometheus/client_golang/prometheus/promhttp"

	// "github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/server"
	// albumRepo "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/repository"
	// albumUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/usecase"
)

func BindRoutes(s *server.Server) {
	// artistRepo := artistRepo.NewArtistPGRepository(s.pg)
	// albumRepo := albumRepo.NewAlbumPGRepository(s.pg)
	// albumUsecase := albumUsecase.NewAlbumUsecase(albumRepo, artistRepo, s.logger)
	// albumHandleres := albumHandlers.NewAlbumHandlers(albumUsecase, s.logger)

	// s.mux.HandleFunc("/api/v1/albums/search", albumHandleres.SearchAlbum).Methods("GET")
	// s.mux.HandleFunc("/api/v1/albums/{id:[0-9]+}", albumHandleres.ViewAlbum).Methods("GET")
	// s.mux.HandleFunc("/api/v1/albums", albumHandleres.GetAll).Methods("GET")
	// s.mux.HandleFunc("/api/v1/albums/byArtistId/{artistId:[0-9]+}", albumHandleres.GetAllByArtistID).Methods("GET")
}