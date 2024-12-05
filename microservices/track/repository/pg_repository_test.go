package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTrackRepositoryCreate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)
	mockTrack := &models.Track{
		ID:           1,
		Name:         "ok im cool",
		Duration:     167,
		FilePath:     "/songs/track_1.mp4",
		Image:        "/imgs/tracks/track_1.jpg",
		ArtistID:     1,
		AlbumID:      1,
		OrderInAlbum: 1,
		ReleaseDate:  time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "track_order_in_album", "release", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockTrack.ID,
		mockTrack.Name,
		mockTrack.Duration,
		mockTrack.FilePath,
		mockTrack.Image,
		mockTrack.ArtistID,
		mockTrack.AlbumID,
		mockTrack.OrderInAlbum,
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
		mockTrack.OrderInAlbum,
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

	trackPGRepository := NewTrackPGRepository(db)
	mockTrack := &models.Track{
		ID:           1,
		Name:         "ok im cool",
		Duration:     167,
		FilePath:     "/songs/track_1.mp4",
		Image:        "/imgs/tracks/track_1.jpg",
		ArtistID:     1,
		AlbumID:      1,
		OrderInAlbum: 1,
		ReleaseDate:  time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "track_order_in_album", "release", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockTrack.ID,
		mockTrack.Name,
		mockTrack.Duration,
		mockTrack.FilePath,
		mockTrack.Image,
		mockTrack.ArtistID,
		mockTrack.AlbumID,
		mockTrack.OrderInAlbum,
		mockTrack.ReleaseDate,
		time.Now(),
		time.Now(),
	)

	stmt := mock.ExpectPrepare(findByIDQuery)
	stmt.ExpectQuery().WithArgs(mockTrack.ID).WillReturnRows(rows)

	foundTrack, err := trackPGRepository.FindById(context.Background(), mockTrack.ID)
	require.NoError(t, err)
	require.NotNil(t, foundTrack)
	require.Equal(t, mockTrack.ID, foundTrack.ID)
}

func TestTrackRepositoryFindByQuery(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)
	tracks := []models.Track{
		{
			ID:           1,
			Name:         "test song 1",
			Duration:     123,
			FilePath:     "/songs/track_1.mp4",
			Image:        "/imgs/tracks/track_1.jpg",
			ArtistID:     1,
			AlbumID:      1,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "another song",
			Duration:     93,
			FilePath:     "/songs/track_2.mp4",
			Image:        "/imgs/tracks/track_2.jpg",
			ArtistID:     2,
			AlbumID:      2,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           3,
			Name:         "song test",
			Duration:     99,
			FilePath:     "/songs/track_3.mp4",
			Image:        "/imgs/tracks/track_3.jpg",
			ArtistID:     3,
			AlbumID:      3,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "track_order_in_album", "release", "created_at", "updated_at"}
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
			track.OrderInAlbum,
			track.ReleaseDate,
			track.CreatedAt,
			track.UpdatedAt,
		)
	}

	findName := "test"
	expectedTracks := []*models.Track{&tracks[0], &tracks[1], &tracks[2]}
	mock.ExpectPrepare(findByQuery).ExpectQuery().WithArgs(utils.MakeSearchQuery(findName)).WillReturnRows(rows)

	foundTracks, err := trackPGRepository.FindByQuery(context.Background(), findName)
	require.NoError(t, err)
	require.NotNil(t, foundTracks)
	require.Equal(t, foundTracks, expectedTracks)
}

func TestTrackRepositoryGetAll(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)
	tracks := []models.Track{
		{
			ID:           1,
			Name:         "test song 1",
			Duration:     123,
			FilePath:     "/songs/track_1.mp4",
			Image:        "/imgs/tracks/track_1.jpg",
			ArtistID:     1,
			AlbumID:      1,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "another song",
			Duration:     93,
			FilePath:     "/songs/track_2.mp4",
			Image:        "/imgs/tracks/track_2.jpg",
			ArtistID:     2,
			AlbumID:      2,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           3,
			Name:         "song test",
			Duration:     99,
			FilePath:     "/songs/track_3.mp4",
			Image:        "/imgs/tracks/track_3.jpg",
			ArtistID:     3,
			AlbumID:      3,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "track_order_in_album", "release", "created_at", "updated_at"}
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
			track.OrderInAlbum,
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

	trackPGRepository := NewTrackPGRepository(db)

	tracks := []models.Track{
		{
			ID:           1,
			Name:         "test song 1",
			Duration:     123,
			FilePath:     "/songs/track_1.mp4",
			Image:        "/imgs/tracks/track_1.jpg",
			ArtistID:     1,
			AlbumID:      1,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "another song",
			Duration:     93,
			FilePath:     "/songs/track_2.mp4",
			Image:        "/imgs/tracks/track_2.jpg",
			ArtistID:     1,
			AlbumID:      2,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "track_order_in_album", "release", "created_at", "updated_at"}
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
			track.OrderInAlbum,
			track.ReleaseDate,
			track.CreatedAt,
			track.UpdatedAt,
		)
	}

	expectedTracks := []*models.Track{&tracks[0], &tracks[1]}
	stmt := mock.ExpectPrepare(getByArtistIDQuery)
	stmt.ExpectQuery().WithArgs(uint64(1)).WillReturnRows(rows)

	foundTracks, err := trackPGRepository.GetAllByArtistID(context.Background(), uint64(1))
	require.NoError(t, err)
	require.NotNil(t, foundTracks)
	require.Equal(t, foundTracks, expectedTracks)
}

func TestTrackRepositoryGetAllByAlbumID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)

	tracks := []models.Track{
		{
			ID:           1,
			Name:         "test song 1",
			Duration:     123,
			FilePath:     "/songs/track_1.mp4",
			Image:        "/imgs/tracks/track_1.jpg",
			ArtistID:     1,
			AlbumID:      1,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "another song",
			Duration:     93,
			FilePath:     "/songs/track_2.mp4",
			Image:        "/imgs/tracks/track_2.jpg",
			ArtistID:     1,
			AlbumID:      1,
			OrderInAlbum: 2,
			ReleaseDate:  time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "track_order_in_album", "release", "created_at", "updated_at"}
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
			track.OrderInAlbum,
			track.ReleaseDate,
			track.CreatedAt,
			track.UpdatedAt,
		)
	}

	expectedTracks := []*models.Track{&tracks[0], &tracks[1]}
	mock.ExpectPrepare(getByAlbumIDQuery).ExpectQuery().WithArgs(1).WillReturnRows(rows)

	foundTracks, err := trackPGRepository.GetAllByAlbumID(context.Background(), uint64(1))
	require.NoError(t, err)
	require.NotNil(t, foundTracks)
	require.Equal(t, foundTracks, expectedTracks)
}

func TestTrackRepositoryAddFavoriteTrack(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)

	userID := uuid.New()
	trackID := uint64(12345)

	mock.ExpectPrepare(addFavoriteTrackQuery).ExpectExec().WithArgs(userID, trackID).WillReturnResult(sqlmock.NewResult(0, 0))
	err = trackPGRepository.AddFavoriteTrack(context.Background(), userID, trackID)

	require.NoError(t, err)
}

func TestTrackRepositoryDeleteFavoriteTrack(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)

	userID := uuid.New()
	trackID := uint64(12345)

	mock.ExpectPrepare(deleteFavoriteTrackQuery).ExpectExec().WithArgs(userID, trackID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = trackPGRepository.DeleteFavoriteTrack(context.Background(), userID, trackID)

	require.NoError(t, err)
}

func TestTrackRepositoryIsFavoriteTrack(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)

	userID := uuid.New()
	trackID := uint64(12345)

	mock.ExpectPrepare(isFavoriteTrackQuery).ExpectQuery().WithArgs(userID, trackID).WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	exists, err := trackPGRepository.IsFavoriteTrack(context.Background(), userID, trackID)

	require.NoError(t, err)
	require.True(t, exists)
}

func TestTrackRepositoryGetFavoriteTracks(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)

	tracks := []models.Track{
		{
			ID:           1,
			Name:         "test song 1",
			Duration:     123,
			FilePath:     "/songs/track_1.mp4",
			Image:        "/imgs/tracks/track_1.jpg",
			ArtistID:     1,
			AlbumID:      1,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "another song",
			Duration:     93,
			FilePath:     "/songs/track_2.mp4",
			Image:        "/imgs/tracks/track_2.jpg",
			ArtistID:     1,
			AlbumID:      2,
			OrderInAlbum: 1,
			ReleaseDate:  time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	columns := []string{"id", "name", "duration", "filepath", "image", "artist_id", "album_id", "track_order_in_album", "release", "created_at", "updated_at"}
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
			track.OrderInAlbum,
			track.ReleaseDate,
			track.CreatedAt,
			track.UpdatedAt,
		)
	}

	userID := uuid.New()

	expectedTracks := []*models.Track{&tracks[0], &tracks[1]}
	mock.ExpectPrepare(getFavoriteQuery).ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	foundTracks, err := trackPGRepository.GetFavoriteTracks(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, foundTracks)
	require.Equal(t, foundTracks, expectedTracks)
}

func TestTrackRepositoryGetTracksFromPlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	trackPGRepository := NewTrackPGRepository(db)

	playlistTracks := []*models.PlaylistTrack{
		{
			ID:                   1,
			PlaylistID:           1,
			TrackOrderInPlaylist: 1,
			TrackID:              1,
			CreatedAt:            time.Now(),
		},
		{
			ID:                   2,
			PlaylistID:           1,
			TrackOrderInPlaylist: 2,
			TrackID:              2,
			CreatedAt:            time.Now(),
		},
	}

	columns := []string{"id", "playlist_id", "track_order_in_playlist", "track_id", "created_at"}
	rows := sqlmock.NewRows(columns)
	for _, track := range playlistTracks {
		rows.AddRow(
			track.ID,
			track.PlaylistID,
			track.TrackOrderInPlaylist,
			track.TrackID,
			track.CreatedAt,
		)
	}

	mock.ExpectPrepare(GetTracksFromPlaylistQuery).ExpectQuery().WithArgs(uint64(1)).WillReturnRows(rows)

	foundPlaylistTracks, err := trackPGRepository.GetTracksFromPlaylist(context.Background(), uint64(1))
	require.NoError(t, err)
	require.NotNil(t, foundPlaylistTracks)
	require.Equal(t, foundPlaylistTracks, playlistTracks)
}
