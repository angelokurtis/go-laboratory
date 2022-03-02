//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/go-logr/logr"
	"github.com/google/wire"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/dynamic"

	"github.com/angelokurtis/go-laboratory/internal/kustomize"
)

func ZapLogger() (*zap.Logger, error) {
	wire.Build(Providers)
	return nil, nil
}

func LogrLogger() (logr.Logger, error) {
	wire.Build(Providers)
	return logr.Logger{}, nil
}

func DynamicClient() (dynamic.Interface, error) {
	wire.Build(Providers)
	return nil, nil
}

type dynamicClientAndMapper struct {
	Dynamic    dynamic.Interface
	RESTMapper meta.RESTMapper
}

func DynamicClientAndMapper() (*dynamicClientAndMapper, error) {
	wire.Build(Providers, wire.Struct(new(dynamicClientAndMapper), "*"))
	return nil, nil
}

func KustomizeClient() (*kustomize.Client, error) {
	wire.Build(Providers)
	return nil, nil
}

type clientAndMapping struct {
	RESTClient resource.RESTClient
	RESTMapper meta.RESTMapper
}

func RESTClientAndRESTMapping() (*clientAndMapping, error) {
	wire.Build(Providers, wire.Struct(new(clientAndMapping), "*"))
	return nil, nil
}
