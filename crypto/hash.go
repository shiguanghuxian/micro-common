package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

/* 获取hash值 */

// Md5 获取字符串md5值
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 获取字符串sha1值
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256 获取字符串sha256值
func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512 获取字符串sha512值
func Sha512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// PasswordHash 密码hash
func PasswordHash(password string, salt ...string) string {
	if len(salt) == 0 {
		salt = append(salt, "")
	}
	return Md5(Sha512(Sha256(password)+salt[0]) + salt[0])
}
