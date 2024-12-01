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
)
