package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

/* 获取md5值 */

// Md5String 获取字符串md5值
func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
