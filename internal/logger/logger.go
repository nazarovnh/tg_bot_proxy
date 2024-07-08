package logger

import (

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger() (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zapConfig.DisableStacktrace = true
	// zapConfig.DisableCaller = true // uncomment for debug
	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	logger.Info("success setup zap logger")

	return logger, nil
}