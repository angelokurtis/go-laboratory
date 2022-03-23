package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/angelokurtis/go-laboratory/internal/log"
)

func main() {
	config, err := config.GetConfigWithContext("arn:aws:eks:us-east-1:671982376808:cluster/horus-dev")
	//config, err := config.GetConfigWithContext("kind-flux")
	if err != nil {
		log.Fatal(err)
	}

	httpClient, err := newHttpClient(config)
	if err != nil {
		log.Fatal(err)
	}

	client, _ := kubernetes.NewForConfigAndClient(config, httpClient)
	namespace := "default"
	pods, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("%d pods were found on namespace %q\n", len(pods.Items), namespace)
}
