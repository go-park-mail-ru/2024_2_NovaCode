package server

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/middleware"
	userHandlers "github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/delivery/http"
	userRepo "github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/repository/postgres"
	userUsecase "github.com/go-park-mail-ru/2024_2_NovaCode/internal/user/usecase"

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
	s.BindUser()
	s.BindTrack()
	s.BindArtist()
	s.BindAlbum()
}

func (s *Server) BindTrack() {
	trackRepo := trackRepo.NewTrackPGRepository(s.db)
	albumRepo := albumRepo.NewAlbumPGRepository(s.db)
	artistRepo := artistRepo.NewArtistPGRepository(s.db)
	trackUsecase := trackUsecase.NewTrackUsecase(trackRepo, albumRepo, artistRepo, s.logger)
	trackHandleres := trackHandlers.NewTrackHandlers(trackUsecase, s.logger)

	s.mux.HandleFunc("/track/{id}", trackHandleres.ViewTrack)
	s.mux.HandleFunc("/api/v1/track/all", trackHandleres.GetAll)
	s.mux.HandleFunc("/track/search/", trackHandleres.SearchTrack)
}

func (s *Server) BindArtist() {
	artistRepo := artistRepo.NewArtistPGRepository(s.db)
	artistUsecase := artistUsecase.NewArtistUsecase(artistRepo, s.logger)
	artistHandlers := artistHandlers.NewArtistHandlers(artistUsecase, s.logger)

	s.mux.HandleFunc("/artist/{id}", artistHandlers.ViewArtist)
	s.mux.HandleFunc("/api/v1/artist/all", artistHandlers.GetAll)
	s.mux.HandleFunc("/artist/search/", artistHandlers.SearchArtist)
}

func (s *Server) BindAlbum() {
	artistRepo := artistRepo.NewArtistPGRepository(s.db)
	albumRepo := albumRepo.NewAlbumPGRepository(s.db)
	albumUsecase := albumUsecase.NewAlbumUsecase(albumRepo, artistRepo, s.logger)
	albumHandleres := albumHandlers.NewAlbumHandlers(albumUsecase, s.logger)

	s.mux.HandleFunc("/album/{id}", albumHandleres.ViewAlbum)
	s.mux.HandleFunc("/api/v1/album/all", albumHandleres.GetAll)
	s.mux.HandleFunc("/album/search/", albumHandleres.SearchAlbum)
}

func (s *Server) BindUser() {
	userRepo := userRepo.NewUserPostgresRepository(s.db)
	userUsecase := userUsecase.NewUserUsecase(&s.cfg.Auth, userRepo, s.logger)
	userHandleres := userHandlers.NewUserHandlers(&s.cfg.Auth, userUsecase, s.logger)

	s.mux.HandleFunc("GET /api/v1/health", userHandleres.Health)

	s.mux.HandleFunc("POST /api/v1/auth/register", userHandleres.Register)
	s.mux.HandleFunc("POST /api/v1/auth/login", userHandleres.Login)
	s.mux.HandleFunc("POST /api/v1/auth/logout", userHandleres.Logout)

	// auth middleware usage example
	s.mux.Handle(
		"GET /api/v1/auth/health",
		middleware.AuthMiddleware(&s.cfg.Auth, s.logger, http.HandlerFunc(userHandleres.Health)),
	)
}
