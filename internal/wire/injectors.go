//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/go-logr/logr"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func ZapLogger() (*zap.Logger, error) {
	wire.Build(Providers)
	return nil, nil
}

func LogrLogger() (logr.Logger, error) {
	wire.Build(Providers)
	return logr.Logger{}, nil
}
