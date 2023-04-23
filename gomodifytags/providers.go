package main

import "github.com/google/wire"

var providers = wire.NewSet(
	wire.Struct(new(Service), "*"),
)

type Service struct {
	// find me
	Name string
}
