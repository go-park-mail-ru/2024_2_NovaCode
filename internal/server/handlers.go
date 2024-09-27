package server

import (
	"net/http"

	trackHandlers "github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/delivery/http"
	trackRepo "github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/repository"
	trackUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/usecase"

	artistHandlers "github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/delivery/http"
	artistRepo "github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/repository"
	artistUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/usecase"

	albumHandlers "github.com/go-park-mail-ru/2024_2_NovaCode/internal/album/delivery/http"
	albumRepo "github.com/go-park-mail-ru/2024_2_NovaCode/internal/album/repository"
	albumUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/internal/album/usecase"
)

func (s *Server) BindRoutes() {
	s.BindTrack()
	s.BindArtist()
	s.BindAlbum()
}

func (s *Server) BindTrack() {
	trackRepo := trackRepo.NewTrackPGRepository(s.db)
	albumRepo := albumRepo.NewAlbumPGRepository(s.db)
	artistRepo := artistRepo.NewArtistPGRepository(s.db)
	trackUsecase := trackUsecase.NewTrackUsecase(s.cfg, trackRepo, albumRepo, artistRepo, s.logger)
	trackHandleres := trackHandlers.NewTrackHandlers(s.cfg, trackUsecase, s.logger)

	s.mux.HandleFunc("/api/v1/track/{id}", trackHandleres.ViewTrack)
	s.mux.HandleFunc("/api/v1/track/all", trackHandleres.GetAll)
	s.mux.HandleFunc("/api/v1/track/", trackHandleres.SearchTrack)
}

func (s *Server) BindArtist() {
	artistRepo := artistRepo.NewArtistPGRepository(s.db)
	artistUsecase := artistUsecase.NewArtistUsecase(s.cfg, artistRepo, s.logger)
	artistHandlers := artistHandlers.NewArtistHandlers(s.cfg, artistUsecase, s.logger)

	s.mux.HandleFunc("/api/v1/artist/{id}", artistHandlers.ViewArtist)
	s.mux.HandleFunc("/api/v1/artist/all", artistHandlers.GetAll)
	s.mux.HandleFunc("/api/v1/artist/", artistHandlers.SearchArtist)
}

func (s *Server) BindAlbum() {
	artistRepo := artistRepo.NewArtistPGRepository(s.db)
	albumRepo := albumRepo.NewAlbumPGRepository(s.db)
	albumUsecase := albumUsecase.NewAlbumUsecase(s.cfg, albumRepo, artistRepo, s.logger)
	albumHandleres := albumHandlers.NewAlbumHandlers(s.cfg, albumUsecase, s.logger)

	s.mux.HandleFunc("/api/v1/album/{id}", albumHandleres.ViewAlbum)
	s.mux.HandleFunc("/api/v1/album/all", albumHandleres.GetAll)
	s.mux.HandleFunc("/api/v1/album/", albumHandleres.SearchAlbum)
}
