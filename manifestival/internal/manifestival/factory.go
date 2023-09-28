package manifestival

import (
	"github.com/go-logr/logr"
	mf "github.com/manifestival/manifestival"
	"github.com/pkg/errors"
)

type Factory struct {
	client mf.Client
	logger logr.Logger
}

func NewFactory(client mf.Client, logger logr.Logger) *Factory {
	return &Factory{client: client, logger: logger}
}

func (f *Factory) ManifestFrom(src mf.Source, opts ...mf.Option) (mf.Manifest, error) {
	opts = append(opts, mf.UseClient(f.client), mf.UseLogger(f.logger))

	m, err := mf.ManifestFrom(src, opts...)
	if err != nil {
		return mf.Manifest{}, errors.WithStack(err)
	}

	return m, nil
}
