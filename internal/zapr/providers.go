//go:build wireinject
// +build wireinject

package zapr

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewLogger,
)
