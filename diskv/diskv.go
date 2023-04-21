package main

import (
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

func NewDiskv(directory CurrentDirectory) *diskv.Diskv {
	return diskv.New(diskv.Options{
		BasePath:          string(directory),
		CacheSizeMax:      1024 * 1024,
		AdvancedTransform: AdvancedTransform,
		InverseTransform:  InverseTransform,
	})
}

func AdvancedTransform(key string) *diskv.PathKey {
	path := strings.Split(key, "/")
	last := len(path) - 1

	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last] + ".txt",
	}
}

func InverseTransform(pathKey *diskv.PathKey) (key string) {
	txt := pathKey.FileName[len(pathKey.FileName)-4:]
	if txt != ".txt" {
		panic("Invalid file found in storage folder!")
	}

	return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)-4]
}
