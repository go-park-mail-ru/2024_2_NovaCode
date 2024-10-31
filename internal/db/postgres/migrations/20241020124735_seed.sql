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
    ('Luotathan', 123, 'tracks/test_track_1.mp3', 'tracks/Luotathan.jpeg', 1, 1),
    ('Satama', 123, 'tracks/test_track_2.mp3', 'tracks/Satama.jpeg', 2, 2),
    ('Rallikansa', 123, 'tracks/test_track_3.mp3', 'tracks/Rallikansa.jpeg', 3, 3),
    ('Kolmistaan', 123, 'tracks/test_track_4.mp3', 'tracks/Kolmistaan.jpeg', 4, 4),
    ('Houdini', 123, 'tracks/test_track_5.mp3', 'tracks/Houdini.jpeg', 5, 5);

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
DELETE FROM genre;
DELETE FROM artist;
DELETE FROM album;
DELETE FROM track;
DELETE FROM genre_artist;
DELETE FROM genre_track;
DELETE FROM genre_album;
-- +goose StatementEnd
