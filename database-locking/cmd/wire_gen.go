// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/metrics"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/mysql"
	"github.com/angelokurtis/go-laboratory/database-locking/internal/persistence"
	"github.com/google/wire"
)

// Injectors from wire.go:

func initialize() (*X, func(), error) {
	db, err := mysql.NewDB()
	if err != nil {
		return nil, nil, err
	}

	queries := persistence.New(db)
	x := &X{
		DB:      db,
		Queries: queries,
	}

	return x, func() {
	}, nil
}

// wire.go:

var providers = wire.NewSet(metrics.NewHandler, mysql.NewDB, persistence.New, wire.Struct(new(X), "*"), wire.Bind(new(persistence.DBTX), new(*sql.DB)))

type X struct {
	DB      *sql.DB
	Queries *persistence.Queries
}
