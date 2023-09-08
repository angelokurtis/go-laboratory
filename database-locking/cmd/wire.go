//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"

	"github.com/angelokurtis/go-laboratory/database-locking/internal/metrics"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/mysql"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/persistence"
)

var providers = wire.NewSet(
	metrics.NewHandler,
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
