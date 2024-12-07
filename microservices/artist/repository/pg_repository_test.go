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

func TestArtistRepositoryCreate(t *testing.T) {
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

	artistPGRepository := NewArtistPGRepository(db, logger)
	mockArtist := &models.Artist{
		ID:      1,
		Name:    "quinn",
		Bio:     "Some random bio",
		Country: "USA",
		Image:   "/imgs/artists/artist_1.jpg",
	}

	columns := []string{"id", "name", "bio", "country", "image", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockArtist.ID,
		mockArtist.Name,
		mockArtist.Bio,
		mockArtist.Country,
		mockArtist.Image,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(createArtistQuery).WithArgs(
		mockArtist.Name,
		mockArtist.Bio,
		mockArtist.Country,
		mockArtist.Image,
	).WillReturnRows(rows)

	createdArtist, err := artistPGRepository.Create(context.Background(), mockArtist)
	require.NoError(t, err)
	require.NotNil(t, createdArtist)
}

func TestArtistRepositoryFindById(t *testing.T) {
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

	artistPGRepository := NewArtistPGRepository(db, logger)
	mockArtist := &models.Artist{
		ID:      1,
		Name:    "quinn",
		Bio:     "Some random bio",
		Country: "USA",
		Image:   "/imgs/artists/artist_1.jpg",
	}

	columns := []string{"id", "name", "bio", "country", "image", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockArtist.ID,
		mockArtist.Name,
		mockArtist.Bio,
		mockArtist.Country,
		mockArtist.Image,
		time.Now(),
		time.Now(),
	)

	mock.ExpectPrepare(findByIDQuery).
		ExpectQuery().
		WithArgs(mockArtist.ID).
		WillReturnRows(rows)

	foundArtist, err := artistPGRepository.FindById(context.Background(), mockArtist.ID)
	require.NoError(t, err)
	require.NotNil(t, foundArtist)
	require.Equal(t, mockArtist.ID, foundArtist.ID)
}

func TestArtistRepositoryFindByQuery(t *testing.T) {
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

	artistPGRepository := NewArtistPGRepository(db, logger)
	artists := []models.Artist{
		{
			ID:        1,
			Name:      "First artist",
			Bio:       "Some random bio",
			Country:   "USA",
			Image:     "/imgs/artists/artist_1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Artist for test 1",
			Bio:       "Some random bio",
			Country:   "USA",
			Image:     "/imgs/artists/artist_2.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "Artist for test 2",
			Bio:       "Some random bio",
			Country:   "USA",
			Image:     "/imgs/artists/artist_3.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	columns := []string{"id", "name", "bio", "country", "image", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, artist := range artists {
		rows.AddRow(
			artist.ID,
			artist.Name,
			artist.Bio,
			artist.Country,
			artist.Image,
			artist.CreatedAt,
			artist.UpdatedAt,
		)
	}

	findName := "test"
	expectedArtists := []*models.Artist{&artists[0], &artists[1], &artists[2]}
	mock.ExpectPrepare(findByQuery).
		ExpectQuery().
		WithArgs(utils.MakeSearchQuery(findName)).
		WillReturnRows(rows)

	foundArtists, err := artistPGRepository.FindByQuery(context.Background(), findName)
	require.NoError(t, err)
	require.NotNil(t, foundArtists)
	require.Equal(t, expectedArtists, foundArtists)
}

func TestArtistRepositoryGetAll(t *testing.T) {
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

	artistPGRepository := NewArtistPGRepository(db, logger)
	artists := []models.Artist{
		{
			ID:        1,
			Name:      "First artist",
			Bio:       "Some random bio",
			Country:   "USA",
			Image:     "/imgs/artists/artist_1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Artist for test 1",
			Bio:       "Some random bio",
			Country:   "USA",
			Image:     "/imgs/artists/artist_2.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "Artist for test 2",
			Bio:       "Some random bio",
			Country:   "USA",
			Image:     "/imgs/artists/artist_3.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	columns := []string{"id", "name", "bio", "country", "image", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, artist := range artists {
		rows.AddRow(
			artist.ID,
			artist.Name,
			artist.Bio,
			artist.Country,
			artist.Image,
			artist.CreatedAt,
			artist.UpdatedAt,
		)
	}

	expectedArtists := []*models.Artist{&artists[0], &artists[1], &artists[2]}
	mock.ExpectQuery(getAllQuery).WillReturnRows(rows)

	foundArtists, err := artistPGRepository.GetAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, foundArtists)
	require.Equal(t, foundArtists, expectedArtists)
}
