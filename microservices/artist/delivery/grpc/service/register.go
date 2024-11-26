package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
	"google.golang.org/grpc"
)

func RegisterArtistService(usecase artist.Usecase, logger logger.Logger) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		artistsServer := NewArtistsService(usecase, logger)
		artistService.RegisterArtistServiceServer(server, artistsServer)
	}
}
