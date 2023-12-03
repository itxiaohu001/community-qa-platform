package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5ByStr(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
