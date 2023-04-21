package main

import "github.com/google/wire"

var providers = wire.NewSet(
	NewBoltDB,
	NewCurrentDirectory,
	wire.Struct(new(Service), "*"),
)
