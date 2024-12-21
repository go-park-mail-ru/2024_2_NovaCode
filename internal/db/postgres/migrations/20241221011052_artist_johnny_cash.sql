-- +goose Up
-- +goose StatementBegin
INSERT INTO artist
(name, bio, country, image)
VALUES
('Johnny Cash', 'American singer-songwriter of Scottish descent, a key figure in country music, and one of the most influential musicians of the 20th century.', 'USA', 'johnny_cash.webp');

INSERT INTO album
(name, image, artist_id)
VALUES
('Cry, Cry, Cry', 'johnny_cash_cry_cry_cry.webp', (SELECT id FROM artist WHERE name = 'Johnny Cash'));

INSERT INTO track
(name, duration, filepath, image, artist_id, album_id, track_order_in_album)
VALUES
('Don''t Make Me Go', 149, 'johnny_cash_cry_cry_cry_1.mp3', 'johnny_cash_cry_cry_cry.webp', (SELECT id FROM artist WHERE name = 'Johnny Cash'), (SELECT id FROM album WHERE name = 'Cry, Cry, Cry'), 1),     
('Hey Porter', 141, 'johnny_cash_cry_cry_cry_2.mp3', 'johnny_cash_cry_cry_cry.webp', (SELECT id FROM artist WHERE name = 'Johnny Cash'), (SELECT id FROM album WHERE name = 'Cry, Cry, Cry'), 2),     
('Ballad Of A Teenage Queen', 134, 'johnny_cash_cry_cry_cry_3.mp3', 'johnny_cash_cry_cry_cry.webp', (SELECT id FROM artist WHERE name = 'Johnny Cash'), (SELECT id FROM album WHERE name = 'Cry, Cry, Cry'), 3),     
('Doin'' My Time', 157, 'johnny_cash_cry_cry_cry_4.mp3', 'johnny_cash_cry_cry_cry.webp', (SELECT id FROM artist WHERE name = 'Johnny Cash'), (SELECT id FROM album WHERE name = 'Cry, Cry, Cry'), 4),
('Cry, Cry, Cry', 147, 'johnny_cash_cry_cry_cry_5.mp3', 'johnny_cash_cry_cry_cry.webp', (SELECT id FROM artist WHERE name = 'Johnny Cash'), (SELECT id FROM album WHERE name = 'Cry, Cry, Cry'), 5),
('Straight A''s In Love', 141, 'johnny_cash_cry_cry_cry_6.mp3', 'johnny_cash_cry_cry_cry.webp', (SELECT id FROM artist WHERE name = 'Johnny Cash'), (SELECT id FROM album WHERE name = 'Cry, Cry, Cry'), 6);

INSERT INTO genre_artist (genre_id, artist_id) VALUES
((SELECT id FROM genre WHERE name = 'country'), (SELECT id FROM artist WHERE name = 'Johnny Cash'));

INSERT INTO genre_track (genre_id, track_id) VALUES
((SELECT id FROM genre WHERE name = 'country'), (SELECT id FROM track WHERE name = 'Don''t Make Me Go')),   
((SELECT id FROM genre WHERE name = 'country'), (SELECT id FROM track WHERE name = 'Hey Porter')),   
((SELECT id FROM genre WHERE name = 'country'), (SELECT id FROM track WHERE name = 'Ballad Of A Teenage Queen')),   
((SELECT id FROM genre WHERE name = 'country'), (SELECT id FROM track WHERE name = 'Doin'' My Time')),
((SELECT id FROM genre WHERE name = 'country'), (SELECT id FROM track WHERE name = 'Cry, Cry, Cry')),
((SELECT id FROM genre WHERE name = 'country'), (SELECT id FROM track WHERE name = 'Straight A''s In Love'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE artist CASCADE;
TRUNCATE TABLE album CASCADE;
TRUNCATE TABLE track CASCADE;
TRUNCATE TABLE genre_artist CASCADE;
TRUNCATE TABLE genre_track CASCADE;
-- +goose StatementEnd