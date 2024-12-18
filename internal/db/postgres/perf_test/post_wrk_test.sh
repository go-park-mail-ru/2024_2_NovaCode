JWT_TOKEN=$(
  curl -X POST \
       -H "Content-Type: application/json" \
       -d '{"username": "testt", "password": "12345678"}' \
       https://nova-music.ru/api/v1/auth/login | \
  jq -r '.token'
)

wrk -c150 -d3000s -t100 -spost_wrk_test.lua \
    -H "Cookie: jwt-token=${JWT_TOKEN}" \
    https://nova-music.ru/api/v1/playlists