-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "album" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE
    CONSTRAINT album_name_length CHECK (char_length(name) <= 31),
  genre TEXT
    CONSTRAINT album_genre_length CHECK (char_length(genre) <= 31), 
  track_count INT DEFAULT 0,
  release_date TIMESTAMPTZ DEFAULT NOW(),
  image TEXT
    CONSTRAINT album_image_length CHECK (char_length(image) <= 255),
  artist_id INT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp,
  FOREIGN KEY (artist_id) REFERENCES artist (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "album";
-- +goose StatementEnd
