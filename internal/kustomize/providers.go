package kustomize

import (
	"github.com/google/wire"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

var Providers = wire.NewSet(
	NewClient,
	NewFileSystem,
	NewKustomizer,
)

func NewKustomizer() *krusty.Kustomizer {
	return krusty.MakeKustomizer(krusty.MakeDefaultOptions())
}

func NewFileSystem() filesys.FileSystem {
	return filesys.MakeFsOnDisk()
}
