package repository

const (
	CreatePlaylistQuery = `INSERT INTO playlist (name, image, owner_id)
VALUES ($1, $2, $3)
RETURNING id, name, image, owner_id, is_private, created_at, updated_at`

	GetAllPlaylistsQuery = `SELECT id, name, image, owner_id, is_private, created_at, updated_at FROM playlist`

	GetPlaylistQuery = `SELECT id, name, image, owner_id, is_private, created_at, updated_at FROM playlist WHERE id = $1`

	GetLengthPlaylistsQuery = `SELECT COUNT(id) FROM playlist_track WHERE id = $1`

	GetUserPlaylistsQuery = "SELECT id, name, image, owner_id, is_private, created_at, updated_at FROM playlist WHERE owner_id = $1 ORDER BY created_at DESC"

	AddToPlaylistQuery = `INSERT INTO playlist_track (playlist_id, track_order_in_playlist, track_id)
VALUES ($1, $2, $3)
RETURNING id, playlist_id, track_order_in_playlist, track_id, created_at`

	RemoveFromPlaylistQuery = `DELETE FROM playlist_track WHERE playlist_id = $1 AND track_id = $2`

	DeletePlaylistQuery = `DELETE FROM playlist WHERE id = $1`

	addFavoritePlaylistQuery = `
    INSERT INTO favorite_playlist (user_id, playlist_id) 
    VALUES ($1, $2)
    ON CONFLICT (user_id, playlist_id) DO NOTHING`

	deleteFavoritePlaylistQuery = `
    DELETE FROM favorite_playlist
    WHERE user_id = $1 AND playlist_id = $2`

	isFavoritePlaylistQuery = `
    SELECT 1 
    FROM favorite_playlist
    WHERE user_id = $1 AND playlist_id = $2`

	getFavoriteQuery = `
    SELECT p.id AS id, name, image, owner_id, is_private, p.created_at AS created_at, p.updated_at AS updated_at
    FROM playlist AS p
      JOIN favorite_playlist AS fp
      ON p.id = fp.playlist_id
    WHERE fp.user_id = $1`

	getFavoriteCountQuery = `
    SELECT COUNT(*)
    FROM playlist AS p
      JOIN favorite_playlist AS fp
      ON p.id = fp.playlist_id
    WHERE fp.user_id = $1`

	getLikesCountQuery = `
    SELECT COUNT(*)
    FROM favorite_playlist WHERE playlist_id = $1`

	getPopularPlaylistsQuery = `SELECT 
    p.id, 
    p.name, 
    p.image, 
    p.owner_id, 
    p.is_private, 
    p.created_at, 
    p.updated_at
    FROM 
        playlist p
    LEFT JOIN 
        favorite_playlist fp ON p.id = fp.playlist_id
    GROUP BY 
        p.id
    ORDER BY 
        COUNT(fp.playlist_id) DESC
    LIMIT 50;
    `

	updatePlaylistQuery = `
		UPDATE "playlist"
		SET name = $1, image = $2, is_private = $3
		WHERE id = $4
		RETURNING id, name, image, owner_id, is_private
	`
)
