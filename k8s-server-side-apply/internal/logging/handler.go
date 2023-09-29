package logging

import "log/slog"

func NewHandler(logger *slog.Logger) slog.Handler {
	return logger.Handler()
}
