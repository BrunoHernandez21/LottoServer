package mysha

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(toencode string) string {
	h := sha256.New()
	h.Write([]byte(toencode))
	i := hex.EncodeToString(h.Sum(nil))
	return i
}

func SHA1(toencode string) string {
	h := sha1.New()
	h.Write([]byte(toencode))
	i := hex.EncodeToString(h.Sum(nil))
	return i
}
