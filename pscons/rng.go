package pscons

import (
	"crypto/rand"
	"encoding/base64"
)

func genBytes(len int) ([]byte, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func genPassword() (string, error) {
	b, err := genBytes(30)
	return base64.URLEncoding.EncodeToString(b), err
}
