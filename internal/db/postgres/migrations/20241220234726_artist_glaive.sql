-- +goose Up
-- +goose StatementBegin
INSERT INTO artist
    (name, bio, country, image)
VALUES
    ('glaive', 'Ash Blue Gutierrez (born January 20, 2005), known professionally as Glaive (stylized as glaive), is an American singer-songwriter.', 'North Carolina', 'glaive.webp');

INSERT INTO album
    (name, image, artist_id)
VALUES
    ('a bit of a mad one', 'glaive_a_bit_of_a_mad_one.webp', (SELECT id FROM artist WHERE name = 'glaive'));

INSERT INTO track
    (name, duration, filepath, image, artist_id, album_id, track_order_in_album)
VALUES
    ('even when the sun is dead, will you tell them how hard i tried', 137, 'glaive_a_bit_1.mp3', 'glaive_a_bit_of_a_mad_one.webp', (SELECT id FROM artist WHERE name = 'glaive'), (SELECT id FROM album WHERE name = 'a bit of a mad one'), 1),
    ('i don''t really feel it anymore', 120, 'glaive_a_bit_2.mp3', 'glaive_a_bit_of_a_mad_one.webp', (SELECT id FROM artist WHERE name = 'glaive'), (SELECT id FROM album WHERE name = 'a bit of a mad one'), 2),
    ('huh', 107, 'glaive_a_bit_3.mp3', 'glaive_a_bit_of_a_mad_one.webp', (SELECT id FROM artist WHERE name = 'glaive'), (SELECT id FROM album WHERE name = 'a bit of a mad one'), 3),
    ('hope alaska national anthem', 119, 'glaive_a_bit_4.mp3', 'glaive_a_bit_of_a_mad_one.webp', (SELECT id FROM artist WHERE name = 'glaive'), (SELECT id FROM album WHERE name = 'a bit of a mad one'), 4),
    ('god is dead', 131, 'glaive_a_bit_5.mp3', 'glaive_a_bit_of_a_mad_one.webp', (SELECT id FROM artist WHERE name = 'glaive'), (SELECT id FROM album WHERE name = 'a bit of a mad one'), 5),
    ('living proof (that it hurts)', 103, 'glaive_a_bit_6.mp3', 'glaive_a_bit_of_a_mad_one.webp', (SELECT id FROM artist WHERE name = 'glaive'), (SELECT id FROM album WHERE name = 'a bit of a mad one'), 6),
    ('phobie d''impulsion', 116, 'glaive_a_bit_7.mp3', 'glaive_a_bit_of_a_mad_one.webp', (SELECT id FROM artist WHERE name = 'glaive'), (SELECT id FROM album WHERE name = 'a bit of a mad one'), 7);

INSERT INTO genre_artist (genre_id, artist_id) VALUES
  ((SELECT id FROM genre WHERE name = 'alternative'), (SELECT id FROM artist WHERE name = 'glaive')),
  ((SELECT id FROM genre WHERE name = 'indie'), (SELECT id FROM artist WHERE name = 'glaive'));

INSERT INTO genre_track (genre_id, track_id) VALUES
  ((SELECT id FROM genre WHERE name = 'indie'), (SELECT id FROM track WHERE name = 'even when the sun is dead, will you tell them how hard i tried')),
  ((SELECT id FROM genre WHERE name = 'indie'), (SELECT id FROM track WHERE name = 'i don''t really feel it anymore')),
  ((SELECT id FROM genre WHERE name = 'indie'), (SELECT id FROM track WHERE name = 'huh')),
  ((SELECT id FROM genre WHERE name = 'indie'), (SELECT id FROM track WHERE name = 'hope alaska national anthem')),
  ((SELECT id FROM genre WHERE name = 'indie'), (SELECT id FROM track WHERE name = 'god is dead')),
  ((SELECT id FROM genre WHERE name = 'indie'), (SELECT id FROM track WHERE name = 'living proof (that it hurts)')),
  ((SELECT id FROM genre WHERE name = 'indie'), (SELECT id FROM track WHERE name = 'phobie d''impulsion')),
  ((SELECT id FROM genre WHERE name = 'alternative'), (SELECT id FROM track WHERE name = 'even when the sun is dead, will you tell them how hard i tried')),
  ((SELECT id FROM genre WHERE name = 'alternative'), (SELECT id FROM track WHERE name = 'i don''t really feel it anymore')),
  ((SELECT id FROM genre WHERE name = 'alternative'), (SELECT id FROM track WHERE name = 'huh')),
  ((SELECT id FROM genre WHERE name = 'alternative'), (SELECT id FROM track WHERE name = 'hope alaska national anthem')),
  ((SELECT id FROM genre WHERE name = 'alternative'), (SELECT id FROM track WHERE name = 'god is dead')),
  ((SELECT id FROM genre WHERE name = 'alternative'), (SELECT id FROM track WHERE name = 'living proof (that it hurts)')),
  ((SELECT id FROM genre WHERE name = 'alternative'), (SELECT id FROM track WHERE name = 'phobie d''impulsion'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE artist CASCADE;
TRUNCATE TABLE album CASCADE;
TRUNCATE TABLE track CASCADE;
TRUNCATE TABLE genre_artist CASCADE;
TRUNCATE TABLE genre_track CASCADE;
-- +goose StatementEnd


