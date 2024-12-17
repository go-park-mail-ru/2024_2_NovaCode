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

	mock.ExpectQuery(GetPlaylistQuery).WithArgs(mockPlaylistID).WillReturnRows(row)

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

	mock.ExpectQuery(GetLengthPlaylistsQuery).WithArgs(mockPlaylistID).
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

	mock.ExpectQuery(GetUserPlaylistsQuery).WithArgs(mockUserID).WillReturnRows(mockRows)

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

	mock.ExpectQuery(AddToPlaylistQuery).WithArgs(mockPlaylistID, mockTrackOrder, mockTrackID).WillReturnRows(row)

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

	mock.ExpectExec(RemoveFromPlaylistQuery).WithArgs(mockPlaylistID, mockTrackID).
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

	mock.ExpectExec(DeletePlaylistQuery).WithArgs(mockPlaylistID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := playlistRepository.DeletePlaylist(context.Background(), mockPlaylistID)
	require.NoError(t, err)
	rowsAffected, _ := res.RowsAffected()
	require.Equal(t, int64(1), rowsAffected)
}

func TestPlaylistRepositoryAddFavoritePlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)

	userID := uuid.New()
	playlistID := uint64(12345)

	mock.ExpectExec(addFavoritePlaylistQuery).WithArgs(userID, playlistID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = playlistRepository.AddFavoritePlaylist(context.Background(), userID, playlistID)

	require.NoError(t, err)
}

func TestPlaylistRepositoryDeleteFavoritePlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)

	userID := uuid.New()
	playlistID := uint64(12345)

	mock.ExpectExec(deleteFavoritePlaylistQuery).WithArgs(userID, playlistID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = playlistRepository.DeleteFavoritePlaylist(context.Background(), userID, playlistID)

	require.NoError(t, err)
}

func TestPlaylistRepositoryIsFavoritePlaylist(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)

	userID := uuid.New()
	playlistID := uint64(12345)

	mock.ExpectQuery(isFavoritePlaylistQuery).WithArgs(userID, playlistID).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	exists, err := playlistRepository.IsFavoritePlaylist(context.Background(), userID, playlistID)

	require.NoError(t, err)
	require.True(t, exists)
}

func TestPlaylistRepositoryGetFavoritePlaylists(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	playlistRepository := NewPlaylistRepository(db)

	playlists := []models.Playlist{
		{
			ID:        1,
			Name:      "Playlist 1",
			Image:     "/imgs/playlists/playlist_1.jpg",
			OwnerID:   uuid.New(),
			IsPrivate: false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Playlist 2",
			Image:     "/imgs/playlists/playlist_2.jpg",
			OwnerID:   uuid.New(),
			IsPrivate: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	columns := []string{"id", "name", "image", "owner_id", "is_private", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns)
	for _, playlist := range playlists {
		rows.AddRow(
			playlist.ID,
			playlist.Name,
			playlist.Image,
			playlist.OwnerID,
			playlist.IsPrivate,
			playlist.CreatedAt,
			playlist.UpdatedAt,
		)
	}

	userID := uuid.New()

	expectedPlaylists := []*models.Playlist{&playlists[0], &playlists[1]}
	mock.ExpectQuery(getFavoriteQuery).WithArgs(userID).WillReturnRows(rows)

	foundPlaylists, err := playlistRepository.GetFavoritePlaylists(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, foundPlaylists)
	require.Equal(t, foundPlaylists, expectedPlaylists)
}
