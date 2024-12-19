package repository

const (
	createTrackQuery = `INSERT INTO track (name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date, created_at, updated_at`

	findByIDQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date, created_at, updated_at FROM track WHERE id = $1`

	getAllQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date, created_at, updated_at FROM track`

	findByQuery = `
    SELECT id, name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date, created_at, updated_at
    FROM "track"
    WHERE fts @@ to_tsquery('english', $1 || ':*') 
        OR fts @@ to_tsquery('russian_hunspell', $1 || ':*')`

	getByArtistIDQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date, created_at, updated_at FROM track WHERE artist_id = $1`

	getByAlbumIDQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date, created_at, updated_at FROM track WHERE artist_id = $1 ORDER BY track_order_in_album ASC`

	addFavoriteTrackQuery = `
    INSERT INTO favorite_track (user_id, track_id) 
    VALUES ($1, $2)
    ON CONFLICT (user_id, track_id) DO NOTHING`

	deleteFavoriteTrackQuery = `
    DELETE FROM favorite_track
    WHERE user_id = $1 AND track_id = $2`

	isFavoriteTrackQuery = `
    SELECT 1 
    FROM favorite_track 
    WHERE user_id = $1 AND track_id = $2`

	getFavoriteQuery = `
    SELECT t.id AS id, name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date, t.created_at AS created_at, t.updated_at AS updated_at 
    FROM track AS t 
      JOIN favorite_track AS ft 
      ON t.id = ft.track_id
    WHERE ft.user_id = $1`

	getFavoriteCountQuery = `
    SELECT COUNT(*)
    FROM track AS a 
      JOIN favorite_track AS fa
      ON a.id = fa.track_id
    WHERE fa.user_id = $1`

	getTracksFromPlaylistQuery = `SELECT id, playlist_id, track_order_in_playlist, track_id, created_at FROM playlist_track WHERE playlist_id = $1 ORDER BY created_at DESC`

	getPopularTracksQuery = `SELECT 
    t.id, 
    t.name, 
    t.duration, 
    t.filepath, 
    t.image, 
    t.artist_id, 
    t.album_id, 
    t.track_order_in_album, 
    t.release_date, 
    t.created_at, 
    t.updated_at
    FROM 
        track t
    LEFT JOIN 
        favorite_track ft ON t.id = ft.track_id
    GROUP BY 
        t.id
    ORDER BY 
        COUNT(ft.track_id) DESC
    LIMIT 50;
  `

	getTracksByGenre = `
    SELECT t.id AS id, name, duration, filepath, image, artist_id, album_id, track_order_in_album, release_date, t.created_at AS created_at, t.updated_at AS updated_at 
    FROM track AS t
      JOIN genre_track gt ON t.id = gt.track_id
      JOIN genre g ON gt.genre_id = g.id
    WHERE g.name = $1`
)
