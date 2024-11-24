package repository

const (
	CreatePlaylistQuery = `INSERT INTO playlist (name, image, owner_id)
VALUES ($1, $2, $3)
RETURNING id, name, image, owner_id, is_private, created_at, updated_at`

	GetAllPlaylistsQuery = `SELECT id, name, image, owner_id, is_private, created_at, updated_at FROM playlist`

	GetPlaylistQuery = `SELECT id, name, image, owner_id, is_private, created_at, updated_at FROM playlist WHERE id = $1`

	GetLengthPlaylistsQuery = `SELECT COUNT(id) FROM playlist_track WHERE id = $1`

	GetTracksFromPlaylistQuery = `SELECT id, playlist_id, track_order_in_playlist, track_id, created_at FROM playlist_track WHERE playlist_id = $1 ORDER BY created_at DESC`

	GetUserPlaylistsQuery = "SELECT id, name, image, owner_id, is_private, created_at, updated_at FROM playlist WHERE owner_id = $1 ORDER BY created_at DESC"

	AddToPlaylistQuery = `INSERT INTO playlist_track (playlist_id, track_order_in_playlist, track_id)
VALUES ($1, $2, $3)
RETURNING id, playlist_id, track_order_in_playlist, track_id, created_at`

	RemoveFromPlaylistQuery = `DELETE FROM playlist_track WHERE playlist_id = $1 AND track_id = $2`

	DeletePlaylistQuery = `DELETE FROM playlist WHERE id = $1`
)
