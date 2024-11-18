-- +goose Up
-- +goose StatementBegin
ALTER TABLE "track"
    ADD COLUMN fts tsvector;

ALTER TABLE "album"
    ADD COLUMN fts tsvector;

ALTER TABLE "artist"
    ADD COLUMN fts tsvector;

UPDATE "track"
SET fts = to_tsvector(name);

UPDATE "artist"
SET fts = setweight(to_tsvector(name), 'A') ||
    setweight(to_tsvector(country), 'B') ||
    setweight(to_tsvector(bio), 'C');

UPDATE "album"
SET fts = to_tsvector(name);

CREATE INDEX track_fts_idx
    ON "track"
    USING GIN (fts);

CREATE INDEX artist_fts_idx
    ON "artist"
    USING GIN (fts);

CREATE INDEX album_fts_idx
    ON "album"
    USING GIN (fts);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS album_fts_idx;
DROP INDEX IF EXISTS artist_fts_idx;
DROP INDEX IF EXISTS track_fts_idx;

ALTER TABLE "album"
DROP COLUMN IF EXISTS fts;

ALTER TABLE "artist"
DROP COLUMN IF EXISTS fts;

ALTER TABLE "track"
DROP COLUMN IF EXISTS fts;
-- +goose StatementEnd