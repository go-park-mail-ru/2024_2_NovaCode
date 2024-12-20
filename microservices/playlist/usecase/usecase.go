package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
)

type PlaylistUsecase struct {
	playlistRepo playlist.Repository
	userClient   userService.UserServiceClient
	s3Repo       s3.S3Repo
	logger       logger.Logger
}

func NewPlaylistUsecase(
	playlistRepo playlist.Repository,
	userClient userService.UserServiceClient,
	s3Repo s3.S3Repo,
	logger logger.Logger,
) playlist.Usecase {
	return &PlaylistUsecase{playlistRepo, userClient, s3Repo, logger}
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

func (usecase *PlaylistUsecase) GetFavoritePlaylistsCount(ctx context.Context, userID uuid.UUID) (uint64, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	count, err := usecase.playlistRepo.GetFavoritePlaylistsCount(ctx, userID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load playlists count by user ID %v: %v", userID, err), requestID)
		return 0, fmt.Errorf("Can't load playlists by user ID %v", userID)
	}
	usecase.logger.Infof("Found %d playlists for user ID %v", count, userID)

	return count, nil
}

func (usecase *PlaylistUsecase) GetPlaylistLikesCount(ctx context.Context, playlistID uint64) (uint64, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	likesCount, err := usecase.playlistRepo.GetPlaylistLikesCount(ctx, playlistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load playlist likes count by playlist ID %v: %v", playlistID, err), requestID)
		return 0, fmt.Errorf("Can't load playlist likes count by playlist ID %v", playlistID)
	}
	usecase.logger.Infof("Found %d likes for playlist ID %v", likesCount, playlistID)

	return likesCount, nil
}

func (u *PlaylistUsecase) GetPopularPlaylists(ctx context.Context) ([]*dto.PlaylistDTO, error) {
	playlists, err := u.playlistRepo.GetPopularPlaylists(ctx)
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

func (u *PlaylistUsecase) Update(ctx context.Context, playlist *models.Playlist) (*dto.PlaylistDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	currentPlaylist, err := u.playlistRepo.GetPlaylist(ctx, playlist.ID)
	if err != nil {
		u.logger.Warn(fmt.Sprintf("playlist not found: %v", err), requestID)
		return nil, fmt.Errorf("failed to find playlist")
	}

	currentPlaylist.Name = playlist.Name
	currentPlaylist.Image = playlist.Image
	currentPlaylist.OwnerID = playlist.OwnerID
	currentPlaylist.IsPrivate = playlist.IsPrivate

	updatedPlaylist, err := u.playlistRepo.Update(ctx, currentPlaylist)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error updating playlist: %v", err), requestID)
		return nil, fmt.Errorf("failed to update playlist")
	}
	u.logger.Infof("playlist '%s' successfully updated", updatedPlaylist.ID)

	playlistDTO := dto.NewPlaylistToPlaylistDTO(updatedPlaylist)
	return playlistDTO, nil
}

func (u *PlaylistUsecase) UploadImage(ctx context.Context, playlistID uint64, file s3.Upload) (*dto.PlaylistDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	playlist, err := u.playlistRepo.GetPlaylist(ctx, playlistID)
	if err != nil {
		u.logger.Warn(fmt.Sprintf("playlist not found: %v", err), requestID)
		return nil, fmt.Errorf("playlist not found")
	}

	uploadInfo, err := u.s3Repo.Put(ctx, file)
	if err != nil {
		u.logger.Warn(fmt.Sprintf("failed to save playlist image: %v", err), requestID)
		return nil, fmt.Errorf("failed to save playlist image")
	}

	imageURL := uploadInfo.Key

	updatedPlaylistDTO, err := u.Update(ctx, &models.Playlist{
		ID:        playlist.ID,
		Name:      playlist.Name,
		Image:     imageURL,
		IsPrivate: playlist.IsPrivate,
	})
	if err != nil {
		u.logger.Warn(fmt.Sprintf("failed to update playlist model: %v", err), requestID)
		return nil, fmt.Errorf("failed to update playlist model")
	}

	return updatedPlaylistDTO, nil
}
