-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE role_type AS ENUM ('regular', 'admin');

CREATE TABLE IF NOT EXISTS "user" (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role role_type NOT NULL DEFAULT 'regular',
    username TEXT NOT NULL UNIQUE
        CONSTRAINT username_length CHECK (char_length(username) <= 31),
    email TEXT NOT NULL UNIQUE
        CONSTRAINT email_length CHECK (char_length(email) <= 255),
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
DROP TYPE IF EXISTS role_type;
-- +goose StatementEnd
