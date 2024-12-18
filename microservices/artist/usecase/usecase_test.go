package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	uuid "github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/dto"
	mockArtist "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_View_FoundArtist(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	artist := &models.Artist{
		ID:      1,
		Name:    "quinn",
		Bio:     "Some random bio",
		Country: "USA",
		Image:   "/imgs/artists/artist_1.jpg",
	}

	ctx := context.Background()
	artistRepoMock.EXPECT().FindById(ctx, artist.ID).Return(artist, nil)

	dtoArtist, err := artistUsecase.View(ctx, artist.ID)

	require.NoError(t, err)
	require.NotNil(t, dtoArtist)
	require.Equal(t, artist.Name, dtoArtist.Name)
}

func TestUsecase_View_NotFoundArtist(t *testing.T) {
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
	ctx := context.Background()

	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	artistRepoMock.EXPECT().FindById(ctx, uint64(1)).Return(nil, errors.New("Artist wasn't found"))
	dtoArtist, err := artistUsecase.View(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoArtist)
	require.EqualError(t, err, "Artist wasn't found")
}

func TestUsecase_Search_FoundArtists(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	now := time.Now()
	artists := []*models.Artist{
		{ID: uint64(1), Name: "test", Bio: "1", Country: "1", Image: "1", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(2), Name: "artist", Bio: "2", Country: "2", Image: "2", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(3), Name: "artist", Bio: "3", Country: "3", Image: "3", CreatedAt: now, UpdatedAt: now},
	}

	ctx := context.Background()
	artistRepoMock.EXPECT().FindByQuery(ctx, "artist").Return([]*models.Artist{artists[1], artists[2]}, nil)

	dtoArtists, err := artistUsecase.Search(ctx, "artist")

	require.NoError(t, err)
	require.NotNil(t, dtoArtists)
	require.Equal(t, 2, len(dtoArtists))
	require.Equal(t, artists[1].Name, dtoArtists[0].Name)
	require.Equal(t, artists[2].Name, dtoArtists[1].Name)
}

func TestUsecase_Search_NotFoundArtists(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	ctx := context.Background()
	artistRepoMock.EXPECT().FindByQuery(ctx, "artist").Return(nil, errors.New("Can't find artist"))

	dtoArtists, err := artistUsecase.Search(ctx, "artist")

	require.Error(t, err)
	require.Nil(t, dtoArtists)
	require.EqualError(t, err, "Can't find artist")
}

func TestUsecase_GetAll_FoundArtists(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	now := time.Now()
	artists := []*models.Artist{
		{ID: uint64(1), Name: "artist1", Bio: "1", Country: "1", Image: "1", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(2), Name: "artist2", Bio: "2", Country: "2", Image: "2", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(3), Name: "artist3", Bio: "3", Country: "3", Image: "3", CreatedAt: now, UpdatedAt: now},
	}

	ctx := context.Background()
	artistRepoMock.EXPECT().GetAll(ctx).Return(artists, nil)

	dtoArtists, err := artistUsecase.GetAll(ctx)

	require.NoError(t, err)
	require.NotNil(t, dtoArtists)
	require.Equal(t, len(artists), len(dtoArtists))

	for i := 0; i < len(artists); i++ {
		require.Equal(t, artists[i].Name, dtoArtists[i].Name)
	}
}

func TestUsecase_GetAll_NotFoundArtists(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	ctx := context.Background()
	artistRepoMock.EXPECT().GetAll(ctx).Return(nil, errors.New("Can't load artists"))
	dtoArtists, err := artistUsecase.GetAll(ctx)

	require.Error(t, err)
	require.Nil(t, dtoArtists)
	require.EqualError(t, err, "Can't load artists")
}

func TestArtistUsecase_AddFavoriteArtist(t *testing.T) {
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

	mockArtistRepo := mockArtist.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	artistUsecase := &artistUsecase{
		artistRepo: mockArtistRepo,
		logger:     logger,
	}

	userID := uuid.New()
	artistID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockArtistRepo.EXPECT().AddFavoriteArtist(ctx, userID, artistID).Return(nil)
		err := artistUsecase.AddFavoriteArtist(ctx, userID, artistID)
		require.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockArtistRepo.EXPECT().AddFavoriteArtist(ctx, userID, artistID).Return(mockError)

		err := artistUsecase.AddFavoriteArtist(ctx, userID, artistID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
	})
}

func TestArtistUsecase_DeleteFavoriteArtist(t *testing.T) {
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

	mockArtistRepo := mockArtist.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	artistUsecase := &artistUsecase{
		artistRepo: mockArtistRepo,
		logger:     logger,
	}

	userID := uuid.New()
	artistID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockArtistRepo.EXPECT().DeleteFavoriteArtist(ctx, userID, artistID).Return(nil)
		err := artistUsecase.DeleteFavoriteArtist(ctx, userID, artistID)
		require.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockArtistRepo.EXPECT().DeleteFavoriteArtist(ctx, userID, artistID).Return(mockError)

		err := artistUsecase.DeleteFavoriteArtist(ctx, userID, artistID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
	})
}

func TestArtistUsecase_IsFavoriteArtist(t *testing.T) {
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

	mockArtistRepo := mockArtist.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	artistUsecase := &artistUsecase{
		artistRepo: mockArtistRepo,
		logger:     logger,
	}

	userID := uuid.New()
	artistID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockArtistRepo.EXPECT().IsFavoriteArtist(ctx, userID, artistID).Return(true, nil)
		exists, err := artistUsecase.IsFavoriteArtist(ctx, userID, artistID)
		require.NoError(t, err)
		require.True(t, exists)
	})

	t.Run("artist not found", func(t *testing.T) {
		mockArtistRepo.EXPECT().IsFavoriteArtist(ctx, userID, artistID).Return(false, nil)
		exists, err := artistUsecase.IsFavoriteArtist(ctx, userID, artistID)
		require.NoError(t, err)
		require.False(t, exists)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockArtistRepo.EXPECT().IsFavoriteArtist(ctx, userID, artistID).Return(false, mockError)

		exists, err := artistUsecase.IsFavoriteArtist(ctx, userID, artistID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
		require.False(t, exists)
	})
}

func TestUsecase_GetFavoriteArtistsCount_FoundArtists(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	userID := uuid.New()
	ctx := context.Background()
	expectedCount := uint64(5)
	artistRepoMock.EXPECT().GetFavoriteArtistsCount(ctx, userID).Return(expectedCount, nil)

	count, err := artistUsecase.GetFavoriteArtistsCount(ctx, userID)

	require.NoError(t, err)
	require.Equal(t, expectedCount, count)
}

func TestUsecase_GetFavoriteArtistsCount_ErrorGettingCount(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	userID := uuid.New()
	ctx := context.Background()
	expectedError := fmt.Errorf("Can't load artists by user ID %v", userID)
	artistRepoMock.EXPECT().GetFavoriteArtistsCount(ctx, userID).Return(uint64(0), expectedError)

	count, err := artistUsecase.GetFavoriteArtistsCount(ctx, userID)

	require.Error(t, err)
	require.EqualError(t, err, expectedError.Error())
	require.Equal(t, uint64(0), count)
}

func TestUsecase_GetArtistLikesCount_Success(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	ctx := context.Background()
	ctx = context.WithValue(ctx, utils.RequestIDKey{}, "test-request-id")

	artistID := uint64(123)
	expectedLikesCount := uint64(10)

	artistRepoMock.EXPECT().GetArtistLikesCount(ctx, artistID).Return(expectedLikesCount, nil)

	likesCount, err := artistUsecase.GetArtistLikesCount(ctx, artistID)
	require.NoError(t, err)
	require.Equal(t, expectedLikesCount, likesCount)
}

func TestUsecase_GetArtistLikesCount_ErrorGettingLikesCount(t *testing.T) {
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
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	artistUsecase := NewArtistUsecase(artistRepoMock, logger)

	ctx := context.Background()
	ctx = context.WithValue(ctx, utils.RequestIDKey{}, "test-request-id")

	artistID := uint64(123)
	expectedError := fmt.Errorf("Can't load artist likes count by artist ID %v", artistID)

	artistRepoMock.EXPECT().GetArtistLikesCount(ctx, artistID).Return(uint64(0), expectedError)

	likesCount, err := artistUsecase.GetArtistLikesCount(ctx, artistID)
	require.Error(t, err)
	require.EqualError(t, err, expectedError.Error())
	require.Equal(t, uint64(0), likesCount)
}

func TestArtistUsecase_GetPopular(t *testing.T) {
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

	mockArtistRepo := mockArtist.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	artistUsecase := &artistUsecase{
		artistRepo: mockArtistRepo,
		logger:     logger,
	}

	requestID := uuid.New()
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	artists := []*models.Artist{
		{
			ID:        1,
			Name:      "Artist 1",
			Bio:       "Bio 1",
			Country:   "Country A",
			Image:     "/images/artists/artist_1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Artist 2",
			Bio:       "Bio 2",
			Country:   "Country B",
			Image:     "/images/artists/artist_2.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	dtoArtists := []*dto.ArtistDTO{
		{
			ID:      1,
			Name:    "Artist 1",
			Bio:     "Bio 1",
			Country: "Country A",
			Image:   "/images/artists/artist_1.jpg",
		},
		{
			ID:      2,
			Name:    "Artist 2",
			Bio:     "Bio 2",
			Country: "Country B",
			Image:   "/images/artists/artist_2.jpg",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockArtistRepo.EXPECT().GetPopular(ctx).Return(artists, nil)

		for i, artist := range artists {
			artistUsecase.logger.Info(fmt.Sprintf("Convert artist %s to DTO", artist.Name), requestID)
			dtoArtist := dtoArtists[i]
			require.Equal(t, artist.ID, dtoArtist.ID)
			require.Equal(t, artist.Name, dtoArtist.Name)
			require.Equal(t, artist.Bio, dtoArtist.Bio)
			require.Equal(t, artist.Country, dtoArtist.Country)
			require.Equal(t, artist.Image, dtoArtist.Image)
		}

		result, err := artistUsecase.GetPopular(ctx)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, len(dtoArtists), len(result))

		for i, dtoArtist := range dtoArtists {
			require.Equal(t, dtoArtist.ID, result[i].ID)
			require.Equal(t, dtoArtist.Name, result[i].Name)
			require.Equal(t, dtoArtist.Bio, result[i].Bio)
			require.Equal(t, dtoArtist.Country, result[i].Country)
			require.Equal(t, dtoArtist.Image, result[i].Image)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockArtistRepo.EXPECT().GetPopular(ctx).Return(nil, mockError)

		result, err := artistUsecase.GetPopular(ctx)

		require.Error(t, err)
		require.Nil(t, result)
	})
}