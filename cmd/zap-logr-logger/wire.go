//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-logr/logr"
	"github.com/google/wire"
)

func NewLogger() (logr.Logger, error) {
	wire.Build(providers.Set)
	return logr.Logger{}, nil
}
