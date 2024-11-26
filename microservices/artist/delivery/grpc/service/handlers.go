package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/dto"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *artistsService) FindByID(ctx context.Context, request *artistService.FindByIDRequest) (*artistService.FindByIDResponse, error) {
	artistID := request.GetId()
	artist, err := service.usecase.View(ctx, artistID)
	if err != nil {
		service.logger.Errorf("cannot find artist by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "cannot find artist by id: %v", err)
	}

	return &artistService.FindByIDResponse{Artist: service.artistDTOToProto(artist)}, nil
}

func (service *artistsService) artistDTOToProto(artist *dto.ArtistDTO) *artistService.Artist {
	return &artistService.Artist{
		Id:      artist.ID,
		Name:    artist.Name,
		Bio:     artist.Bio,
		Country: artist.Country,
		Image:   artist.Image,
	}
}
