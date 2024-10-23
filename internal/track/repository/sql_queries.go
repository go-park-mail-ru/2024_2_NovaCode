package repository

const (
	createTrackQuery = `INSERT INTO track (name, duration, filepath, image, artist_id, album_id, release_date) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, name, duration, filepath, image, artist_id, album_id, release_date, created_at, updated_at`

	findByIDQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id,
  release_date, created_at, updated_at FROM track WHERE id = $1`

	getAllQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, 
  release_date, created_at, updated_at FROM track`

	findByNameQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, 
  release_date, created_at, updated_at FROM track WHERE name = $1`

  	getAllByArtistIDQuery = `SELECT id, name, duration, filepath, image, artist_id, album_id, 
  release_date, created_at, updated_at FROM track WHERE artist_id = $1`
)
