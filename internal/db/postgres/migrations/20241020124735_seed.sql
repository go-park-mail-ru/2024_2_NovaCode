-- +goose Up
-- +goose StatementBegin
INSERT INTO genre 
    (name, rus_name) 
VALUES
  ('Pop', 'Поп'),
  ('Rap', 'Рэп'),
  ('Rock', 'Рок'),
  ('Classical', 'Классика'),
  ('Country', 'Кантри'),
  ('Hip-Hop', 'Хип-хоп');

INSERT INTO artist
    (name, bio, country, image)
VALUES
    ('Mirella', 'Artist', 'Finland', 'artists/Mirella.jpeg'),
    ('KUUMAA', 'Artist', 'Finland', 'artists/KUUMAA.jpeg'),
    ('JVG', 'Artist', 'Finland', 'artists/JVG.jpeg'),
    ('Eminem', 'Artist', 'USA', 'artists/Eminem.jpeg'),
    ('Robin Packalen', 'Artist', 'Finland', 'artists/Robin_Packalen.jpeg');


INSERT INTO album
    (name, track_count, image, artist_id)
VALUES
    ('Luotathan', 1, 'albums/Luotathan.jpeg', 1),
    ('Pisara meressä', 1, 'albums/Pisara_meressa.jpeg', 1),
    ('Rallikansa', 1, 'albums/Rallikansa.jpeg', 1),
    ('Kolmistaan', 1, 'albums/Kolmistaan.jpeg', 1),
    ('The Death of Slim Shady', 1, 'albums/The_Death_of_Slim_Shady.jpeg', 1);

INSERT INTO track
    (name, duration, filepath, image, artist_id, album_id)
VALUES
    ('Luotathan', 123, 'test filepath', 'tracks/Luotathan.jpeg', 1, 1),
    ('Satama', 123, 'test filepath', 'tracks/Satama.jpeg', 2, 2),
    ('Rallikansa', 123, 'test filepath', 'tracks/Rallikansa.jpeg', 3, 3),
    ('Kolmistaan', 123, 'test filepath', 'tracks/Kolmistaan.jpeg', 4, 4),
    ('Houdini', 123, 'test filepath', 'tracks/Houdini.jpeg', 5, 5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE artist CASCADE;
TRUNCATE TABLE album CASCADE;
TRUNCATE TABLE track CASCADE;
TRUNCATE TABLE genre CASCADE;
-- +goose StatementEnd
