package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
)

type albumsService struct {
	usecase album.Usecase
	logger  logger.Logger

	albumService.UnimplementedAlbumServiceServer
}

func NewAlbumsService(usecase album.Usecase, logger logger.Logger) *albumsService {
	return &albumsService{usecase, logger, albumService.UnimplementedAlbumServiceServer{}}
}
