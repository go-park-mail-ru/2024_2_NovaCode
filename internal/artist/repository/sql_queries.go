package repository

const (
	createArtistQuery = `INSERT INTO artist (name, bio, country, image) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, bio, country, image, created_at, updated_at`

	findByIDQuery = `SELECT id, name, bio, country, image, created_at, updated_at FROM artist WHERE id = $1`

	getAllQuery = `SELECT id, name, bio, country, image, created_at, updated_at FROM artist`

	findByNameQuery = `SELECT id, name, bio, country, image, created_at, updated_at FROM artist WHERE name = $1`
)
