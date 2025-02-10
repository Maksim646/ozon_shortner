package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BuildLogger(level string) error {
	var loggerLevel zapcore.Level
	if err := loggerLevel.UnmarshalText([]byte(level)); err != nil {
		return fmt.Errorf("cannot parse logger level: %w", err)
	}

	loggerCfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(loggerLevel),
		Encoding:         "json",
		OutputPaths:      []string{"./zap.log"},
		ErrorOutputPaths: []string{"./zap.log"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:      "level",
			TimeKey:       "time",
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			CallerKey:     "caller",
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	logger, err := loggerCfg.Build()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	return nil
}
