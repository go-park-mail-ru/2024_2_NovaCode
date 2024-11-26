package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/user"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"
)

type usersService struct {
	cfg     *config.AuthConfig
	usecase user.Usecase
	logger  logger.Logger

	userService.UnimplementedUserServiceServer
}

func NewUsersService(cfg *config.AuthConfig, usecase user.Usecase, logger logger.Logger) *usersService {
	return &usersService{cfg, usecase, logger, userService.UnimplementedUserServiceServer{}}
}
