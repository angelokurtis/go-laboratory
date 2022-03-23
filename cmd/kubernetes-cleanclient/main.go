package main

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client/config"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	config, err := config.GetConfigWithContext("arn:aws:eks:us-east-1:839895334915:cluster/charlescd-sandbox-dev")
	//config, err := config.GetConfigWithContext("kind-flux")
	if err != nil {
		panic(err)
	}

	httpClient, err := newHttpClient(config)
	if err != nil {
		panic(err)
	}

	client, _ := kubernetes.NewForConfigAndClient(config, httpClient)
	namespace := "default"
	pods, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d pods were found on namespace %q\n", len(pods.Items), namespace)
}
