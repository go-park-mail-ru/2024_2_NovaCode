package repository

import (
	"context"
	"testing"
	"time"

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
		TrackCount:  12,
		ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
		Image:       "/imgs/albums/album_1.jpg",
		ArtistID:    1,
	}

	columns := []string{"id", "name", "track_count", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockAlbum.ID,
		mockAlbum.Name,
		mockAlbum.TrackCount,
		mockAlbum.ReleaseDate,
		mockAlbum.Image,
		mockAlbum.ArtistID,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(createAlbumQuery).WithArgs(
		mockAlbum.Name,
		mockAlbum.TrackCount,
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
		TrackCount:  12,
		ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
		Image:       "/imgs/albums/album_1.jpg",
		ArtistID:    1,
	}

	columns := []string{"id", "name", "track_count", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockAlbum.ID,
		mockAlbum.Name,
		mockAlbum.TrackCount,
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

func TestAlbumRepositoryFindByName(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	albumPGRepository := NewAlbumPGRepository(db)
	albums := []models.Album{
		{
			ID:          1,
			Name:        "Album for test 1",
			TrackCount:  12,
			ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_1.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Album for test 2",
			TrackCount:  9,
			ReleaseDate: time.Date(2021, 02, 3, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_2.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "Another album",
			TrackCount:  4,
			ReleaseDate: time.Date(2019, 01, 5, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_3.jpg",
			ArtistID:    3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "track_count", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, album := range albums {
		rows.AddRow(
			album.ID,
			album.Name,
			album.TrackCount,
			album.ReleaseDate,
			album.Image,
			album.ArtistID,
			album.CreatedAt,
			album.UpdatedAt,
		)
	}

	findName := "test"
	expectedAlbums := []*models.Album{&albums[0], &albums[1]}
	mock.ExpectQuery(findByNameQuery).WithArgs(findName).WillReturnRows(rows)

	foundAlbums, err := albumPGRepository.FindByName(context.Background(), findName)
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
			TrackCount:  12,
			ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_1.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Album for test 2",
			TrackCount:  9,
			ReleaseDate: time.Date(2021, 02, 3, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_2.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "Another album",
			TrackCount:  4,
			ReleaseDate: time.Date(2019, 01, 5, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_3.jpg",
			ArtistID:    3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "track_count", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, album := range albums {
		rows.AddRow(
			album.ID,
			album.Name,
			album.TrackCount,
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
			TrackCount:  12,
			ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_1.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Album for test 2",
			TrackCount:  9,
			ReleaseDate: time.Date(2021, 02, 3, 0, 0, 0, 0, time.UTC),
			Image:       "/imgs/albums/album_2.jpg",
			ArtistID:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	columns := []string{"id", "name", "track_count", "release", "image", "artist_id", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, album := range albums {
		rows.AddRow(
			album.ID,
			album.Name,
			album.TrackCount,
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
