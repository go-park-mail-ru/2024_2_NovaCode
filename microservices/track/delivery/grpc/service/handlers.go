package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/dto"
	trackService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/track"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *tracksService) FindByID(ctx context.Context, request *trackService.FindByIDRequest) (*trackService.FindByIDResponse, error) {
	trackID, err := request.GetId()
	track, err := service.usecase.View(ctx, trackID)
	if err != nil {
		service.logger.Errorf("cannot find track by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "cannot find track by id: %v", err)
	}

	return &trackService.FindByIDResponse{User: service.trackDTOToProto(track)}, nil
}

func (service *tracksService) trackDTOToProto(track *dto.TrackDTO) *trackService.Track {
	return &trackService.Track{
		Id:           track.ID,
		Name:         track.Name,
		Duration:     track.Duration,
		FilePath:     track.FilePath,
		Image:        track.Image,
		ArtistId:     track.ArtistID,
		AlbumId:      track.AlbumID,
		OrderInAlbum: track.OrderInAlbum,
		ReleaseDate:  track.ReleaseDate,
	}
}
