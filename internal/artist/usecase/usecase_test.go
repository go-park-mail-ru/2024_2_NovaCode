package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	mockArtist "github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
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
	artistRepoMock.EXPECT().FindByName(ctx, "artist").Return([]*models.Artist{artists[1], artists[2]}, nil)

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
	artistRepoMock.EXPECT().FindByName(ctx, "artist").Return(nil, errors.New("Can't find artist"))

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
