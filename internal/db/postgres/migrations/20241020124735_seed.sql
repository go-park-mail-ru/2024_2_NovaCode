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
  ('Hip-Hop', 'Хип-хоп'),
  ('Indie', 'Инди');

INSERT INTO artist
    (name, bio, country, image)
VALUES
    ('EKKSTACY', 'Khyree Zienty, professionally known as Ekkstacy, is a Canadian singer-songwriter from Vancouver, British Columbia.', 'Canada', 'ekkstacy.jpg'),
    ('Sueco', 'William Henry Victor Schultz, better known by his stage name Sueco or SuecoTheChild, is an American rapper and singer-songwriter from Los Angeles.', 'California', 'sueco.jpeg'),
    ('Причастие', 'Калужский дуэт, в который входят две вокалистки: Мария Пенская и Ксения Кузина.', 'Russia', 'prichastie.jpeg'),
    ('Брюки бри', 'Группа из Калуги, бывшее название вб.', 'Russia', 'bryuki_bri.jpeg');

INSERT INTO album
    (name, track_count, image, artist_id)
VALUES
    ('misery', 5, 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY')),
    ('Attempted Lover', 5, 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco')),
    ('Перламутр', 5, 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие')),
    ('Больше никогда', 5, 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'));

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
    ('Never Even Left', 139, 'sueco_attempted_lover_5.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover')),
    ('Бренди', 151, 'prichastie_perlamytr_1.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр')),
    ('Вечеринка', 208, 'prichastie_perlamytr_2.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр')),
    ('Город', 201, 'prichastie_perlamytr_3.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр')),
    ('Мокрую щеку', 202, 'prichastie_perlamytr_4.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр')),
    ('Пустое утро', 180, 'prichastie_perlamytr_5.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр')),
    ('Последнее свидание', 181, 'bryuki_bri_bolshe_nikogda_1.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда')),
    ('Мини купер', 147, 'bryuki_bri_bolshe_nikogda_2.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда')),
    ('Мамина дочь', 161, 'bryuki_bri_bolshe_nikogda_3.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда')),
    ('Права', 130, 'bryuki_bri_bolshe_nikogda_4.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда')),
    ('Ы', 141, 'bryuki_bri_bolshe_nikogda_5.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда'));

INSERT INTO genre_artist (genre_id, artist_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM artist WHERE name = 'EKKSTACY')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM artist WHERE name = 'Sueco')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM artist WHERE name = 'Причастие')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM artist WHERE name = 'Брюки бри'));

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
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM track WHERE name = 'Never Even Left')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM track WHERE name = 'Бренди')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM track WHERE name = 'Вечеринка')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM track WHERE name = 'Город')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM track WHERE name = 'Мокрую щеку')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM track WHERE name = 'Пустое утро')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM track WHERE name = 'Последнее свидание')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM track WHERE name = 'Мини купер')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM track WHERE name = 'Мамина дочь')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM track WHERE name = 'Права')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM track WHERE name = 'Ы'));

INSERT INTO genre_album (genre_id, album_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM album WHERE name = 'misery')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM album WHERE name = 'Attempted Lover')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM album WHERE name = 'Перламутр')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM album WHERE name = 'Больше никогда'));
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
