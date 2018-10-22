package log

import (
	"log"

	"github.com/shiguanghuxian/micro-common/config"
	"go.uber.org/zap"
)

var (
	Logger *Log
)

func init() {
	mode := config.GetMode()
	sugar, err := InitLogger("./logs", mode == "dev")
	if err != nil {
		log.Panicln(err)
	}
	Logger = NewLog(sugar)
}

// Log 实现go-kit日志接口
type Log struct {
	*zap.SugaredLogger
}

// NewLog 创建日志对象
func NewLog(logger *zap.SugaredLogger) *Log {
	return &Log{logger}
}

// Log 输出日志
func (l *Log) Log(keyvals ...interface{}) error {
	l.Error(keyvals...)
	return nil
}
