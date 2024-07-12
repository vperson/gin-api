package util

import (
	"crypto/md5"
	"encoding/hex"
)

// StringToMD5 将字符串转换为MD5
func StringToMD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
