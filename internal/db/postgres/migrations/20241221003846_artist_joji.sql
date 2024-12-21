-- +goose Up
-- +goose StatementBegin
INSERT INTO artist
    (name, bio, country, image)
VALUES
    ('joji', 'Joji and formerly as Filthy Frank and Pink Guy, is a Japanese-Australian comedian, singer-songwriter, rapper, and record producer.', 'Japan', 'joji.webp');

INSERT INTO album
    (name, image, artist_id)
VALUES
    ('SMITHEREENS', 'joji_smithereens.webp', (SELECT id FROM artist WHERE name = 'joji'));

INSERT INTO track
    (name, duration, filepath, image, artist_id, album_id, track_order_in_album)
VALUES
    ('Die For You', 212, 'joji_smithereens_1.mp3', 'joji_smithereens.webp', (SELECT id FROM artist WHERE name = 'joji'), (SELECT id FROM album WHERE name = 'SMITHEREENS'), 1),
    ('Dissolve', 177, 'joji_smithereens_2.mp3', 'joji_smithereens.webp', (SELECT id FROM artist WHERE name = 'joji'), (SELECT id FROM album WHERE name = 'SMITHEREENS'), 2),
    ('NIGHT RIDER', 128, 'joji_smithereens_3.mp3', 'joji_smithereens.webp', (SELECT id FROM artist WHERE name = 'joji'), (SELECT id FROM album WHERE name = 'SMITHEREENS'), 3),
    ('BLAHBLAHBLAH DEMO', 143, 'joji_smithereens_4.mp3', 'joji_smithereens.webp', (SELECT id FROM artist WHERE name = 'joji'), (SELECT id FROM album WHERE name = 'SMITHEREENS'), 4),
    ('YUKON (INTERLUDE)', 141, 'joji_smithereens_5.mp3', 'joji_smithereens.webp', (SELECT id FROM artist WHERE name = 'joji'), (SELECT id FROM album WHERE name = 'SMITHEREENS'), 5),
    ('1AM FREESTYLE', 113, 'joji_smithereens_6.mp3', 'joji_smithereens.webp', (SELECT id FROM artist WHERE name = 'joji'), (SELECT id FROM album WHERE name = 'SMITHEREENS'), 6);

INSERT INTO genre_artist (genre_id, artist_id) VALUES
  ((SELECT id FROM genre WHERE name = 'pop'), (SELECT id FROM artist WHERE name = 'joji')),
  ((SELECT id FROM genre WHERE name = 'hip-hop'), (SELECT id FROM artist WHERE name = 'joji'));

INSERT INTO genre_track (genre_id, track_id) VALUES
  ((SELECT id FROM genre WHERE name = 'pop'), (SELECT id FROM track WHERE name = 'Die For You')),
  ((SELECT id FROM genre WHERE name = 'pop'), (SELECT id FROM track WHERE name = 'Dissolve')),
  ((SELECT id FROM genre WHERE name = 'pop'), (SELECT id FROM track WHERE name = 'NIGHT RIDER')),
  ((SELECT id FROM genre WHERE name = 'pop'), (SELECT id FROM track WHERE name = 'BLAHBLAHBLAH DEMO')),
  ((SELECT id FROM genre WHERE name = 'pop'), (SELECT id FROM track WHERE name = 'YUKON (INTERLUDE)')),
  ((SELECT id FROM genre WHERE name = 'pop'), (SELECT id FROM track WHERE name = '1AM FREESTYLE')),
  ((SELECT id FROM genre WHERE name = 'hip-hop'), (SELECT id FROM track WHERE name = 'Die For You')),
  ((SELECT id FROM genre WHERE name = 'hip-hop'), (SELECT id FROM track WHERE name = 'Dissolve')),
  ((SELECT id FROM genre WHERE name = 'hip-hop'), (SELECT id FROM track WHERE name = 'NIGHT RIDER')),
  ((SELECT id FROM genre WHERE name = 'hip-hop'), (SELECT id FROM track WHERE name = 'BLAHBLAHBLAH DEMO')),
  ((SELECT id FROM genre WHERE name = 'hip-hop'), (SELECT id FROM track WHERE name = 'YUKON (INTERLUDE)')),
  ((SELECT id FROM genre WHERE name = 'hip-hop'), (SELECT id FROM track WHERE name = '1AM FREESTYLE'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE artist CASCADE;
TRUNCATE TABLE album CASCADE;
TRUNCATE TABLE track CASCADE;
TRUNCATE TABLE genre_artist CASCADE;
TRUNCATE TABLE genre_track CASCADE;
-- +goose StatementEnd
