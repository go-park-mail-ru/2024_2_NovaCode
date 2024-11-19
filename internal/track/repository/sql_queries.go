package repository

const (
	createTrackQuery = `INSERT INTO track (name, duration, filepath, image, artist_id, album_id, release_date) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, name, duration, filepath, image, artist_id, album_id, release_date, created_at, updated_at`

	findByIDQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, release_date, created_at, updated_at FROM track WHERE id = $1`

	getAllQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, release_date, created_at, updated_at FROM track`

	findByNameQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, release_date, created_at, updated_at FROM track WHERE name = $1`

	getByArtistIDQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, release_date, created_at, updated_at FROM track WHERE artist_id = $1`

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
    SELECT t.id AS id, name, duration, filepath, image, artist_id, album_id, release_date, t.created_at AS created_at, t.updated_at AS updated_at 
    FROM track AS t 
      JOIN favorite_track AS ft 
      ON t.id = ft.track_id
    WHERE ft.user_id = $1`
)
