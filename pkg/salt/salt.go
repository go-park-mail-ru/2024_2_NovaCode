package salt

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func Generate(length int) (string, error) {
	if length < 1 {
		return "", errors.New("length must be a positive integer")
	}

	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(salt), nil
}
