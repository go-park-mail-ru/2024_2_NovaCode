local wrk = require "wrk"

-- Инициализируем генератор случайных чисел
math.randomseed(os.clock() * 1000000)

-- Функция для генерации случайного имени плейлиста
local function generate_playlist_name()
  local characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  local length = math.random(5, 15) -- Длина имени от 5 до 15 символов
  local playlist_name = ""
  
  for i = 1, length do
    local random_index = math.random(1, #characters)
    local random_char = characters:sub(random_index, random_index)
    playlist_name = playlist_name .. random_char
  end
  
  return playlist_name
end

-- Ограничение кличества запросов (публикуем 100_000 сущностей)
local counter = 1

function response()
   if counter == 1000 then
      wrk.thread:stop()
   end
   counter = counter + 1
end

wrk.method = "POST"
wrk.body   = '{"name": "' .. generate_playlist_name() .. '"}'
wrk.headers["Content-Type"] = "application/json"
