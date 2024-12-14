local wrk = require "wrk"

-- -- Функция для генерации случайного имени плейлиста
-- local function generate_playlist_name()
--   local characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
--   local length = math.random(5, 15) -- Длина имени от 5 до 15 символов
--   local playlist_name = ""
  
--   for i = 1, length do
--     local random_index = math.random(1, #characters)
--     local random_char = characters:sub(random_index, random_index)
--     playlist_name = playlist_name .. random_char
--   end
  
--   return playlist_name
-- end

-- wrk.requests = 100000

wrk.method = "POST"
-- wrk.body   = '{"name": "' .. generate_playlist_name() .. '"}'
wrk.body   = '{"name": "wrk_test"}'
wrk.headers["Content-Type"] = "application/json"