package kustomize

import (
	"os"
	"path"

	"github.com/pkg/errors"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

type Client struct {
	fs         filesys.FileSystem
	kustomizer *krusty.Kustomizer
}

func NewClient(fs filesys.FileSystem, kustomizer *krusty.Kustomizer) *Client {
	return &Client{fs: fs, kustomizer: kustomizer}
}

func (c *Client) GenerateConfigMap(dir string, generator *ConfigMapGenerator) error {
	content, err := generator.Marshal()
	if err != nil {
		return errors.Wrap(err, "config map generation failed")
	}

	file := path.Join(dir, "kustomization.yaml")
	err = os.WriteFile(file, content, 0600)
	if err != nil {
		return errors.Wrap(err, "config map generation failed")
	}
	defer os.Remove(file)

	m, err := c.kustomizer.Run(c.fs, dir)
	if err != nil {
		return errors.Wrap(err, "config map generation failed")
	}

	yml, err := m.AsYaml()
	if err != nil {
		return errors.Wrap(err, "config map generation failed")
	}

	return nil
}
