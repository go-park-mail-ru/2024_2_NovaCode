package repository

const (
	createGenreQuery = `INSERT INTO genres (name, rus_name, created_at, updated_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id, name, rus_name, created_at, updated_at`

	getAllGenresQuery = `SELECT id, name, rus_name FROM genre`

	getAllGenresByArtistIDQuery = `SELECT g.id, g.name, g.rus_name 
	FROM genre g
	JOIN genre_artist ga ON g.id = ga.genre_id
	WHERE ga.artist_id = $1`

	getAllGenresByAlbumIDQuery = `SELECT g.id, g.name, g.rus_name
	FROM genre g
	JOIN genre_album ga ON g.id = ga.genre_id
	WHERE ga.album_id = $1`

	getAllGenresByTrackIDQuery = `SELECT g.id, g.name, g.rus_name 
	FROM genre g
	JOIN genre_track gt ON g.id = gt.genre_id
	WHERE gt.track_id = $1`
)
