package maxprocs

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"time"

	"github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/automaxprocs/maxprocs"
)

func SetupWithLogger(logger *slog.Logger) error {
	opts := maxprocs.Logger(func(format string, args ...interface{}) {
		infof(logger, format, args)
	})
	if _, err := maxprocs.Set(opts); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func infof(logger *slog.Logger, format string, args ...any) {
	if !logger.Enabled(context.Background(), slog.LevelInfo) {
		return
	}

	var pcs [1]uintptr

	runtime.Callers(3, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(format, args...), pcs[0])
	_ = logger.Handler().Handle(context.Background(), r)
}
