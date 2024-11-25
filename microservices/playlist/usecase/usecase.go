package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist"
	pldto "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"

	// "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track"
	// tdto "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type PlaylistUsecase struct {
	// trackUsecase track.Usecase
	playlistRepo playlist.Repository
	// trackRepo    track.Repo
	// userRepo user.PostgresRepo
	logger logger.Logger
}

// func NewPlaylistUsecase(trackUsecase track.Usecase, playlistRepo playlist.Repository, trackRepo track.Repo, userRepo user.PostgresRepo, logger logger.Logger) playlist.Usecase {
// 	return &PlaylistUsecase{trackUsecase, playlistRepo, trackRepo, userRepo, logger}
// }

func NewPlaylistUsecase(playlistRepo playlist.Repository, logger logger.Logger) playlist.Usecase {
	return &PlaylistUsecase{playlistRepo, logger}
}

// func (u *PlaylistUsecase) CreatePlaylist(ctx context.Context, newPlaylistDTO *pldto.PlaylistDTO) (*pldto.PlaylistDTO, error) {
// 	playlist := pldto.NewPlaylistFromPlaylistDTO(newPlaylistDTO)
// 	playlist, err := u.playlistRepo.CreatePlaylist(ctx, playlist)
// 	if err != nil {
// 		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 		return nil, err
// 	}

// 	owner, err := u.userRepo.FindByID(ctx, playlist.OwnerID)
// 	if err != nil {
// 		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 		return nil, err
// 	}
// 	playlistDTO := pldto.NewPlaylistToPlaylistDTO(playlist)
// 	playlistDTO.OwnerName = owner.Username

// 	return playlistDTO, nil
// }

// func (u *PlaylistUsecase) GetPlaylist(ctx context.Context, playlistID uint64) (*pldto.PlaylistDTO, error) {
// 	playlist, err := u.playlistRepo.GetPlaylist(ctx, playlistID)
// 	if err != nil {
// 		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 		return nil, err
// 	}

// 	owner, err := u.userRepo.FindByID(ctx, playlist.OwnerID)
// 	if err != nil {
// 		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 		return nil, err
// 	}
// 	playlistDTO := pldto.NewPlaylistToPlaylistDTO(playlist)
// 	playlistDTO.OwnerName = owner.Username

// 	return playlistDTO, nil
// }

// func (u *PlaylistUsecase) GetAllPlaylists(ctx context.Context) ([]*pldto.PlaylistDTO, error) {
// 	playlists, err := u.playlistRepo.GetAllPlaylists(ctx)
// 	if err != nil {
// 		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 		return nil, err
// 	}

// 	playlistsDTO := []*pldto.PlaylistDTO{}
// 	for _, playlist := range playlists {
// 		owner, err := u.userRepo.FindByID(ctx, playlist.OwnerID)
// 		if err != nil {
// 			u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 			return nil, err
// 		}
// 		playlistDTO := pldto.NewPlaylistToPlaylistDTO(playlist)
// 		playlistDTO.OwnerName = owner.Username
// 		playlistsDTO = append(playlistsDTO, playlistDTO)
// 	}

// 	return playlistsDTO, nil
// }

// func (u *PlaylistUsecase) GetTracksFromPlaylist(ctx context.Context, playlistID uint64) ([]*tdto.TrackDTO, error) {
// 	tracks := []*models.Track{}
// 	playlistTracks, err := u.playlistRepo.GetTracksFromPlaylist(ctx, playlistID)
// 	if err != nil {
// 		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 		return nil, err
// 	}

// 	for _, playlistTrack := range playlistTracks {
// 		track, err := u.trackRepo.FindById(ctx, playlistTrack.TrackID)
// 		if err != nil {
// 			u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 			return nil, err
// 		}
// 		tracks = append(tracks, track)
// 	}

// 	tracksDTO := []*tdto.TrackDTO{}
// 	for _, track := range tracks {
// 		trackDTO, err := u.trackUsecase.ConvertTrackToDTO(ctx, track)
// 		if err != nil {
// 			u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 			return nil, err
// 		}
// 		tracksDTO = append(tracksDTO, trackDTO)
// 	}

// 	return tracksDTO, nil
// }

// func (u *PlaylistUsecase) GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*pldto.PlaylistDTO, error) {
// 	playlists, err := u.playlistRepo.GetUserPlaylists(ctx, userID)
// 	if err != nil {
// 		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 		return nil, err
// 	}

// 	owner, err := u.userRepo.FindByID(ctx, userID)
// 	if err != nil {
// 		u.logger.Error(err.Error(), ctx.Value(utils.RequestIDKey{}))
// 		return nil, err
// 	}

// 	playlistsDTO := []*pldto.PlaylistDTO{}
// 	for _, playlist := range playlists {
// 		playlistDTO := pldto.NewPlaylistToPlaylistDTO(playlist)
// 		playlistDTO.OwnerName = owner.Username
// 		playlistsDTO = append(playlistsDTO, playlistDTO)
// 	}

// 	return playlistsDTO, nil
// }

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
