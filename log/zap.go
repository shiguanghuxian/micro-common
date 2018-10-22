package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger go.uber.org/zap 初始化
func InitLogger(lp string, isDebug bool) (*zap.SugaredLogger, error) {
	cfg := &zap.Config{
		Encoding: "json",
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	atom := zap.NewAtomicLevel()
	if isDebug == true {
		atom.SetLevel(zapcore.DebugLevel)
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stdout"}
	} else {
		atom.SetLevel(zapcore.WarnLevel)
		cfg.OutputPaths = []string{lp}
		cfg.ErrorOutputPaths = []string{lp}
	}
	cfg.Level = atom
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
