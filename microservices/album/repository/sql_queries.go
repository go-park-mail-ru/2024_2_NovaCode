package repository

const (
	createAlbumQuery = `INSERT INTO album (name, release_date, image, artist_id) 
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, name, release_date, image, artist_id, created_at, updated_at`

	findByIDQuery = `SELECT id, name, release_date, image, artist_id, created_at, updated_at FROM album WHERE id = $1`

	getAllQuery = `SELECT id, name, release_date, image, artist_id, created_at, updated_at FROM album`

	findByQuery = `
	SELECT id, name, release_date, image, artist_id, created_at, updated_at
	FROM album
    WHERE fts @@ to_tsquery('english', $1 || ':*') 
        OR fts @@ to_tsquery('russian_hunspell', $1 || ':*')`

	getByArtistIDQuery = `SELECT id, name, release_date, image, artist_id, created_at, updated_at FROM album WHERE artist_id = $1`

	addFavoriteAlbumQuery = `
    INSERT INTO favorite_album (user_id, album_id) 
    VALUES ($1, $2)
    ON CONFLICT (user_id, album_id) DO NOTHING`

	deleteFavoriteAlbumQuery = `
    DELETE FROM favorite_album
    WHERE user_id = $1 AND album_id = $2`

	isFavoriteAlbumQuery = `
    SELECT 1 
    FROM favorite_album
    WHERE user_id = $1 AND album_id = $2`

	getFavoriteQuery = `
    SELECT a.id AS id, name, release_date, image, artist_id, a.created_at AS created_at, a.updated_at AS updated_at
    FROM album AS a 
      JOIN favorite_album AS fa
      ON a.id = fa.album_id
    WHERE fa.user_id = $1`
)
