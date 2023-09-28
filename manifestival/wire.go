//go:build wireinject
// +build wireinject

package main

import (
	"log/slog"

	"github.com/google/wire"

	"github.com/angelokurtis/go-laboratory/manifestival/internal/fatality"
	"github.com/angelokurtis/go-laboratory/manifestival/internal/kubernetes"
	"github.com/angelokurtis/go-laboratory/manifestival/internal/logging"
	"github.com/angelokurtis/go-laboratory/manifestival/internal/manifestival"
	"github.com/angelokurtis/go-laboratory/manifestival/internal/maxprocs"
)

var providers = wire.NewSet(
	kubernetes.NewConfig,
	logging.NewHandler,
	manifestival.NewClient,
	manifestival.NewLogger,
	manifestival.NewFactory,
	slog.Default,
)

type Application struct {
	*manifestival.Factory
}

func init() {
	logger := logging.Setup()
	if err := maxprocs.SetupWithLogger(logger); err != nil {
		fatality.With(err)
	}
}

func initialize() (*Application, func(), error) {
	wire.Build(
		providers,
		wire.Struct(new(Application), "*"),
	)

	return nil, nil, nil
}
