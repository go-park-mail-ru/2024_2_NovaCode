package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestAlbumRepositoryCreate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	albumPGRepository := NewAlbumPGRepository(db, logger)
	mockAlbum := &models.Album{
		ID:          1,
		Name:        "Attempted Lover",
		ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
		Image:       "/imgs/albums/album_1.jpg",
		ArtistID:    1,
	}

	columns := []string{"id", "name", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockAlbum.ID,
		mockAlbum.Name,
		mockAlbum.ReleaseDate,
		mockAlbum.Image,
		mockAlbum.ArtistID,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(createAlbumQuery).WithArgs(
		mockAlbum.Name,
		mockAlbum.ReleaseDate,
		mockAlbum.Image,
		mockAlbum.ArtistID,
	).WillReturnRows(rows)

	createdAlbum, err := albumPGRepository.Create(context.Background(), mockAlbum)
	require.NoError(t, err)
	require.NotNil(t, createdAlbum)
}

func TestAlbumRepositoryFindById(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	albumPGRepository := NewAlbumPGRepository(db, logger)
	mockAlbum := &models.Album{
		ID:          1,
		Name:        "Attempted Lover",
		ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
		Image:       "/imgs/albums/album_1.jpg",
		ArtistID:    1,
	}

	columns := []string{"id", "name", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockAlbum.ID,
		mockAlbum.Name,
		mockAlbum.ReleaseDate,
		mockAlbum.Image,
		mockAlbum.ArtistID,
		time.Now(),
		time.Now(),
	)

	mock.ExpectPrepare(findByIDQuery).
		ExpectQuery().
		WithArgs(mockAlbum.ID).
		WillReturnRows(rows)

	foundAlbum, err := albumPGRepository.FindById(context.Background(), mockAlbum.ID)
	require.NoError(t, err)
	require.NotNil(t, foundAlbum)
	require.Equal(t, mockAlbum.ID, foundAlbum.ID)
}

func TestAlbumRepositoryFindByQuery(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	albumPGRepository := NewAlbumPGRepository(db, logger)
	albums := []models.Album{
		{
			ID:          1,
			Name:        "Album for test 1",
			ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_1.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Album for test 2",
			ReleaseDate: time.Date(2021, 02, 3, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_2.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "Another album",
			ReleaseDate: time.Date(2019, 01, 5, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_3.jpg",
			ArtistID:    3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, album := range albums {
		rows.AddRow(
			album.ID,
			album.Name,
			album.ReleaseDate,
			album.Image,
			album.ArtistID,
			album.CreatedAt,
			album.UpdatedAt,
		)
	}

	findName := "test"
	expectedAlbums := []*models.Album{&albums[0], &albums[1], &albums[2]}
	mock.ExpectPrepare(findByQuery).
		ExpectQuery().
		WithArgs(utils.MakeSearchQuery(findName)).
		WillReturnRows(rows)

	foundAlbums, err := albumPGRepository.FindByQuery(context.Background(), findName)
	require.NoError(t, err)
	require.NotNil(t, foundAlbums)
	require.Equal(t, expectedAlbums, foundAlbums)
}

func TestAlbumRepositoryGetAll(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	albumPGRepository := NewAlbumPGRepository(db, logger)
	albums := []models.Album{
		{
			ID:          1,
			Name:        "Album for test 1",
			ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_1.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Album for test 2",
			ReleaseDate: time.Date(2021, 02, 3, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_2.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "Another album",
			ReleaseDate: time.Date(2019, 01, 5, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_3.jpg",
			ArtistID:    3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, album := range albums {
		rows.AddRow(
			album.ID,
			album.Name,
			album.ReleaseDate,
			album.Image,
			album.ArtistID,
			album.CreatedAt,
			album.UpdatedAt,
		)
	}

	expectedAlbums := []*models.Album{&albums[0], &albums[1], &albums[2]}
	mock.ExpectQuery(getAllQuery).WillReturnRows(rows)

	foundAlbums, err := albumPGRepository.GetAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, foundAlbums)
	require.Equal(t, foundAlbums, expectedAlbums)
}

func TestAlbumRepositoryGetAllByArtistID(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)

	albumPGRepository := NewAlbumPGRepository(db, logger)
	albums := []models.Album{
		{
			ID:          1,
			Name:        "Album for test 1",
			ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_1.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Album for test 2",
			ReleaseDate: time.Date(2021, 02, 3, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_2.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, album := range albums {
		rows.AddRow(
			album.ID,
			album.Name,
			album.ReleaseDate,
			album.Image,
			album.ArtistID,
			album.CreatedAt,
			album.UpdatedAt,
		)
	}

	expectedAlbums := []*models.Album{&albums[0], &albums[1]}
	mock.ExpectPrepare(getByArtistIDQuery).
		ExpectQuery().
		WithArgs(uint64(1)).
		WillReturnRows(rows)

	foundAlbums, err := albumPGRepository.GetAllByArtistID(context.Background(), uint64(1))
	require.NoError(t, err)
	require.NotNil(t, foundAlbums)
	require.Equal(t, expectedAlbums, foundAlbums)
}
