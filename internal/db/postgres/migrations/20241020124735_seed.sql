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
    ('EKKSTACY', 'Khyree Zienty, professionally known as Ekkstacy, is a Canadian singer-songwriter from Vancouver, British Columbia.', 'Canada', 'ekkstacy.jpg'),
    ('Sueco', 'William Henry Victor Schultz, better known by his stage name Sueco or SuecoTheChild, is an American rapper and singer-songwriter from Los Angeles.', 'California', 'sueco.jpeg');

INSERT INTO album
    (name, track_count, image, artist_id)
VALUES
    ('misery', 5, 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY')),
    ('Attempted Lover', 5, 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'));

INSERT INTO track
    (name, duration, filepath, image, artist_id, album_id)
VALUES
    ('i just want to hide my face', 132, 'ekkstacy_misery_1.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery')),
    ('im so happy', 139, 'ekkstacy_misery_2.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery')),
    ('i wish you were pretty on the inside', 127, 'ekkstacy_misery_3.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery')),
    ('christian death', 144, 'ekkstacy_misery_4.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery')),
    ('i want to die in your arms', 136, 'ekkstacy_misery_5.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery')),
    ('Wreck', 167, 'sueco_attempted_lover_1.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover')),
    ('Wanna Feel Something', 203, 'sueco_attempted_lover_2.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover')),
    ('452AM', 172, 'sueco_attempted_lover_3.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover')),
    ('Bad Idea', 151, 'sueco_attempted_lover_4.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover')),
    ('Never Even Left', 139, 'sueco_attempted_lover_5.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover'));

INSERT INTO genre_artist (genre_id, artist_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM artist WHERE name = 'EKKSTACY')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM artist WHERE name = 'Sueco'));

INSERT INTO genre_track (genre_id, track_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'i just want to hide my face')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'im so happy')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'i wish you were pretty on the inside')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'christian death')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'i want to die in your arms')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM track WHERE name = 'Wreck')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM track WHERE name = 'Wanna Feel Something')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM track WHERE name = '452AM')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM track WHERE name = 'Bad Idea')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM track WHERE name = 'Never Even Left'));

INSERT INTO genre_album (genre_id, album_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM album WHERE name = 'misery')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM album WHERE name = 'Attempted Lover'));
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
