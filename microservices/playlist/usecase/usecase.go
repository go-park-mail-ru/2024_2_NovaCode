package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist"
	pldto "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"
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

func (u *PlaylistUsecase) CreatePlaylist(ctx context.Context, newPlaylistDTO *pldto.PlaylistDTO) (*pldto.PlaylistDTO, error) {
	playlist := pldto.NewPlaylistFromPlaylistDTO(newPlaylistDTO)
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
	playlistDTO := pldto.NewPlaylistToPlaylistDTO(playlist)
	playlistDTO.OwnerName = owner.User.Username

	return playlistDTO, nil
}

func (u *PlaylistUsecase) GetPlaylist(ctx context.Context, playlistID uint64) (*pldto.PlaylistDTO, error) {
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
	playlistDTO := pldto.NewPlaylistToPlaylistDTO(playlist)
	playlistDTO.OwnerName = owner.User.Username

	return playlistDTO, nil
}

func (u *PlaylistUsecase) GetAllPlaylists(ctx context.Context) ([]*pldto.PlaylistDTO, error) {
	playlists, err := u.playlistRepo.GetAllPlaylists(ctx)
	if err != nil {
		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
		return nil, err
	}

	playlistsDTO := []*pldto.PlaylistDTO{}
	for _, playlist := range playlists {
		owner, err := u.userClient.FindByID(ctx, &userService.FindByIDRequest{Uuid: playlist.OwnerID.String()})
		if err != nil {
			u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
			return nil, err
		}
		playlistDTO := pldto.NewPlaylistToPlaylistDTO(playlist)
		playlistDTO.OwnerName = owner.User.Username
		playlistsDTO = append(playlistsDTO, playlistDTO)
	}

	return playlistsDTO, nil
}

func (u *PlaylistUsecase) GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*pldto.PlaylistDTO, error) {
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

	playlistsDTO := []*pldto.PlaylistDTO{}
	for _, playlist := range playlists {
		playlistDTO := pldto.NewPlaylistToPlaylistDTO(playlist)
		playlistDTO.OwnerName = owner.User.Username
		playlistsDTO = append(playlistsDTO, playlistDTO)
	}

	return playlistsDTO, nil
}

func (u *PlaylistUsecase) AddToPlaylist(ctx context.Context, playlistTrackDTO *pldto.PlaylistTrackDTO) (*models.PlaylistTrack, error) {
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

func (u *PlaylistUsecase) RemoveFromPlaylist(ctx context.Context, playlistTrackDTO *pldto.PlaylistTrackDTO) error {
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
