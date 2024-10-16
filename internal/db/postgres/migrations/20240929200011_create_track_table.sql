-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "track" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE,
  genre TEXT
    CONSTRAINT track_genre_length CHECK (char_length(genre) <= 31),
  duration INT,
  filepath TEXT
    CONSTRAINT track_filepath_length CHECK (char_length(filepath) <= 255), 
  image TEXT
    CONSTRAINT track_image_length CHECK (char_length(image) <= 255),
  artist_id INT NOT NULL,
  album_id INT NOT NULL,
  release_date TIMESTAMPTZ DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp,
  FOREIGN KEY (artist_id) REFERENCES artist (id) ON DELETE CASCADE,
  FOREIGN KEY (album_id) REFERENCES album (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "track";
-- +goose StatementEnd
