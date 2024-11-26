package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	tarckService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/track"
)

type tracksService struct {
	cfg     *config.AuthConfig
	usecase track.Usecase
	logger  logger.Logger

	trackService.UnimplementedTrackServiceServer
}

func NewUsersService(cfg *config.AuthConfig, usecase track.Usecase, logger logger.Logger) *tarcksService {
	return &tracksService{cfg, usecase, logger, trackService.UnimplementedTrackServiceServer{}}
}