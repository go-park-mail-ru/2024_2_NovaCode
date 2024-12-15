RESPONSE=$(
  curl -X POST \
       -H "Content-Type: application/json" \
       -d '{"username": "testt", "password": "12345678"}' \
       http://localhost:8081/api/v1/auth/login
)

JWT_TOKEN=$(echo $RESPONSE | jq -r '.token')
USER_ID=$(echo $RESPONSE | jq -r '.user.id')

wrk -c150 -d20s -t100 -sget_wrk_test.lua \
    -H "Cookie: jwt-token=${JWT_TOKEN}" \
    http://localhost:8084/api/v1/playlists/${USER_ID}/allPlaylists
