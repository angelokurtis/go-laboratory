package main

import (
	"context"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/dynamic"
)

type Client struct {
	dynamic dynamic.Interface
	client  resource.RESTClient
	mapper  meta.RESTMapper
}

func NewClient(dynamic dynamic.Interface, mapper meta.RESTMapper) *Client {
	return &Client{dynamic: dynamic, mapper: mapper}
}

func (c *Client) Find(ctx context.Context, obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	gvk := obj.GroupVersionKind()
	mapping, err := c.mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if mapping.Scope.Name() == meta.RESTScopeNameRoot {
		result, err := c.dynamic.Resource(mapping.Resource).Get(ctx, obj.GetName(), metav1.GetOptions{})
		if kerrors.IsNotFound(err) {
			return nil, nil
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return result, nil
	}

	result, err := c.dynamic.Resource(mapping.Resource).Namespace(obj.GetNamespace()).Get(ctx, obj.GetName(), metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (c *Client) Apply(ctx context.Context, obj *unstructured.Unstructured) error {
	gvk := obj.GroupVersionKind()
	mapping, err := c.mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return errors.WithStack(err)
	}

	data, err := runtime.Encode(unstructured.UnstructuredJSONScheme, obj)
	if err != nil {
		return errors.WithStack(err)
	}

	force := false
	if mapping.Scope.Name() == meta.RESTScopeNameRoot {
		_, err = c.dynamic.Resource(mapping.Resource).Patch(
			ctx,
			obj.GetName(),
			types.ApplyPatchType,
			data,
			metav1.PatchOptions{
				Force:        &force,
				FieldManager: "go-laboratory",
			},
		)
	} else {
		_, err = c.dynamic.Resource(mapping.Resource).Namespace(obj.GetNamespace()).Patch(
			ctx,
			obj.GetName(),
			types.ApplyPatchType,
			data,
			metav1.PatchOptions{
				Force:        &force,
				FieldManager: "go-laboratory",
			},
		)
	}
	if kerrors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
