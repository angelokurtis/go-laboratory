package kubernetes

import (
	"os"
	"path"

	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewConfig() (*rest.Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", path.Join(homeDir, ".kube/config"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return config, nil
}
