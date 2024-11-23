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
  ('Alternative', 'Альтернатива'),
  ('IDM', 'Электроника');

INSERT INTO artist
    (name, bio, country, image)
VALUES
    ('EKKSTACY', 'Khyree Zienty, professionally known as Ekkstacy, is a Canadian singer-songwriter from Vancouver, British Columbia.', 'Canada', 'ekkstacy.jpg'),
    ('Sueco', 'William Henry Victor Schultz, better known by his stage name Sueco or SuecoTheChild, is an American rapper and singer-songwriter from Los Angeles.', 'California', 'sueco.jpeg'),
    ('Причастие', 'Калужский дуэт, в который входят две вокалистки: Мария Пенская и Ксения Кузина.', 'Russia', 'prichastie.jpeg'),
    ('Брюки бри', 'Группа из Калуги, бывшее название вб.', 'Russia', 'bryuki_bri.jpeg'),
    ('HXVRMXN', 'Pavel Viktorovich Vankovich, better known by his stage name as HXVRMXN, is a Belarusian phonk music producer.', 'Belarus', 'hxvrmxn.jpeg'),
    ('YONAKA', 'Yonaka (stylised as YONAKA) are an English rock band based in Brighton.', 'England', 'yonaka.jpg'),
    ('Yung Lean', 'Yung Lean, is a Swedish rapper. Widely cited as one of the most influential figures in the early cloud rap era, Yung Lean rose to prominence in 2013 with his song "Ginseng Strip 2002", which went viral on YouTube.', 'Sweden', 'yung_lean.jpeg'),
    ('Thom Yorke', 'Thomas Edward Yorke is an English musician who is the main vocalist and songwriter of the rock band Radiohead. Rolling Stone described Yorke as one of the greatest and most influential singers of his generation.', 'England', 'thom_yorke.jpeg');

INSERT INTO album
    (name, image, artist_id)
VALUES
    ('misery', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY')),
    ('Attempted Lover', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco')),
    ('Перламутр', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие')),
    ('Больше никогда', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри')),
    ('Eclipse', 'hxvrmxn_eclipse.jpg', (SELECT id FROM artist WHERE name = 'HXVRMXN')),
    ('Don''t Wait ''Til Tomorrow', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA')),
    ('Stranger', 'yung_lean_stranger.jpeg', (SELECT id FROM artist WHERE name = 'Yung Lean')),
    ('The Eraser', 'thom_yorke_the_eraser.png', (SELECT id FROM artist WHERE name = 'Thom Yorke'));

INSERT INTO track
    (name, duration, filepath, image, artist_id, album_id, track_order_in_album)
VALUES
    ('i just want to hide my face', 132, 'ekkstacy_misery_1.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery'), 1),
    ('im so happy', 139, 'ekkstacy_misery_2.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery'), 2),
    ('i wish you were pretty on the inside', 127, 'ekkstacy_misery_3.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery'), 3),
    ('christian death', 144, 'ekkstacy_misery_4.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery'), 4),
    ('i want to die in your arms', 136, 'ekkstacy_misery_5.mp3', 'ekkstacy_misery.jpeg', (SELECT id FROM artist WHERE name = 'EKKSTACY'), (SELECT id FROM album WHERE name = 'misery'), 5),
    ('Wreck', 167, 'sueco_attempted_lover_1.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover'), 1),
    ('Wanna Feel Something', 203, 'sueco_attempted_lover_2.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover'), 2),
    ('452AM', 172, 'sueco_attempted_lover_3.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover'), 3),
    ('Bad Idea', 151, 'sueco_attempted_lover_4.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover'), 4),
    ('Never Even Left', 139, 'sueco_attempted_lover_5.mp3', 'sueco_attempted_lover.jpeg', (SELECT id FROM artist WHERE name = 'Sueco'), (SELECT id FROM album WHERE name = 'Attempted Lover'), 5),
    ('Бренди', 151, 'prichastie_perlamytr_1.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр'), 1),
    ('Вечеринка', 208, 'prichastie_perlamytr_2.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр'), 2),
    ('Город', 201, 'prichastie_perlamytr_3.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр'), 3),
    ('Мокрую щеку', 202, 'prichastie_perlamytr_4.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр'), 4),
    ('Пустое утро', 180, 'prichastie_perlamytr_5.mp3', 'prichastie_perlamytr.jpeg', (SELECT id FROM artist WHERE name = 'Причастие'), (SELECT id FROM album WHERE name = 'Перламутр'), 5),
    ('Последнее свидание', 181, 'bryuki_bri_bolshe_nikogda_1.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда'), 1),
    ('Мини купер', 147, 'bryuki_bri_bolshe_nikogda_2.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда'), 2),
    ('Мамина дочь', 161, 'bryuki_bri_bolshe_nikogda_3.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда'), 3),
    ('Права', 130, 'bryuki_bri_bolshe_nikogda_4.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда'), 4),
    ('Ы', 141, 'bryuki_bri_bolshe_nikogda_5.mp3', 'bryuki_bri_bolshe_nikogda.jpeg', (SELECT id FROM artist WHERE name = 'Брюки бри'), (SELECT id FROM album WHERE name = 'Больше никогда'), 5),
    ('Eclipse', 174, 'hxvrmxn_eclipse.mp3', 'hxvrmxn_eclipse.jpg', (SELECT id FROM artist WHERE name = 'HXVRMXN'), (SELECT id FROM album WHERE name = 'Eclipse'), 1),
    ('Ecoboost', 132, 'hxvrmxn_ecoboost.mp3', 'hxvrmxn_eclipse.jpg', (SELECT id FROM artist WHERE name = 'HXVRMXN'), (SELECT id FROM album WHERE name = 'Eclipse'), 2),
    ('South', 192, 'hxvrmxn_south.mp3', 'hxvrmxn_eclipse.jpg', (SELECT id FROM artist WHERE name = 'HXVRMXN'), (SELECT id FROM album WHERE name = 'Eclipse'), 3),
    ('Don''t Wait ''Til Tomorrow', 198, 'yonaka_dont_wait_til_tomorrow.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow'), 1),
    ('Creature', 185, 'yonaka_creature.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow'), 2),
    ('Fired Up', 217, 'yonaka_fired_up.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow'), 3),
    ('Lose Our Heads', 200, 'yonaka_lose_our_heads.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow'), 4),
    ('Wake Up', 218, 'yonaka_wake_up.mp3', 'yonaka_dont_wait_til_tomorrow.jpg', (SELECT id FROM artist WHERE name = 'YONAKA'), (SELECT id FROM album WHERE name = 'Don''t Wait ''Til Tomorrow'), 5),
    ('Agony', 213, 'yung_lean_stranger_1.mp3', 'yung_lean_stranger.jpeg', (SELECT id FROM artist WHERE name = 'Yung Lean'), (SELECT id FROM album WHERE name = 'Stranger'), 1),
    ('Drop It / Scooter', 245, 'yung_lean_stranger_2.mp3', 'yung_lean_stranger.jpeg', (SELECT id FROM artist WHERE name = 'Yung Lean'), (SELECT id FROM album WHERE name = 'Stranger'), 2),
    ('Push / Lost Weekend', 173, 'yung_lean_stranger_3.mp3', 'yung_lean_stranger.jpeg', (SELECT id FROM artist WHERE name = 'Yung Lean'), (SELECT id FROM album WHERE name = 'Stranger'), 3),
    ('Red Bottom Sky', 302, 'yung_lean_stranger_4.mp3', 'yung_lean_stranger.jpeg', (SELECT id FROM artist WHERE name = 'Yung Lean'), (SELECT id FROM album WHERE name = 'Stranger'), 4),
    ('Muddy Sea', 164, 'yung_lean_stranger_5.mp3', 'yung_lean_stranger.jpeg', (SELECT id FROM artist WHERE name = 'Yung Lean'), (SELECT id FROM album WHERE name = 'Stranger'), 5),
    ('The Clock', 253, 'thom_yorke_the_eraser_1.mp3', 'thom_yorke_the_eraser.png', (SELECT id FROM artist WHERE name = 'Thom Yorke'), (SELECT id FROM album WHERE name = 'The Eraser'), 1),
    ('Cymbal Rush', 314, 'thom_yorke_the_eraser_2.mp3', 'thom_yorke_the_eraser.png', (SELECT id FROM artist WHERE name = 'Thom Yorke'), (SELECT id FROM album WHERE name = 'The Eraser'), 2),
    ('Atoms For Peace', 313, 'thom_yorke_the_eraser_3.mp3', 'thom_yorke_the_eraser.png', (SELECT id FROM artist WHERE name = 'Thom Yorke'), (SELECT id FROM album WHERE name = 'The Eraser'), 3),
    ('Analyze', 242, 'thom_yorke_the_eraser_4.mp3', 'thom_yorke_the_eraser.png', (SELECT id FROM artist WHERE name = 'Thom Yorke'), (SELECT id FROM album WHERE name = 'The Eraser'), 4),
    ('The Eraser', 295, 'thom_yorke_the_eraser_5.mp3', 'thom_yorke_the_eraser.png', (SELECT id FROM artist WHERE name = 'Thom Yorke'), (SELECT id FROM album WHERE name = 'The Eraser'), 5);

INSERT INTO genre_artist (genre_id, artist_id) VALUES
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM artist WHERE name = 'EKKSTACY')),
  ((SELECT id FROM genre WHERE name = 'Rock'), (SELECT id FROM artist WHERE name = 'Sueco')),
  ((SELECT id FROM genre WHERE name = 'Pop'), (SELECT id FROM artist WHERE name = 'Причастие')),
  ((SELECT id FROM genre WHERE name = 'Indie'), (SELECT id FROM artist WHERE name = 'Брюки бри')),
  ((SELECT id FROM genre WHERE name = 'Hip-Hop'), (SELECT id FROM artist WHERE name = 'HXVRMXN')),
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM artist WHERE name = 'YONAKA')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM artist WHERE name = 'Yung Lean')),
  ((SELECT id FROM genre WHERE name = 'IDM'), (SELECT id FROM artist WHERE name = 'Thom Yorke'));

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
  ((SELECT id FROM genre WHERE name = 'Alternative'), (SELECT id FROM track WHERE name = 'Wake Up')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'Agony')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'Drop It / Scooter')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'Push / Lost Weekend')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'Red Bottom Sky')),
  ((SELECT id FROM genre WHERE name = 'Rap'), (SELECT id FROM track WHERE name = 'Muddy Sea')),
  ((SELECT id FROM genre WHERE name = 'IDM'), (SELECT id FROM track WHERE name = 'The Clock')),
  ((SELECT id FROM genre WHERE name = 'IDM'), (SELECT id FROM track WHERE name = 'Cymbal Rush')),
  ((SELECT id FROM genre WHERE name = 'IDM'), (SELECT id FROM track WHERE name = 'Atoms For Peace')),
  ((SELECT id FROM genre WHERE name = 'IDM'), (SELECT id FROM track WHERE name = 'Analyze')),
  ((SELECT id FROM genre WHERE name = 'IDM'), (SELECT id FROM track WHERE name = 'The Eraser'));

INSERT INTO csat (topic) VALUES
    ('Общие');

INSERT INTO csat_question (title, csat_id) VALUES
  ('Насколько вы убовлетворены удобством NovaCode?', (SELECT id FROM csat WHERE topic = 'Общие')),
  ('Насколько вам лего было найти желаемую музыку?', (SELECT id FROM csat WHERE topic = 'Общие')),
  ('Насколько вы удовлетворены разнообразием музыкальных жанров и артистов?', (SELECT id FROM csat WHERE topic = 'Общие')),
  ('Насколько вам удобно было использовать плейер?', (SELECT id FROM csat WHERE topic = 'Общие')),
  ('Порекомендуете ли вы NovaMusic друзьям?', (SELECT id FROM csat WHERE topic = 'Общие'));

INSERT INTO "user" (username, email, password_hash) VALUES
    ('test', 'mail@mail.ru', 1234),
    ('test2', 'mail2@mail.ru', 1234);

INSERT INTO csat_answer (score, user_id, csat_question_id) VALUES
    (7, (SELECT id FROM "user" WHERE username = 'test'), (SELECT id FROM csat_question WHERE title = 'Насколько вы убовлетворены удобством NovaCode?')),
    (4, (SELECT id FROM "user" WHERE username = 'test'), (SELECT id FROM csat_question WHERE title = 'Насколько вам лего было найти желаемую музыку?')),
    (3, (SELECT id FROM "user" WHERE username = 'test'), (SELECT id FROM csat_question WHERE title = 'Насколько вы удовлетворены разнообразием музыкальных жанров и артистов?')),
    (1, (SELECT id FROM "user" WHERE username = 'test'), (SELECT id FROM csat_question WHERE title = 'Насколько вам удобно было использовать плейер?')),
    (2, (SELECT id FROM "user" WHERE username = 'test'), (SELECT id FROM csat_question WHERE title = 'Порекомендуете ли вы NovaMusic друзьям?')),
    (3, (SELECT id FROM "user" WHERE username = 'test2'), (SELECT id FROM csat_question WHERE title = 'Насколько вы убовлетворены удобством NovaCode?')),
    (5, (SELECT id FROM "user" WHERE username = 'test2'), (SELECT id FROM csat_question WHERE title = 'Насколько вам лего было найти желаемую музыку?')),
    (2, (SELECT id FROM "user" WHERE username = 'test2'), (SELECT id FROM csat_question WHERE title = 'Насколько вы удовлетворены разнообразием музыкальных жанров и артистов?')),
    (8, (SELECT id FROM "user" WHERE username = 'test2'), (SELECT id FROM csat_question WHERE title = 'Насколько вам удобно было использовать плейер?'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE artist CASCADE;
TRUNCATE TABLE album CASCADE;
TRUNCATE TABLE track CASCADE;
TRUNCATE TABLE genre CASCADE;
TRUNCATE TABLE genre_artist CASCADE;
TRUNCATE TABLE genre_track CASCADE;
TRUNCATE TABLE csat CASCADE;
TRUNCATE TABLE csat_question CASCADE;
TRUNCATE TABLE csat_answer CASCADE;
-- +goose StatementEnd
