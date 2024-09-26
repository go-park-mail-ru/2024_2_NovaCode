package repository

const (
	createArtistQuery = `INSERT INTO artist (name, bio, county, image) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, bio, county, image, created_at, updated_at`

	findByIDQuery = `SELECT id, name, bio, county, image, created_at, updated_at FROM artist WHERE id = $1`

	findByNameQuery = `SELECT id, name, bio, county, image, created_at, updated_at FROM artist WHERE name = $1`
)
