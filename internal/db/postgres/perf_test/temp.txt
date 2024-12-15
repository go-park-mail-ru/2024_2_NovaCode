-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_playlist_id ON playlist (id);

CREATE INDEX IF NOT EXISTS idx_playlist_owner_id ON playlist (owner_id);

CREATE INDEX IF NOT EXISTS idx_track_id ON track (id);

CREATE INDEX IF NOT EXISTS idx_track_artist_id ON track (artist_id);

CREATE INDEX IF NOT EXISTS idx_track_album_id ON track (album_id);

CREATE INDEX IF NOT EXISTS idx_favorite_track_track_id_user_id ON favorite_track (track_id, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_playlist_id;

DROP INDEX IF EXISTS idx_playlist_owner_id;

DROP INDEX IF EXISTS idx_track_id;

DROP INDEX IF EXISTS idx_track_artist_id;

DROP INDEX IF EXISTS idx_track_album_id;

DROP INDEX IF EXISTS idx_favorite_track_track_id_user_id;
-- +goose StatementEnd