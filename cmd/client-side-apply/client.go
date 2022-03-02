package main

import (
	"context"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
)

type Client struct {
	client dynamic.Interface
	mapper meta.RESTMapper
}

func NewClient(client dynamic.Interface, mapper meta.RESTMapper) *Client {
	return &Client{client: client, mapper: mapper}
}

func (c *Client) Get(ctx context.Context, obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	gvk := obj.GroupVersionKind()
	mapping, err := c.mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if mapping.Scope.Name() == meta.RESTScopeNameRoot {
		result, err := c.client.Resource(mapping.Resource).
			Get(ctx, obj.GetName(), metav1.GetOptions{})
		if kerrors.IsNotFound(err) {
			return nil, nil
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return result, nil
	}

	result, err := c.client.Resource(mapping.Resource).
		Namespace(obj.GetNamespace()).
		Get(ctx, obj.GetName(), metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (c *Client) Create(ctx context.Context, obj *unstructured.Unstructured) {

}
