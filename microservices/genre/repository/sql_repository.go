package repository

const (
	createGenreQuery = `INSERT INTO genres (name, rus_name, created_at, updated_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id, name, rus_name, created_at, updated_at`

	findByIDQuery = `SELECT id, name, rus_name, created_at, updated_at FROM album WHERE id = $1`

	getAllQuery = `SELECT id, name, rus_name, created_at, updated_at FROM genre`

	getByArtistIDQuery = `SELECT g.id, g.name, g.rus_name, created_at, updated_at
	FROM genre g
	JOIN genre_artist ga ON g.id = ga.genre_id
	WHERE ga.artist_id = $1`

	getByTrackIDQuery = `SELECT g.id, g.name, g.rus_name, created_at, updated_at
	FROM genre g
	JOIN genre_track gt ON g.id = gt.genre_id
	WHERE gt.track_id = $1`
)
