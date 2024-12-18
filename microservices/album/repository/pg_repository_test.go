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

func TestAlbumRepositoryCreate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumPGRepository := NewAlbumPGRepository(db)
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

	albumPGRepository := NewAlbumPGRepository(db)
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

	mock.ExpectQuery(findByIDQuery).WithArgs(mockAlbum.ID).WillReturnRows(rows)

	foundAlbum, err := albumPGRepository.FindById(context.Background(), mockAlbum.ID)
	require.NoError(t, err)
	require.NotNil(t, foundAlbum)
	require.Equal(t, foundAlbum.ID, foundAlbum.ID)
}

func TestAlbumRepositoryFindByQuery(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumPGRepository := NewAlbumPGRepository(db)
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
	mock.ExpectQuery(findByQuery).WithArgs(findName).WillReturnRows(rows)

	foundAlbums, err := albumPGRepository.FindByQuery(context.Background(), findName)
	require.NoError(t, err)
	require.NotNil(t, foundAlbums)
	require.Equal(t, foundAlbums, expectedAlbums)
}

func TestAlbumRepositoryGetAll(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumPGRepository := NewAlbumPGRepository(db)
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

	albumPGRepository := NewAlbumPGRepository(db)
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
	mock.ExpectQuery(getByArtistIDQuery).WithArgs(uint64(1)).WillReturnRows(rows)

	foundAlbums, err := albumPGRepository.GetAllByArtistID(context.Background(), uint64(1))
	require.NoError(t, err)
	require.NotNil(t, foundAlbums)
	require.Equal(t, foundAlbums, expectedAlbums)
}

func TestAlbumRepositoryAddFavoriteAlbum(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumRepository := NewAlbumPGRepository(db)

	userID := uuid.New()
	albumID := uint64(12345)

	mock.ExpectExec(addFavoriteAlbumQuery).WithArgs(userID, albumID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = albumRepository.AddFavoriteAlbum(context.Background(), userID, albumID)

	require.NoError(t, err)
}

func TestAlbumRepositoryDeleteFavoriteAlbum(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumRepository := NewAlbumPGRepository(db)

	userID := uuid.New()
	albumID := uint64(12345)

	mock.ExpectExec(deleteFavoriteAlbumQuery).WithArgs(userID, albumID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = albumRepository.DeleteFavoriteAlbum(context.Background(), userID, albumID)

	require.NoError(t, err)
}

func TestAlbumRepositoryIsFavoriteAlbum(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumRepository := NewAlbumPGRepository(db)

	userID := uuid.New()
	albumID := uint64(12345)

	mock.ExpectQuery(isFavoriteAlbumQuery).WithArgs(userID, albumID).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	exists, err := albumRepository.IsFavoriteAlbum(context.Background(), userID, albumID)

	require.NoError(t, err)
	require.True(t, exists)
}

func TestAlbumRepositoryGetFavoriteAlbums(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumRepository := NewAlbumPGRepository(db)

	albums := []models.Album{
		{
			ID:          1,
			Name:        "Album 1",
			ReleaseDate: time.Now(),
			Image:       "/imgs/albums/album_1.jpg",
			ArtistID:    101,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Album 2",
			ReleaseDate: time.Now(),
			Image:       "/imgs/albums/album_2.jpg",
			ArtistID:    102,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "release_date", "image", "artist_id", "created_at", "updated_at"}
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

	userID := uuid.New()

	expectedAlbums := []*models.Album{&albums[0], &albums[1]}
	mock.ExpectQuery(getFavoriteQuery).WithArgs(userID).WillReturnRows(rows)

	foundAlbums, err := albumRepository.GetFavoriteAlbums(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, foundAlbums)
	require.Equal(t, foundAlbums, expectedAlbums)
}

func TestAlbumRepositoryGetFavoriteAlbumsCount(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumRepository := NewAlbumPGRepository(db)

	userID := uuid.New()
	expectedCount := uint64(2)

	mock.ExpectQuery(getFavoriteCountQuery).WithArgs(userID).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedCount))

	count, err := albumRepository.GetFavoriteAlbumsCount(context.Background(), userID)
	require.NoError(t, err)
	require.Equal(t, expectedCount, count)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestAlbumRepositoryGetAlbumLikesCount(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumRepository := NewAlbumPGRepository(db)

	albumID := uint64(123)
	expectedLikesCount := uint64(10)

	mock.ExpectQuery(getLikesCountQuery).WithArgs(albumID).WillReturnRows(sqlmock.NewRows([]string{"likes_count"}).AddRow(expectedLikesCount))

	likesCount, err := albumRepository.GetAlbumLikesCount(context.Background(), albumID)
	require.NoError(t, err)
	require.Equal(t, expectedLikesCount, likesCount)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
