-- +goose Up
-- +goose StatementBegin
DO
$$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'novamusic_app_role') THEN
        CREATE ROLE novamusic_app_role;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'novamusic_app_user') THEN
        CREATE USER novamusic_app_user WITH PASSWORD 'novamusic_app_password';
        GRANT novamusic_app_role TO novamusic_app_user;
    END IF;

    GRANT SELECT, UPDATE, DELETE, INSERT ON TABLE "user", artist, playlist, playlist_track, genre_artist, genre_track, playlist_user, artist_score, favorite_track TO novamusic_app_role;
    GRANT SELECT, INSERT, DELETE ON TABLE album, track TO novamusic_app_role;
    GRANT SELECT ON TABLE genre, csat, csat_question, csat_answer TO novamusic_app_role;
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DO
$$
BEGIN
    REVOKE SELECT, UPDATE, DELETE, INSERT ON TABLE "user", artist, playlist, playlist_track, genre_artist, genre_track, playlist_user, artist_score, favorite_track FROM novamusic_app_role;
    REVOKE SELECT, INSERT, DELETE ON TABLE album, track FROM novamusic_app_role;
    REVOKE SELECT ON TABLE genre, csat, csat_question, csat_answer FROM novamusic_app_role;

    IF EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'novamusic_app_user') THEN
        DROP USER novamusic_app_user;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'novamusic_app_role') THEN
        DROP ROLE novamusic_app_role;
    END IF;
END;
$$;
-- +goose StatementEnd