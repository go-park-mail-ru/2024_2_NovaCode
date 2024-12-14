JWT_TOKEN=$(
  curl -X POST \
       -H "Content-Type: application/json" \
       -d '{"username": "testt", "password": "12345678"}' \
       http://localhost:8081/api/v1/auth/login | \
  jq -r '.token'
)

wrk -c200 -d30s -t100 -sget_wrk_test.lua \
    -H "Cookie: jwt-token=${JWT_TOKEN}" \
    http://localhost:8084/api/v1/playlists