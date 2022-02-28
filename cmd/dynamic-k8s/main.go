package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"

	"github.com/angelokurtis/go-laboratory/internal/wire"
)

func main() {
	client, err := wire.DynamicClient()
	dieIfNotNil(err)

	managed, err := labels.NewRequirement("app.kubernetes.io/managed-by", selection.Equals, []string{"apirator"})
	dieIfNotNil(err)

	name, err := labels.NewRequirement("app.kubernetes.io/name", selection.Equals, []string{"petstore-expanded"})
	dieIfNotNil(err)

	selector := labels.NewSelector()
	selector = selector.Add(*managed, *name)

	resources, err := client.Resource(schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}).Namespace("default").List(context.TODO(), metav1.ListOptions{LabelSelector: selector.String()})
	dieIfNotNil(err)

	_ = resources
}

func dieIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
