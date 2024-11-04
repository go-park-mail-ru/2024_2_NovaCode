package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/track"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/httpErrors"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type trackUsecase struct {
	trackRepo  track.Repo
	albumRepo  album.Repo
	artistRepo artist.Repo
	logger     logger.Logger
}

func NewTrackUsecase(trackRepo track.Repo, albumRepo album.Repo, artistRepo artist.Repo, logger logger.Logger) track.Usecase {
	return &trackUsecase{trackRepo, albumRepo, artistRepo, logger}
}

func (usecase *trackUsecase) View(ctx context.Context, trackID uint64) (*dto.TrackDTO, error) {
	requestId := ctx.Value(utils.RequestIdKey{}).(uuid.UUID)

	foundTrack, err := usecase.trackRepo.FindById(ctx, trackID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Track wasn't found: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Track wasn't found", err)
	}
	usecase.logger.Info("Track found", zap.String("request_id", requestId.String()))

	dtoTrack, err := usecase.convertTrackToDTO(ctx, foundTrack)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", foundTrack.Name, err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
	}

	return dtoTrack, nil
}

func (usecase *trackUsecase) Search(ctx context.Context, name string) ([]*dto.TrackDTO, error) {
	requestId := ctx.Value(utils.RequestIdKey{}).(uuid.UUID)

	foundTracks, err := usecase.trackRepo.FindByName(ctx, name)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Tracks with name '%s' were not found: %v", name, err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't find tracks", err)
	}
	usecase.logger.Info("Tracks found", zap.String("request_id", requestId.String()))

	var dtoTracks []*dto.TrackDTO
	for _, track := range foundTracks {
		dtoTrack, err := usecase.convertTrackToDTO(ctx, track)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", track.Name, err), zap.String("request_id", requestId.String()))
			return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
		}
		dtoTracks = append(dtoTracks, dtoTrack)
	}

	return dtoTracks, nil
}

func (usecase *trackUsecase) GetAll(ctx context.Context) ([]*dto.TrackDTO, error) {
	requestId := ctx.Value(utils.RequestIdKey{}).(uuid.UUID)

	tracks, err := usecase.trackRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load tracks: %v", err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't load tracks", err)
	}
	usecase.logger.Info("Found tracks", zap.String("request_id", requestId.String()))

	var dtoTracks []*dto.TrackDTO
	for _, track := range tracks {
		dtoTrack, err := usecase.convertTrackToDTO(ctx, track)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", track.Name, err), zap.String("request_id", requestId.String()))
			return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
		}
		dtoTracks = append(dtoTracks, dtoTrack)
	}

	return dtoTracks, nil
}

func (usecase *trackUsecase) convertTrackToDTO(ctx context.Context, track *models.Track) (*dto.TrackDTO, error) {
	requestId := ctx.Value(utils.RequestIdKey{}).(uuid.UUID)

	artist, err := usecase.artistRepo.FindById(ctx, track.ArtistID)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't find artist for track %s: %v", track.Name, err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't find artist for track", err)
	}

	album, err := usecase.albumRepo.FindById(ctx, track.AlbumID)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't find album for track %s: %v", track.Name, err), zap.String("request_id", requestId.String()))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't find album for track", err)
	}

	return dto.NewTrackDTO(track, artist, album), nil
}
