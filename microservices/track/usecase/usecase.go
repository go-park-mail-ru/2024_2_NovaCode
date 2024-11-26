package usecase

import (
	"context"
	"fmt"

	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	uuid "github.com/google/uuid"
)

type trackUsecase struct {
	trackRepo    track.Repo
	artistClient artistService.ArtistServiceClient
	albumClient  albumService.AlbumServiceClient
	logger       logger.Logger
}

func NewTrackUsecase(
	trackRepo track.Repo,
	artistClient artistService.ArtistServiceClient,
	albumClient albumService.AlbumServiceClient,
	logger logger.Logger,
) track.Usecase {
	return &trackUsecase{trackRepo, artistClient, albumClient, logger}
}

func (usecase *trackUsecase) View(ctx context.Context, trackID uint64) (*dto.TrackDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	foundTrack, err := usecase.trackRepo.FindById(ctx, trackID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Track wasn't found: %v", err), requestID)
		return nil, fmt.Errorf("Track wasn't found")
	}
	usecase.logger.Info("Track found", requestID)

	dtoTrack, err := usecase.ConvertTrackToDTO(ctx, foundTrack)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", foundTrack.Name, err), requestID)
		return nil, fmt.Errorf("Can't create DTO")
	}

	return dtoTrack, nil
}

func (usecase *trackUsecase) Search(ctx context.Context, query string) ([]*dto.TrackDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	foundTracks, err := usecase.trackRepo.FindByQuery(ctx, query)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Tracks with name '%s' were not found: %v", query, err), requestID)
		return nil, fmt.Errorf("Can't find tracks")
	}
	usecase.logger.Info("Tracks found", requestID)

	var dtoTracks []*dto.TrackDTO
	for _, track := range foundTracks {
		dtoTrack, err := usecase.ConvertTrackToDTO(ctx, track)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", track.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoTracks = append(dtoTracks, dtoTrack)
	}

	return dtoTracks, nil
}

func (usecase *trackUsecase) GetAll(ctx context.Context) ([]*dto.TrackDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	tracks, err := usecase.trackRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load tracks: %v", err), requestID)
		return nil, fmt.Errorf("Can't load tracks")
	}
	usecase.logger.Info("Found tracks", requestID)

	var dtoTracks []*dto.TrackDTO
	for _, track := range tracks {
		dtoTrack, err := usecase.ConvertTrackToDTO(ctx, track)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", track.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoTracks = append(dtoTracks, dtoTrack)
	}

	return dtoTracks, nil
}

func (usecase *trackUsecase) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.TrackDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	tracks, err := usecase.trackRepo.GetAllByArtistID(ctx, artistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load tracks by artist ID %d: %v", artistID, err), requestID)
		return nil, fmt.Errorf("Can't load tracks by artist ID %d", artistID)
	}
	usecase.logger.Infof("Found %d tracks for artist ID %d", len(tracks), artistID)

	var dtoTracks []*dto.TrackDTO
	for _, track := range tracks {
		dtoTrack, err := usecase.ConvertTrackToDTO(ctx, track)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", track.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO for track")
		}
		dtoTracks = append(dtoTracks, dtoTrack)
	}

	return dtoTracks, nil
}

func (usecase *trackUsecase) GetAllByAlbumID(ctx context.Context, albumID uint64) ([]*dto.TrackDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	tracks, err := usecase.trackRepo.GetAllByAlbumID(ctx, albumID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load tracks by album ID %d: %v", albumID, err), requestID)
		return nil, fmt.Errorf("Can't load tracks by album ID %d", albumID)
	}
	usecase.logger.Infof("Found %d tracks for album ID %d", len(tracks), albumID)

	var dtoTracks []*dto.TrackDTO
	for _, track := range tracks {
		dtoTrack, err := usecase.ConvertTrackToDTO(ctx, track)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", track.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO for track")
		}
		dtoTracks = append(dtoTracks, dtoTrack)
	}

	return dtoTracks, nil
}

func (usecase *trackUsecase) AddFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	if err := usecase.trackRepo.AddFavoriteTrack(ctx, userID, trackID); err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't add track %d to favorite for user %v: %v", trackID, userID, err), requestID)
		return fmt.Errorf("Can't add track %d to favorite for user %v: %v", trackID, userID, err)
	}

	return nil
}

func (usecase *trackUsecase) DeleteFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	if err := usecase.trackRepo.DeleteFavoriteTrack(ctx, userID, trackID); err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't delete track %d from favorite for user %v: %v", trackID, userID, err), requestID)
		return fmt.Errorf("Can't delete track %d from favorite for user %v: %v", trackID, userID, err)
	}

	return nil
}

func (usecase *trackUsecase) IsFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) (bool, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	exists, err := usecase.trackRepo.IsFavoriteTrack(ctx, userID, trackID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't find track %d in favorite for user %v: %v", trackID, userID, err), requestID)
		return false, fmt.Errorf("Can't find track %d in favorite for user %v: %v", trackID, userID, err)
	}

	return exists, nil
}

func (usecase *trackUsecase) GetFavoriteTracks(ctx context.Context, userID uuid.UUID) ([]*dto.TrackDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	tracks, err := usecase.trackRepo.GetFavoriteTracks(ctx, userID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load tracks by user ID %v: %v", userID, err), requestID)
		return nil, fmt.Errorf("Can't load tracks by user ID %v", userID)
	}
	usecase.logger.Infof("Found %d tracks for user ID %v", len(tracks), userID)

	var dtoTracks []*dto.TrackDTO
	for _, track := range tracks {
		dtoTrack, err := usecase.ConvertTrackToDTO(ctx, track)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s track: %v", track.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO for track")
		}
		dtoTracks = append(dtoTracks, dtoTrack)
	}

	return dtoTracks, nil
}

func (usecase *trackUsecase) ConvertTrackToDTO(ctx context.Context, track *models.Track) (*dto.TrackDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	artist, err := usecase.artistClient.FindByID(ctx, &artistService.FindByIDRequest{Id: track.ArtistID})
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't find artist for track %s: %v", track.Name, err), requestID)
		return nil, fmt.Errorf("Can't find artist for track")
	}

	album, err := usecase.albumClient.FindByID(ctx, &albumService.FindByIDRequest{Id: track.AlbumID})
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't find album for track %s: %v", track.Name, err), requestID)
		return nil, fmt.Errorf("Can't find album for track")
	}

	trackDTO := dto.NewTrackDTO(track)
	trackDTO.AlbumID = album.Album.Id
	trackDTO.AlbumName = album.Album.Name
	trackDTO.ArtistID = artist.Artist.Id
	trackDTO.ArtistName = artist.Artist.Name
	return trackDTO, nil
}
