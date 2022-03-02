package k8s

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

var Providers = wire.NewSet(
	NewAuthorizationClient,
	NewConfig,
	NewDiscoveryClient,
	NewDynamicClient,
	NewDynamicRESTMapper,
	NewRESTClient,
)

func NewConfig() *rest.Config {
	config := ctrl.GetConfigOrDie()
	return config
}

func NewDynamicClient(config *rest.Config) dynamic.Interface {
	client := dynamic.NewForConfigOrDie(config)
	return client
}

func NewDiscoveryClient(config *rest.Config) discovery.DiscoveryInterface {
	client := discovery.NewDiscoveryClientForConfigOrDie(config)
	return client
}

func NewAuthorizationClient(config *rest.Config) authorizationv1.AuthorizationV1Interface {
	client := kubernetes.NewForConfigOrDie(config).AuthorizationV1()
	return client
}

func NewDynamicRESTMapper(config *rest.Config) (meta.RESTMapper, error) {
	mapper, err := apiutil.NewDynamicRESTMapper(config)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return mapper, nil
}

func NewRESTClient(config *rest.Config) (resource.RESTClient, error) {
	//client, err := rest.RESTClientFor(config)
	//if err != nil {
	//	return nil, errors.WithStack(err)
	//}
	return nil, nil
}
