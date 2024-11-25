package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"
	"google.golang.org/grpc"
)

func RegisterUserService(cfg *config.AuthConfig, usecase user.Usecase, logger logger.Logger) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		usersServer := NewUsersService(cfg, usecase, logger)
		userService.RegisterUserServiceServer(server, usersServer)
	}
}
