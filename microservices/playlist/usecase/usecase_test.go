package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	userService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/user"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPlaylistUsecaseCreatePlaylist_Success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	userClientMock := mock.NewMockUserServiceClient(ctrl)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, userClientMock, logger)

	ownerId := uuid.New()

	findByIDResponseUser := &userService.FindByIDResponse{
		User: &userService.User{
			Uuid:     ownerId.String(),
			Username: "user",
			Email:    "email@example.com",
			Password: "new_password",
		},
	}

	mockPlaylist := &models.Playlist{
		ID:        0,
		Name:      "gym training playlist",
		Image:     "/images/playlists/playlist_1.jpg",
		OwnerID:   ownerId,
		IsPrivate: false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	mockPlaylistDTO := &dto.PlaylistDTO{
		Name:    mockPlaylist.Name,
		Image:   mockPlaylist.Image,
		OwnerID: mockPlaylist.OwnerID,
	}

	ctx := context.Background()
	userClientMock.EXPECT().FindByID(ctx, &userService.FindByIDRequest{Uuid: mockPlaylistDTO.OwnerID.String()}).Return(findByIDResponseUser, nil)
	playlistRepoMock.EXPECT().CreatePlaylist(ctx, mockPlaylist).Return(mockPlaylist, nil)

	dtoPlaylist, err := playlistUsecase.CreatePlaylist(ctx, mockPlaylistDTO)

	require.NoError(t, err)
	require.NotNil(t, dtoPlaylist)
	require.Equal(t, mockPlaylist.Name, mockPlaylistDTO.Name)
	require.Equal(t, mockPlaylist.Image, mockPlaylistDTO.Image)
	require.Equal(t, mockPlaylist.OwnerID, mockPlaylistDTO.OwnerID)
}

func TestPlaylistUsecaseGetPlaylist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	userClientMock := mock.NewMockUserServiceClient(ctrl)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, userClientMock, logger)

	playlistID := uint64(1)
	ownerID := uuid.New()

	mockPlaylist := &models.Playlist{
		ID:        playlistID,
		Name:      "chill vibes",
		Image:     "/images/playlists/playlist_2.jpg",
		OwnerID:   ownerID,
		IsPrivate: false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	findByIDResponseUser := &userService.FindByIDResponse{
		User: &userService.User{
			Uuid:     ownerID.String(),
			Username: "user2",
		},
	}

	ctx := context.Background()
	playlistRepoMock.EXPECT().GetPlaylist(ctx, playlistID).Return(mockPlaylist, nil)
	userClientMock.EXPECT().FindByID(ctx, &userService.FindByIDRequest{Uuid: ownerID.String()}).Return(findByIDResponseUser, nil)

	dtoPlaylist, err := playlistUsecase.GetPlaylist(ctx, playlistID)

	require.NoError(t, err)
	require.NotNil(t, dtoPlaylist)
	require.Equal(t, mockPlaylist.Name, dtoPlaylist.Name)
	require.Equal(t, "user2", dtoPlaylist.OwnerName)
}

func TestPlaylistUsecaseGetPlaylist_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	userClientMock := mock.NewMockUserServiceClient(ctrl)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, userClientMock, logger)

	playlistID := uint64(1)
	ctx := context.Background()

	playlistRepoMock.EXPECT().GetPlaylist(ctx, playlistID).Return(nil, sql.ErrNoRows)

	dtoPlaylist, err := playlistUsecase.GetPlaylist(ctx, playlistID)

	require.Error(t, err)
	require.Nil(t, dtoPlaylist)
	require.Equal(t, err, sql.ErrNoRows)
}

func TestPlaylistUsecaseGetAllPlaylists_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	userClientMock := mock.NewMockUserServiceClient(ctrl)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, userClientMock, logger)

	ownerID := uuid.New()
	playlists := []*models.Playlist{
		{ID: 1, Name: "chill vibes", Image: "/images/playlist_1.jpg", OwnerID: ownerID, IsPrivate: false},
		{ID: 2, Name: "workout beats", Image: "/images/playlist_2.jpg", OwnerID: ownerID, IsPrivate: false},
	}

	findByIDResponseUser := &userService.FindByIDResponse{
		User: &userService.User{
			Uuid:     ownerID.String(),
			Username: "user2",
		},
	}

	ctx := context.Background()
	playlistRepoMock.EXPECT().GetAllPlaylists(ctx).Return(playlists, nil)
	userClientMock.EXPECT().FindByID(ctx, gomock.Any()).Return(findByIDResponseUser, nil).Times(len(playlists))

	dtoPlaylists, err := playlistUsecase.GetAllPlaylists(ctx)

	require.NoError(t, err)
	require.Len(t, dtoPlaylists, len(playlists))
	require.Equal(t, playlists[0].Name, dtoPlaylists[0].Name)
}

func TestPlaylistUsecaseGetAllPlaylists_ConnDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	userClientMock := mock.NewMockUserServiceClient(ctrl)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, userClientMock, logger)

	ctx := context.Background()
	playlistRepoMock.EXPECT().GetAllPlaylists(ctx).Return(nil, sql.ErrConnDone)

	dtoPlaylists, err := playlistUsecase.GetAllPlaylists(ctx)

	require.Error(t, err)
	require.Nil(t, dtoPlaylists)
}

func TestPlaylistUsecaseAddToPlaylist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, nil, logger)

	playlistID := uint64(1)
	trackID := uint64(42)
	length := uint64(5)

	mockPlaylistTrack := &models.PlaylistTrack{
		ID:                   10,
		PlaylistID:           playlistID,
		TrackOrderInPlaylist: length + 1,
		TrackID:              trackID,
	}

	playlistRepoMock.EXPECT().GetLengthPlaylist(context.Background(), playlistID).Return(length, nil)
	playlistRepoMock.EXPECT().AddToPlaylist(context.Background(), playlistID, length+1, trackID).Return(mockPlaylistTrack, nil)

	result, err := playlistUsecase.AddToPlaylist(context.Background(), &dto.PlaylistTrackDTO{
		PlaylistID: playlistID,
		TrackID:    trackID,
	})

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, mockPlaylistTrack.TrackID, result.TrackID)
}

func TestPlaylistUsecaseAddToPlaylist_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, nil, logger)

	playlistID := uint64(1)
	trackID := uint64(42)

	playlistRepoMock.EXPECT().GetLengthPlaylist(context.Background(), playlistID).Return(uint64(0), sql.ErrNoRows)

	result, err := playlistUsecase.AddToPlaylist(context.Background(), &dto.PlaylistTrackDTO{
		PlaylistID: playlistID,
		TrackID:    trackID,
	})

	require.Error(t, err)
	require.Nil(t, result)
}

func TestPlaylistUsecaseGetUserPlaylists_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	userClientMock := mock.NewMockUserServiceClient(ctrl)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, userClientMock, logger)

	userID := uuid.New()
	playlists := []*models.Playlist{
		{ID: 1, Name: "chill vibes", Image: "/images/playlist_1.jpg", OwnerID: userID, IsPrivate: false},
		{ID: 2, Name: "workout beats", Image: "/images/playlist_2.jpg", OwnerID: userID, IsPrivate: false},
	}

	findByIDResponseUser := &userService.FindByIDResponse{
		User: &userService.User{
			Uuid:     userID.String(),
			Username: "user2",
		},
	}

	ctx := context.Background()
	playlistRepoMock.EXPECT().GetUserPlaylists(ctx, userID).Return(playlists, nil)
	userClientMock.EXPECT().FindByID(ctx, &userService.FindByIDRequest{Uuid: userID.String()}).Return(findByIDResponseUser, nil)

	dtoPlaylists, err := playlistUsecase.GetUserPlaylists(ctx, userID)

	require.NoError(t, err)
	require.Len(t, dtoPlaylists, len(playlists))
	require.Equal(t, "user2", dtoPlaylists[0].OwnerName)
}

func TestPlaylistUsecaseGetUserPlaylists_ConnDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	userClientMock := mock.NewMockUserServiceClient(ctrl)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, userClientMock, logger)

	userID := uuid.New()
	ctx := context.Background()

	playlistRepoMock.EXPECT().GetUserPlaylists(ctx, userID).Return(nil, sql.ErrConnDone)

	dtoPlaylists, err := playlistUsecase.GetUserPlaylists(ctx, userID)

	require.Error(t, err)
	require.Nil(t, dtoPlaylists)
}

func TestPlaylistUsecaseRemoveFromPlaylist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, nil, logger)

	playlistID := uint64(1)
	trackID := uint64(42)

	ctx := context.Background()
	playlistRepoMock.EXPECT().RemoveFromPlaylist(ctx, playlistID, trackID).Return(sqlmock.NewResult(1, 1), nil)

	err := playlistUsecase.RemoveFromPlaylist(ctx, &dto.PlaylistTrackDTO{
		PlaylistID: playlistID,
		TrackID:    trackID,
	})

	require.NoError(t, err)
}

func TestPlaylistUsecaseRemoveFromPlaylist_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, nil, logger)

	playlistID := uint64(1)
	trackID := uint64(42)

	ctx := context.Background()
	playlistRepoMock.EXPECT().RemoveFromPlaylist(ctx, playlistID, trackID).Return(nil, sql.ErrNoRows)

	err := playlistUsecase.RemoveFromPlaylist(ctx, &dto.PlaylistTrackDTO{
		PlaylistID: playlistID,
		TrackID:    trackID,
	})

	require.Error(t, err)
}

func TestPlaylistUsecaseDeletePlaylist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, nil, logger)

	playlistID := uint64(1)

	ctx := context.Background()
	playlistRepoMock.EXPECT().DeletePlaylist(ctx, playlistID).Return(sqlmock.NewResult(1, 1), nil)

	err := playlistUsecase.DeletePlaylist(ctx, playlistID)

	require.NoError(t, err)
}

func TestPlaylistUsecaseDeletePlaylist_ConnDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	playlistRepoMock := mock.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepoMock, nil, logger)

	playlistID := uint64(1)

	ctx := context.Background()
	playlistRepoMock.EXPECT().DeletePlaylist(ctx, playlistID).Return(nil, sql.ErrConnDone)

	err := playlistUsecase.DeletePlaylist(ctx, playlistID)

	require.Error(t, err)
}
