package main

import (
	"log/slog"

	mf "github.com/manifestival/manifestival"
	"github.com/pkg/errors"

	"github.com/angelokurtis/go-laboratory/manifestival/internal/fatality"
)

func main() {
	slog.Info("Starting")

	app, cleanup, err := initialize()
	if err != nil {
		fatality.With(err)
	}

	defer cleanup()

	m, err := app.ManifestFrom(mf.Path("./manifests/"))
	if err != nil {
		fatality.With(err)
	}

	m, err = m.Transform(mf.InjectNamespace("tks"))
	if err != nil {
		fatality.With(errors.WithStack(err))
	}

	if err = m.Apply(); err != nil {
		fatality.With(err)
	}
}
