package service

import (
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
)

type artistsService struct {
	usecase artist.Usecase
	logger  logger.Logger

	artistService.UnimplementedArtistServiceServer
}

func NewArtistsService(usecase artist.Usecase, logger logger.Logger) *artistsService {
	return &artistsService{usecase, logger, artistService.UnimplementedArtistServiceServer{}}
}
