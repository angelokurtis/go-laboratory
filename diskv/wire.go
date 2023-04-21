//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func initialize() (*Service, error) {
	wire.Build(providers)
	return nil, nil
}
