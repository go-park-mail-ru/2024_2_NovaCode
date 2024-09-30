-- +goose Up
-- +goose StatementBegin
INSERT INTO artist
    (name, bio, country, image)
VALUES
    ('Mirella', 'Artist', 'Finland', 'artists/Mirella.jpeg'),
    ('KUUMAA', 'Artist', 'Finland', 'artists/KUUMAA.jpeg'),
    ('JVG', 'Artist', 'Finland', 'artists/JVG.jpeg'),
    ('Eminem', 'Artist', 'USA', 'artists/Eminem.jpeg'),
    ('Robin Packalen', 'Artist', 'Finland', 'artists/Robin_Packalen.jpeg');


INSERT INTO album
    (name, genre, track_count, image, artist_id)
VALUES
    ('Luotathan', 'Pop', 1, 'albums/Luotathan.jpeg', 1),
    ('Pisara meressä', 'Rap', 1, 'albums/Pisara_meressa.jpeg', 1),
    ('Rallikansa', 'Pop', 1, 'albums/Rallikansa.jpeg', 1),
    ('Kolmistaan', 'Country', 1, 'albums/Kolmistaan.jpeg', 1),
    ('The Death of Slim Shady', 'Hip-Hop', 1, 'albums/The_Death_of_Slim_Shady.jpeg', 1);

INSERT INTO track
    (name, genre, duration, filepath, image, artist_id, album_id)
VALUES
    ('Luotathan', 'Pop', 123, 'test filepath', 'tracks/Luotathan.jpeg', 1, 1),
    ('Satama', 'Rap', 123, 'test filepath', 'tracks/Satama.jpeg', 2, 2),
    ('Rallikansa', 'Pop', 123, 'test filepath', 'tracks/Rallikansa.jpeg', 3, 3),
    ('Kolmistaan', 'Country', 123, 'test filepath', 'tracks/Kolmistaan.jpeg', 4, 4),
    ('Houdini', 'Hip-Hop', 123, 'test filepath', 'tracks/Houdini.jpeg', 5, 5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE artist;
TRUNCATE TABLE album;
TRUNCATE TABLE track;
-- +goose StatementEnd
