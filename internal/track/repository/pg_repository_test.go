package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestTrackRepositoryCreate(t *testing.T) {
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

	trackPGRepository := NewTrackPGRepository(db, logger)
	mockTrack := &models.Track{
		ID:          1,
		Name:        "ok im cool",
		Duration:    167,
		FilePath:    "/songs/track_1.mp4",
		Image:       "/imgs/tracks/track_1.jpg",
		ArtistID:    1,
		AlbumID:     1,
		ReleaseDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "release", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockTrack.ID,
		mockTrack.Name,
		mockTrack.Duration,
		mockTrack.FilePath,
		mockTrack.Image,
		mockTrack.ArtistID,
		mockTrack.AlbumID,
		mockTrack.ReleaseDate,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(createTrackQuery).WithArgs(
		mockTrack.Name,
		mockTrack.Duration,
		mockTrack.FilePath,
		mockTrack.Image,
		mockTrack.ArtistID,
		mockTrack.AlbumID,
		mockTrack.ReleaseDate,
	).WillReturnRows(rows)

	createdTrack, err := trackPGRepository.Create(context.Background(), mockTrack)
	require.NoError(t, err)
	require.NotNil(t, createdTrack)
}

func TestTrackRepositoryFindById(t *testing.T) {
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

	trackPGRepository := NewTrackPGRepository(db, logger)
	mockTrack := &models.Track{
		ID:          1,
		Name:        "ok im cool",
		Duration:    167,
		FilePath:    "/songs/track_1.mp4",
		Image:       "/imgs/tracks/track_1.jpg",
		ArtistID:    1,
		AlbumID:     1,
		ReleaseDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "release", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockTrack.ID,
		mockTrack.Name,
		mockTrack.Duration,
		mockTrack.FilePath,
		mockTrack.Image,
		mockTrack.ArtistID,
		mockTrack.AlbumID,
		mockTrack.ReleaseDate,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(findByIDQuery).WithArgs(mockTrack.ID).WillReturnRows(rows)

	foundTrack, err := trackPGRepository.FindById(context.Background(), mockTrack.ID)
	require.NoError(t, err)
	require.NotNil(t, foundTrack)
	require.Equal(t, foundTrack.ID, foundTrack.ID)
}

func TestTrackRepositoryFindByName(t *testing.T) {
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

	trackPGRepository := NewTrackPGRepository(db, logger)
	tracks := []models.Track{
		{
			ID:          1,
			Name:        "test song 1",
			Duration:    123,
			FilePath:    "/songs/track_1.mp4",
			Image:       "/imgs/tracks/track_1.jpg",
			ArtistID:    1,
			AlbumID:     1,
			ReleaseDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "another song",
			Duration:    93,
			FilePath:    "/songs/track_2.mp4",
			Image:       "/imgs/tracks/track_2.jpg",
			ArtistID:    2,
			AlbumID:     2,
			ReleaseDate: time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "song test",
			Duration:    99,
			FilePath:    "/songs/track_3.mp4",
			Image:       "/imgs/tracks/track_3.jpg",
			ArtistID:    3,
			AlbumID:     3,
			ReleaseDate: time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "release", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, track := range tracks {
		rows.AddRow(
			track.ID,
			track.Name,
			track.Duration,
			track.FilePath,
			track.Image,
			track.ArtistID,
			track.AlbumID,
			track.ReleaseDate,
			track.CreatedAt,
			track.UpdatedAt,
		)
	}

	findName := "test"
	expectedTracks := []*models.Track{&tracks[0], &tracks[2]}
	mock.ExpectQuery(findByNameQuery).WithArgs(findName).WillReturnRows(rows)

	foundTracks, err := trackPGRepository.FindByName(context.Background(), findName)
	require.NoError(t, err)
	require.NotNil(t, foundTracks)
	require.Equal(t, foundTracks, expectedTracks)
}

func TestTrackRepositoryGetAll(t *testing.T) {
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

	trackPGRepository := NewTrackPGRepository(db, logger)
	tracks := []models.Track{
		{
			ID:          1,
			Name:        "test song 1",
			Duration:    123,
			FilePath:    "/songs/track_1.mp4",
			Image:       "/imgs/tracks/track_1.jpg",
			ArtistID:    1,
			AlbumID:     1,
			ReleaseDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "another song",
			Duration:    93,
			FilePath:    "/songs/track_2.mp4",
			Image:       "/imgs/tracks/track_2.jpg",
			ArtistID:    2,
			AlbumID:     2,
			ReleaseDate: time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "song test",
			Duration:    99,
			FilePath:    "/songs/track_3.mp4",
			Image:       "/imgs/tracks/track_3.jpg",
			ArtistID:    3,
			AlbumID:     3,
			ReleaseDate: time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "release", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, track := range tracks {
		rows.AddRow(
			track.ID,
			track.Name,
			track.Duration,
			track.FilePath,
			track.Image,
			track.ArtistID,
			track.AlbumID,
			track.ReleaseDate,
			track.CreatedAt,
			track.UpdatedAt,
		)
	}

	expectedTracks := []*models.Track{&tracks[0], &tracks[1], &tracks[2]}
	mock.ExpectQuery(getAllQuery).WillReturnRows(rows)

	foundTracks, err := trackPGRepository.GetAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, foundTracks)
	require.Equal(t, foundTracks, expectedTracks)
}

func TestTrackRepositoryGetAllByArtistID(t *testing.T) {
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

	trackPGRepository := NewTrackPGRepository(db, logger)

	tracks := []models.Track{
		{
			ID:          1,
			Name:        "test song 1",
			Duration:    123,
			FilePath:    "/songs/track_1.mp4",
			Image:       "/imgs/tracks/track_1.jpg",
			ArtistID:    1,
			AlbumID:     1,
			ReleaseDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "another song",
			Duration:    93,
			FilePath:    "/songs/track_2.mp4",
			Image:       "/imgs/tracks/track_2.jpg",
			ArtistID:    1,
			AlbumID:     2,
			ReleaseDate: time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "release", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, track := range tracks {
		rows.AddRow(
			track.ID,
			track.Name,
			track.Duration,
			track.FilePath,
			track.Image,
			track.ArtistID,
			track.AlbumID,
			track.ReleaseDate,
			track.CreatedAt,
			track.UpdatedAt,
		)
	}

	expectedTracks := []*models.Track{&tracks[0], &tracks[1]}
	mock.ExpectQuery(getByArtistIDQuery).WithArgs(1).WillReturnRows(rows)

	foundTracks, err := trackPGRepository.GetAllByArtistID(context.Background(), uint64(1))
	require.NoError(t, err)
	require.NotNil(t, foundTracks)
	require.Equal(t, foundTracks, expectedTracks)
}
