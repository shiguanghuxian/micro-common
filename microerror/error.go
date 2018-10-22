package microerror

import (
	"encoding/json"
)

/* 微服务返回错误代码列表 */

var (
	// ErrRecordNotFound 数据未查询到
	ErrRecordNotFound = NewMicroError(10001, "record not found")
	// ErrRedisCmd redis 操作错误
	ErrRedisCmd = NewMicroError(10002, "record not found")

	// ErrUnknownServerError 服务端错误
	ErrUnknownServerError = NewMicroError(30000, "Unknown server error")
)

// MicroError 错误类型
type MicroError struct {
	Msg  string `json:"msg"`  // 错误信息
	Code int16  `json:"code"` // 错误代码
}

// NewMicroError 创建MicroError
func NewMicroError(code int16, msg string) *MicroError {
	return &MicroError{
		Code: code,
		Msg:  msg,
	}
}

// Error 实现error接口
func (err *MicroError) Error() string {
	js, _ := json.Marshal(err)
	return string(js)
}
