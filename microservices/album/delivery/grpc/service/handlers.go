package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/dto"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *albumsService) FindByID(ctx context.Context, request *albumService.FindByIDRequest) (*albumService.FindByIDResponse, error) {
	albumID := request.GetId()
	album, err := service.usecase.View(ctx, albumID)
	if err != nil {
		service.logger.Errorf("cannot find album by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "cannot find album by id: %v", err)
	}

	return &albumService.FindByIDResponse{User: service.albumDTOToProto(album)}, nil
}

func (service *albumsService) albumDTOToProto(album *dto.AlbumDTO) *albumService.Album {
	return &albumService.Album{
		Id:          album.ID,
		Name:        album.Name,
		ReleaseDate: album.ReleaseDate,
		Image:       album.Image,
		ArtistId:    album.ArtistID,
	}
}
