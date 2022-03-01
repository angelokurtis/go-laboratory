package wire

import (
	"github.com/google/wire"

	"github.com/angelokurtis/go-laboratory/internal/k8s"
	"github.com/angelokurtis/go-laboratory/internal/kustomize"
	"github.com/angelokurtis/go-laboratory/internal/zap"
	"github.com/angelokurtis/go-laboratory/internal/zapr"
)

var Providers = wire.NewSet(
	k8s.Providers,
	kustomize.Providers,
	zap.Providers,
	zapr.Providers,
)
