package zap

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg zap.Config) (*zap.Logger, error) {
	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("error while creating the logger: %w", err)
	}
	return logger, nil
}

func NewConfig(encoder zapcore.EncoderConfig) zap.Config {
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableStacktrace: true,
		Encoding:          "console",
		EncoderConfig:     encoder,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	return config
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
