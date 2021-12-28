//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/angelokurtis/go-laboratory/internal/providers"
)

func NewLogger() (*zap.Logger, error) {
	wire.Build(providers.Set)
	return nil, nil
}
