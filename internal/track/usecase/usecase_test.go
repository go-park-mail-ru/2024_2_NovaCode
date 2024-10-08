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
	mockTrack "github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_View_FoundTrack(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, albumRepoMock, artistRepoMock, logger)

	track := &models.Track{
		ID:          1,
		Name:        "ok im cool",
		Genre:       "Rap",
		Duration:    167,
		FilePath:    "/songs/track_1.mp4",
		Image:       "/imgs/tracks/track_1.jpg",
		ArtistID:    1,
		AlbumID:     1,
		ReleaseDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
	}

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
	trackRepoMock.EXPECT().FindById(ctx, track.ID).Return(track, nil)
	artistRepoMock.EXPECT().FindById(ctx, track.ArtistID).Return(artist, nil)
	albumRepoMock.EXPECT().FindById(ctx, track.AlbumID).Return(album, nil)

	dtoTrack, err := trackUsecase.View(ctx, track.ID)

	require.NoError(t, err)
	require.NotNil(t, dtoTrack)
	require.Equal(t, track.Name, dtoTrack.Name)
	require.Equal(t, artist.Name, dtoTrack.Artist)
	require.Equal(t, album.Name, dtoTrack.Album)
}

func TestUsecase_View_NotFoundTrack(t *testing.T) {
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

	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, albumRepoMock, artistRepoMock, logger)

	trackRepoMock.EXPECT().FindById(ctx, uint64(1)).Return(nil, errors.New("Track wasn't found"))
	dtoTrack, err := trackUsecase.View(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoTrack)
	require.EqualError(t, err, "Track wasn't found")
}

func TestUsecase_Search_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, albumRepoMock, artistRepoMock, logger)

	now := time.Now()
	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test", Genre: "1", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "none", Genre: "2", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(2), AlbumID: uint64(2), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "test", Genre: "3", Duration: uint64(3), FilePath: "3", Image: "3",
			ArtistID: uint64(3), AlbumID: uint64(3), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

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
	artistRepoMock.EXPECT().FindById(ctx, artists[0].ID).Return(artists[0], nil)
	artistRepoMock.EXPECT().FindById(ctx, artists[2].ID).Return(artists[2], nil)
	albumRepoMock.EXPECT().FindById(ctx, albums[0].ID).Return(albums[0], nil)
	albumRepoMock.EXPECT().FindById(ctx, albums[2].ID).Return(albums[2], nil)
	trackRepoMock.EXPECT().FindByName(ctx, "test").Return([]*models.Track{tracks[0], tracks[2]}, nil)

	dtoTracks, err := trackUsecase.Search(ctx, "test")

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, 2, len(dtoTracks))
	require.Equal(t, artists[0].Name, dtoTracks[0].Artist)
	require.Equal(t, albums[0].Name, dtoTracks[0].Album)
	require.Equal(t, artists[2].Name, dtoTracks[1].Artist)
	require.Equal(t, albums[2].Name, dtoTracks[1].Album)
}

func TestUsecase_Search_NotFoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, albumRepoMock, artistRepoMock, logger)

	ctx := context.Background()
	trackRepoMock.EXPECT().FindByName(ctx, "song").Return(nil, errors.New("Can't find tracks"))

	dtoTracks, err := trackUsecase.Search(ctx, "song")

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, "Can't find tracks")
}

func TestUsecase_GetAll_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, albumRepoMock, artistRepoMock, logger)

	now := time.Now()
	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test1", Genre: "1", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test2", Genre: "2", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(2), AlbumID: uint64(2), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "test3", Genre: "3", Duration: uint64(3), FilePath: "3", Image: "3",
			ArtistID: uint64(3), AlbumID: uint64(3), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

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
	for i := 0; i < len(tracks); i++ {
		artistRepoMock.EXPECT().FindById(ctx, artists[i].ID).Return(artists[i], nil)
		albumRepoMock.EXPECT().FindById(ctx, albums[i].ID).Return(albums[i], nil)
	}
	trackRepoMock.EXPECT().GetAll(ctx).Return(tracks, nil)

	dtoTracks, err := trackUsecase.GetAll(ctx)

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, len(tracks), len(dtoTracks))

	for i := 0; i < len(tracks); i++ {
		require.Equal(t, tracks[i].Name, dtoTracks[i].Name)
		require.Equal(t, artists[i].Name, dtoTracks[i].Artist)
		require.Equal(t, albums[i].Name, dtoTracks[i].Album)
	}
}

func TestUsecase_GetAll_NotFoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistRepoMock := mockArtist.NewMockRepo(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, albumRepoMock, artistRepoMock, logger)

	ctx := context.Background()
	trackRepoMock.EXPECT().GetAll(ctx).Return(nil, errors.New("Can't load tracks"))
	dtoTracks, err := trackUsecase.GetAll(ctx)

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, "Can't load tracks")
}
