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
  password_hash TEXT NOT NULL,
  image TEXT DEFAULT 'default.webp',
    CONSTRAINT profile_image_length CHECK (char_length(image) <= 255), 
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
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
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "genre" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE,
    CONSTRAINT genre_name_length CHECK (char_length(name) <= 31),
  rus_name TEXT NOT NULL UNIQUE,
    CONSTRAINT genre_rus_name_length CHECK (char_length(name) <= 31),
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "album" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE,
    CONSTRAINT album_name_length CHECK (char_length(name) <= 31),
  release_date TIMESTAMPTZ DEFAULT NOW(),
  image TEXT,
    CONSTRAINT album_image_length CHECK (char_length(image) <= 255),
  artist_id INT REFERENCES artist (id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "playlist" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL UNIQUE
    CONSTRAINT playlist_name_length CHECK (char_length(name) <= 31),
  image TEXT DEFAULT 'default.webp'
    CONSTRAINT playlist_image_length CHECK (char_length(image) <= 255),
  owner_id UUID REFERENCES "user" (id) ON DELETE CASCADE,
  is_private BOOL NOT NULL DEFAULT false,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
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
  track_order_in_album INT,
  release_date TIMESTAMPTZ DEFAULT NOW(),
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "playlist_track" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  playlist_id INT NOT NULL REFERENCES playlist (id) ON DELETE CASCADE,
  track_order_in_playlist INT,
  track_id INT NOT NULL REFERENCES track (id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  UNIQUE (playlist_id, track_id)
);

CREATE TABLE IF NOT EXISTS "genre_artist" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  genre_id INT NOT NULL REFERENCES genre (id) ON DELETE CASCADE,
  artist_id INT NOT NULL REFERENCES artist (id) ON DELETE CASCADE,
  UNIQUE (genre_id, artist_id)
);

CREATE TABLE IF NOT EXISTS "genre_track" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  genre_id INT NOT NULL REFERENCES genre (id) ON DELETE CASCADE,
  track_id INT NOT NULL REFERENCES track (id) ON DELETE CASCADE,
  UNIQUE (genre_id, track_id)
);

CREATE TABLE IF NOT EXISTS "playlist_user" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  playlist_id INT NOT NULL REFERENCES playlist (id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
  UNIQUE (playlist_id, user_id)
);

CREATE TABLE IF NOT EXISTS "artist_score" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  artist_id INT NOT NULL REFERENCES artist (id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
  score INT NOT NULL CHECK (score IN (-1, 1)),
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  UNIQUE (artist_id, user_id)
);

CREATE TABLE IF NOT EXISTS "favorite_track" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  user_id UUID REFERENCES "user" (id) ON DELETE CASCADE,
  track_id INT NOT NULL REFERENCES track (id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "csat" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  topic TEXT NOT NULL
    CONSTRAINT csat_title_length CHECK (char_length(topic) <= 31)
);

CREATE TABLE IF NOT EXISTS "csat_question" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  title TEXT NOT NULL
    CONSTRAINT csat_question_title_length CHECK (char_length(title) <= 255),
  csat_id INT NOT NULL REFERENCES csat (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "csat_answer" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  score INT NOT NULL,
  user_id UUID NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
  csat_question_id INT NOT NULL REFERENCES csat_question (id) ON DELETE CASCADE
  -- UNIQUE (user_id, csat_question_id)
);


CREATE UNIQUE INDEX favorite_track_unique ON favorite_track (user_id, track_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS favorite_track_unique;

DROP TABLE IF EXISTS "csat_answer" CASCADE;
DROP TABLE IF EXISTS "csat_question" CASCADE;
DROP TABLE IF EXISTS "csat" CASCADE;
DROP TABLE IF EXISTS "favorite_track" CASCADE;
DROP TABLE IF EXISTS "artist_score" CASCADE;
DROP TABLE IF EXISTS "playlist_user" CASCADE;
DROP TABLE IF EXISTS "genre_track" CASCADE;
DROP TABLE IF EXISTS "genre_artist" CASCADE;
DROP TABLE IF EXISTS "playlist_track" CASCADE;
DROP TABLE IF EXISTS "track" CASCADE;
DROP TABLE IF EXISTS "playlist" CASCADE;
DROP TABLE IF EXISTS "album" CASCADE;
DROP TABLE IF EXISTS "genre" CASCADE;
DROP TABLE IF EXISTS "artist" CASCADE;
DROP TABLE IF EXISTS "user" CASCADE;
-- +goose StatementEnd
