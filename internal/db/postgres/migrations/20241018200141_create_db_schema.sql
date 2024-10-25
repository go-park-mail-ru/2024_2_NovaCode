-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  role TEXT NOT NULL DEFAULT 'regular',
    CONSTRAINT role_type_enum CHECK (role IN ('regular', 'admin')),
  username TEXT NOT NULL UNIQUE,
    CONSTRAINT username_length CHECK (char_length(username) <= 31),
  email TEXT NOT NULL UNIQUE,
    CONSTRAINT email_length CHECK (char_length(email) <= 255),
  password TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS "profile" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  user_id UUID NOT NULL REFERENCES "user" (id) ON DELETE CASCADE UNIQUE,
  avatar_filename TEXT,
    CONSTRAINT profile_avatar_filename_length CHECK (char_length(avatar_filename) <= 255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp 
);

CREATE OR REPLACE FUNCTION create_profile()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO profile (user_id)
  VALUES (NEW.id);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_user_insert
AFTER INSERT ON "user"
FOR EACH ROW
EXECUTE FUNCTION create_profile();

CREATE TABLE IF NOT EXISTS "genre" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE,
    CONSTRAINT genre_name_length CHECK (char_length(name) <= 31),
  rus_name TEXT NOT NULL UNIQUE,
    CONSTRAINT genre_rus_name_length CHECK (char_length(name) <= 31)
);

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

CREATE TABLE IF NOT EXISTS "album" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE,
    CONSTRAINT album_name_length CHECK (char_length(name) <= 31),
  track_count INT DEFAULT 0,
  release_date TIMESTAMPTZ DEFAULT NOW(),
  image TEXT,
    CONSTRAINT album_image_length CHECK (char_length(image) <= 255),
  artist_id INT REFERENCES artist (id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS "playlist" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE
    CONSTRAINT playlist_name_length CHECK (char_length(name) <= 31),
  track_count INT DEFAULT 0,
  image TEXT
    CONSTRAINT playlist_image_length CHECK (char_length(image) <= 255),
  owner_id UUID REFERENCES "user" (id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS "track" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE,
  duration INT,
  filepath TEXT
    CONSTRAINT track_filepath_length CHECK (char_length(filepath) <= 255), 
  image TEXT
    CONSTRAINT track_image_length CHECK (char_length(image) <= 255),
  artist_id INT NOT NULL REFERENCES artist (id) ON DELETE CASCADE,
  album_id INT NOT NULL REFERENCES album (id) ON DELETE CASCADE,
  release_date TIMESTAMPTZ DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS "playlist_track" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  playlist_id INT NOT NULL REFERENCES playlist (id) ON DELETE CASCADE,
  track_id INT NOT NULL REFERENCES track (id) ON DELETE CASCADE,
  UNIQUE (playlist_id, track_id)
);

CREATE TABLE IF NOT EXISTS "genre_artist" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  genre_id INT NOT NULL REFERENCES genre (id) ON DELETE CASCADE,
  artist_id INT NOT NULL REFERENCES artist (id) ON DELETE CASCADE,
  UNIQUE (genre_id, artist_id)
);

CREATE TABLE IF NOT EXISTS "genre_album" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  genre_id INT NOT NULL REFERENCES genre (id) ON DELETE CASCADE,
  album_id INT NOT NULL REFERENCES album (id) ON DELETE CASCADE,
  UNIQUE (genre_id, album_id)
);

CREATE TABLE IF NOT EXISTS "genre_track" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  genre_id INT NOT NULL REFERENCES genre (id) ON DELETE CASCADE,
  track_id INT NOT NULL REFERENCES track (id) ON DELETE CASCADE,
  UNIQUE (genre_id, track_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "genre_track";
DROP TABLE IF EXISTS "genre_album";
DROP TABLE IF EXISTS "genre_artist";
DROP TABLE IF EXISTS "playlist_track";
DROP TABLE IF EXISTS "track";
DROP TABLE IF EXISTS "playlist";
DROP TABLE IF EXISTS "album";
DROP TABLE IF EXISTS "genre";
DROP TABLE IF EXISTS "artist";
DROP TABLE IF EXISTS "profile";
DROP TABLE IF EXISTS "user";

DROP TRIGGER IF EXISTS after_user_insert ON "user";
DROP FUNCTION IF EXISTS create_profile();
-- +goose StatementEnd
