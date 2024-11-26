package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
	"google.golang.org/grpc"
)

func RegisterAlbumService(usecase album.Usecase, logger logger.Logger) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		albumsServer := NewAlbumsService(usecase, logger)
		albumService.RegisterAlbumServiceServer(server, albumsServer)
	}
}
