package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPlaylistRepositoryCreatePlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)
	mockPlaylist := &models.Playlist{
		ID:        1,
		Name:      "gym training playlist",
		Image:     "/images/playlists/playlist_1.jpg",
		OwnerID:   uuid.New(),
		IsPrivate: false,
		CreatedAt: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
	}

	columns := []string{"id", "name", "image", "owner_id", "is_private", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		mockPlaylist.ID,
		mockPlaylist.Name,
		mockPlaylist.Image,
		mockPlaylist.OwnerID,
		mockPlaylist.IsPrivate,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(CreatePlaylistQuery).WithArgs(
		mockPlaylist.Name,
		mockPlaylist.Image,
		mockPlaylist.OwnerID,
	).WillReturnRows(rows)

	createdPlaylist, err := playlistRepository.CreatePlaylist(context.Background(), mockPlaylist)
	require.NoError(t, err)
	require.NotNil(t, createdPlaylist)
}

func TestPlaylistRepositoryGetAllPlaylists(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)
	columns := []string{"id", "name", "image", "owner_id", "is_private", "created_at", "updated_at"}
	mockPlaylists := sqlmock.NewRows(columns).
		AddRow(1, "Playlist 1", "/images/playlists/1.jpg", uuid.New(), false, time.Now(), time.Now()).
		AddRow(2, "Playlist 2", "/images/playlists/2.jpg", uuid.New(), false, time.Now(), time.Now())

	mock.ExpectQuery(GetAllPlaylistsQuery).WillReturnRows(mockPlaylists)

	playlists, err := playlistRepository.GetAllPlaylists(context.Background())
	require.NoError(t, err)
	require.NotNil(t, playlists)
	require.Len(t, playlists, 2)
}

func TestPlaylistRepositoryGetPlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)
	mockPlaylistID := uint64(1)
	mockPlaylist := &models.Playlist{
		ID:        mockPlaylistID,
		Name:      "My Playlist",
		Image:     "/images/playlists/playlist_1.jpg",
		OwnerID:   uuid.New(),
		IsPrivate: false,
	}

	columns := []string{"id", "name", "image", "owner_id", "is_private", "created_at", "updated_at"}
	row := sqlmock.NewRows(columns).AddRow(
		mockPlaylist.ID,
		mockPlaylist.Name,
		mockPlaylist.Image,
		mockPlaylist.OwnerID,
		mockPlaylist.IsPrivate,
		time.Now(),
		time.Now(),
	)

	mock.ExpectPrepare(GetPlaylistQuery).
		ExpectQuery().
		WithArgs(mockPlaylistID).
		WillReturnRows(row)

	playlist, err := playlistRepository.GetPlaylist(context.Background(), mockPlaylistID)
	require.NoError(t, err)
	require.NotNil(t, playlist)
	require.Equal(t, mockPlaylistID, playlist.ID)
}

func TestPlaylistRepositoryGetLengthPlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)
	mockPlaylistID := uint64(1)
	mockLength := uint64(10)

	mock.ExpectPrepare(GetLengthPlaylistsQuery).
		ExpectQuery().
		WithArgs(mockPlaylistID).
		WillReturnRows(sqlmock.NewRows([]string{"length"}).AddRow(mockLength))

	length, err := playlistRepository.GetLengthPlaylist(context.Background(), mockPlaylistID)
	require.NoError(t, err)
	require.Equal(t, mockLength, length)
}

func TestPlaylistRepositoryGetUserPlaylists(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)
	mockUserID := uuid.New()
	columns := []string{"id", "name", "image", "owner_id", "is_private", "created_at", "updated_at"}
	mockRows := sqlmock.NewRows(columns).
		AddRow(1, "Playlist 1", "/images/playlists/1.jpg", mockUserID, false, time.Now(), time.Now()).
		AddRow(2, "Playlist 2", "/images/playlists/2.jpg", mockUserID, false, time.Now(), time.Now())

	mock.ExpectPrepare(GetUserPlaylistsQuery).
		ExpectQuery().
		WithArgs(mockUserID).
		WillReturnRows(mockRows)

	playlists, err := playlistRepository.GetUserPlaylists(context.Background(), mockUserID)
	require.NoError(t, err)
	require.Len(t, playlists, 2)
	require.Equal(t, mockUserID, playlists[0].OwnerID)
}

func TestPlaylistRepositoryAddToPlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)
	mockPlaylistID := uint64(1)
	mockTrackID := uint64(42)
	mockTrackOrder := uint64(1)

	columns := []string{"id", "playlist_id", "track_order_in_playlist", "track_id", "created_at"}
	row := sqlmock.NewRows(columns).AddRow(1, mockPlaylistID, mockTrackOrder, mockTrackID, time.Now())

	mock.ExpectPrepare(AddToPlaylistQuery).
		ExpectQuery().
		WithArgs(mockPlaylistID, mockTrackOrder, mockTrackID).
		WillReturnRows(row)

	track, err := playlistRepository.AddToPlaylist(context.Background(), mockPlaylistID, mockTrackOrder, mockTrackID)
	require.NoError(t, err)
	require.NotNil(t, track)
	require.Equal(t, mockPlaylistID, track.PlaylistID)
	require.Equal(t, mockTrackID, track.TrackID)
}

func TestPlaylistRepositoryRemoveFromPlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)
	mockPlaylistID := uint64(1)
	mockTrackID := uint64(42)

	mock.ExpectPrepare(RemoveFromPlaylistQuery).
		ExpectExec().
		WithArgs(mockPlaylistID, mockTrackID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := playlistRepository.RemoveFromPlaylist(context.Background(), mockPlaylistID, mockTrackID)
	require.NoError(t, err)
	rowsAffected, _ := res.RowsAffected()
	require.Equal(t, int64(1), rowsAffected)
}

func TestPlaylistRepositoryDeletePlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)
	mockPlaylistID := uint64(1)

	mock.ExpectPrepare(DeletePlaylistQuery).
		ExpectExec().
		WithArgs(mockPlaylistID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := playlistRepository.DeletePlaylist(context.Background(), mockPlaylistID)
	require.NoError(t, err)
	rowsAffected, _ := res.RowsAffected()
	require.Equal(t, int64(1), rowsAffected)
}
