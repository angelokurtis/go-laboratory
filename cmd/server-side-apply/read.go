package main

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
)

//go:embed grafana.yaml
var grafana []byte

func getObjects() ([]*unstructured.Unstructured, error) {
	reader := bytes.NewReader(grafana)
	decoder := yaml.NewYAMLToJSONDecoder(reader)
	var objs []*unstructured.Unstructured
	var err error
	for {
		out := &unstructured.Unstructured{}
		err = decoder.Decode(out)
		if err == io.EOF {
			break
		}
		if err != nil || len(out.Object) == 0 {
			continue
		}
		objs = append(objs, out)
	}
	if err != io.EOF {
		return nil, errors.WithStack(err)
	}
	return objs, nil
}
