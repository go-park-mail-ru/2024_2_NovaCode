package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
)

type PlaylistUsecase struct {
	playlistRepo playlist.Repository
	userClient   userService.UserServiceClient
	logger       logger.Logger
}

func NewPlaylistUsecase(
	playlistRepo playlist.Repository,
	userClient userService.UserServiceClient,
	logger logger.Logger,
) playlist.Usecase {
	return &PlaylistUsecase{playlistRepo, userClient, logger}
}

func (u *PlaylistUsecase) CreatePlaylist(ctx context.Context, newPlaylistDTO *dto.PlaylistDTO) (*dto.PlaylistDTO, error) {
	playlist := dto.NewPlaylistFromPlaylistDTO(newPlaylistDTO)
	playlist, err := u.playlistRepo.CreatePlaylist(ctx, playlist)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}

	owner, err := u.userClient.FindByID(ctx, &userService.FindByIDRequest{Uuid: playlist.OwnerID.String()})
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}
	playlistDTO := dto.NewPlaylistToPlaylistDTO(playlist)
	playlistDTO.OwnerName = owner.User.Username

	return playlistDTO, nil
}

func (u *PlaylistUsecase) GetPlaylist(ctx context.Context, playlistID uint64) (*dto.PlaylistDTO, error) {
	playlist, err := u.playlistRepo.GetPlaylist(ctx, playlistID)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}

	owner, err := u.userClient.FindByID(ctx, &userService.FindByIDRequest{Uuid: playlist.OwnerID.String()})
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}
	playlistDTO := dto.NewPlaylistToPlaylistDTO(playlist)
	playlistDTO.OwnerName = owner.User.Username

	return playlistDTO, nil
}

func (u *PlaylistUsecase) GetAllPlaylists(ctx context.Context) ([]*dto.PlaylistDTO, error) {
	playlists, err := u.playlistRepo.GetAllPlaylists(ctx)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}

	playlistsDTO := []*dto.PlaylistDTO{}
	for _, playlist := range playlists {
		owner, err := u.userClient.FindByID(ctx, &userService.FindByIDRequest{Uuid: playlist.OwnerID.String()})
		if err != nil {
			u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
			return nil, err
		}
		playlistDTO := dto.NewPlaylistToPlaylistDTO(playlist)
		playlistDTO.OwnerName = owner.User.Username
		playlistsDTO = append(playlistsDTO, playlistDTO)
	}

	return playlistsDTO, nil
}

func (u *PlaylistUsecase) GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*dto.PlaylistDTO, error) {
	playlists, err := u.playlistRepo.GetUserPlaylists(ctx, userID)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}

	owner, err := u.userClient.FindByID(ctx, &userService.FindByIDRequest{Uuid: userID.String()})
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}

	playlistsDTO := []*dto.PlaylistDTO{}
	for _, playlist := range playlists {
		playlistDTO := dto.NewPlaylistToPlaylistDTO(playlist)
		playlistDTO.OwnerName = owner.User.Username
		playlistsDTO = append(playlistsDTO, playlistDTO)
	}

	return playlistsDTO, nil
}

func (u *PlaylistUsecase) AddToPlaylist(ctx context.Context, playlistTrackDTO *dto.PlaylistTrackDTO) (*models.PlaylistTrack, error) {
	length, err := u.playlistRepo.GetLengthPlaylist(ctx, playlistTrackDTO.PlaylistID)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}
	playlistTrack, err := u.playlistRepo.AddToPlaylist(ctx, playlistTrackDTO.PlaylistID, length+1, playlistTrackDTO.TrackID)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}

	return playlistTrack, nil
}

func (u *PlaylistUsecase) RemoveFromPlaylist(ctx context.Context, playlistTrackDTO *dto.PlaylistTrackDTO) error {
	_, err := u.playlistRepo.RemoveFromPlaylist(ctx, playlistTrackDTO.PlaylistID, playlistTrackDTO.TrackID)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return err
	}
	return nil
}

func (u *PlaylistUsecase) DeletePlaylist(ctx context.Context, playlistID uint64) error {
	_, err := u.playlistRepo.DeletePlaylist(ctx, playlistID)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return err
	}
	return nil
}

func (u *PlaylistUsecase) AddFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	if err := u.playlistRepo.AddFavoritePlaylist(ctx, userID, playlistID); err != nil {
		u.logger.Warn(fmt.Sprintf("Can't add playlist %d to favorite for user %v: %v", playlistID, userID, err), requestID)
		return fmt.Errorf("Can't add playlist %d to favorite for user %v: %v", playlistID, userID, err)
	}

	return nil
}

func (u *PlaylistUsecase) DeleteFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	if err := u.playlistRepo.DeleteFavoritePlaylist(ctx, userID, playlistID); err != nil {
		u.logger.Warn(fmt.Sprintf("Can't delete playlist %d from favorite for user %v: %v", playlistID, userID, err), requestID)
		return fmt.Errorf("Can't delete playlist %d from favorite for user %v: %v", playlistID, userID, err)
	}

	return nil
}

func (u *PlaylistUsecase) IsFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) (bool, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	exists, err := u.playlistRepo.IsFavoritePlaylist(ctx, userID, playlistID)
	if err != nil {
		u.logger.Warn(fmt.Sprintf("Can't find playlist %d in favorite for user %v: %v", playlistID, userID, err), requestID)
		return false, fmt.Errorf("Can't find playlist %d in favorite for user %v: %v", playlistID, userID, err)
	}

	return exists, nil
}

func (u *PlaylistUsecase) GetFavoritePlaylists(ctx context.Context, userID uuid.UUID) ([]*dto.PlaylistDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	playlists, err := u.playlistRepo.GetFavoritePlaylists(ctx, userID)
	if err != nil {
		u.logger.Warn(fmt.Sprintf("Can't load playlists by user ID %v: %v", userID, err), requestID)
		return nil, fmt.Errorf("Can't load playlists by user ID %v", userID)
	}
	u.logger.Infof("Found %d playlists for user ID %v", len(playlists), userID)

	var dtoPlaylists []*dto.PlaylistDTO
	for _, playlist := range playlists {
		dtoPlaylist := dto.NewPlaylistToPlaylistDTO(playlist)
		dtoPlaylists = append(dtoPlaylists, dtoPlaylist)
	}

	return dtoPlaylists, nil
}
