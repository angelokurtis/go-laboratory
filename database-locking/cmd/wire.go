//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	_ "go.uber.org/automaxprocs"

	"github.com/angelokurtis/go-laboratory/database-locking/internal/account"
	_ "github.com/angelokurtis/go-laboratory/database-locking/internal/logging"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/metrics"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/mysql"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/persistence"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/redis"
)

var providers = wire.NewSet(
	account.NewDefaultRepository,
	account.NewDistributedRepository,
	account.NewOptimisticRepository,
	account.NewPessimisticRepository,
	metrics.NewHandler,
	mysql.NewDB,
	persistence.New,
	redis.NewClient,
	redis.NewPool,
	redis.NewRedsync,
)

func initialize() (account.Repository, func(), error) {
	wire.Build(
		providers,
		wire.Bind(new(account.Repository), new(*account.DefaultRepository)),
		wire.Bind(new(persistence.DBTX), new(*sql.DB)),
	)

	return nil, nil, nil
}

func initializeOptimistic() (account.Repository, func(), error) {
	wire.Build(
		providers,
		wire.Bind(new(account.Repository), new(*account.OptimisticRepository)),
		wire.Bind(new(persistence.DBTX), new(*sql.DB)),
	)

	return nil, nil, nil
}

func initializePessimistic() (account.Repository, func(), error) {
	wire.Build(
		providers,
		wire.Bind(new(account.Repository), new(*account.PessimisticRepository)),
		wire.Bind(new(persistence.DBTX), new(*sql.DB)),
	)

	return nil, nil, nil
}

func initializeDistributed() (account.Repository, func(), error) {
	wire.Build(
		providers,
		wire.Bind(new(account.Repository), new(*account.DistributedRepository)),
		wire.Bind(new(persistence.DBTX), new(*sql.DB)),
	)

	return nil, nil, nil
}
