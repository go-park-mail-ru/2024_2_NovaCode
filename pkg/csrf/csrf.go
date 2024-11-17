package csrf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Generate(key, salt string) string {
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(key))
	token := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(token)
}

func Validate(token, key, salt string) bool {
	return token == Generate(key, salt)
}
