package main

import (
	"context"
	_ "embed"
	"os"
	"path"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/angelokurtis/go-laboratory/k8s-server-side-apply/internal/fatality"
	"github.com/angelokurtis/go-laboratory/k8s-server-side-apply/internal/logging"
	"github.com/angelokurtis/go-laboratory/k8s-server-side-apply/internal/manifest"
)

//go:embed manifests/nginx.yaml
var data []byte

func main() {
	_ = logging.Setup()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fatality.With(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", path.Join(homeDir, ".kube/config"))
	if err != nil {
		fatality.With(err)
	}

	mr, err := manifest.NewReader(config)
	if err != nil {
		fatality.With(err)
	}

	m, err := mr.FromBytes(data)
	if err != nil {
		fatality.With(err)
	}

	m, err = m.Transform(func(u *unstructured.Unstructured) error {
		u.SetNamespace("default")
		return nil
	})
	if err != nil {
		fatality.With(err)
	}

	if err = m.Apply(context.Background()); err != nil {
		fatality.With(err)
	}
}
