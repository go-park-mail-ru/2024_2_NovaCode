package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	trackService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/track"
	"google.golang.org/grpc"
)

func RegisterTrackService(cfg *config.AuthConfig, usecase track.Usecase, logger logger.Logger) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		tracksServer := NewTracksService(cfg, usecase, logger)
		trackService.RegisterTrackServiceServer(server, tracksServer)
	}
}
