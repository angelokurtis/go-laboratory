//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"

	"dev/labs/database-locking/internal/mysql"
	"dev/labs/database-locking/internal/persistence"
)

var providers = wire.NewSet(
	mysql.NewDB,
	persistence.New,
	wire.Struct(new(X), "*"),
	wire.Bind(new(persistence.DBTX), new(*sql.DB)),
)

type X struct {
	DB      *sql.DB
	Queries *persistence.Queries
}

func initialize() (*X, func(), error) {
	wire.Build(
		providers,
	)
	return nil, nil, nil
}
