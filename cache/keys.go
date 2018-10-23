package cache

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shiguanghuxian/micro-common/config"
)

/* 格式化redis key */

// GetPrefixKey 给key添加前缀
func GetPrefixKey(key string) string {
	return "vivi/" + config.GetSvcName("") + key
}

// GetUserLoginToken 获取token存储key
func GetUserLoginToken() (key, token string) {
	tokenUUID, _ := uuid.NewV4()
	token = tokenUUID.String()
	key = GetPrefixKey("login/token/" + token)
	return
}
