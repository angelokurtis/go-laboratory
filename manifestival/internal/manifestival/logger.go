package manifestival

import (
	"log/slog"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/slogr"
)

func NewLogger(handler slog.Handler) logr.Logger {
	return slogr.NewLogr(handler)
}
