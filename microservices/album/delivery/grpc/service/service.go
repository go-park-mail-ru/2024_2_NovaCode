package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
)

type albumsService struct {
	cfg     *config.AuthConfig
	usecase album.Usecase
	logger  logger.Logger

	albumService.UnimplementedAlbumServiceServer
}

func NewUsersService(cfg *config.AuthConfig, usecase album.Usecase, logger logger.Logger) *albumsService {
	return &albumsService{cfg, usecase, logger, albumService.UnimplementedAlbumServiceServer{}}
}