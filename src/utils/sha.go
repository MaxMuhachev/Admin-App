package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Encode(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Decode(str string) ([]byte, error) {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.DecodeString(str)
}
