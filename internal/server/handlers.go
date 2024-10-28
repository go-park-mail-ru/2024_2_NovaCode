package server

import (
	"net/http"

	s3Repo "github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3/repository/s3"

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
	trackRepo := trackRepo.NewTrackPGRepository(s.pg)
	albumRepo := albumRepo.NewAlbumPGRepository(s.pg)
	artistRepo := artistRepo.NewArtistPGRepository(s.pg)
	trackUsecase := trackUsecase.NewTrackUsecase(trackRepo, albumRepo, artistRepo, s.logger)
	trackHandleres := trackHandlers.NewTrackHandlers(trackUsecase, s.logger)

	s.mux.HandleFunc("/api/v1/tracks/search", trackHandleres.SearchTrack).Methods("GET")
	s.mux.HandleFunc("/api/v1/tracks/{id:[0-9]+}", trackHandleres.ViewTrack).Methods("GET")
	s.mux.HandleFunc("/api/v1/tracks", trackHandleres.GetAll).Methods("GET")
}

func (s *Server) BindArtist() {
	artistRepo := artistRepo.NewArtistPGRepository(s.pg)
	artistUsecase := artistUsecase.NewArtistUsecase(artistRepo, s.logger)
	artistHandlers := artistHandlers.NewArtistHandlers(artistUsecase, s.logger)

	s.mux.HandleFunc("/api/v1/artists/search", artistHandlers.SearchArtist).Methods("GET")
	s.mux.HandleFunc("/api/v1/artists/{id:[0-9]+}", artistHandlers.ViewArtist).Methods("GET")
	s.mux.HandleFunc("/api/v1/artists", artistHandlers.GetAll).Methods("GET")
}

func (s *Server) BindAlbum() {
	artistRepo := artistRepo.NewArtistPGRepository(s.pg)
	albumRepo := albumRepo.NewAlbumPGRepository(s.pg)
	albumUsecase := albumUsecase.NewAlbumUsecase(albumRepo, artistRepo, s.logger)
	albumHandleres := albumHandlers.NewAlbumHandlers(albumUsecase, s.logger)

	s.mux.HandleFunc("/api/v1/albums/search", albumHandleres.SearchAlbum).Methods("GET")
	s.mux.HandleFunc("/api/v1/albums/{id:[0-9]+}", albumHandleres.ViewAlbum).Methods("GET")
	s.mux.HandleFunc("/api/v1/albums", albumHandleres.GetAll).Methods("GET")
}

func (s *Server) BindUser() {
	userPGRepo := userRepo.NewUserPostgresRepository(s.pg, s.logger)
	userS3Repo := s3Repo.NewS3Repository(s.s3, s.logger)
	userUsecase := userUsecase.NewUserUsecase(&s.cfg.Service.Auth, &s.cfg.Minio, userPGRepo, userS3Repo, s.logger)
	userHandleres := userHandlers.NewUserHandlers(&s.cfg.Service.Auth, userUsecase, s.logger)

	s.mux.HandleFunc("/api/v1/health", userHandleres.Health).Methods("GET")

	s.mux.HandleFunc("/api/v1/auth/register", userHandleres.Register).Methods("POST")
	s.mux.HandleFunc("/api/v1/auth/login", userHandleres.Login).Methods("POST")

	s.mux.Handle(
		"/api/v1/auth/logout",
		middleware.AuthMiddleware(&s.cfg.Service.Auth, s.logger, http.HandlerFunc(userHandleres.Logout)),
	).Methods("POST")

	s.mux.HandleFunc("/api/v1/users/{user_id:[0-9a-fA-F-]+}", userHandleres.GetUserByID).Methods("GET")

	s.mux.Handle(
		"/api/v1/users/me",
		middleware.AuthMiddleware(&s.cfg.Service.Auth, s.logger, http.HandlerFunc(userHandleres.GetMe)),
	).Methods("GET")

	s.mux.Handle(
		"/api/v1/users/{user_id:[0-9a-fA-F-]+}",
		middleware.AuthMiddleware(&s.cfg.Service.Auth, s.logger, http.HandlerFunc(userHandleres.Update)),
	).Methods("PUT")

	s.mux.Handle(
		"/api/v1/users/{user_id:[0-9a-fA-F-]+}/image",
		middleware.AuthMiddleware(&s.cfg.Service.Auth, s.logger, http.HandlerFunc(userHandleres.UploadImage)),
	).Methods("POST")

	s.mux.Handle(
		"/api/v1/auth/health",
		middleware.AuthMiddleware(&s.cfg.Service.Auth, s.logger, http.HandlerFunc(userHandleres.Health)),
	).Methods("GET")
}
