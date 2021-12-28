package zap

import (
	"fmt"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var Providers = wire.NewSet(
	NewLogger,
)

func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("error while creating the logger: %w", err)
	}
	return logger, nil
}
