package main

import "github.com/google/wire"

var providers = wire.NewSet(
	NewCurrentDirectory,
	NewDiskv,
	wire.Struct(new(Service), "*"),
)
