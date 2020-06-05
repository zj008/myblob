package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) (string, error) {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:]), nil
}
