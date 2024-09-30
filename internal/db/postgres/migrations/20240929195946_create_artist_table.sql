-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "artist" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE,
    CONSTRAINT artist_name_length CHECK (char_length(name) <= 31),
  bio TEXT,
    CONSTRAINT artist_bio_length CHECK (char_length(bio) <= 255),
  country TEXT,
    CONSTRAINT artist_country_length CHECK (char_length(country) <= 31),
  image TEXT,
    CONSTRAINT artist_image_length CHECK (char_length(image) <= 255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "artist";
-- +goose StatementEnd
