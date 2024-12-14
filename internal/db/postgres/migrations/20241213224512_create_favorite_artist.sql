-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "favorite_artist" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  user_id UUID REFERENCES "user" (id) ON DELETE CASCADE,
  artist_id INT NOT NULL REFERENCES artist (id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX favorite_artist_unique ON favorite_track (user_id, artist_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "favorite_artist" CASCADE;
-- +goose StatementEnd
