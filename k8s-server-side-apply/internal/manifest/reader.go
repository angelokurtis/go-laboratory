package manifest

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

type Reader struct {
	client dynamic.Interface
	mapper meta.RESTMapper
}

func NewReader(config *rest.Config) (*Reader, error) {
	httpClient, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize the manifest reader: %w", err)
	}

	mapper, err := apiutil.NewDynamicRESTMapper(config, httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize the manifest reader: %w", err)
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize the manifest reader: %w", err)
	}

	return &Reader{client: client, mapper: mapper}, nil
}

func (s *Reader) FromBytes(data []byte) (List, error) {
	reader := bytes.NewReader(data)
	decoder := yaml.NewYAMLToJSONDecoder(reader)

	var resources []*unstructured.Unstructured

	var err error

	for {
		out := &unstructured.Unstructured{}

		err = decoder.Decode(out)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil || len(out.Object) == 0 {
			continue
		}

		resources = append(resources, out)
	}

	if !errors.Is(err, io.EOF) {
		return &list{}, fmt.Errorf("unable to parse manifest from bytes: %w", err)
	}

	return &list{resources: resources, client: s.client, mapper: s.mapper}, nil
}
