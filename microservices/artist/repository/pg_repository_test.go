package repository

import (
	"context"
	"testing"
	"time"

	uuid "github.com/google/uuid"

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

func TestArtistRepositoryFindByQuery(t *testing.T) {
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
	expectedArtists := []*models.Artist{&artists[0], &artists[1], &artists[2]}
	mock.ExpectQuery(findByQuery).WithArgs(findName).WillReturnRows(rows)

	foundArtists, err := artistPGRepository.FindByQuery(context.Background(), findName)
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

func TestArtistRepositoryAddFavoriteArtist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistRepository := NewArtistPGRepository(db)

	userID := uuid.New()
	artistID := uint64(12345)

	mock.ExpectExec(addFavoriteArtistQuery).WithArgs(userID, artistID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = artistRepository.AddFavoriteArtist(context.Background(), userID, artistID)

	require.NoError(t, err)
}

func TestArtistRepositoryDeleteFavoriteArtist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistRepository := NewArtistPGRepository(db)

	userID := uuid.New()
	artistID := uint64(12345)

	mock.ExpectExec(deleteFavoriteArtistQuery).WithArgs(userID, artistID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = artistRepository.DeleteFavoriteArtist(context.Background(), userID, artistID)

	require.NoError(t, err)
}

func TestArtistRepositoryIsFavoriteArtist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistRepository := NewArtistPGRepository(db)

	userID := uuid.New()
	artistID := uint64(12345)

	mock.ExpectQuery(isFavoriteArtistQuery).WithArgs(userID, artistID).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	exists, err := artistRepository.IsFavoriteArtist(context.Background(), userID, artistID)

	require.NoError(t, err)
	require.True(t, exists)
}

func TestArtistRepositoryGetFavoriteArtists(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistRepository := NewArtistPGRepository(db)

	artists := []models.Artist{
		{
			ID:        1,
			Name:      "Artist 1",
			Bio:       "Bio 1",
			Country:   "Country 1",
			Image:     "/imgs/artists/artist_1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Artist 2",
			Bio:       "Bio 2",
			Country:   "Country 2",
			Image:     "/imgs/artists/artist_2.jpg",
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

	userID := uuid.New()

	expectedArtists := []*models.Artist{&artists[0], &artists[1]}
	mock.ExpectQuery(getFavoriteQuery).WithArgs(userID).WillReturnRows(rows)

	foundArtists, err := artistRepository.GetFavoriteArtists(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, foundArtists)
	require.Equal(t, foundArtists, expectedArtists)
}

func TestArtistRepositoryGetFavoriteArtistsCount(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistRepository := NewArtistPGRepository(db)

	userID := uuid.New()
	expectedCount := uint64(3)

	mock.ExpectQuery(getFavoriteCountQuery).WithArgs(userID).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedCount))

	count, err := artistRepository.GetFavoriteArtistsCount(context.Background(), userID)
	require.NoError(t, err)
	require.Equal(t, expectedCount, count)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestAlbumRepositoryGetArtistLikesCount(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumRepository := NewArtistPGRepository(db)

	artistID := uint64(456)
	expectedLikesCount := uint64(20)

	mock.ExpectQuery(getLikesCountQuery).WithArgs(artistID).WillReturnRows(sqlmock.NewRows([]string{"likes_count"}).AddRow(expectedLikesCount))

	likesCount, err := albumRepository.GetArtistLikesCount(context.Background(), artistID)
	require.NoError(t, err)
	require.Equal(t, expectedLikesCount, likesCount)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestArtistRepositoryGetPopular(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistRepository := NewArtistPGRepository(db)

	artists := []models.Artist{
		{
			ID:        1,
			Name:      "Artist 1",
			Bio:       "Bio of Artist 1",
			Country:   "Country A",
			Image:     "/images/artists/artist_1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Artist 2",
			Bio:       "Bio of Artist 2",
			Country:   "Country B",
			Image:     "/images/artists/artist_2.jpg",
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

	mock.ExpectQuery(getPopularArtistsQuery).WillReturnRows(rows)

	ctx := context.Background()
	foundArtists, err := artistRepository.GetPopular(ctx)

	require.NoError(t, err)
	require.NotNil(t, foundArtists)
	require.Equal(t, len(artists), len(foundArtists))

	for i, artist := range artists {
		require.Equal(t, artist.ID, foundArtists[i].ID)
		require.Equal(t, artist.Name, foundArtists[i].Name)
		require.Equal(t, artist.Bio, foundArtists[i].Bio)
		require.Equal(t, artist.Country, foundArtists[i].Country)
		require.Equal(t, artist.Image, foundArtists[i].Image)
		require.WithinDuration(t, artist.CreatedAt, foundArtists[i].CreatedAt, time.Second)
		require.WithinDuration(t, artist.UpdatedAt, foundArtists[i].UpdatedAt, time.Second)
	}

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestArtistRepositoryGetPopular_Error(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	artistRepository := NewArtistPGRepository(db)

	mock.ExpectQuery(getPopularArtistsQuery).WillReturnError(sqlmock.ErrCancelled)

	ctx := context.Background()
	foundArtists, err := artistRepository.GetPopular(ctx)

	require.Error(t, err)
	require.Nil(t, foundArtists)

	require.NoError(t, mock.ExpectationsWereMet())
}
