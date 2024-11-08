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
  ('Indie', 'Инди'),
  ('Alternative', 'Альтернатива');

INSERT INTO artist
    (name, bio, country, image)
VALUES
    ('EKKSTACY', 'Khyree Zienty, professionally known as Ekkstacy, is a Canadian singer-songwriter from Vancouver, British Columbia.', 'Canada', 'ekkstacy.jpg'),
    ('Sueco', 'William Henry Victor Schultz, better known by his stage name Sueco or SuecoTheChild, is an American rapper and singer-songwriter from Los Angeles.', 'California', 'sueco.jpeg'),
    ('Причастие', 'Калужский дуэт, в который входят две вокалистки: Мария Пенская и Ксения Кузина.', 'Russia', 'prichastie.jpeg'),
    ('Брюки бри', 'Группа из Калуги, бывшее название вб.', 'Russia', 'bryuki_bri.jpeg'),
    ('HXVRMXN', 'Pavel Viktorovich Vankovich, better known by his stage name as HXVRMXN, is a Belarusian phonk music producer.', 'Belarus', 'hxvrmxn.jpeg'),
    ('YONAKA', 'Yonaka (stylised as YONAKA) are an English rock band based in Brighton.', 'England', 'yonaka.jpg');

INSERT INTO album
    (name, track_count, image, artist_id)
VALUES
    ('misery', 5, 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY')),
    ('Attempted Lover', 5, 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco')),
    ('Перламутр', 5, 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие')),
    ('Больше никогда', 5, 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри')),
    ('Eclipse', 3, 'hxvrmxn_eclipse.jpg', (SELECT id FROM artist WHERE name = 'HXVRMXN')),
    ('Don''t Wait ''Til Tomorrow', 5, 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'));

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
    ('Ы', 141, 'bryuki_bri_bolshe_nikogda_5.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда')),
    ('Eclipse', 174, 'hxvrmxn_eclipse.mp3', 'hxvrmxn_eclipse.jpg', (SELECT id FROM artist WHERE name = 'HXVRMXN'), (SELECT id FROM album WHERE name = 'Eclipse')),
    ('Ecoboost', 132, 'hxvrmxn_ecoboost.mp3', 'hxvrmxn_eclipse.jpg', (SELECT id FROM artist WHERE name = 'HXVRMXN'), (SELECT id FROM album WHERE name = 'Eclipse')),
    ('South', 192, 'hxvrmxn_south.mp3', 'hxvrmxn_eclipse.jpg', (SELECT id FROM artist WHERE name = 'HXVRMXN'), (SELECT id FROM album WHERE name = 'Eclipse')),
    ('Don''t Wait ''Til Tomorrow', 198, 'yonaka_dont_wait_til_tomorrow.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow')),
    ('Creature', 185, 'yonaka_creature.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow')),
    ('Fired Up', 217, 'yonaka_fired_up.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow')),
    ('Lose Our Heads', 200, 'yonaka_lose_our_heads.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow')),
    ('Wake Up', 218, 'yonaka_wake_up.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow'));

INSERT INTO genre_artist (genre_id, artist_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM artist WHERE name = 'EKKSTACY')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM artist WHERE name = 'Sueco')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM artist WHERE name = 'Причастие')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM artist WHERE name = 'Брюки бри')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM artist WHERE name = 'HXVRMXN')),
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM artist WHERE name = 'YONAKA'));

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
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM track WHERE name = 'Ы')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM track WHERE name = 'Eclipse')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM track WHERE name = 'Ecoboost')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM track WHERE name = 'South')),
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM track WHERE name = 'Don''t Wait ''Til Tomorrow')),
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM track WHERE name = 'Creature')),
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM track WHERE name = 'Fired Up')),
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM track WHERE name = 'Lose Our Heads')),
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM track WHERE name = 'Wake Up'));

INSERT INTO genre_album (genre_id, album_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM album WHERE name = 'misery')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM album WHERE name = 'Attempted Lover')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM album WHERE name = 'Перламутр')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM album WHERE name = 'Больше никогда')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM album WHERE name = 'Eclipse')),
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow'));
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
