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
    ('Mirella', 'Artist', 'Finland', 'Mirella.jpeg'),
    ('KUUMAA', 'Artist', 'Finland', 'KUUMAA.jpeg'),
    ('JVG', 'Artist', 'Finland', 'JVG.jpeg'),
    ('Eminem', 'Artist', 'USA', 'Eminem.jpeg'),
    ('Robin Packalen', 'Artist', 'Finland', 'Robin_Packalen.jpeg');

INSERT INTO album
    (name, track_count, image, artist_id)
VALUES
    ('Luotathan', 1, 'Luotathan.jpeg', (SELECT id FROM artist WHERE name = 'Mirella')),
    ('Pisara meressä', 1, 'Pisara_meressa.jpeg', (SELECT id FROM artist WHERE name = 'KUUMAA')),
    ('Rallikansa', 1, 'Rallikansa.jpeg', (SELECT id FROM artist WHERE name = 'JVG')),
    ('Kolmistaan', 1, 'Kolmistaan.jpeg', (SELECT id FROM artist WHERE name = 'Robin Packalen')),
    ('The Death of Slim Shady', 1, 'The_Death_of_Slim_Shady.jpeg', (SELECT id FROM artist WHERE name = 'Eminem'));

INSERT INTO track
    (name, duration, filepath, image, artist_id, album_id)
VALUES
    ('Luotathan', 123, 'test_track_1.mp3', 'Luotathan.jpeg', (SELECT id FROM artist WHERE name = 'Mirella'), (SELECT id FROM album WHERE name = 'Luotathan')),
    ('Satama', 123, 'test_track_2.mp3', 'Satama.jpeg', (SELECT id FROM artist WHERE name = 'KUUMAA'), (SELECT id FROM album WHERE name = 'Pisara meressä')),
    ('Rallikansa', 123, 'test_track_3.mp3', 'Rallikansa.jpeg', (SELECT id FROM artist WHERE name = 'JVG'), (SELECT id FROM album WHERE name = 'Rallikansa')),
    ('Kolmistaan', 123, 'test_track_4.mp3', 'Kolmistaan.jpeg', (SELECT id FROM artist WHERE name = 'Robin Packalen'), (SELECT id FROM album WHERE name = 'Kolmistaan')),
    ('Houdini', 123, 'test_track_5.mp3', 'Houdini.jpeg', (SELECT id FROM artist WHERE name = 'Eminem'), (SELECT id FROM album WHERE name = 'The Death of Slim Shady'));

INSERT INTO genre_artist (genre_id, artist_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM artist WHERE name = 'Mirella')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM artist WHERE name = 'KUUMAA')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM artist WHERE name = 'JVG')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM artist WHERE name = 'Eminem')),
  ((SELECT id FROM genre WHERE name = 'Country'), (SELECT id FROM artist WHERE name = 'Robin Packalen'));

INSERT INTO genre_track (genre_id, track_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM track WHERE name = 'Luotathan')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'Satama')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM track WHERE name = 'Rallikansa')),
  ((SELECT id FROM genre WHERE name = 'Country'), (SELECT id FROM track WHERE name = 'Kolmistaan')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM track WHERE name = 'Houdini'));

INSERT INTO genre_album (genre_id, album_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM album WHERE name = 'Luotathan')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM album WHERE name = 'Pisara meressä')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM album WHERE name = 'Rallikansa')),
  ((SELECT id FROM genre WHERE name = 'Country'), (SELECT id FROM album WHERE name = 'Kolmistaan')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM album WHERE name = 'The Death of Slim Shady'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE artist CASCADE;
TRUNCATE TABLE album CASCADE;
TRUNCATE TABLE track CASCADE;
TRUNCATE TABLE genre CASCADE;
TRUNCATE TABLE genre_artist CASCADE;
TRUNCATE TABLE genre_track CASCADE;
TRUNCATE TABLE genre_album CASCADE;
-- +goose StatementEnd
