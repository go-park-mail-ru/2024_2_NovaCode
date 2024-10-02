package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	mockAlbum "github.com/go-park-mail-ru/2024_2_NovaCode/internal/album/mock"
	mockArtist "github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_View_FoundAlbum(t *testing.T) {
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
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistRepoMock, logger)

	artist := &models.Artist{
		ID:      1,
		Name:    "quinn",
		Bio:     "Some random bio",
		Country: "USA",
		Image:   "/imgs/artists/artist_1.jpg",
	}

	album := &models.Album{
		ID:          1,
		Name:        "Attempted Lover",
		Genre:       "Rock",
		TrackCount:  12,
		ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
		Image:       "/imgs/albums/album_1.jpg",
		ArtistID:    1,
	}

	ctx := context.Background()
	artistRepoMock.EXPECT().FindById(ctx, artist.ID).Return(artist, nil)
	albumRepoMock.EXPECT().FindById(ctx, album.ID).Return(album, nil)

	dtoAlbum, err := albumUsecase.View(ctx, album.ID)

	require.NoError(t, err)
	require.NotNil(t, dtoAlbum)
	require.Equal(t, artist.Name, dtoAlbum.Artist)
	require.Equal(t, album.Name, dtoAlbum.Name)
}

func TestUsecase_View_NotFoundAlbums(t *testing.T) {
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
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistRepoMock, logger)

	albumRepoMock.EXPECT().FindById(ctx, uint64(1)).Return(nil, errors.New("Album wasn't found"))
	dtoAlbum, err := albumUsecase.View(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoAlbum)
	require.EqualError(t, err, "Album wasn't found")
}

func TestUsecase_Search_FoundAlbums(t *testing.T) {
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
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistRepoMock, logger)

	now := time.Now()
	artists := []*models.Artist{
		{ID: uint64(1), Name: "artist1", Bio: "1", Country: "1", Image: "1", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(2), Name: "artist2", Bio: "2", Country: "2", Image: "2", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(3), Name: "artist3", Bio: "3", Country: "3", Image: "3", CreatedAt: now, UpdatedAt: now},
	}

	albums := []*models.Album{
		{
			ID: uint64(1), Name: "test", Genre: "1", TrackCount: uint64(1), ReleaseDate: now, Image: "1",
			ArtistID: uint64(1), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test", Genre: "2", TrackCount: uint64(2), ReleaseDate: now, Image: "2",
			ArtistID: uint64(2), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "album3", Genre: "3", TrackCount: uint64(3), ReleaseDate: now, Image: "3",
			ArtistID: uint64(3), CreatedAt: now, UpdatedAt: now,
		},
	}

	ctx := context.Background()
	artistRepoMock.EXPECT().FindById(ctx, artists[0].ID).Return(artists[0], nil)
	artistRepoMock.EXPECT().FindById(ctx, artists[1].ID).Return(artists[1], nil)
	albumRepoMock.EXPECT().FindByName(ctx, "test").Return([]*models.Album{albums[0], albums[1]}, nil)

	dtoAlbums, err := albumUsecase.Search(ctx, "test")

	require.NoError(t, err)
	require.NotNil(t, dtoAlbums)
	require.Equal(t, 2, len(dtoAlbums))
	require.Equal(t, albums[0].Name, dtoAlbums[0].Name)
	require.Equal(t, albums[1].Name, dtoAlbums[1].Name)
	require.Equal(t, artists[0].Name, dtoAlbums[0].Artist)
	require.Equal(t, artists[1].Name, dtoAlbums[1].Artist)
}

func TestUsecase_Search_NotFoundAlbums(t *testing.T) {
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
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistRepoMock, logger)

	ctx := context.Background()
	albumRepoMock.EXPECT().FindByName(ctx, "album").Return(nil, errors.New("Can't find albums"))

	dtoAlbums, err := albumUsecase.Search(ctx, "album")

	require.Error(t, err)
	require.Nil(t, dtoAlbums)
	require.EqualError(t, err, "Can't find albums")
}

func TestUsecase_GetAll_FoundAlbums(t *testing.T) {
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
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistRepoMock, logger)

	now := time.Now()
	artists := []*models.Artist{
		{ID: uint64(1), Name: "artist1", Bio: "1", Country: "1", Image: "1", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(2), Name: "artist2", Bio: "2", Country: "2", Image: "2", CreatedAt: now, UpdatedAt: now},
		{ID: uint64(3), Name: "artist3", Bio: "3", Country: "3", Image: "3", CreatedAt: now, UpdatedAt: now},
	}

	albums := []*models.Album{
		{
			ID: uint64(1), Name: "album1", Genre: "1", TrackCount: uint64(1), ReleaseDate: now, Image: "1",
			ArtistID: uint64(1), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "album2", Genre: "2", TrackCount: uint64(2), ReleaseDate: now, Image: "2",
			ArtistID: uint64(2), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "album3", Genre: "3", TrackCount: uint64(3), ReleaseDate: now, Image: "3",
			ArtistID: uint64(3), CreatedAt: now, UpdatedAt: now,
		},
	}

	ctx := context.Background()
	for i := 0; i < len(albums); i++ {
		artistRepoMock.EXPECT().FindById(ctx, artists[i].ID).Return(artists[i], nil)
	}
	albumRepoMock.EXPECT().GetAll(ctx).Return(albums, nil)

	dtoAlbums, err := albumUsecase.GetAll(ctx)

	require.NoError(t, err)
	require.NotNil(t, dtoAlbums)
	require.Equal(t, len(albums), len(dtoAlbums))

	for i := 0; i < len(albums); i++ {
		require.Equal(t, artists[i].Name, dtoAlbums[i].Artist)
		require.Equal(t, albums[i].Name, dtoAlbums[i].Name)
	}
}

func TestUsecase_GetAll_NotFoundAlbums(t *testing.T) {
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
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistRepoMock, logger)

	ctx := context.Background()
	albumRepoMock.EXPECT().GetAll(ctx).Return(nil, errors.New("Can't find albums"))
	dtoAlbums, err := albumUsecase.GetAll(ctx)

	require.Error(t, err)
	require.Nil(t, dtoAlbums)
	require.EqualError(t, err, "Can't find albums")
}
