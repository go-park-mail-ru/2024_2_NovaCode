package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/stretchr/testify/require"
)

func TestArtistRepositoryCreate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistPGRepository := NewArtistPGRepository(db)
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

	artistPGRepository := NewArtistPGRepository(db)
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

	mock.ExpectQuery(findByIDQuery).WithArgs(mockArtist.ID).WillReturnRows(rows)

	foundArtist, err := artistPGRepository.FindById(context.Background(), mockArtist.ID)
	require.NoError(t, err)
	require.NotNil(t, foundArtist)
	require.Equal(t, foundArtist.ID, foundArtist.ID)
}

func TestArtistRepositoryFindByName(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistPGRepository := NewArtistPGRepository(db)
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
	expectedArtists := []*models.Artist{&artists[1], &artists[2]}
	mock.ExpectQuery(findByNameQuery).WithArgs(findName).WillReturnRows(rows)

	foundArtists, err := artistPGRepository.FindByName(context.Background(), findName)
	require.NoError(t, err)
	require.NotNil(t, foundArtists)
	require.Equal(t, foundArtists, expectedArtists)
}

func TestArtistRepositoryGetAll(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistPGRepository := NewArtistPGRepository(db)
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
