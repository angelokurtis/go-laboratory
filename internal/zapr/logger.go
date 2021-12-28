package zapr

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

func NewLogger(zapLog *zap.Logger) logr.Logger {
	return zapr.NewLogger(zapLog)
}
