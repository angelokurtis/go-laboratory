package manifestival

import (
	mfc "github.com/manifestival/client-go-client"
	mf "github.com/manifestival/manifestival"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
)

func NewClient(config *rest.Config) (mf.Client, error) {
	client, err := mfc.NewClient(config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return client, nil
}
