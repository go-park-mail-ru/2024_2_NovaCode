package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
	"google.golang.org/grpc"
)

func RegisterUserService(cfg *config.AuthConfig, usecase album.Usecase, logger logger.Logger) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		albumsServer := NewUsersService(cfg, usecase, logger)
		albumService.RegisterUserServiceServer(server, albumsServer)
	}
}
