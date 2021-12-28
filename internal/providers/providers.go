package providers

import (
	"github.com/google/wire"

	"github.com/angelokurtis/go-laboratory/internal/zap"
)

var Set = wire.NewSet(
	zap.Providers,
)
