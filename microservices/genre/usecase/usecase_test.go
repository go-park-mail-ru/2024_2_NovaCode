package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	mockGenre "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/genre/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_GetAll_FoundGenres(t *testing.T) {
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
	genreRepoMock := mockGenre.NewMockRepo(ctrl)
	genreUsecase := NewGenreUsecase(genreRepoMock, logger)

	now := time.Now()
	genres := []*models.Genre{
		{ID: uint64(1), Name: "genre1", RusName: "жанр1", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(2), Name: "genre2", RusName: "жанр2", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(3), Name: "genre3", RusName: "жанр3", CreatedAt: now, UpdatedAt: now},
	}

	ctx := context.Background()
	genreRepoMock.EXPECT().GetAll(ctx).Return(genres, nil)

	dtoGenres, err := genreUsecase.GetAll(ctx)

	require.NoError(t, err)
	require.NotNil(t, dtoGenres)
	require.Equal(t, len(genres), len(dtoGenres))

	for i := 0; i < len(genres); i++ {
		require.Equal(t, genres[i].Name, dtoGenres[i].Name)
		require.Equal(t, genres[i].RusName, dtoGenres[i].RusName)
	}
}

func TestUsecase_GetAll_NotFoundGenres(t *testing.T) {
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
	genreRepoMock := mockGenre.NewMockRepo(ctrl)
	genreUsecase := NewGenreUsecase(genreRepoMock, logger)

	ctx := context.Background()
	genreRepoMock.EXPECT().GetAll(ctx).Return(nil, errors.New("Can't find genres"))
	dtoGenres, err := genreUsecase.GetAll(ctx)

	require.Error(t, err)
	require.Nil(t, dtoGenres)
	require.EqualError(t, err, "Can't find genres")
}

func TestUsecase_GetAllByArtistID_FoundGenres(t *testing.T) {
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
	genreRepoMock := mockGenre.NewMockRepo(ctrl)
	genreUsecase := NewGenreUsecase(genreRepoMock, logger)

	now := time.Now()

	genres := []*models.Genre{
		{
			ID: uint64(1), Name: "genre1", RusName: "жанр1", CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "genre2", RusName: "жанр2", CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "genre3", RusName: "жанр3", CreatedAt: now, UpdatedAt: now,
		},
	}

	ctx := context.Background()
	genreRepoMock.EXPECT().GetAllByArtistID(ctx, uint64(1)).Return(genres, nil)

	dtoGenres, err := genreUsecase.GetAllByArtistID(ctx, uint64(1))

	require.NoError(t, err)
	require.NotNil(t, dtoGenres)
	require.Equal(t, len(genres), len(dtoGenres))

	for i := 0; i < len(genres); i++ {
		require.Equal(t, genres[i].Name, dtoGenres[i].Name)
		require.Equal(t, genres[i].RusName, dtoGenres[i].RusName)
	}
}

func TestUsecase_GetAllByArtistID_NotFoundGenres(t *testing.T) {
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
	genreRepoMock := mockGenre.NewMockRepo(ctrl)
	genreUsecase := NewGenreUsecase(genreRepoMock, logger)

	ctx := context.Background()
	genreRepoMock.EXPECT().GetAllByArtistID(ctx, uint64(1)).Return(nil, errors.New("Can't load genres by artist ID 1"))

	dtoGenres, err := genreUsecase.GetAllByArtistID(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoGenres)
	require.EqualError(t, err, "Can't load genres by artist ID 1")
}

func TestUsecase_GetAllByTrackID_FoundGenres(t *testing.T) {
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
	genreRepoMock := mockGenre.NewMockRepo(ctrl)
	genreUsecase := NewGenreUsecase(genreRepoMock, logger)

	now := time.Now()

	genres := []*models.Genre{
		{
			ID: uint64(1), Name: "genre1", RusName: "жанр1", CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "genre2", RusName: "жанр2", CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "genre3", RusName: "жанр3", CreatedAt: now, UpdatedAt: now,
		},
	}

	ctx := context.Background()
	genreRepoMock.EXPECT().GetAllByTrackID(ctx, uint64(1)).Return(genres, nil)

	dtoGenres, err := genreUsecase.GetAllByTrackID(ctx, uint64(1))

	require.NoError(t, err)
	require.NotNil(t, dtoGenres)
	require.Equal(t, len(genres), len(dtoGenres))

	for i := 0; i < len(genres); i++ {
		require.Equal(t, genres[i].Name, dtoGenres[i].Name)
		require.Equal(t, genres[i].RusName, dtoGenres[i].RusName)
	}
}

func TestUsecase_GetAllByTrackID_NotFoundGenres(t *testing.T) {
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
	genreRepoMock := mockGenre.NewMockRepo(ctrl)
	genreUsecase := NewGenreUsecase(genreRepoMock, logger)

	ctx := context.Background()
	genreRepoMock.EXPECT().GetAllByTrackID(ctx, uint64(1)).Return(nil, errors.New("Can't load genres by track ID 1"))

	dtoGenres, err := genreUsecase.GetAllByTrackID(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoGenres)
	require.EqualError(t, err, "Can't load genres by track ID 1")
}
