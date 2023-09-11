package main

import (
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hako/durafmt"
	cli "github.com/urfave/cli/v2"
)

func pessimisticAction(cCtx *cli.Context) error {
	user := cCtx.String("user")
	amount := cCtx.Float64("amount")
	executions := cCtx.Int("executions")

	repo, cleanup, err := initializePessimistic()
	if err != nil {
		return err
	}

	defer cleanup()

	var retries, errs uint64

	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(executions)

	for i := 0; i < executions; i++ {
		go func() {
			defer wg.Done()

			if derr := repo.Deposit(cCtx.Context, user, amount); derr != nil {
				atomic.AddUint64(&errs, 1)
			}
		}()
	}
	wg.Wait()
	slog.Info(fmt.Sprintf("During %v executions, we had %v errors, even after retrying %v times in just %v.", executions, errs, retries, durafmt.ParseShort(time.Since(start))))

	return nil
}
