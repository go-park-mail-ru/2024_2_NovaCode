package repository

const (
	createArtistQuery = `INSERT INTO artist (name, bio, country, image) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, bio, country, image, created_at, updated_at`

	findByIDQuery = `SELECT id, name, bio, country, image, created_at, updated_at FROM artist WHERE id = $1`

	getAllQuery = `SELECT id, name, bio, country, image, created_at, updated_at FROM artist`

	findByQuery = `
	SELECT id, name, bio, country, image, created_at, updated_at
	FROM artist
    WHERE fts @@ to_tsquery('english', $1 || ':*') 
        OR fts @@ to_tsquery('russian_hunspell', $1 || ':*')`

	addFavoriteArtistQuery = `
    INSERT INTO favorite_artist (user_id, artist_id) 
    VALUES ($1, $2)
    ON CONFLICT (user_id, artist_id) DO NOTHING`

	deleteFavoriteArtistQuery = `
    DELETE FROM favorite_artist
    WHERE user_id = $1 AND artist_id = $2`

	isFavoriteArtistQuery = `
    SELECT 1 
    FROM favorite_artist 
    WHERE user_id = $1 AND artist_id = $2`

	getFavoriteQuery = `
    SELECT a.id AS id, name, bio, country, image, a.created_at AS created_at, a.updated_at AS updated_at
    FROM artist AS a 
      JOIN favorite_artist AS fa
      ON a.id = fa.artist_id
    WHERE fa.user_id = $1`

	getFavoriteCountQuery = `
    SELECT COUNT(*)
    FROM artist AS a 
      JOIN favorite_artist AS fa
      ON a.id = fa.artist_id
    WHERE fa.user_id = $1`

	getLikesCountQuery = `
    SELECT COUNT(*)
    FROM favorite_artist WHERE artist_id = $1`
)
