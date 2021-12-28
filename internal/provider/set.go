//go:build wireinject
// +build wireinject

package providers

import (
	"github.com/google/wire"

	"github.com/angelokurtis/go-laboratory/internal/zap"
	"github.com/angelokurtis/go-laboratory/internal/zapr"
)

var Set = wire.NewSet(
	zap.Providers,
	zapr.Providers,
)
