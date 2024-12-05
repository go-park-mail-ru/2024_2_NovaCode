package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/stretchr/testify/require"
)

func TestGenreRepositoryCreate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	genrePGRepository := NewGenrePGRepository(db)
	mockGenre := &models.Genre{
		ID:        1,
		Name:      "Rock",
		RusName:   "Рок",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	columns := []string{"id", "name", "rus_name", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockGenre.ID,
		mockGenre.Name,
		mockGenre.RusName,
		mockGenre.CreatedAt,
		mockGenre.UpdatedAt,
	)

	mock.ExpectQuery(createGenreQuery).WithArgs(
		mockGenre.Name,
		mockGenre.RusName,
	).WillReturnRows(rows)

	createdGenre, err := genrePGRepository.Create(context.Background(), mockGenre)
	require.NoError(t, err)
	require.NotNil(t, createdGenre)
	require.Equal(t, mockGenre.ID, createdGenre.ID)
	require.Equal(t, mockGenre.Name, createdGenre.Name)
	require.Equal(t, mockGenre.RusName, createdGenre.RusName)
	require.Equal(t, mockGenre.CreatedAt, createdGenre.CreatedAt)
	require.Equal(t, mockGenre.UpdatedAt, createdGenre.UpdatedAt)
}

func TestGenreRepositoryFindById(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	genrePGRepository := NewGenrePGRepository(db)
	mockGenre := &models.Genre{
		ID:        1,
		Name:      "Rock",
		RusName:   "Рок",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	columns := []string{"id", "name", "rus_name", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockGenre.ID,
		mockGenre.Name,
		mockGenre.RusName,
		mockGenre.CreatedAt,
		mockGenre.UpdatedAt,
	)

	mock.ExpectPrepare(findByIDQuery).
		ExpectQuery().
		WithArgs(mockGenre.ID).
		WillReturnRows(rows)

	foundGenre, err := genrePGRepository.FindById(context.Background(), mockGenre.ID)
	require.NoError(t, err)
	require.NotNil(t, foundGenre)
	require.Equal(t, mockGenre.ID, foundGenre.ID)
	require.Equal(t, mockGenre.Name, foundGenre.Name)
	require.Equal(t, mockGenre.RusName, foundGenre.RusName)
	require.Equal(t, mockGenre.CreatedAt, foundGenre.CreatedAt)
	require.Equal(t, mockGenre.UpdatedAt, foundGenre.UpdatedAt)
}

func TestGenreRepositoryGetAll(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	genrePGRepository := NewGenrePGRepository(db)
	genres := []models.Genre{
		{
			ID:        1,
			Name:      "Rock",
			RusName:   "Рок",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Pop",
			RusName:   "Поп",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "Hip-Hop",
			RusName:   "Хип-Хоп",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	columns := []string{"id", "name", "rus_name", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, genre := range genres {
		rows.AddRow(
			genre.ID,
			genre.Name,
			genre.RusName,
			genre.CreatedAt,
			genre.UpdatedAt,
		)
	}

	expectedGenres := []*models.Genre{&genres[0], &genres[1], &genres[2]}
	mock.ExpectQuery(getAllQuery).WillReturnRows(rows)

	foundGenres, err := genrePGRepository.GetAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, foundGenres)
	require.Equal(t, foundGenres, expectedGenres)
}

func TestGenreRepositoryGetByArtistID(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	genrePGRepository := NewGenrePGRepository(db)
	genres := []models.Genre{
		{
			ID:        1,
			Name:      "Rock",
			RusName:   "Рок",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Pop",
			RusName:   "Поп",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "Hip-Hop",
			RusName:   "Хип-Хоп",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	columns := []string{"id", "name", "rus_name", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, genre := range genres {
		rows.AddRow(
			genre.ID,
			genre.Name,
			genre.RusName,
			genre.CreatedAt,
			genre.UpdatedAt,
		)
	}

	expectedGenres := []*models.Genre{&genres[0], &genres[1], &genres[2]}
	mock.ExpectPrepare(getByArtistIDQuery).
		ExpectQuery().
		WithArgs(uint64(1)).
		WillReturnRows(rows)

	foundGenres, err := genrePGRepository.GetAllByArtistID(context.Background(), uint64(1))
	require.NoError(t, err)
	require.NotNil(t, foundGenres)
	require.Equal(t, foundGenres, expectedGenres)
}

func TestGenreRepositoryGetAllByTrackID(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	genrePGRepository := NewGenrePGRepository(db)
	genres := []models.Genre{
		{
			ID:        1,
			Name:      "Rock",
			RusName:   "Рок",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Pop",
			RusName:   "Поп",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "Hip-Hop",
			RusName:   "Хип-Хоп",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	columns := []string{"id", "name", "rus_name", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, genre := range genres {
		rows.AddRow(
			genre.ID,
			genre.Name,
			genre.RusName,
			genre.CreatedAt,
			genre.UpdatedAt,
		)
	}

	expectedGenres := []*models.Genre{&genres[0], &genres[1], &genres[2]}
	mock.ExpectPrepare(getByTrackIDQuery).
		ExpectQuery().
		WithArgs(uint64(1)).
		WillReturnRows(rows)

	foundGenres, err := genrePGRepository.GetAllByTrackID(context.Background(), uint64(1))
	require.NoError(t, err)
	require.NotNil(t, foundGenres)
	require.Equal(t, expectedGenres, foundGenres)
}
