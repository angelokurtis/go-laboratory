//go:build wireinject
// +build wireinject

package zap

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewLogger,
	NewConfig,
	NewEncoderConfig,
)
