package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	trackService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/track"
)

type tracksService struct {
	cfg     *config.AuthConfig
	usecase track.Usecase
	logger  logger.Logger

	trackService.UnimplementedTrackServiceServer
}

func NewTracksService(cfg *config.AuthConfig, usecase track.Usecase, logger logger.Logger) *tracksService {
	return &tracksService{cfg, usecase, logger, trackService.UnimplementedTrackServiceServer{}}
}