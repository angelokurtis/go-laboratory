package main

import (
	"context"
	"log"

	"github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"

	"github.com/angelokurtis/go-laboratory/database-locking/internal/persistence"
)

func main() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	if err := run(); err != nil {
		if sterr, ok := err.(stackTracer); ok {
			st := sterr.StackTrace()
			log.Fatalf("%s%+v\n", err.Error(), st[:len(st)-2])
		} else {
			log.Fatalln(err)
		}
	}
}

func run() error {
	ctx := context.Background()

	x, cleanup, err := initialize()
	if err != nil {
		return errors.WithStack(err)
	}

	defer cleanup()

	tx, err := x.DB.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		if rerr := tx.Rollback(); rerr != nil {
			log.Printf("[WARN] %s\n", rerr.Error())
		}
	}()

	qtx := x.Queries.WithTx(tx)

	account, err := qtx.GetAccountAndLockForUpdates(ctx, "kurtis")
	if err != nil {
		return errors.WithStack(err)
	}

	res, err := qtx.UpdateAccountBalance(ctx, persistence.UpdateAccountBalanceParams{
		ID:      account.ID,
		Balance: 125,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.WithStack(err)
	}

	log.Printf("%d rows have been impacted in total.\n", rowsAffected)

	if err = tx.Commit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
