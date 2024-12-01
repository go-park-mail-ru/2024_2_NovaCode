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
)
