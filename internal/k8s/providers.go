package k8s

import (
	"github.com/google/wire"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
)

var Providers = wire.NewSet(
	NewAuthorizationClient,
	NewConfig,
	NewDiscoveryClient,
	NewDynamicClient,
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
